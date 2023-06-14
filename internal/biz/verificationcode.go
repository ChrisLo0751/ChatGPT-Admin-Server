package biz

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// VerificationCodeRepo is the repository interface for verification code.
type VerificationCodeRepo interface {
	Create(ctx context.Context, code *VerificationCode) error
	FindByUserID(ctx context.Context, userID int64) (*VerificationCode, error)
}

// EmailService is the service for sending emails.
type EmailRepo interface {
	Send(to, subject, body string) error
}

type VerificationCodeUsecase struct {
	repo       VerificationCodeRepo
	emailRepo  EmailRepo
	log        *log.Helper
	codeLength int
}

// VerificationCode represents a verification code entity.
type VerificationCode struct {
	ID        int64
	UserID    int64
	Code      string
	ExpiresAt time.Time
}

// GenerateCode generates a random verification code.
func (v *VerificationCodeUsecase) GenerateCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(code)
}

// SendVerificationCode sends the verification code to the user's email.
func (v *VerificationCodeUsecase) SendVerificationCode(ctx context.Context, userID int64, email string) error {
	// Check if a verification code already exists for the user
	existingCode, err := v.repo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if existingCode != nil {
		return errors.New("verification code already sent")
	}

	// Generate a verification code
	code := v.GenerateCode(v.codeLength)
	expiresAt := time.Now().Add(10 * time.Minute)

	// Save the verification code to the database
	verificationCode := &VerificationCode{
		UserID:    userID,
		Code:      code,
		ExpiresAt: expiresAt,
	}
	err = v.repo.Create(ctx, verificationCode)
	if err != nil {
		return err
	}

	// Send the verification code via email
	subject := "Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", code)
	err = v.emailRepo.Send(email, subject, body)
	if err != nil {
		// If sending the email fails, you can choose to delete the verification code from the database.
		return err
	}

	return nil
}
