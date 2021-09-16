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

type LocationsController interface {
	PostLocation(http.ResponseWriter, *http.Request)
	GetLocation(http.ResponseWriter, *http.Request)
	GetLocations(http.ResponseWriter, *http.Request)
	PutLocation(http.ResponseWriter, *http.Request)
	DeleteLocation(http.ResponseWriter, *http.Request)
}

type locationsControllerImpl struct {
	locationsRepository repository.LocationsRepository
}

func NewLocationsRepository(locationsRepository repository.LocationsRepository) *locationsControllerImpl {
	return &locationsControllerImpl{locationsRepository}
}

func (c *locationsControllerImpl) PostLocation(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	location := &models.Location{}
	err = json.Unmarshal(bytes, location)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = location.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	location, err = c.locationsRepository.Save(location)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, location.ID))
	utils.WriteAsJson(w, location)
}

func (c *locationsControllerImpl) GetLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	location_id, err := strconv.ParseUint(params["location_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	location, err := c.locationsRepository.Find(location_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, location)
}

func (c *locationsControllerImpl) GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := c.locationsRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, locations)
}

func (c *locationsControllerImpl) PutLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	location_id, err := strconv.ParseUint(params["location_id"], 10, 64)
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

	location := &models.Location{}
	err = json.Unmarshal(bytes, location)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	location.ID = location_id

	err = location.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.locationsRepository.Update(location)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *locationsControllerImpl) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	location_id, err := strconv.ParseUint(params["location_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.locationsRepository.Delete(location_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, location_id)
	utils.WriteAsJson(w, "{}")
}
