package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sewakantor/sw-be/helpers"
	"net/http"
	"time"
)

type JwtCustomClaims struct {
	UserID uint      `json:"id"`
	Roles  string    `json:"roles"`
	Name string      `json:"name"`
	Company  string    `json:"company"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return c.JSON(http.StatusForbidden, helpers.BuildErrorResponse("Forbidden Access!", e, helpers.EmptyObj{}))
		},
	}
}

func (jwtConf *ConfigJWT) GenerateToken(userID uint, roles, name, company string) string {
	claims := JwtCustomClaims{
		userID,
		roles,
		name,
		company,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}