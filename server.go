package main

import (
	"github.com/go-chi/chi"
	customMiddleware "meetmeup/middleware"

	//"github.com/go-pg/pg/v9"
	"log"
	"meetmeup/database"

	//"meetmeup/database"
	"meetmeup/graph"
	"meetmeup/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	// TODO: I should come back to fix this line of code beneath
	DB := database.ConnectDB()
	//defer DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:8080"},
	//	AllowCredentials: true,
	//	Debug:            true,
	//}).Handler)
	//router.Use(middleware.RequestID)
	//router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(database.UsersRepo{}))

	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", graph.DataLoaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
