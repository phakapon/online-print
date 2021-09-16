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

type ProductsorderController interface {
	PostProductorder(http.ResponseWriter, *http.Request)
	GetProductorder(http.ResponseWriter, *http.Request)
	GetProductsorder(http.ResponseWriter, *http.Request)
	PutProductorder(http.ResponseWriter, *http.Request)
	DeleteProductorder(http.ResponseWriter, *http.Request)
}

type productsorderControllerImpl struct {
	productsorderRepository repository.ProductsorderRepository
}

func NewProductsorderRepository(productsorderRepository repository.ProductsorderRepository) *productsorderControllerImpl {
	return &productsorderControllerImpl{productsorderRepository}
}

func (c *productsorderControllerImpl) PostProductorder(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	productorder := &models.Productorder{}
	err = json.Unmarshal(bytes, productorder)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = productorder.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	productorder, err = c.productsorderRepository.Save(productorder)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, productorder.ID))
	utils.WriteAsJson(w, productorder)
}

func (c *productsorderControllerImpl) GetProductorder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productorder_id, err := strconv.ParseUint(params["productorder_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	productorder, err := c.productsorderRepository.Find(productorder_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, productorder)
}

func (c *productsorderControllerImpl) GetProductsorder(w http.ResponseWriter, r *http.Request) {
	productsorder, err := c.productsorderRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, productsorder)
}

func (c *productsorderControllerImpl) PutProductorder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productorder_id, err := strconv.ParseUint(params["productorder_id"], 10, 64)
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

	productorder := &models.Productorder{}
	err = json.Unmarshal(bytes, productorder)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	productorder.ID = productorder_id

	err = productorder.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsorderRepository.Update(productorder)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *productsorderControllerImpl) DeleteProductorder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	productorder_id, err := strconv.ParseUint(params["productorder_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsorderRepository.Delete(productorder_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, productorder_id)
	utils.WriteAsJson(w, "{}")
}
