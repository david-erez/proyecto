package handlers

import (
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	_ "golang.org/x/image/webp"

	"backend/services"
)

type ConversionHandler struct {
	service *services.ConversionService
}

func NewConversionHandler(service *services.ConversionService) *ConversionHandler {
	return &ConversionHandler{service: service}
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errorResponse{Error: msg})
}

func (h *ConversionHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	// Límite de 10 MB para el archivo subido
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		writeError(w, http.StatusBadRequest, "no se pudo procesar el formulario")
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		writeError(w, http.StatusBadRequest, "falta el archivo 'image'")
		return
	}
	defer file.Close()

	tipo := r.FormValue("type")
	if tipo == "" {
		tipo = "ascii"
	}

	img, _, err := image.Decode(file)
	if err != nil {
		writeError(w, http.StatusBadRequest, "el archivo no es una imagen válida")
		return
	}

	record, err := h.service.ProcesarImagen(r.Context(), img, tipo, header.Filename)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, record)
}

func (h *ConversionHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	records, err := h.service.ListarConversiones(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, records)
}
