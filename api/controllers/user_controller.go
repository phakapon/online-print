package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"online-print/api/models"
	"online-print/api/repository"
	"online-print/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UsersController interface {
	PostUser(http.ResponseWriter, *http.Request)
	GetUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	PutUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type usersControllerImpl struct {
	usersRepository repository.UsersRepository
}

func NewUsersRepository(usersRepository repository.UsersRepository) *usersControllerImpl {
	return &usersControllerImpl{usersRepository}
}

func (c *usersControllerImpl) PostUser(w http.ResponseWriter, r *http.Request) {
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

	err = user.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	user, err = c.usersRepository.Save(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, user.ID))
	utils.WriteAsJson(w, user)
}

func (c *usersControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	user, err := c.usersRepository.Find(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, user)
}

func (c *usersControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.usersRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, users)
}

func (c *usersControllerImpl) PutUser(w http.ResponseWriter, r *http.Request) {
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

	err = user.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.usersRepository.Update(user)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *usersControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user_id, err := strconv.ParseUint(params["user_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.usersRepository.Delete(user_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, user_id)
	utils.WriteAsJson(w, "{}")
}
