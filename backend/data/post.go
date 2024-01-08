package data

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type Post struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Body      string    `json:"body"`
	UserId    int       `json:"userId"`
	ThreadId  int       `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewPostInfo struct {
	Body       string `json:"body"`
	UserId     int    `json:"userId"`
	ThreadUuid string `json:"threadUuid"`
}

// get posts to a thread, given a thread id
func RetrievePostsFromThreadId(threadId int) ([]Post, error) {
	posts := make([]Post, 0)
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_id = $1 ORDER BY created_at", threadId)
	if err != nil {
		return nil, err
	}
	// make list of posts
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, nil
}

// get a post from a post uuid
func RetrievePostFromUuid(postUuid string) (Post, error) {
	post := Post{}
	err := Db.QueryRow("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts WHERE uuid = $1", postUuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	if err != nil {
		return post, err
	}
	return post, nil
}

// get a thread id from a thread uuid
func RetrieveThreadIdFromUuid(threadUuid string) (int, error) {
	var threadId int
	err := Db.QueryRow("SELECT id FROM threads WHERE uuid = $1", threadUuid).Scan(&threadId)
	if err != nil {
		return -1, err
	}
	return threadId, nil
}

// get posts to a thread
// GET /posts?thread_uuid=1
func GetPosts(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	threadUuid := vals.Get("thread_uuid")
	threadId, err := RetrieveThreadIdFromUuid(threadUuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := RetrievePostsFromThreadId(threadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

// create a new post
// POST /posts
func CreatePost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var newPostInfo NewPostInfo
	json.Unmarshal(body, &newPostInfo)
	bodyText := newPostInfo.Body
	userId := newPostInfo.UserId
	// fmt.Println(newPostInfo.ThreadUuid)
	threadId, err := RetrieveThreadIdFromUuid(newPostInfo.ThreadUuid)
	// fmt.Println(threadId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// generate uuid for new post
	u4, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// insert new post into database
	uuid := u4.String()
	_, err = Db.Query("INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES ($1, $2, $3, $4, $5)", uuid, bodyText, userId, threadId, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get new post information
	post, err := RetrievePostFromUuid(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(post)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseJson)
	return
}

// handle function for /posts
func HandlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetPosts(w, r)
	case "POST":
		CreatePost(w, r)
	}
	return
}
