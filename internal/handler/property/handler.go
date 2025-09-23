package propertyhandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Oleja123/dcaa-property/internal/domain/property"
	propertydto "github.com/Oleja123/dcaa-property/internal/dto/property"
)

type PropertyHandler struct {
	service property.Service
}

func NewHandler(s property.Service) *PropertyHandler {
	return &PropertyHandler{service: s}
}

func (h *PropertyHandler) Create(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "неразрешенный метод", http.StatusMethodNotAllowed)
		return
	}

	var dto propertydto.PropertyDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(rw, "неправильное тело запроса", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(r.Context(), dto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]int{"id": id})
}

func (h *PropertyHandler) FindAll(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(rw, "неразрешенный метод", http.StatusMethodNotAllowed)
		return
	}

	list, err := h.service.FindAll(r.Context())

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(list)
}

func (h *PropertyHandler) FindOne(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(rw, "неразрешенный метод", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	dto, err := h.service.FindOne(r.Context(), id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(rw).Encode(dto)
}

func (h *PropertyHandler) Update(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(rw, "неразрешенный метод", http.StatusMethodNotAllowed)
		return
	}

	var dto propertydto.PropertyDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(rw, "invalid request body", http.StatusBadRequest)
		return
	}

	if _, err := h.service.FindOne(r.Context(), dto.Id); err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.Update(r.Context(), dto); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (h *PropertyHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(rw, "неразрешенный метод", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	if _, err := h.service.FindOne(r.Context(), id); err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
