package propertyhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
	myErrors "github.com/Oleja123/dcaa-property/pkg/errors"
)

type PropertyHandler struct {
	service property.Service
}

func NewHandler(s property.Service) *PropertyHandler {
	return &PropertyHandler{service: s}
}

func (h *PropertyHandler) Handle(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.FindAll(rw, r)
	case http.MethodPost:
		h.Create(rw, r)
	case http.MethodPut:
		h.Update(rw, r)
	default:
		http.Error(rw, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *PropertyHandler) HandleWithId(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.FindOne(rw, r)
	case http.MethodDelete:
		h.Delete(rw, r)
	default:
		http.Error(rw, "метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func (h *PropertyHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var dto propertydto.PropertyDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(rw, "неправильное тело запроса", http.StatusBadRequest)
		return
	}

	if !dto.Validate(false) {
		http.Error(rw, "неправильное тело запроса", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(r.Context(), dto)
	if err != nil {
		switch {
		case errors.Is(err, myErrors.ErrInternalError):
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		case errors.Is(err, myErrors.ErrNotFound):
			http.Error(rw, err.Error(), http.StatusNotFound)
		}
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]int{"id": id})
}

func (h *PropertyHandler) FindAll(rw http.ResponseWriter, r *http.Request) {
	list, err := h.service.FindAll(r.Context())

	if err != nil {
		switch {
		case errors.Is(err, myErrors.ErrInternalError):
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		case errors.Is(err, myErrors.ErrNotFound):
			http.Error(rw, err.Error(), http.StatusNotFound)
		}
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(list)
}

func (h *PropertyHandler) FindOne(rw http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	dto, err := h.service.FindOne(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, myErrors.ErrInternalError):
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		case errors.Is(err, myErrors.ErrNotFound):
			http.Error(rw, err.Error(), http.StatusNotFound)
		}
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(dto)
}

func (h *PropertyHandler) Update(rw http.ResponseWriter, r *http.Request) {
	var dto propertydto.PropertyDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(rw, "неправильное тело запроса", http.StatusBadRequest)
		return
	}

	if !dto.Validate(true) {
		http.Error(rw, "неправильное тело запроса", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(r.Context(), dto); err != nil {
		switch {
		case errors.Is(err, myErrors.ErrInternalError):
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		case errors.Is(err, myErrors.ErrNotFound):
			http.Error(rw, err.Error(), http.StatusNotFound)
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (h *PropertyHandler) Delete(rw http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.service.Delete(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, myErrors.ErrInternalError):
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		case errors.Is(err, myErrors.ErrNotFound):
			http.Error(rw, err.Error(), http.StatusNotFound)
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
