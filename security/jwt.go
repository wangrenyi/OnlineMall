package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinemall/common"
	"onlinemall/model"
	"time"
)

var signingKey = []byte("d00cb3eb2c24433894f031e913f0b6e5")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(loginUser model.MstUserInfo) (string, error) {

	expireTime := time.Now().Add(3 * time.Hour)
	//username := util.EncodeMD5(loginUser.LoginName)
	//password := util.EncodeMD5(loginUser.Password)

	claims := Claims{
		loginUser.LoginName,
		loginUser.Password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Renyi,Wang",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(signingKey)

	return token, err
}

//jwt middleware
func JWTMiddleware(context *gin.Context) {
	authToken := context.GetHeader("Authorization")
	if authToken == "" {
		context.JSON(http.StatusOK, common.AuthError())
		context.Abort()
		return
	}

	claims, err := parseToken(authToken)
	if err != nil {
		context.JSON(http.StatusOK, common.Error(http.StatusUnauthorized, err.Error()))
		context.Abort()
		return
	}

	context.Set("claims", claims)
	context.Next()
}

// ParseToken parsing token
func parseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
