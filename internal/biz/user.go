package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u *UserUsecase) Register(ctx context.Context, user *User) error {
	// Check if the email is already registered
	existingUser, err := u.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already registered")
	}

	// Save the user to the database
	err = u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// AuthenticateUser authenticates a user based on the provided credentials.
func (s *UserUsecase) AuthenticateUser(ctx context.Context, username, password string) (*User, error) {
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	// Validate the password
	if user.Password != password {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

// GenerateToken generates a JWT token for the authenticated user.
// func (s *UserUsecase) GenerateToken(user *User) (string, error) {
// 	// Set the token expiration time
// 	expiration := time.Now().Add(24 * time.Hour)

// 	// Create a new token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": user.ID,
// 		"exp":     expiration.Unix(),
// 	})

// 	// Sign the token with a secret key
// 	tokenString, err := token.SignedString([]byte(ChatGPTSecretToken))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
