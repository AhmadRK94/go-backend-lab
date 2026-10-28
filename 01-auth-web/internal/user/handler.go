package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.ParseUint(strID, 10, 32)
	if err != nil || id == 0 {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
		return
	}

	user, err := h.service.GetByID(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			JSONResponse(w, http.StatusNotFound, map[string]string{
				"error": "user not found",
			})
			return
		}
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}
	JSONResponse(w, http.StatusOK, user)

}

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "email is required",
		})
		return
	}

	user, err := h.service.GetByEmail(r.Context(), email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			JSONResponse(w, http.StatusNotFound, map[string]string{
				"error": "users not found",
			})
			return
		}
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}
	JSONResponse(w, http.StatusOK, user)

}

// to do: add pagination
func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}
	JSONResponse(w, http.StatusOK, users)
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.ParseUint(strID, 10, 32)
	if err != nil || id == 0 {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
		return
	}
	if err := h.service.DeleteAccount(r.Context(), uint(id)); err != nil {
		if errors.Is(err, ErrUserNotFound) {
			JSONResponse(w, http.StatusNotFound, map[string]string{
				"error": "user not found",
			})
			return
		}
		JSONResponse(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.ParseUint(strID, 10, 32)
	if err != nil || id == 0 {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
		return
	}
	var req UpdateUserProfile
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

	user, err := h.service.UpdateProfile(r.Context(), uint(id), req)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			JSONResponse(w, http.StatusNotFound, map[string]string{
				"error": "user not found",
			})
		case errors.Is(err, ErrDuplicateEmail):
			JSONResponse(w, http.StatusConflict, map[string]string{
				"error": "email already exists",
			})
		default:
			JSONResponse(w, http.StatusInternalServerError, map[string]string{
				"error": "internal server error",
			})
		}
		return
	}

	JSONResponse(w, http.StatusOK, user)

}
func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.ParseUint(strID, 10, 32)
	if err != nil || id == 0 {
		JSONResponse(w, http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
		return
	}
	var req UpdatePassword
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

	err = h.service.ChangePassword(r.Context(), uint(id), req.CurrentPassword, req.NewPassword)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			JSONResponse(w, http.StatusNotFound, map[string]string{
				"error": "user not found",
			})
		case errors.Is(err, ErrInvalidCurrentPassword):
			JSONResponse(w, http.StatusUnauthorized, map[string]string{
				"error": "current password is incorrect",
			})
		default:
			JSONResponse(w, http.StatusInternalServerError, map[string]string{
				"error": "internal server error",
			})
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
