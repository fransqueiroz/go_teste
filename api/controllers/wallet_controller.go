package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/fransqueiroz/go_teste/api/services"
	"github.com/fransqueiroz/go_teste/api/utils"
	"github.com/gorilla/mux"
)

type WalletController interface {
	PostWallet(http.ResponseWriter, *http.Request)
	GetWallet(http.ResponseWriter, *http.Request)
	GetWalletByUserId(http.ResponseWriter, *http.Request)
	GetWallets(http.ResponseWriter, *http.Request)
	PutWallet(http.ResponseWriter, *http.Request)
	DeleteWallet(http.ResponseWriter, *http.Request)
}

type walletControllerImpl struct {
	walletService services.WalletService
}

type teste struct{}

func NewWalletController(walletService services.WalletService) *walletControllerImpl {
	return &walletControllerImpl{walletService}
}

func (c *walletControllerImpl) PostWallet(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	wallet := &models.Wallet{}
	err = json.Unmarshal(bytes, wallet)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	wallet, err = c.walletService.Post(wallet)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	buildCreatedResponse(w, buildLocation(r, wallet.ID))
	utils.WriteAsJson(w, wallet)
}

func (c *walletControllerImpl) GetWallets(w http.ResponseWriter, r *http.Request) {
	wallet, err := c.walletService.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, wallet)
}

func (c *walletControllerImpl) GetWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	wallet_id, err := strconv.ParseUint(params["wallet_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	wallet, err := c.walletService.Find(wallet_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, wallet)
}

func (c *walletControllerImpl) GetWalletByUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	wallet, err := c.walletService.FindByUserId(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, wallet)
}

func (c *walletControllerImpl) PutWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	wallet := &models.Wallet{}
	err = json.Unmarshal(bytes, wallet)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	wallet.User_id = user_id

	err = c.walletService.Update(wallet)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *walletControllerImpl) DeleteWallet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	wallet_id, err := strconv.ParseUint(params["wallet_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.walletService.Delete(wallet_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, "{}")
}
