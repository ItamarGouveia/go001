package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
)

// Estrutura para resposta de erro em formato JSON
type ErrorResponse struct {
	Message string `json:"message"`
}

func HtmlRendering(w http.ResponseWriter, r *http.Request) {
	// Caminho para o arquivo HTML na pasta "public"
	htmlFile := filepath.Join("go001", "public", "seu_arquivo.html")

	// Log para verificar o caminho do arquivo
	log.Printf("Tentando servir o arquivo: %s\n", htmlFile)

	// Verificar se o arquivo existe
	_, err := filepath.Glob(htmlFile)
	if err != nil {
		log.Printf("Erro ao encontrar o arquivo: %v", err)
		// Retornando erro como JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Erro interno do servidor"})
		return
	}

	// Log para confirmar que o arquivo foi servido corretamente
	log.Printf("Arquivo %s servido com sucesso.\n", htmlFile)
}
