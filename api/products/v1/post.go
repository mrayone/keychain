package v1

import (
	"net/http"

	"github.com/mrayone/golang-learn/api"
)

type ProductHTTPPostHandler struct {
	Route string
}

const productRoute = "/v1/products"

func NewPostProductHandler() *ProductHTTPPostHandler {
	return &ProductHTTPPostHandler{
		Route: productRoute,
	}
}

type Product struct {
	Message string `json:"message"`
}

func (h *ProductHTTPPostHandler) Handler(w http.ResponseWriter, r *http.Request) {
	p := Product{
		Message: "teste",
	}
	api.RespondWithJSON(w, http.StatusCreated, p)
}
