package providers

import (
	"errors"
	"fmt"
	"boilerplate/internal/utils"
	"log"
	"strings"
	"sync"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const UserId = "userId"

type Success struct {
	Code   int    `json:"code" example:"200"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

var once sync.Once

var mw *jwtAuthMiddleware

func NewJwtAuth(db *gorm.DB) JwtAuthMiddleware {
	once.Do(func() {
		var err error

		mw = &jwtAuthMiddleware{
			databaseDriver: db,
		}

		mw.authMiddleware, err = jwt.New(mw.prepareMiddleware())
		if err != nil {
			log.Println("JWT error")
		}
	})

	return mw
}

type JwtAuthMiddleware interface {
	Middleware() *jwt.GinJWTMiddleware
	Refresh(c *gin.Context)
}

type jwtAuthMiddleware struct {
	databaseDriver *gorm.DB
	authMiddleware *jwt.GinJWTMiddleware
}

func (mw *jwtAuthMiddleware) Middleware() *jwt.GinJWTMiddleware {
	return mw.authMiddleware
}

func (mw *jwtAuthMiddleware) prepareMiddleware() *jwt.GinJWTMiddleware {
	jwtSettings, err := utils.NewJwtEnvVars()
	fmt.Println("jwtSettings: ", jwtSettings)

	if err != nil {
		log.Println(err.Error())
	}

	middleware := &jwt.GinJWTMiddleware{
		Realm:                 jwtSettings.Realm(),
		Key:                   []byte(jwtSettings.Secret()),
		Timeout:               jwtSettings.Expiration(),
		MaxRefresh:            jwtSettings.RefreshTime(),
		IdentityKey:           UserId,
		IdentityHandler:       extractIdentityKeyFromClaims,
		Authorizator:          mw.isUserValid,
		Authenticator:         mw.authenticate,
		HTTPStatusMessageFunc: takeAppropriateErrorMessage,
		TimeFunc:              time.Now,
		TokenLookup:           "header: Authorization",
		TokenHeadName:         "Bearer",
	}

	return middleware
}

func (mw jwtAuthMiddleware) authenticate(c *gin.Context) (interface{}, error) {
	fmt.Println("authenticate")
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		return nil, errors.New("invalid request, apiKey is required")
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return nil, errors.New("invalid request, apiKey is required")
	}

	token := authorization[len("Bearer "):]
	jwtSettings, _ := utils.NewJwtEnvVars()
	claims, _ := utils.DecodeJWTToken(token, jwtSettings.Secret())
	c.Set("JWT_PAYLOAD", claims)
	return claims[UserId], nil
}

func (mw jwtAuthMiddleware) Refresh(c *gin.Context) {
	mw.Middleware().RefreshHandler(c)
}

func (mw jwtAuthMiddleware) isUserValid(data interface{}, _ *gin.Context) bool {
	return true
}

func extractIdentityKeyFromClaims(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return claims[UserId].(string)
}

func takeAppropriateErrorMessage(err error, _ *gin.Context) string {
	return err.Error()
}
