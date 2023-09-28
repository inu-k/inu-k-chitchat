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
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Body      string    `json:"body"`
	UserId    int       `json:"userId"`
	ThreadId  int       `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
}

// get all threads in the database and returns it
// GET /threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
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

// get posts to a thread
// GET /posts?thread_uuid=1
func GetPosts(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	threadUuid := vals.Get("thread_uuid")
	var threadId int
	err := Db.QueryRow("SELECT id FROM threads WHERE uuid = $1", threadUuid).Scan(&threadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts := make([]Post, 0)
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1 ORDER BY created_at", threadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}
	rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return
}
