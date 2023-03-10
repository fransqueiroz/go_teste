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

type UserController interface {
	PostUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	PutUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type userControllerImpl struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userControllerImpl {
	return &userControllerImpl{userService}
}

func (c *userControllerImpl) PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(bytes, user)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err = c.userService.Post(user)

	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	buildCreatedResponse(w, buildLocation(r, user.ID))
	utils.WriteAsJson(w, user)
}

func (c *userControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {

	user, err := c.userService.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, user)
}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	user, err := c.userService.GetUser(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, user)
}

func (c *userControllerImpl) PutUser(w http.ResponseWriter, r *http.Request) {
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

	user := &models.User{}
	err = json.Unmarshal(bytes, user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user.ID = user_id

	err = c.userService.Update(user)

	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.userService.Delete(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, "{}")
}
