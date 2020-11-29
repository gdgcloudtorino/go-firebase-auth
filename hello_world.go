// Package gdgcloudtorino provides a set of Cloud Functions samples.
package gdgcloudtorino

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		fmt.Fprint(w, "Missing authorization")
		w.WriteHeader(400)
		return
	}

	idToken := strings.Replace(authorization, "Bearer ", "", -1)
	log.Printf("Authorization Bearer %s\n", idToken)
	token, err := firebaseAuth.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		log.Println("Error authorization", err)
		w.WriteHeader(401)
		errorMsg := err.Error()
		fmt.Fprintf(w, "Invalid authorization: %s", errorMsg)
		return
	}

	log.Printf("Verified ID token: %v\n", token)

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	/*w.Header().Set("Content-Type", "application/json")
	var response struct {
		Message string `json:"message"`
		UID     string `json:"uid"`
		Email   string `json:"email"`
	}*/
	fmt.Fprintf(w, "Hello, %s!, %s", html.EscapeString(d.Name), token.Firebase.Identities["email"])
}

// [END functions_helloworld_http]
