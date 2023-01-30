package routes

import (
	"net/http"

	"github.com/fransqueiroz/go_teste/api/controllers"
)

type TransactionRoutes interface {
	Routes() []*Route
}

type transactionRoutesImpl struct {
	transactionController controllers.TransactionController
}

func NewTransactionRoutes(transactionController controllers.TransactionController) *transactionRoutesImpl {
	return &transactionRoutesImpl{transactionController}
}

func (r *transactionRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:    "/transaction",
			Method:  http.MethodPost,
			Handler: r.transactionController.PostTransaction,
		},
		{
			Path:    "/transaction",
			Method:  http.MethodGet,
			Handler: r.transactionController.GetTransactions,
		},
		{
			Path:    "/transaction/{transaction_id}",
			Method:  http.MethodGet,
			Handler: r.transactionController.GetTransaction,
		},
		{
			Path:    "/transaction/{payee_id}",
			Method:  http.MethodGet,
			Handler: r.transactionController.GetTransactionPayeeId,
		},
		{
			Path:    "/transaction/{payer_id}",
			Method:  http.MethodGet,
			Handler: r.transactionController.GetTransactionPayerId,
		},
		{
			Path:    "/transaction/{transaction_id}",
			Method:  http.MethodDelete,
			Handler: r.transactionController.Delete,
		},
	}
}
