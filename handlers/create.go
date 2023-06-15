package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/celiobj/todoApi/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"MEssage": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"MEssage": fmt.Sprintf("Todo inserido com sucesso: %d", id),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
