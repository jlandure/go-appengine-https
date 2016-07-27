package https

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/.well-known/acme-challenge/", challengeHandler)
}

var challenges = map[string]string{
	"<KEY>": "<VALUE>",
}

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	challenge := strings.Split(r.RequestURI, "/.well-known/acme-challenge/")[1]
	if responseToChallenge, ok := challenges[challenge]; !ok {
		http.Error(w, "Error", http.StatusNotFound)
		return
	} else {
		fmt.Fprintf(w, "%s", responseToChallenge)
	}
}
