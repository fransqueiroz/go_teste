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
	"github.com/fransqueiroz/go_teste/api/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Repositorys
	userRepository := repository.NewUserRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	//Services
	userService := services.NewUserService(userRepository, walletRepository)
	walletService := services.NewWalletService(walletRepository)
	transactionService := services.NewTransactionService(transactionRepository, userRepository, walletRepository)

	//Controllers
	userController := controllers.NewUserController(userService)
	walletController := controllers.NewWalletController(walletService)
	transactionController := controllers.NewTransactionController(transactionService)

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
