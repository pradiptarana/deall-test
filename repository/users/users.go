package users

import (
	"deal-test/model"
	"fmt"
)

// type UsersRepository struct {
// 	db *sql.DB
// }

// func NewUserRepository(db *sql.DB) *UsersRepository {
// 	return &UsersRepository{db}
// }

// User represents a user in the system
type User struct {
	Username string
	Password string
}

// UserRepository is responsible for storing and retrieving user information
type UserRepository struct {
	users map[string]*model.UserSocial
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*model.UserSocial),
	}
}

// SignUp adds a new user to the repository
func (ur *UserRepository) SignUp(us *model.UserSocial) error {
	// Check if the username already exists
	if _, exists := ur.users[us.Username]; exists {
		return fmt.Errorf("username %s already exists", us.Username)
	}

	// Store the user information
	ur.users[us.Username] = us

	fmt.Printf("User %s signed up successfully!\n", us.Username)
	return nil
}

// Login verifies the username and password for login
func (ur *UserRepository) Login(username, password string) (bool, error) {
	// Check if the username exists
	user, exists := ur.users[username]
	if !exists {
		return false, fmt.Errorf("user %s not found", username)
	}

	// Check if the password matches
	if user.Password == password {
		fmt.Printf("User %s logged in successfully!\n", username)
		return true, nil
	}

	return false, fmt.Errorf("incorrect password for user %s", username)
}
