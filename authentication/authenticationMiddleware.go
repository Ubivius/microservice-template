package authentication

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var oauth2Config oauth2.Config
var state string
var verifier *oidc.IDTokenVerifier
var ctx context.Context
var originUrl string
var method string

func init() {
	//Authentication setup
	configURL := "http://localhost:8080/auth/realms/ubivius"
	ctx = context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		log.Println("Auth panic")
		panic(err)
	}

	clientID := "ubivius-client"
	clientSecret := "ef14f638-98ad-4c5b-9320-a223077e0797"

	redirectURL := "http://localhost:9090/ubivius/callback"
	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state = "verysafestate"

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier = provider.Verifier(oidcConfig)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		originUrl = "http://localhost:9090" + r.URL.Path
		method = r.Method
		log.Println(rawAccessToken)
		if rawAccessToken == "" {
			log.Println("No access token provided")
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			w.WriteHeader(400)
			return
		}
		_, err := verifier.Verify(ctx, parts[1])
		log.Println(err)
		if err != nil {
			log.Println("error redirecting " + oauth2Config.AuthCodeURL(state))
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}
		log.Println("serving http")
		next.ServeHTTP(w, r)
	})
}

func AuthCallback(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("start auth callback")

	// Verify state and errors.
	if request.URL.Query().Get("state") != state {
		log.Println("state did not match")
		http.Error(responseWriter, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := oauth2Config.Exchange(ctx, request.URL.Query().Get("code"))
	if err != nil {
		log.Println("Failed to exchange token")
		http.Error(responseWriter, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		log.Println("No id_token field in oauth2 token.")
		http.Error(responseWriter, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	// Parse and verify ID Token payload.
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		log.Println("Auth Failed to verify ID Token")
		http.Error(responseWriter, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract custom claims
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		log.Println("Auth idToken claim error")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println("Auth marshal indent error")
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	val, err := responseWriter.Write(data)
	RedirectToInitialUrl(rawIDToken)

	request.Header.Set("Authorization", rawIDToken)
	if err != nil {
		log.Println("Auth write data error")
		http.Error(responseWriter, err.Error(), val)
		return
	}
}

func RedirectToInitialUrl(accessToken string) {
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + accessToken

	// Create a new request using http
	req, err := http.NewRequest(method, originUrl, nil)
	if err != nil {
		log.Println("Error while creating new request")
		return
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}
