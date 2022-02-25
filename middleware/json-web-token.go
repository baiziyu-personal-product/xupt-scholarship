package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
)

const YYYY_MM_DD_HH_mm_ss = "2006-01-02 15:04:05"

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
	encKey = []byte("GCM_AES_256_secret_shared_key_32")
)

type UserClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	Time  string `json:"time"`
}

func JwtVerify() context.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, sigKey)
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(UserClaims)
	})
	return verifyMiddleware
}

func GenerateToken(email string) string {
	signer := jwt.NewSigner(jwt.HS256, sigKey, 2*24*time.Hour)
	claims := UserClaims{
		Email: email,
		Time:  time.Now().Format(YYYY_MM_DD_HH_mm_ss),
	}
	token, err := signer.Sign(claims)
	if err != nil {
		return ""
	}
	return string(token)
}

func ParseToken(ctx iris.Context) *UserClaims {
	claims := jwt.Get(ctx).(*UserClaims)
	return claims
}
