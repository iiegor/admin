package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"alsur/admin"
)

// TODO: Abstraer los modelos.
// Se podría usar un config para cargarlos
// y evitaría tener que recompilar el código
// para cualquier cambio relacionado con estos.
type Course struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Province string `json:"province"`
	Location string `json:"location"`
	Country  string `json:"country"`
}

var (
	port  int
	debug bool
)

func main() {
	flag.IntVar(&port, "port", 4000, "set listening port")
	flag.BoolVar(&debug, "debug", false, "set debug level for logs")
	flag.Parse()

	Admin := admin.New(&admin.AdminConfig{
		Prefix: "/",
		Debug:  debug,
		UI:     true,
		DB:     admin.NewDB("mysql", "root:iegor@/example_db?charset=utf8&parseTime=True"),
	})
	Admin.AddResource(new(Course), admin.ResourceConfig{
		Methods: []string{"read", "create", "update", "delete"},
	})

	mux := http.NewServeMux()

	Admin.MountTo(mux)

	log.Printf("Listening at http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), mux))
}
