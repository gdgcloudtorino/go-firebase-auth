// Package gdgcloudtorino provides a set of Cloud Functions samples.
package gdgcloudtorino

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// MockLogin is an HTTP Cloud Function with a request parameter.
func MockLogin(w http.ResponseWriter, r *http.Request) {
	var d struct {
		UID   string `json:"uid"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return
	}
	if d.UID == "" {
		return
	}
	token, err := firebaseAuth.CustomToken(r.Context(), d.UID)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	fmt.Fprint(w, token)
	return
}
