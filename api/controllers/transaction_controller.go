package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/fransqueiroz/go_teste/api/services"
	"github.com/fransqueiroz/go_teste/api/utils"
	"github.com/gorilla/mux"
)

type TransactionController interface {
	PostTransaction(http.ResponseWriter, *http.Request)
	GetTransaction(http.ResponseWriter, *http.Request)
	GetTransactions(http.ResponseWriter, *http.Request)
	GetTransactionPayerId(http.ResponseWriter, *http.Request)
	GetTransactionPayeeId(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type transactionControllerImpl struct {
	transactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) *transactionControllerImpl {
	return &transactionControllerImpl{transactionService}
}

func (c *transactionControllerImpl) PostTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	transaction := &models.Transaction{}
	err = json.Unmarshal(bytes, transaction)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	mockValidation := utils.TransactionValidate()

	if mockValidation {
		transaction, err = c.transactionService.Post(transaction)
	} else {
		err = errors.New("Unauthorized Transaction")
		utils.WriteError(w, err, http.StatusForbidden)
		return
	}

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, transaction.ID))
	utils.WriteAsJson(w, transaction)
}

func (c *transactionControllerImpl) GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	transaction_id, err := strconv.ParseUint(params["transaction_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	transaction, err := c.transactionService.Find(transaction_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, transaction)
}

func (c *transactionControllerImpl) GetTransactionPayerId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	payer_id, err := strconv.ParseUint(params["payer_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	transaction, err := c.transactionService.FindByPayerId(payer_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, transaction)
}

func (c *transactionControllerImpl) GetTransactionPayeeId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	payee_id, err := strconv.ParseUint(params["payee_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	transaction, err := c.transactionService.FindByPayeeId(payee_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, transaction)
}

func (c *transactionControllerImpl) GetTransactions(w http.ResponseWriter, r *http.Request) {

	transaction, err := c.transactionService.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, transaction)
}

func (c *transactionControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	transaction_id, err := strconv.ParseUint(params["transaction_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.transactionService.Delete(transaction_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, "{}")
}
