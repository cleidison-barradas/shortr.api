package signin

import (
	"encoding/json"
	"net/http"

	"github.com/cleidison-barradas/shortr.api/internal/auth"
	"github.com/cleidison-barradas/shortr.api/internal/utils"
)

type SigninHandler struct {
	SignInSrv *siginService
	JwtSrv    *auth.JWTService
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponse struct {
	UserID      string `json:"user_id"`
	OrgID       string `json:"org_id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func NewSigninHandler(signinsrv *siginService, jwtsrv *auth.JWTService) *SigninHandler {
	return &SigninHandler{
		JwtSrv:    jwtsrv,
		SignInSrv: signinsrv,
	}
}

func (s *SigninHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var body SigninRequest

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&body); err != nil {
		utils.Error(w, r, err)
	}

	result, err := s.SignInSrv.SignInUser(r.Context(), body)
	if err != nil {
		utils.Error(w, r, err)
		return
	}

	accessToken, err := s.JwtSrv.GenerateToken(&auth.GenerateTokenParams{
		UserID: result.UserID,
		OrgID:  result.OrgID,
		Email:  body.Email,
	})
	if err != nil {
		utils.Error(w, r, err)
		return
	}

	result.AccessToken = accessToken

	utils.Success(w, r, utils.ApiResponse{
		Result:     result,
		Success:    true,
		StatusCode: http.StatusOK,
	})
}
