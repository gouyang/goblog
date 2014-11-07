package auth

import (
	"html/template"
	"log"
	"net/http"
	"path"

	httpauth "github.com/abbot/go-http-auth"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Secret(user, realm string) string {
	if user == "admin" {
		return "$1$HRJLR.AX$cqPG8rm2J51.WKfgL15/H1"
	}
	return ""
}

func LoginAdmin(w http.ResponseWriter, r *httpauth.AuthenticatedRequest) {
	tmpl := path.Join("templates", "admin.html")
	t, err := template.ParseFiles(tmpl)
	CheckErr(err)
	err = t.Execute(w, "")
	CheckErr(err)
}
