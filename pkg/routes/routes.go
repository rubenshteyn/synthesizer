package routes

import (
	"database/sql"
	"synthesizer/pkg/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login(db)).Methods("POST")
	router.HandleFunc("/authorization", controllers.Authentication(db)).Methods("POST")

	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users", controllers.UpdateUser(db)).Methods("PUT")

	router.HandleFunc("/tasks/user_id/{id}", controllers.GetTasks(db)).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask(db)).Methods("POST")
	router.HandleFunc("/tasks/sequence", controllers.UpdateTaskSequence(db)).Methods("PUT")
	router.HandleFunc("/tasks/status", controllers.UpdateTaskStatus(db)).Methods("PUT")
	router.HandleFunc("/tasks", controllers.DeleteTask(db)).Methods("DELETE")

	router.HandleFunc("/monomers", controllers.GetMonomers(db)).Methods("GET")
	router.HandleFunc("/monomers/{name}", controllers.GetMonomerCount(db)).Methods("GET")
	router.HandleFunc("/monomers", controllers.CreateMonomer(db)).Methods("POST")
	router.HandleFunc("/monomers", controllers.UpdateMonomerCount(db)).Methods("PUT")
	return router
}
