package services

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	models "github.com/bogdanvv/mabooks-api"
	"github.com/bogdanvv/mabooks-api/pkg/repository"
	"github.com/golang-jwt/jwt"
)

type AuthorizationService struct {
	repo repository.Authorization
}

type jwtTokenClaims struct {
	jwt.Claims
	Id         string `json:"id"`
	Role       string `json:"role"`
	Expiration int64  `json:"expiration"`
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) SignUp(signUpData models.SignUpBody) (string, error) {
	signUpData.Password = createHashPassword(signUpData.Password)
	signUpData.Role = "user"

	id, err := s.repo.CreateUser(signUpData)
	if err != nil {
		fmt.Println("err>>>", err.Error())
	}

	return id, nil
}

func (s *AuthorizationService) Login(loginBody models.LoginBody) (models.LoginResponse, error) {
	var loginResponse models.LoginResponse
	hashedPassword := createHashPassword(loginBody.Password)
	user, err := s.repo.GetUserByLoginData(loginBody.Email, hashedPassword)
	if err != nil {
		return loginResponse, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtTokenClaims{
		Id:         user.Id,
		Role:       user.Role,
		Expiration: time.Now().Add(time.Hour * 720).Unix(),
	})

	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return loginResponse, err
	}

	loginResponse.Id = user.Id
	loginResponse.Email = user.Email
	loginResponse.Username = user.Username
	loginResponse.Phone = user.Phone
	loginResponse.Role = user.Role
	loginResponse.AccessToken = token
	loginResponse.IsActive = user.IsActive

	return loginResponse, nil
}

func (s *AuthorizationService) HandleToken(token string) {
	tokenParsed, err := jwt.Parse(token, validateToken)

	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
		fmt.Printf("claims[id]>>> %s \n", claims["id"])
		fmt.Printf("claims[role]>>> %s \n", claims["role"])
	} else {
		fmt.Println(err)
	}
}

func (s *AuthorizationService) ReissueTokens(refreshToken string) (models.TokenPair, error) {
	tokenPair := models.TokenPair{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}

	return tokenPair, nil
}

func validateToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	fmt.Println("token.Header[alg]", token.Header["alg"])

	return []byte(os.Getenv("JWT_SECRET")), nil
}

func createHashPassword(password string) string {
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(password))
	return fmt.Sprintf("%x", hashedPassword.Sum([]byte(os.Getenv("HASH_SALT"))))
}
