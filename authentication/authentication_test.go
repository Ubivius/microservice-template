package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareWithoutToken(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	response := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	handlerToTest := Middleware(nextHandler)
	handlerToTest.ServeHTTP(response, request)

	if response.Code != 403 {
		t.Errorf("Expected status code 403 but got : %d", response.Code)
	}
}

func TestMiddlewareWithIncompleteToken(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	request.Header.Add("Authorization", "Bearer 12345abcde")
	response := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	handlerToTest := Middleware(nextHandler)
	handlerToTest.ServeHTTP(response, request)

	if response.Code != 400 {
		t.Errorf("Expected status code 400 but got : %d", response.Code)
	}
}

func TestMiddlewareWithExpiredToken(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/products", nil)
	request.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJwZDhIeFB3bzY5ZmVBYnNDdGZhQXQyWjdxQUpvM2NpSnpxXzJ5Y0ZpRFBrIn0.eyJleHAiOjE2MTUyMjE3MDIsImlhdCI6MTYxNTIyMTQwMiwianRpIjoiNzM0YTcxNGYtNWFmYi00ZDM0LTg5MGYtODU2MDkyNzRmZGRkIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL2F1dGgvcmVhbG1zL3ViaXZpdXMiLCJhdWQiOiJ1Yml2aXVzLWNsaWVudCIsInN1YiI6ImRhM2M1MDFmLTA3ZDMtNDc0Zi1iMjAwLTVkN2JlZThjZmU3YiIsInR5cCI6IkJlYXJlciIsImF6cCI6InViaXZpdXMtY2xpZW50Iiwic2Vzc2lvbl9zdGF0ZSI6IjI1ZDkzZmRmLWRkYjUtNGIwYi1iZTQwLTVjODMwOTNhNTZkOSIsImFjciI6IjEiLCJhbGxvd2VkLW9yaWdpbnMiOlsiaHR0cHM6Ly93d3cua2V5Y2xvYWsub3JnIl0sInNjb3BlIjoib3BlbmlkIGdvb2Qtc2VydmljZSBwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJuYW1lIjoiSmVyZW1pIFNhdmFyZCIsInByZWZlcnJlZF91c2VybmFtZSI6InNpY2tib3kiLCJnaXZlbl9uYW1lIjoiSmVyZW1pIiwiZmFtaWx5X25hbWUiOiJTYXZhcmQifQ.KNu_23uXRbYtuHwweFGKalyUVVLsxuqd9ZhLKVJVdHrFHZx1qOgKiG5Y2ODzZmXb1qGXGjb_xwt4F9DnazItvYt4QlWz1LjJTUOBMhUx72Y2M-4hacxnEcYgjNHyFXdZlxuyuKH0JvT5drhWR7kjDHumT-hSHlXNLzAe5B0foH-i-5tcPYNJGZOAHKqGAhJsLhbSm4jDdI6RfjUOrZ2CYmPYKlg5uQ9VXA3lKJBCMSrKHXQeqUCKs4miUvcK3ZAc2-Y6Wdw6mbzVCQ3u3KS9v7gQldKjxSbt-tXbmy5PbQLs_ztB2Nq1hcn-m5MclhBsMuJtsTQg1pj7rvZD4QrNgQ")
	response := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	handlerToTest := Middleware(nextHandler)
	handlerToTest.ServeHTTP(response, request)

	if response.Code != 400 {
		t.Errorf("Expected status code 400 but got : %d", response.Code)
	}
}
