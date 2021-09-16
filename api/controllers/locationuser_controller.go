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

type LocationusersController interface {
	PostLocationuser(http.ResponseWriter, *http.Request)
	GetLocationuser(http.ResponseWriter, *http.Request)
	GetLocationusers(http.ResponseWriter, *http.Request)
	PutLocationuser(http.ResponseWriter, *http.Request)
	DeleteLocationuser(http.ResponseWriter, *http.Request)
}

type locationusersControllerImpl struct {
	locationusersRepository repository.LocationusersRepository
}

func NewLocationusersRepository(locationusersRepository repository.LocationusersRepository) *locationusersControllerImpl {
	return &locationusersControllerImpl{locationusersRepository}
}

func (c *locationusersControllerImpl) PostLocationuser(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	locationuser := &models.Locationuser{}
	err = json.Unmarshal(bytes, locationuser)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = locationuser.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	locationuser, err = c.locationusersRepository.Save(locationuser)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, locationuser.ID))
	utils.WriteAsJson(w, locationuser)
}

func (c *locationusersControllerImpl) GetLocationuser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	locationuser_id, err := strconv.ParseUint(params["locationuser_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	locationuser, err := c.locationusersRepository.Find(locationuser_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, locationuser)
}

func (c *locationusersControllerImpl) GetLocationusers(w http.ResponseWriter, r *http.Request) {
	locationusers, err := c.locationusersRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, locationusers)
}

func (c *locationusersControllerImpl) PutLocationuser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	locationuser_id, err := strconv.ParseUint(params["locationuser_id"], 10, 64)
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

	locationuser := &models.Locationuser{}
	err = json.Unmarshal(bytes, locationuser)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	locationuser.ID = locationuser_id

	err = locationuser.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.locationusersRepository.Update(locationuser)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *locationusersControllerImpl) DeleteLocationuser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	locationuser_id, err := strconv.ParseUint(params["locationuser_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.locationusersRepository.Delete(locationuser_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, locationuser_id)
	utils.WriteAsJson(w, "{}")
}
