package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nomansalhab/golang_training_project/pkg/config"
	"github.com/nomansalhab/golang_training_project/pkg/handlers"
	"github.com/nomansalhab/golang_training_project/pkg/render"
)

const portNumber = ":8016"

// addValues adds two integers and returns the result
// func addValues(x, y int) int {
// 	return x + y
// }
// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 00.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot Divide By Zero")
// 	} else {
// 		fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
// 	}
// }
// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by equals or less than zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	// ? When in Development Mode UseCache is false
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, World!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Bytes Written: %d", n))
	// })
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting Applicationon Port %s", portNumber)

	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err1 := srv.ListenAndServe()
	if err1 != nil {
		log.Fatal(err1)
	}
}
