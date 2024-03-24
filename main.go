package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RLungWu/Tiny-REST-API/internal/taskstore"
	"github.com/gin-gonic/gin"
)

type taskServer struct {
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

func (ts *taskServer) deleteAllTasksHandler(c *gin.Context){
	ts.store.DeleteAllTasks()
}

func (ts *taskServer) createTaskHandler(c *gin.Context){
	type RequestTask struct{
		Text string `json:"text"`
		Tags []string `json:"tags"`
		Due time.Time `json:"due"`
	}

	var rt RequestTask
	if err := c.ShouldBindJSON(&rt); err != nil{
		c.String(http.StatusBadRequest , err.Error())
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	c.JSON(http.StatusOK, gin.H{"Id": id})
}

func (ts *taskServer) getTaskHandler(c *gin.Context){
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	task, err := ts.store.GetTask(id)
	if err != nil{
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ts *taskServer) deleteTaskHandler(c *gin.Context){
	id , err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil{
		c.String(http.StatusBadRequest, err.Error())
	}

	if err = ts.store.DeleteTask(id); err != nil{
		c.String(http.StatusNotFound, err.Error())
	}
}

func (ts *taskServer) tagHandler(c *gin.Context){
	tag := c.Params.ByName("tags")
	tasks := ts.store.GetTasksByTag(tag)
	c.JSON(http.StatusOK, tasks)
}

func (ts *taskServer) dueHandler(c *gin.Context){
	badRequestError := func(){
		c.String(http.StatusBadRequest, "expect /due/<year>/<month>/<day>, got %v", c.FullPath())
	}

	year, err := strconv.Atoi(c.Params.ByName("year"))
	if err != nil{
		badRequestError()
		return
	}

	month, err := strconv.Atoi(c.Params.ByName("month"))
	if err != nil{
		badRequestError()
		return
	}

	day, err := strconv.Atoi(c.Params.ByName("day"))
	if err != nil{
		badRequestError()
		return
	}

	tasks := ts.store.GetTasksByDueDate(year, time.Month(month), day)
	c.JSON(http.StatusOK, tasks)
}



func main(){
	router := gin.Default()
	server := NewTaskServer()

	router.POST("/task/", server.createTaskHandler)
	router.GET("/task", server.getAllTasksHandler)
	router.DELETE("/task/", server.deleteAllTasksHandler)
	router.GET("/task/:id", server.getTaskHandler)
	router.DELETE("/task/:id", server.deleteTaskHandler)
	router.GET("/tag/:tag", server.tagHandler)
	router.GET("/due/:year/:montyh/:day", server.dueHandler)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))
}



/*
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

*/