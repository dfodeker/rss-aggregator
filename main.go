package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/dfodeker/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not Load Env", err)
	}

	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Failed to load DB URL")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("failed to load Port from ENV")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	router.Mount("/v1", v1Router)

	svr := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server starting on port: %v\n", port)
	err = svr.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
