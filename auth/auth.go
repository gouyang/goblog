package auth

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"net/http"
)

func Secret(user, realm string) string {
	if user == "admin" {
		return "$1$HRJLR.AX$cqPG8rm2J51.WKfgL15/H1"
	}
	return ""
}

func LoginAdmin(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	fmt.Fprint(w, "<h1><a href=\"/\">Home</a></h1>")
	fmt.Fprint(w, "<h1><a href=\"/blog/new/\">Add New Post</a></h1>")
	fmt.Fprint(w, "<h1><a href=\"/blogs/delete\">Delete</a></h1>")
}
