package domain

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	MinUsernameLength = 4
	MaxUsernameLength = 10
	MinPasswordLength = 4
	MaxPasswordLength = 30
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func ValidateUsername(username string) error {
	if len(username) < MinUsernameLength {
		return errors.New("username must be at least " + strconv.Itoa(MinUsernameLength) + " characters long")
	}

	if len(username) > MaxUsernameLength {
		return errors.New("username must be less than " + strconv.Itoa(MaxUsernameLength) + " characters long")
	}

	matched, err := regexp.MatchString(`^[a-z0-9]+$`, username)
	if err != nil {
		return errors.New("failed to validate username")
	}

	if !matched {
		return errors.New("username must contain only lowercase letters and numbers")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return errors.New("password must be at least " + strconv.Itoa(MinPasswordLength) + " characters long")
	}

	if len(password) > MaxPasswordLength {
		return errors.New("password must be less than " + strconv.Itoa(MaxPasswordLength) + " characters long")
	}

	matched, err := regexp.MatchString(`^[a-zA-Z0-9[:punct:]]+$`, password)
	if err != nil {
		return errors.New("failed to validate password")
	}

	if !matched {
		return errors.New("password must contain only letters, numbers, and special characters")
	}

	return nil
}
