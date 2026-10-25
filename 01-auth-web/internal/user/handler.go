package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	service  *UserService
	validate *validator.Validate
}

func NewUserHandler(service *UserService, validate *validator.Validate) *userHandler {
	return &userHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {

}
func (h *userHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {

}
func (h *userHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {

}
