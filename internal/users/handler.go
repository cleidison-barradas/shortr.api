package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cleidison-barradas/shortr.api/internal/utils"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreatedUserResult struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserHandler struct {
	userSrv *UserService
}

func NewUserHandler(userSrv *UserService) *UserHandler {
	return &UserHandler{
		userSrv: userSrv,
	}
}

func (h *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.userSrv.FindUsers(r.Context())

	if err != nil {
		utils.Error(w, r, err)
	}

	utils.Success(w, r, utils.ApiResponse{
		Result:     users,
		Success:    true,
		StatusCode: http.StatusOK,
	})
}

func (h *UserHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var body CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.Error(w, r, err)
		return
	}

	user, err := h.userSrv.CreateUser(r.Context(), body)

	if err != nil {
		utils.Error(w, r, err)
		return
	}

	utils.Success(w, r, utils.ApiResponse{
		Result:     user,
		Success:    true,
		StatusCode: http.StatusCreated,
	})
}

func (h *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
