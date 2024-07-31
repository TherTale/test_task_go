package services

import (
	"contact-center-system/internal/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"time"
)

func ValidateOperator(operator *models.Operator) error {
	// Валидация имени
	if !isValidName(operator.FirstName) || !isValidName(operator.LastName) || !isValidName(operator.MiddleName) {
		return errors.New("invalid name format")
	}
	// Валидация города
	if operator.City == "" {
		return errors.New("city cannot be empty")
	}
	// Валидация номера телефона
	if !isValidPhoneNumber(operator.PhoneNumber) {
		return errors.New("invalid phone number format")
	}
	// Валидация Email
	if !isValidEmail(operator.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

// Валидация формата имени
func isValidName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ\s-]+$`)
	return re.MatchString(name) && len(name) > 0 && len(name) <= 255
}

// Валидация формата номера телефона (8**********)
func isValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^8\d{10}$`)
	return re.MatchString(phone)
}

// Валидация формата Email
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Генерация случайного пароля
func GeneratePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(password)
}

// Хэширование пароля
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
