package routes

import (
	"net/http"

	"github.com/fransqueiroz/go_teste/api/controllers"
)

type WalletRoutes interface {
	Routes() []*Route
}

type walletRoutesImpl struct {
	walletController controllers.WalletController
}

func NewWalletRoutes(walletController controllers.WalletController) *walletRoutesImpl {
	return &walletRoutesImpl{walletController}
}
func (r *walletRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/wallet",
			Method:  http.MethodPost,
			Handler: r.walletController.PostWallet,
		},
		{
			Path:    "/wallet",
			Method:  http.MethodGet,
			Handler: r.walletController.GetWallets,
		},
		{
			Path:    "/wallet/{wallet_id}",
			Method:  http.MethodGet,
			Handler: r.walletController.GetWallet,
		},
		{
			Path:    "/wallet/user/{user_id}",
			Method:  http.MethodGet,
			Handler: r.walletController.GetWalletByUserId,
		},
		{
			Path:    "/wallet/{wallet_id}",
			Method:  http.MethodPut,
			Handler: r.walletController.PutWallet,
		},
		{
			Path:    "/wallet/{wallet_id}",
			Method:  http.MethodDelete,
			Handler: r.walletController.DeleteWallet,
		},
	}
}
