package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/hootsuite/google-calendar-app-sample/api/planned-content"
	"github.com/hootsuite/google-calendar-app-sample/api/status"
	"github.com/joho/godotenv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	cal "google.golang.org/api/calendar/v3"
)

func main() {
	env, _ := godotenv.Read(".env")
	CLIENT_ID := os.Getenv("CLIENT_ID")
	if CLIENT_ID == "" {
		CLIENT_ID = env["CLIENT_ID"]
		if CLIENT_ID == "" {
			log.Fatal("CLIENT_ID is required")
		}
	}
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")
	if CLIENT_SECRET == "" {
		CLIENT_SECRET = env["CLIENT_SECRET"]
		if CLIENT_SECRET == "" {
			log.Fatal("CLIENT_SECRET is required")
		}
	}

	// Get the client id and secret from google cloud credential.
	// url: https://console.cloud.google.com/apis/credentials
	// access them from environment variable

	// Oauth config manage the OAuth flow. You have to register
	// the redirect url in the OAuth provider. For the endpoint,
	// there are many provider specific package inside the
	// golang.org/x/oauth2 package
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		//:  "http://localhost:8000/auth/callback",
		Scopes:   []string{cal.CalendarReadonlyScope},
		Endpoint: google.Endpoint,
	}

	api_server := api.NewServer(conf)
	status_server := status.NewServer()

	mux := http.NewServeMux()

	mux.Handle("/v1/", http.StripPrefix("/v1", api.HandlerFromMux(api_server, http.NewServeMux())))
	mux.Handle("/", status.HandlerFromMux(status_server, http.NewServeMux()))

	s := &http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8000",
	}

	log.Fatal(s.ListenAndServe())
}
