package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/myservice/pkg/config"
	"github.com/myservice/pkg/handlers"
	"github.com/myservice/pkg/render"
)

func main() {
	portNumber := ":8080"

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)
	Repo := handlers.NewRepo(&app)
	handlers.NewHandlers(Repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/login", handlers.Repo.Login)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(portNumber, nil)
}
