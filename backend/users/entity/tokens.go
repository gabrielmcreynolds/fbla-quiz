package entity

import (
	"backend/errorCodes"
	"backend/helpers"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

type RefreshToken struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	CreatedByIp string             `json:"createdByIp" bson:"createdByIp"`
}

func (token *RefreshToken) GenerateJWT() (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["userId"] = token.UserId.Hex()
	claims["id"] = token.ID.Hex()
	claims["createdByIp"] = token.CreatedByIp
	tokenString, err := refreshToken.SignedString([]byte(os.Getenv("refreshSecret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewRefreshTokenFromJWT(tokenString string) (*RefreshToken, error) {
	token, err := helpers.DecodeRefreshToken(tokenString)
	if err != nil {
		return nil, errors.New(errorCodes.InvalidRefreshToken)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, _ := primitive.ObjectIDFromHex(claims["userId"].(string))
		id, _ := primitive.ObjectIDFromHex(claims["id"].(string))
		return &RefreshToken{
			ID:          id,
			UserId:      userId,
			CreatedByIp: claims["createdByIp"].(string),
		}, nil
	} else {
		return nil, errors.New(errorCodes.InvalidRefreshToken)
	}
}

type AccessToken struct {
	UserID primitive.ObjectID `json:"userId"`
	Email  string             `json:"email"`
}

func (token *AccessToken) GenerateJWT() (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["userId"] = token.UserID
	claims["email"] = token.Email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tokenString, err := accessToken.SignedString([]byte(os.Getenv("accessSecret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
