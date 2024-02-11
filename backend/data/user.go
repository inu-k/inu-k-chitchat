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
	Uuid  string `json:"uuid"`
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

// create a new session, save session info into the database
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

	cookie, err := r.Cookie("_cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		Uuid:  user.Uuid,
	}

	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return
}

// check email and password
// /POST /sessions
func Authenticate(w http.ResponseWriter, r *http.Request) (err error) {
	// get user info form
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := RetrieveUserFromEmail(email)
	if err != nil {
		http.Error(w, "cannot login", http.StatusForbidden)
		return err
	}

	if user.Password != CalcHash(password) {
		http.Error(w, "cannot login", http.StatusForbidden)
		return err
	} else {
		// generate uuid for new session
		session, err := CreateSession(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		cookie := http.Cookie{
			Name:  "_cookie",
			Value: session.Uuid,
			// HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		return err
	}
}

// delete session by uuid
func DeleteSessionByUuid(uuid string) (err error) {
	_, err = Db.Exec("DELETE FROM sessions WHERE uuid = $1", uuid)
	return
}

// delete session
// DELETE /sessions/me
func DeleteSessionMe(w http.ResponseWriter, r *http.Request) (err error) {
	// get session uuid from cookie
	cookie, err := r.Cookie("_cookie")
	if err == http.ErrNoCookie {
		fmt.Println("Warning: No cookie")
		w.WriteHeader(http.StatusOK)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uuid := cookie.Value

	// delete cookie
	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1
	cookie.Path = "/"
	cookie.Domain = "localhost"

	// delete session
	err = DeleteSessionByUuid(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	return
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

func HandleSessionsMe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		w.WriteHeader(http.StatusNoContent)
	case "DELETE":
		DeleteSessionMe(w, r)
	}
}
