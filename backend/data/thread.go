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

type TopicInfo struct {
	Topic string `json:"topic"`
}

func RetrieveAllThreads() ([]ThreadInfo, error) {
	rows, err := Db.Query("select A.*, COUNT(B.id) as post_num from threads as A left join posts as B on B.thread_id=A.id group by A.id, A.uuid, A.topic, A.user_id, A.created_at ORDER BY A.created_at DESC")
	if err != nil {
		return nil, err
	}

	threads := make([]ThreadInfo, 0)
	for rows.Next() {
		thread := ThreadInfo{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt, &thread.PostsNum); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	rows.Close()
	return threads, nil
}

func RetrieveThreadInfoFromUuid(threadUuid string) (ThreadInfo, error) {
	thread_info := ThreadInfo{}
	err := Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", threadUuid).Scan(&thread_info.Id, &thread_info.Uuid, &thread_info.Topic, &thread_info.UserId, &thread_info.CreatedAt)
	if err != nil {
		return thread_info, err
	}
	return thread_info, nil
}

// get all threads in the database with a number of posts and returns it
// GET /threads
func GetThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	threads, err := RetrieveAllThreads()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	thread_info, err := RetrieveThreadInfoFromUuid(thread_uuid)
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
	ThreadInfo, err := RetrieveThreadInfoFromUuid(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseJson, err := json.Marshal(ThreadInfo)
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
