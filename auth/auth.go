package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"net/http"
)

func Secret(user, realm string) string {
	if user == "guohua" {
		return "$1$HRJLR.AX$cqPG8rm2J51.WKfgL15/H1"
	}
	return ""
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprint(w, "<html><body><h1>Hello %s</h1></body></html>", r.Username)
}

func main() {
	authen := auth.NewBasicAuthenticator("localhost", Secret)
	http.HandleFunc("/", authen.Wrap(handle))
	http.ListenAndServe(":8080", nil)
}
