package auth

import (
	"errors"
	"strings"
	"unicode"

	"github.com/alexedwards/argon2id"
	"github.com/tsraveling/dog-fight/server/internal/repositories"
)

const (
	DefaultStartingMoney 	= 1000
	DefaultRank          	= "Zeo I"
	MinUsernameLength		= 3
	MaxUsernameLength		= 32
	MinPasswordLength		= 8
)

// validateCredentials validates the username and password.
func validateCredentials(username, password string) error {
	// Step 1: Trim and check lengths.
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	if len(username) < MinUsernameLength || len(username) > MaxUsernameLength {
		return errors.New("username must be between 3 and 32 characters")
	}
	if len(password) < MinPasswordLength {
		return errors.New("password must be at least 8 characters")
	}

	// Step 2: Check for at least one letter and one number in the password.
	var hasLetter, hasNumber bool
	for _, ch := range password {
		if unicode.IsLetter(ch) {
			hasLetter = true
		}
		if unicode.IsNumber(ch) {
			hasNumber = true
		}
	}
	if !hasLetter || !hasNumber {
		return errors.New("password must contain at least one letter and one number")
	}

	return nil
}

// Enlist creates a new captain account on this server.
func Enlist(repo repositories.CaptainRepository, username, password string) (string, error) {
	// Step 1: Validate credentials.
	if err := validateCredentials(username, password); err != nil {
		return "", err
	}

	// Step 1 (restart): Normalize the username.
	username = strings.ToLower(strings.TrimSpace(username))

	// Step 2: Check if the username already exists.
	_, err := repo.GetByUsername(username)
	if err == nil {
		return "", errors.New("username already taken")
	}

	// Step 3: Hash the provided password.
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	// Step 4: Create a new captain record with default rank and starting money.
	captain := repositories.Captain {
		Username: username,
		PasswordHash: hash,
		Rank: DefaultRank,
		Money: DefaultStartingMoney,
	}

	// Step 5: Insert the new captain record into the repository.
	id, err := repo.Create(captain)
	if err != nil {
		return "", errors.New("failed to create captain record")
	}
	return id, nil
}

// Login authenticates a captain using the provided username and password.
func Login(repo repositories.CaptainRepository, username, password string) (*repositories.SafeCaptain, error) {
	// Step 1: Normalize the username.
	username = strings.ToLower(strings.TrimSpace(username))

	// Step 2: Retrieve the captain record.
	captain, err := repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Step 3: Compare the provided password with the stored hash.
	match, err := argon2id.ComparePasswordAndHash(password, captain.PasswordHash)
	if err != nil || !match {
		return nil, errors.New("invalid credentials")
	}
	
	// Step 4: Build a safe copy of the captain
	safeCaptain := &repositories.SafeCaptain {
		ID: captain.ID,
		Username: captain.Username,
		Rank: captain.Rank,
		Money: captain.Money,
	}

	// Step 5: Return the sanitized captain record.
	return safeCaptain, nil
}