package product

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	ErrParseInput = errors.New("invalid input format")
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Router() *chi.Mux {
	r := chi.NewMux()

	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Get("/", listProductsHandler)
	r.Post("/", submitProductHandler)
	r.Get("/{productID}", getProductHandler)
	r.Put("/{productID}", updateProductHandler)
	r.Post("/{productID}/reviews", submitProductReviewHandler)
	r.Post("/{productID}/images", submitProductImageHandler)
	return r
}

func writeResponse(w http.ResponseWriter, res Response, status int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func mapError(err error) (res Response, httpStatus int) {
	res = Response{
		Message: err.Error(),
	}
	switch err {
	case ErrParseInput, ErrTitleMissing, ErrSkuMissing, ErrPriceMissing, ErrRatingMissing, ErrRatingInvalid:
		httpStatus = http.StatusBadRequest
	case ErrNotFound:
		httpStatus = http.StatusNotFound
	case ErrNoRows, ErrSkuDuplicate:
		httpStatus = http.StatusUnprocessableEntity
	default:
		res.Message = "internal server error"
		httpStatus = http.StatusInternalServerError
	}
	return
}

func writeError(w http.ResponseWriter, err error) {
	res, status := mapError(err)
	writeResponse(w, res, status)
}

func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	val, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err := ErrParseInput
		writeError(w, err)
		return
	}

	offset, _ := strconv.ParseUint(val.Get("offset"), 10, 64)
	limit, _ := strconv.ParseUint(val.Get("limit"), 10, 64)
	if limit == 0 {
		limit = 5
	}
	searchQuery := ProductSearchQuery{
		SKU:      val.Get("sku"),
		Title:    val.Get("title"),
		Category: val.Get("category"),
		SortBy:   convertSortColumn(val.Get("sort_by")),
		Sort:     convertSortDirection(val.Get("sort")),
		Offset:   offset,
		Limit:    limit,
	}
	products, err := listProducts(r.Context(), searchQuery)
	if err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, Response{Data: products}, http.StatusOK)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	productIDstr := chi.URLParam(r, "productID")
	if productIDstr == "" {
		err := ErrParseInput
		writeError(w, err)
		return
	}

	productID, _ := strconv.ParseUint(productIDstr, 10, 64)
	product, err := getProductDetails(r.Context(), productID)
	if err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, Response{Data: product}, http.StatusOK)
}

func submitProductHandler(w http.ResponseWriter, r *http.Request) {
	product := Product{}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		err = ErrParseInput
		writeError(w, err)
		return
	}

	if err := validateProduct(product); err != nil {
		writeError(w, err)
		return
	}

	if err := submitProduct(r.Context(), &product); err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, Response{Data: product}, http.StatusCreated)
}

func submitProductImageHandler(w http.ResponseWriter, r *http.Request) {
	productIDstr := chi.URLParam(r, "productID")
	if productIDstr == "" {
		err := ErrParseInput
		writeError(w, err)
		return
	}

	productID, _ := strconv.ParseUint(productIDstr, 10, 64)
	file, handler, err := r.FormFile("image")
	if err != nil {
		err = ErrParseInput
		writeError(w, err)
		return
	}

	defer file.Close()
	fileParts := strings.Split(handler.Filename, ".")

	imgData := ProductImageUpload{
		ProductID:   productID,
		Extension:   fileParts[len(fileParts)-1],
		File:        file,
		Description: r.FormValue("description"),
	}
	if err = submitProductImage(r.Context(), imgData); err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, Response{Message: "product image submitted"}, http.StatusCreated)
}

func submitProductReviewHandler(w http.ResponseWriter, r *http.Request) {
	productReview := ProductReview{}
	if err := json.NewDecoder(r.Body).Decode(&productReview); err != nil {
		err = ErrParseInput
		writeError(w, err)
		return
	}

	productIDstr := chi.URLParam(r, "productID")
	if productIDstr == "" {
		err := ErrParseInput
		writeError(w, err)
		return
	}

	productID, _ := strconv.ParseUint(productIDstr, 10, 64)
	productReview.ProductID = productID

	if err := validateProductReview(productReview); err != nil {
		writeError(w, err)
		return
	}

	if err := submitProductReview(r.Context(), productReview); err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, Response{Message: "product review submitted"}, http.StatusCreated)
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	product := Product{}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		err = ErrParseInput
		writeError(w, err)
		return
	}

	productIDstr := chi.URLParam(r, "productID")
	if productIDstr == "" {
		err := ErrParseInput
		writeError(w, err)
		return
	}

	product.ID, _ = strconv.ParseUint(productIDstr, 10, 64)

	if err := validateProduct(product); err != nil {
		writeError(w, err)
		return
	}

	err := submitProduct(r.Context(), &product)
	if err != nil {
		writeError(w, err)
		return
	}

	writeResponse(w, Response{Data: product}, http.StatusOK)
}
