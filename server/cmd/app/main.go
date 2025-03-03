package main

import (
	"encoding/json"
	"log"
	"net"
	"strings"

	"github.com/tsraveling/dog-fight/server/internal/auth"
	"github.com/tsraveling/dog-fight/server/internal/db"
	"github.com/tsraveling/dog-fight/server/internal/repositories"
)

// Update AuthResponse to use SafeCaptain.
type AuthResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message,omitempty"`
	Captain *repositories.SafeCaptain    `json:"captain,omitempty"`
}

type AuthRequest struct {
	Action   string `json:"action"`   // "login" or "enlist"
	Username string `json:"username"` // the user's username
	Password string `json:"password"` // the user's password
}

var captainRepo repositories.CaptainRepository

func main() {
	// Initialize the database and repository.
	sqliteDB, err := db.OpenDB("internal/db/data.db")
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer sqliteDB.Close()

	captainRepo, err = repositories.NewCaptainRepository(sqliteDB)
	if err != nil {
		log.Fatalf("Error initializing Captain Repository: %v", err)
	}

	// Set up a TCP socket server on port 8080.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	log.Println("Socket server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		// Handle each connection concurrently.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	// Step 1: Decode the incoming JSON request.
	var req AuthRequest
	if err := decoder.Decode(&req); err != nil {
		log.Println("Error decoding request:", err)
		return
	}

	// Step 2: Process the request based on the action.
	var resp AuthResponse
	action := strings.ToLower(req.Action)
	switch action {
	case "enlist":
		id, err := auth.Enlist(captainRepo, req.Username, req.Password)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
		} else {
			resp.Success = true
			resp.Message = "Enlist successful. Captain ID: " + id
		}
	case "login":
		captain, err := auth.Login(captainRepo, req.Username, req.Password)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
		} else {
			resp.Success = true
			resp.Message = "Login successful"
			resp.Captain = captain
		}
	default:
		resp.Success = false
		resp.Message = "Unknown action: " + req.Action
	}

	// Step 3: Encode and send the JSON response.
	if err := encoder.Encode(resp); err != nil {
		log.Println("Error encoding response:", err)
		return
	}
}
