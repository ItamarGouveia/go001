package handler

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	// Alterado para servir o arquivo index.html
	server.GET("/", func(context *Context) {
		// Caminho relativo ao arquivo index.html
		http.ServeFile(w, r, "index.html")
	})

	server.Handle(w, r)
}
