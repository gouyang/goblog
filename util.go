package main

import (
	"html/template"
	"path"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

const timeFormat = "2006-01-02 15:04:05"

var stripHTMLReplacer = strings.NewReplacer("\n", "<br />")

func time2String(t time.Time) string {
	return t.Format(timeFormat)
}

func str2html(raw []byte) template.HTML {
	return template.HTML(string(raw))
}

func markdown2HtmlTemplate(raw []byte) template.HTML {
	//str := string(raw)
	//s := stripHTMLReplacer.Replace(str)
	//unsafe := blackfriday.MarkdownCommon(raw)
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	html := markdownRender(raw)
	return template.HTML(html)
}

func markdownRender(content []byte) string {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS

	renderer := blackfriday.HtmlRenderer(htmlFlags, "", "")

	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK

	return string(blackfriday.Markdown(content, renderer, extensions))
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

func (p *page) renderTemplate() error {
	t := compileTemplate(p.Tmpl)
	err := t.ExecuteTemplate(p.W, "base", p.Post)
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
