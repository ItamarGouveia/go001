package handler

import (
	"net/http"
	"path/filepath"
)

func HtmlRendering(w http.ResponseWriter, r *http.Request) {
	// Caminho para o arquivo HTML dentro da pasta "public"
	htmlFile := filepath.Join("go001", "public", "teste.html")

	// Serve o arquivo HTML da pasta public
	http.ServeFile(w, r, htmlFile)
}
