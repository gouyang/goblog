package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/russross/blackfriday"
)

const timeFormat = "2006-01-02 15:04:05"

func time2String(t time.Time) string {
	return t.Format(timeFormat)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func str2html(raw []byte) template.HTML {
	return template.HTML(string(raw))
}

func markdown2HtmlTemplate(raw []byte) template.HTML {
	return template.HTML(string(blackfriday.MarkdownCommon(raw)))
}

var funcMap = template.FuncMap{
	"time2string":           time2String,
	"str2html":              str2html,
	"markdown2htmltemplate": markdown2HtmlTemplate,
}

func compileTemplate(tmpl string) *template.Template {
	base := path.Join("templates", "base.html")
	rendertmpl := path.Join("templates", tmpl+".html")
	t := template.New("")
	t = template.Must(t.Funcs(funcMap).ParseGlob(base))
	return template.Must(t.ParseFiles(rendertmpl))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) error {
	t := compileTemplate(tmpl)
	err := t.ExecuteTemplate(w, "base", p)
	if err != nil {
		return err
	}
	return nil
}

func secret(user, realm string) string {
	if user == "admin" {
		return "$1$HRJLR.AX$cqPG8rm2J51.WKfgL15/H1"
	}
	return ""
}
