package data

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// for response
type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// get a user from email
func RetrieveUserFromEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users where email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// get a session from uuid
func RetrieveSessionFromUuid(uuid string) (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions where uuid = $1", uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func CalcHash(plaintext string) string {
	r := sha256.Sum256([]byte(plaintext))
	return fmt.Sprintf("%x", r)
}

// create a new user, save user info into the database
// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) (err error) {
	// get user info from form
	name := r.FormValue("user-name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// generate uuid for new user
	u4, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create new user
	// insert into users table
	uuid := u4.String()
	_, err = Db.Exec("INSERT INTO users (uuid, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)", uuid, name, email, CalcHash(password), time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

// get user info from session uuid in cookie
// GET /users/me
// NOTE: Does this function suit for RESTful API?
func GetMyInfo(w http.ResponseWriter, r *http.Request) (err error) {
	// get session uuid from cookie
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
		if err == http.ErrNoCookie {
			fmt.Println("No cookie")
			return
		}
		return
	}
	uuid := cookie.Value

	// get session information
	session, err := RetrieveSessionFromUuid(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get user information
	user, err := RetrieveUserFromEmail(session.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create response json
	userInfo := UserInfo{
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return
}

func CreateSession(user User) (session Session, err error) {
	// generate uuid for new session
	u4, err := uuid.NewV4()
	if err != nil {
		return
	}
	uuid := u4.String()

	// insert into sessions table
	_, err = Db.Exec("INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4)", uuid, user.Email, user.Id, time.Now())
	if err != nil {
		return
	}

	// get new session information
	session, err = RetrieveSessionFromUuid(uuid)
	return
}

// check email and password
// /POST /sessions
func Authenticate(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// get user info form
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := RetrieveUserFromEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if user.Password != CalcHash(password) {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else {
		// generate uuid for new session
		session, err := CreateSession(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		cookie := http.Cookie{
			Name:  "cookie",
			Value: session.Uuid,
			// HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		fmt.Println(cookie)
		fmt.Println("Set cookie")
		return err
	}
}

func HandleUsersMe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetMyInfo(w, r)
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreateUser(w, r)
	}
}

func HandleSessions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		Authenticate(w, r)
	}
}
