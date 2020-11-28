// Package gdgcloudtorino provides a set of Cloud Functions samples.
package gdgcloudtorino

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func initializeAppDefault() *firebase.App {
	// [START initialize_app_default_golang]
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// [END initialize_app_default_golang]

	return app, client
}

var firebaseApp *firebase.App
var firebaseAuth *auth.Client

// la funzione init servira ad inizalizzare le componenti fondamentali della function in comune a tutte le richieste
func init() {
	firebaseApp = initializeAppDefault()
	client, err := firebaseApp.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	firebaseAuth = client
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

// [END functions_helloworld_http]
