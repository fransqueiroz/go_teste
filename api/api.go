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

	//Repositorys
	userRepository := repository.NewUserRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	//Controllers
	userController := controllers.NewUserController(userRepository, walletRepository)
	walletController := controllers.NewWalletController(walletRepository)
	transactionController := controllers.NewTransactionController(transactionRepository, userRepository, walletRepository)

	//Routes
	userRoutes := routes.NewUserRoutes(userController)
	walletRoutes := routes.NewWalletRoutes(walletController)
	transactionRoutes := routes.NewTransactionRoutes(transactionController)

	router := mux.NewRouter().StrictSlash(true)

	routes.InstallUserRoute(router, userRoutes)
	routes.InstallWalletRoute(router, walletRoutes)
	routes.InstallTransactionRoute(router, transactionRoutes)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location"})
	methods := handlers.AllowedHeaders([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("API Running and listening on ", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handlers.CORS(headers, methods, origins)(router)))
}
