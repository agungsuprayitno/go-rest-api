package responses

type AuthorizationResponse struct {
	AuthorizationCode  string      `json:"authorization_code"`
}

func (authorizationResponse AuthorizationResponse) SetResponse(authorizationCode string,) AuthorizationResponse {
	response := AuthorizationResponse{
		AuthorizationCode:  authorizationCode,
	}

	return response
}
