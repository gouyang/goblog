package core

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/russross/blackfriday"
)

const FORMAT = "2006-01-02 15:04:05"

func Time2String(t time.Time) string {
	return t.Format(FORMAT)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Str2html(raw []byte) template.HTML {
	return template.HTML(string(raw))
}

func Markdown2HtmlTemplate(raw []byte) template.HTML {
	return template.HTML(string(blackfriday.MarkdownCommon(raw)))
}

var FuncMap = template.FuncMap{
	"time2string":           Time2String,
	"str2html":              Str2html,
	"markdown2htmltemplate": Markdown2HtmlTemplate,
}

func compileTemplate(tmpl string) *template.Template {
	base := path.Join("templates", "base.html")
	rendertmpl := path.Join("templates", tmpl+".html")
	t := template.New("")
	t = template.Must(t.Funcs(FuncMap).ParseGlob(base))
	return template.Must(t.ParseFiles(rendertmpl))
}

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t := compileTemplate(tmpl)
	err := t.ExecuteTemplate(w, "base", p)
	CheckErr(err)
}

func Secret(user, realm string) string {
	if user == "admin" {
		return "$1$HRJLR.AX$cqPG8rm2J51.WKfgL15/H1"
	}
	return ""
}
