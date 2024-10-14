package routes

import (
	"database/sql"
	"synthesizer/pkg/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users", controllers.UpdateUser(db)).Methods("PUT")

	router.HandleFunc("/tasks", controllers.GetTasks(db)).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask(db)).Methods("POST")
	router.HandleFunc("/tasks", controllers.UpdateTask(db)).Methods("PUT")
	router.HandleFunc("/tasks", controllers.DeleteTask(db)).Methods("DELETE")
	return router
}
