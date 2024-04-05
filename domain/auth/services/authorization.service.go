package services

import (
	"go-rest-postgres/domain/auth/repositories"
	"go-rest-postgres/domain/auth/responses"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"context"

	"golang.org/x/oauth2"
)

func GenerateAuthorization(ctx *gin.Context) (authorizationResp responses.AuthorizationResponse, err error) {

	//	1. get Client ID from request and do query to find record on database
	//	2. set client_secret response from database as a new variable called as "clientSecret"
	//	3. create variable oauth2 config
	//	4. create variable verifier
	//	5. generate random string as a "authorization code" variable
	//	6. set response Authorization Response from authorization code variable 



	authorizationResponse, err := repositories.GenerateAuthorization(ctx)
	ctxBackground := context.Background()

	conf := &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/o/oauth2/auth",
			TokenURL: "https://provider.com/o/oauth2/token",
		},
	}

	// use PKCE to protect against CSRF attacks
	// https://www.ietf.org/archive/id/draft-ietf-oauth-security-topics-22.html#name-countermeasures-6
	verifier := oauth2.GenerateVerifier()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctxBackground, code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctxBackground, tok)
	client.Get("...")

	return authorizationResponse, err
}
