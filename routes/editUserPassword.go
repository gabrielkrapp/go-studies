package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	userModel "estudos-go/model/db"

	"golang.org/x/crypto/bcrypt"
)

func EditPasswordRequest(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userInfosRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error while hashing the password", http.StatusInternalServerError)
			return
		}

		if err := userModel.EditUserPasswordInDatabase(db, req.Username, string(hashedPassword)); err != nil {
			log.Printf("Failed to edit user password: %v", err)
			http.Error(w, "Failed to edit user password", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Password edited successfully"))
	}
}
