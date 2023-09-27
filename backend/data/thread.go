package data

import (
	"encoding/json"
	"net/http"
	"time"
)

type Thread struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Topic     string    `json:"topic"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// get all threads in the database and returns it
func Threads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	threads := make([]Thread, 0)
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		threads = append(threads, thread)
	}
	rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(threads)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return
}
