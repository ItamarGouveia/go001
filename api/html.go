package handler

import (
	"net/http"
)

func HtmlRendering(w http.ResponseWriter, r *http.Request) {
	htmlFile := "../public/blog/teste.html"

	// Serve o arquivo HTML externo
	http.ServeFile(w, r, htmlFile)
}
