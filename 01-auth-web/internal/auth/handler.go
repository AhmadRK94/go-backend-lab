package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/user"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  *Service
	validate *validator.Validate
}

func NewHandler(service *Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		validate: validator,
	}
}

func JSONResponse(w http.ResponseWriter, statusCode int, msg any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(msg)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req user.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "validation failed",
		})
		return
	}

	u, err := h.service.Register(ctx, req)
	if err != nil {
		if errors.Is(err, user.ErrDuplicateEmail) {
			JSONResponse(w, http.StatusConflict, map[string]string{
				"error": "user with this email already exists",
			})
			return
		}

		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "could not register user",
		})
		return
	}

	JSONResponse(w, http.StatusCreated, map[string]any{
		"message": "user registered successfully",
		"data":    u,
	})
}
