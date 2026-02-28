package services

import (
	"errors"
	"os"
	"time"

	"Apps-I_Desa_Backend/dtos"
	"Apps-I_Desa_Backend/models"
	"Apps-I_Desa_Backend/repositories"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(request *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	// Find user by username
	user, err := s.userRepo.FindByUsername(request.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("User not found: ", request.Username)
		return nil, errors.New("user not found")
	} else if err != nil {
		log.Error("Database error: ", err)
		return nil, errors.New("failed to retrieve user")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Error("Error compare", err)
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := generateJWTToken(user)
	if err != nil {
		log.Error("Error generating token: ", err)
		return nil, errors.New("failed to generate token")
	}

	return &dtos.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}

func (s *AuthService) Logout() *dtos.MessageResponse {
	return &dtos.MessageResponse{
		Message: "Logout successful!",
	}
}

// generateJWTToken creates a new JWT token for the authenticated user
func generateJWTToken(user *models.User) (string, error) {
	// Set expiration time
	expTime := time.Now().Add(1 * time.Hour)

	// Create claims
	claims := jwt.MapClaims{
		"village": user.VillageID,
		"exp":     expTime.Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate signed token
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT secret not configured")
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
