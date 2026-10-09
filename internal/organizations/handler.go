package organizations

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
)

type organiztionHandler struct {
	OrganizationService *organizationService
}

type OrganizationRequest struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type OrganizationResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
}

func NewOrganizationHandler(organizationService *organizationService) *organiztionHandler {
	return &organiztionHandler{
		OrganizationService: organizationService,
	}
}

func (h *organiztionHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)

	if !ok {
		utils.Error(w, r, app.ErrMissingUserID)
		return
	}

	var body OrganizationRequest
	body.UserID = userID

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&body); err != nil {
		utils.Error(w, r, app.ErrBodyRequest)
		return
	}

	organization, err := h.OrganizationService.CreateOrganization(r.Context(), body)
	if err != nil {
		utils.Error(w, r, err)
		return
	}

	utils.Success(w, r, utils.ApiResponse{
		Result:     organization,
		Success:    true,
		StatusCode: http.StatusCreated,
	})
}

func (h *organiztionHandler) GetOrganizationByUserHandler(w http.ResponseWriter, r *http.Request) {}
