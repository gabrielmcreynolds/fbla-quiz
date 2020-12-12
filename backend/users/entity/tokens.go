package entity

import (
	"backend/errorCodes"
	"backend/helpers"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

// This is a expanded view of the token stored in the database
type RefreshToken struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	CreatedByIp string             `json:"createdByIp" bson:"createdByIp"`
}

// Generates a usable string for frontend to use in requests
func (token *RefreshToken) GenerateJWT() (string, *errorCodes.Slug) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims := refreshToken.Claims.(jwt.MapClaims)
	claims["userId"] = token.UserId.Hex()
	claims["id"] = token.ID.Hex()
	claims["createdByIp"] = token.CreatedByIp
	tokenString, err := refreshToken.SignedString([]byte(os.Getenv("refreshSecret")))
	if err != nil {
		return "", errorCodes.NewErrTokenToJWT()
	}
	return tokenString, nil
}

func NewRefreshTokenFromJWT(tokenString string) (*RefreshToken, *errorCodes.Slug) {
	token, err := helpers.DecodeRefreshToken(tokenString)
	if err != nil {
		return nil, errorCodes.NewErrInvalidRefreshToken()
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
		return nil, errorCodes.NewErrInvalidRefreshToken()
	}
}

// Token that won't be stored in database and thus can't be revoked
type AccessToken struct {
	UserID primitive.ObjectID `json:"userId"`
	Email  string             `json:"email"`
}

func (token *AccessToken) GenerateJWT() (string, *errorCodes.Slug) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["userId"] = token.UserID
	claims["email"] = token.Email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tokenString, err := accessToken.SignedString([]byte(os.Getenv("accessSecret")))
	if err != nil {
		return "", errorCodes.NewErrTokenToJWT()
	}
	return tokenString, nil
}
