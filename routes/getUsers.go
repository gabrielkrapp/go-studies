package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	userModel "estudos-go/model/db"
)

func GetUsersRequest(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		listOfUsers, err := userModel.GetListOfUsers(db)

		if err != nil {
			log.Printf("Failed to get users: %v", err)
			http.Error(w, "Failed to get users", http.StatusInternalServerError)
			return
		}

		if len(listOfUsers) == 0 {
			log.Println("No users found")
			http.Error(w, "No users found", http.StatusNotFound)
			return
		}

		log.Printf("Get List of Users successfully")
		responseData, err := json.Marshal(listOfUsers)

		if err != nil {
			log.Printf("Error marshalling users list: %v", err)
			http.Error(w, "Error processing data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(responseData); err != nil {
			log.Printf("Error sending response: %v", err)
			http.Error(w, "Failed to send response", http.StatusInternalServerError)
		}
	}
}
