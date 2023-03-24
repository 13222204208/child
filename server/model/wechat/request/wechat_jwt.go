package request

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

// 然后我们定义JWT的过期时间
const TokenExpireDuration = time.Hour * 24 * 30 * 12

// 接下来还需要定义Secret
var MySecret = []byte("tokenyouwillneverguess")
