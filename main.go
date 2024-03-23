package main

import (
	"log"
	"net/http"

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

