// Package gdgcloudtorino provides a set of Cloud Functions samples.
package gdgcloudtorino

import (
	"context"
	"log"

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

	return app
}

var firebaseApp *firebase.App
var firebaseAuth *auth.Client

// la funzione init servira ad inizalizzare le componenti fondamentali della function in comune a tutte le richieste
func init() {
	firebaseApp = initializeAppDefault()
	client, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	firebaseAuth = client
}
