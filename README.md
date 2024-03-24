# Tiny RestFul API Project
This project is a Golang project, it provide some RestFul api for tasks

## How to run
```bash
SERVERPORT=4112 go run .
```

## Features


| Method | Path                | Function                               |
|:------ |:------------------- |:-------------------------------------- |
| POST   | /task/              | create a task, returns ID              |
| GET    | /task/<taskid>      | returns a single task by ID            |
| GET    | /task/              | returns all tasks                      |
| DELETE | /task/<taskid>      | delete a task by ID                    |
| GET    | /tag/<tagname>      | returns list of tasks with this tag    |
| GET    | /due/<yy>/<mm>/<dd> | returns list of tasks due by this date |


## The model
Let's start by discussing the model for our server - the taskstore package (interna./taskstore).Here is its API
```Golang
func New() *TaskStore

// CreateTask creates a new task in the store.
func (ts *TaskStore) CreateTask(text string, tags []string, due time.Time) int

// GetTask retrieves a task from the store, by id. If no such id exists, an
// error is returned.
func (ts *TaskStore) GetTask(id int) (Task, error)

// DeleteTask deletes the task with the given id. If no such id exists, an error
// is returned.
func (ts *TaskStore) DeleteTask(id int) error

// DeleteAllTasks deletes all tasks in the store.
func (ts *TaskStore) DeleteAllTasks() error

// GetAllTasks returns all the tasks in the store, in arbitrary order.
func (ts *TaskStore) GetAllTasks() []Task

// GetTasksByTag returns all the tasks that have the given tag, in arbitrary
// order.
func (ts *TaskStore) GetTasksByTag(tag string) []Task

// GetTasksByDueDate returns all the tasks that have the given due date, in
// arbitrary order.
func (ts *TaskStore) GetTasksByDueDate(year int, month time.Month, day int) []Task
```
And the Task is:
```Golang
type Task struct {
  Id   int       `json:"id"`
  Text string    `json:"text"`
  Tags []string  `json:"tags"`
  Due  time.Time `json:"due"`
}
```
    
Reference from:https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library/


