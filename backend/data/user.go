package data

import (
	"crypto/sha256"
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

func CalcHash(plaintext string) string {
	r := sha256.Sum256([]byte(plaintext))
	return fmt.Sprintf("%x", r)
}

// create a new user, save user info into the database
// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) (err error) {
	// get user info from request body
	name := r.FormValue("name")
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

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreateUser(w, r)
	}
}
