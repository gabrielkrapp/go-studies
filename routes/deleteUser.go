package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	userModel "estudos-go/model/db"
)

type deleteUserInfosRequest struct {
	Username string `json:"username"`
}

func DeleteUserRequest(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteUserInfosRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := userModel.DeleteUserInDatabase(db, req.Username); err != nil {
			log.Printf("Failed to delete user: %v", err)
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}

		log.Printf("User %s deleted successfully", req.Username)

		if _, err := w.Write([]byte("User deleted successfully")); err != nil {
			log.Printf("Error to send a message: %v", err)
		}
	}
}
