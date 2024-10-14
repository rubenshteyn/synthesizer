package main

import (
	"net/http"
	"synthesizer/internal/db"
	"synthesizer/pkg/routes"

	_ "github.com/lib/pq"
)

func main() {
	d := db.Connect()
	defer d.Close()
	r := routes.SetupRouter(d)
	http.ListenAndServe(":8085", r)
}
