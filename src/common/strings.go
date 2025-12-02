package common

import (
	"crypto/rand"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/naeemaei/golang-clean-web-api/config"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func CheckPassword(password string) bool {
	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLength {
		return false
	}

	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}

	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}

	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}

	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}

	return true
}

// secureRandomInt generates a cryptographically secure random integer in [0, max)
func secureRandomInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(nBig.Int64())
}

func GeneratePassword() string {
	var password strings.Builder

	cfg := config.GetConfig()
	passwordLength := cfg.Password.MinLength + 2
	minSpecialChar := 2
	minNum := 3
	if !cfg.Password.IncludeDigits {
		minNum = 0
	}

	minUpperCase := 3
	if !cfg.Password.IncludeUppercase {
		minUpperCase = 0
	}

	minLowerCase := 3
	if !cfg.Password.IncludeLowercase {
		minLowerCase = 0
	}

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := secureRandomInt(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := secureRandomInt(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := secureRandomInt(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := secureRandomInt(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := secureRandomInt(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	
	// Secure shuffle
	for i := len(inRune) - 1; i > 0; i-- {
		j := secureRandomInt(i + 1)
		inRune[i], inRune[j] = inRune[j], inRune[i]
	}
	return string(inRune)
}

func GenerateOtp() string {
	cfg := config.GetConfig()
	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))   // 10^d-1 100000
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1) // 999999 = 1000000 - 1 (10^d) -1

	var num = secureRandomInt(max-min+1) + min
	return strconv.Itoa(num)
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// To snake case : CountryId -> country_id
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
