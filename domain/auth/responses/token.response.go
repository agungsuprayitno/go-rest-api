package responses

import "go-rest-postgres/domain/users/models"

type TokenResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	LoggedIn     bool        `json:"logged_in"`
	User         models.User `json:"user"`
}

func (tokenResponse TokenResponse) SetResponse(accessToken string, refreshToken string, loggedIn bool, user models.User) TokenResponse {
	response := TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		LoggedIn:     loggedIn,
		User:         user,
	}

	return response
}
