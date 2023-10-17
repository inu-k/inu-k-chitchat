package data

import (
	"encoding/json"
	"net/http"
	"path"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type ThreadInfo struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Topic     string    `json:"topic"`
	UserId    int       `json:"userId"`
	PostsNum  int       `json:"postsNum"`
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

// get all threads in the database with a number of posts and returns it
// GET /threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	rows, err := Db.Query("select A.*, COUNT(B.id) as post_num from threads as A left join posts as B on B.thread_id=A.id group by A.id, A.uuid, A.topic, A.user_id, A.created_at ORDER BY A.created_at DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	threads := make([]ThreadInfo, 0)
	for rows.Next() {
		thread := ThreadInfo{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt, &thread.PostsNum); err != nil {
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

// get thread information
// GET /threads/thread_uuid
func GetThread(w http.ResponseWriter, r *http.Request) {
	thread_uuid := path.Base(r.URL.Path)
	thread_info := ThreadInfo{}
	err := Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", thread_uuid).Scan(&thread_info.Id, &thread_info.Uuid, &thread_info.Topic, &thread_info.UserId, &thread_info.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(thread_info)
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

type TopicInfo struct {
	Topic string `json:"topic"`
}

// create a new thread
// POST /threads
func CreateThread(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var topicInfo TopicInfo
	json.Unmarshal(body, &topicInfo)
	topic := topicInfo.Topic
	u4, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	uuid := u4.String()
	_, err = Db.Query("INSERT INTO threads (uuid, topic, user_id, created_at) VALUES ($1, $2, $3, $4)", uuid, topic, 1, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get new thread information
	thread_info := ThreadInfo{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).Scan(&thread_info.Id, &thread_info.Uuid, &thread_info.Topic, &thread_info.UserId, &thread_info.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(thread_info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return
}

func HandleThreads(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetThreads(w, r)
	case "POST":
		CreateThread(w, r)
	}
	return
}
