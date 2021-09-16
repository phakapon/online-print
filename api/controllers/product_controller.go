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

type ProductsController interface {
	PostProduct(http.ResponseWriter, *http.Request)
	GetProduct(http.ResponseWriter, *http.Request)
	GetProducts(http.ResponseWriter, *http.Request)
	PutProduct(http.ResponseWriter, *http.Request)
	DeleteProduct(http.ResponseWriter, *http.Request)
}

type productsControllerImpl struct {
	productsRepository repository.ProductsRepository
}

func NewProductsRepository(productsRepository repository.ProductsRepository) *productsControllerImpl {
	return &productsControllerImpl{productsRepository}
}

func (c *productsControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err = c.productsRepository.Save(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	buildCreatedResponse(w, buildLocation(r, product.ID))
	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	product, err := c.productsRepository.Find(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, product)
}

func (c *productsControllerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.productsRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteAsJson(w, products)
}

func (c *productsControllerImpl) PutProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
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

	product := &models.Product{}
	err = json.Unmarshal(bytes, product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product.ID = product_id

	err = product.Validate()
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsRepository.Update(product)
	if err != nil {
		utils.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *productsControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	product_id, err := strconv.ParseUint(params["product_id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.productsRepository.Delete(product_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	buildDeleteResponse(w, product_id)
	utils.WriteAsJson(w, "{}")
}
