package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/fransqueiroz/go_teste/api/controllers"
	"github.com/fransqueiroz/go_teste/api/database"
	"github.com/fransqueiroz/go_teste/api/repository"
	"github.com/fransqueiroz/go_teste/api/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	port = flag.Int("p", 5000, "set port")
)

func Run() {
	flag.Parse()
	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	userRepository := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)
	userRoutes := routes.NewUserRoutes(userController)

	router := mux.NewRouter().StrictSlash(true)
	routes.InstallUserRoute(router, userRoutes)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Teste"))
	})

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location"})
	methods := handlers.AllowedHeaders([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("API Running and listening on ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}
