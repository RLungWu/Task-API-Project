package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/RLungWu/Tiny-REST-API/internal/taskstore"
	"github.com/gin-gonic/gin"
)

func taskServer struct{
	store *taskstore.TaskStore
}



func NewTaskServer() *taskServer{
	store := taskstore.New()
	return &taskServer{store: store}
}

func (ts *taskServer) getAllTasksHandler(c *gin.Context){
	allTasks := ts.store.GetAllTasks()
	c.JSON(http.StatusOK, allTasks)
}

func (ts *taskServer) getTaskHandler(w http.ResponseWriter, req *http.Request){
	log.Printf("Handling get task at %s \n", req.URL.Path)

	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil{
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	task, err := ts.store.GetTask(id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(task)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}





func main(){
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("POST /task/", server.createTaskHandler)
	mux.HandleFunc("GET /task", server.getAllTasksHandler)
	mux.HandleFunc("GET /task/{id}", server.getTaskHandler)
	mux.HandleFunc("DELETE /task/", server.deleteAllTasksHandler)
	mux.HandleFunc("DELETE /task/{id}", server.deleteTaskHandler)
	mux.HandleFunc("GET /tag/{tag}", server.tagHandler)
	mux.HandleFunc("GET /due/{year}/{month}/{day}", server.dueHandler)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))
}

