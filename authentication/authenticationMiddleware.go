package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
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
	clientSecret := "7d109d2b-524f-4351-bfda-44ecad030eef"

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
		if rawAccessToken == "" {
			log.Println("No access token provided")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("403 Forbidden"))
			SignIn("sickboy", "ubi123")
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			log.Println("Missing token parts")
			w.WriteHeader(400)
			return
		}
		_, err := verifier.Verify(ctx, parts[1])
		log.Println(err)
		if err != nil {
			log.Println("Error while trying to access ressource")
			w.WriteHeader(400)
			w.Write([]byte("400 Bad Request"))
			//log.Println("error redirecting " + oauth2Config.AuthCodeURL(state))
			//http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}
		log.Println("serving http")
		next.ServeHTTP(w, r)
	})
}

func SignIn(username string, password string) []byte {
	urlPath := "http://localhost:8080/auth/realms/ubivius/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("client_id", "ubivius-client")
	data.Set("grant_type", "password")
	data.Set("client_secret", "7d109d2b-524f-4351-bfda-44ecad030eef")
	data.Set("scope", "openid")
	data.Set("username", username)
	data.Set("password", password)

	req, err := http.NewRequest("POST", urlPath, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	access_token := extractValue(string(body), "access_token")
	log.Println("access_token: ", access_token)

	player := map[string]string{"username": username, "access_token": access_token}
	playerJson, _ := json.Marshal(player)

	return playerJson
}

func SignUp(firstName string, lastName string, email string, username string) []byte {
	urlPath := "http://localhost:8080/auth/admin/realms/ubivius/users"

	values := map[string]string{"firstName": firstName, "lastName": lastName, "email": email, "username": username, "enabled": "true"}
	jsonValues, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(jsonValues))
	if err != nil {
		panic(err)
	}
	admin_token := GetAdminAccessToken()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+admin_token)
	log.Println(req.Header)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	access_token := extractValue(string(body), "access_token")
	log.Println("response Body:", string(body))

	player := map[string]string{"username": username, "access_token": access_token}
	playerJson, _ := json.Marshal(player)

	return playerJson
}

func GetAdminAccessToken() string {
	urlPath := "http://localhost:8080/auth/realms/ubivius/protocol/openid-connect/token"

	data := url.Values{}
	data.Set("client_id", "ubivius-client")
	data.Set("grant_type", "client_credentials")
	data.Set("client_secret", "7d109d2b-524f-4351-bfda-44ecad030eef")

	req, err := http.NewRequest("POST", urlPath, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Admintoken response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	admin_token := extractValue(string(body), "access_token")
	return admin_token
}

// extracts the value for a key from a JSON-formatted string
// body - the JSON-response as a string. Usually retrieved via the request body
// key - the key for which the value should be extracted
// returns - the value for the given key
func extractValue(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.Split(match, ":")
	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}

/*func AuthCallback(responseWriter http.ResponseWriter, request *http.Request) {
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
*/
