package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func InstallUserRoute(router *mux.Router, userRoutes UserRoutes) {
	allRoutes := userRoutes.Routes()
	for _, route := range allRoutes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}

func InstallWalletRoute(router *mux.Router, walletRoutes WalletRoutes) {
	allRoutes := walletRoutes.Routes()
	for _, route := range allRoutes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
