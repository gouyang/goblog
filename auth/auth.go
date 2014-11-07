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
	fmt.Fprint(w, "<p><a href=\"/\">Home</a></p>")
	fmt.Fprint(w, "<p><a href=\"/blog/new/\">Add New Post</a></p>")
	fmt.Fprint(w, "<p><a href=\"/blogs/delete\">Delete</a></p>")
}
