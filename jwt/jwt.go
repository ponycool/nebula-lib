package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Jwt struct {
	Secret        []byte
	Payload       *Payload
	SigningMethod jwt.SigningMethod
}

// Payload 载荷
type Payload struct {
	Subject   string
	Issuer    string
	ExpiresAt int64
	Scope     []string
	Admin     bool
}

// PublicClaims 公有声明,JWT 签发方可以自定义的声明，但是为了避免冲突，应该在 Registry 中定义它们
// 声明参考 https://www.iana.org/assignments/jwt/jwt.xhtml
type PublicClaims struct {
	// 作用域
	Scope []string `json:"scope"`
}

// PrivateClaims 私有声明
type PrivateClaims struct {
	Admin bool `json:"admin"`
}

type Claims struct {
	jwt.RegisteredClaims
	PublicClaims
	PrivateClaims
}

// Create 创建Token
func (j Jwt) Create() (string, error) {
	now := time.Now()
	if j.SigningMethod == nil {
		j.SigningMethod = jwt.SigningMethodHS256
	}
	claims := Claims{
		jwt.RegisteredClaims{
			// 受众群体
			Audience: nil,
			// 到期时间
			ExpiresAt: jwt.NewNumericDate(now.Add(60 * time.Second)),
			// 编号
			ID: "",
			// 签发人
			Issuer: "nebula",
			// 签发时间
			IssuedAt: jwt.NewNumericDate(now),
			// 生效时间
			NotBefore: jwt.NewNumericDate(now),
			// 主题
			Subject: "login",
		},
		PublicClaims{},
		PrivateClaims{
			Admin: j.Payload.Admin,
		},
	}

	if j.Payload.ExpiresAt != 0 {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.Payload.ExpiresAt) * time.Second))
	}

	if j.Payload.Subject != "" {
		claims.Subject = j.Payload.Subject
	}

	if j.Payload.Issuer != "" {
		claims.Issuer = j.Payload.Issuer
	}
	if len(j.Payload.Scope) > 0 {
		claims.Scope = j.Payload.Scope
	}

	token := jwt.NewWithClaims(j.SigningMethod, claims)
	return token.SignedString(j.Secret)
}

// Parse 解析JWT Token
func (j Jwt) Parse(token string) (jwt.Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Secret, nil
	})

	if err != nil {
		return Claims{}, err
	}

	if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
		return claims, nil
	} else {
		return Claims{}, err
	}
}
