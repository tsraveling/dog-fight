package auth

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/alexedwards/argon2id"
	"github.com/tsraveling/dog-fight/server/common/pkg/models/player"
)

const (
	DefaultStartingMoney = 1000
	DefaultRank          = "Zeo I"
	MinUsernameLength    = 3
	MaxUsernameLength    = 32
	MinPasswordLength    = 8
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

// Enlist creates a new player account on this server.
func Enlist(repo player.PlayerRepository, username, password string) (string, error) {
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

	// Step 4: Create a new player record with default rank and starting money.
	newPlayer := player.Player{
		Username:     username,
		PasswordHash: hash,
		Rank:         DefaultRank,
		Money:        DefaultStartingMoney,
	}

	// Step 5: Insert the new player record into the repository.
	id, err := repo.Create(newPlayer)
	if err != nil {
		return "", errors.New("failed to create player record")
	}
	return id, nil
}

// Login authenticates a player using the provided username and password.
func Login(repo player.PlayerRepository, username, password string) (*player.SafePlayer, error) {
	// Step 1: Normalize the username.
	username = strings.ToLower(strings.TrimSpace(username))

	// Step 2: Retrieve the player record.
	p, err := repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Step 3: Compare the provided password with the stored hash.
	match, err := argon2id.ComparePasswordAndHash(password, p.PasswordHash)
	if err != nil || !match {
		return nil, errors.New("invalid credentials")
	}

	// Step 4: Build a safe copy of the player.
	safePlayer := &player.SafePlayer{
		ID:       p.ID,
		Username: p.Username,
		Rank:     p.Rank,
		Money:    strconv.Itoa(p.Money),
	}

	// Step 5: Return the sanitized player record.
	return safePlayer, nil
}
