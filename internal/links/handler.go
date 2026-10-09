package links

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cleidison-barradas/shortr.api/internal/app"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
)

type linksHandler struct {
	linkService *linkService
}

type LinkRequest struct {
	OriginalURL string     `json:"original_url"`
	Alias       string     `json:"alias"`
	ExpiresIn   *time.Time `json:"expires_in"`
	OrgID       string     `json:"org_id"`
}

type LinkResponse struct {
	ID          string     `json:"id"`
	ShortURL    string     `json:"short_url"`
	OriginalURL string     `json:"original_url"`
	Alias       string     `json:"alias"`
	ExpiresIn   *time.Time `json:"expires_in"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func newLinksHandler(linkService *linkService) *linksHandler {
	return &linksHandler{
		linkService: linkService,
	}
}

func (h *linksHandler) GetRedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	if code == "" {
		utils.Error(w, r, app.ErrMissingURLCode)
		return
	}

	link, err := h.linkService.FindByShortURL(r.Context(), code)
	if err != nil {
		utils.Error(w, r, err)
		return
	}

	http.Redirect(w, r, link.OriginalURL, http.StatusMovedPermanently)
}

func (h *linksHandler) FindHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Query().Get("short_url")

	link, err := h.linkService.FindByShortURL(r.Context(), shortURL)
	if err != nil {
		utils.Error(w, r, err)
		return
	}

	utils.Success(w, r, utils.ApiResponse{
		Result:     link,
		Success:    true,
		StatusCode: http.StatusOK,
	})
}

func (h *linksHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var req LinkRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&req); err != nil {
		utils.Error(w, r, app.ErrBodyRequest)
		return
	}

	link, err := h.linkService.CreateLink(r.Context(), req)
	if err != nil {
		utils.Error(w, r, err)
		return
	}

	utils.Success(w, r, utils.ApiResponse{
		Result:     link,
		Success:    true,
		StatusCode: http.StatusCreated,
	})
}
