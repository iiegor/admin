package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"iegor/admin"
)

// TODO: Abstraer los modelos.
// Se podría usar un config para cargarlos
// y evitaría tener que recompilar el código
// para cualquier cambio relacionado con estos.
type Course struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Province string  `json:"province"`
	Location string  `json:"location"`
	Country  string  `json:"country"`
}

var (
	port  int
	debug bool
)

func main() {
	flag.IntVar(&port, "port", 4000, "set listening port")
	flag.BoolVar(&debug, "debug", false, "set debug level for logs")
	flag.Parse()

	AdminAuth := &admin.AdminAuth{
		Users: []admin.AuthUser{
			{
				Username:  "iegor",
				Password:  "2008",
				Email:     "iegorazuaga@gmail.com",
				Role:      admin.AdminRole,
			},                                            
			{
				Username:  "guest",
				Password:  "iegor123",
				Email:     "guest+user@gmail.com",
				Role:      admin.GuestRole,
			},
		},
	}

	Admin := admin.New(&admin.AdminConfig{
		Prefix: "/",
		Debug:  debug,
		UI:     !debug,
		DB:     admin.NewDB("mysql", "root:iegor@/example_db?charset=utf8&parseTime=True"),
		Auth:   AdminAuth,
	})
	Admin.AddResource(new(Course), admin.ResourceConfig{
		Methods: []string{"read", "create", "update", "delete"},
	})

	mux := http.NewServeMux()

	Admin.MountTo(mux)

	if !debug {
		println("===========================")
		println("Running on production mode!")
		println("===========================\n")
	}

	log.Printf("Listening at http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), mux))
}
