package main

import (
	"log"

	"github.com/tsraveling/dog-fight/server/internal/db"
	"github.com/tsraveling/dog-fight/server/internal/repositories"
)

func main() {
	// Open or create the SQLite database.
	sqliteDB, err := db.OpenDB("internal/db/data.db")
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer sqliteDB.Close()

	// Initialize the Captain repository.
	captainRepo, err := repositories.NewCaptainRepository(sqliteDB)
	if err != nil {
		log.Fatalf("Error initializing Captain Repository: %v", err)
	}

	// Example: Create a new captain.
	newCaptain := repositories.Captain{
    	Name: "Captain Picard",
    }
	id, err := captainRepo.Create(newCaptain)
	if err != nil {
		log.Fatalf("Error creating captain: %v", err)
	}

	// Example: Retrieve the created captain.
	retrievedCaptain, err := captainRepo.Get(id)
	if err != nil {
		log.Fatalf("Error retrieving captain: %v", err)
	}
	log.Printf("Retrieved Captain: %+v", retrievedCaptain)

	// // Set up a simple HTTP server
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, Dog Fight!")
	// })

	// port := "8080"
	// log.Printf("Server is running on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}