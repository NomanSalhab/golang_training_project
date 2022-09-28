module github.com/nomansalhab/golang_training_project

go 1.19

replace github.com/nomansalhab/golang_training_project/cmd/web/ => ./cmd/web/

require github.com/go-chi/chi v1.5.4

require (
	github.com/alexedwards/scs/v2 v2.5.0 // indirect
	github.com/justinas/nosurf v1.1.1 // indirect
)
