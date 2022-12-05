package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"query/model"
	"query/utils"
	"query/utils/errmsg"
	"strings"
	"time"
)

var (
	JwtKey = []byte(utils.JwtKey)
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 定义错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

// 生成token
func (j *JWT)CreateToken(claims MyClaims) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(j.JwtKey)
}

/**func SetToken(username string) (string,int) {
	expireTime := time.Now().Add(10*time.Hour)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token,err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCESS
}*/

// 解析token
func (j *JWT)ParserToken(tokenString string) (*MyClaims,error) {
	token,err := jwt.ParseWithClaims(tokenString,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey,nil
	})
	if err != nil {
		if ve,ok := err.(*jwt.ValidationError);ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil,TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil,TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil,TokenNotValidYet
			}
		}
	} else if token != nil {
		if claims,ok := token.Claims.(*MyClaims);ok && token.Valid {
			return claims,nil
		}
	}

	return nil,TokenInvalid
}

// 验证token
/**func CheckToken(token string) (*MyClaims,int) {
	setToken,_ := jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key,_ := setToken.Claims.(*MyClaims);setToken.Valid {
		return key,errmsg.SUCCESS
	}
	return nil,errmsg.ERROR
}*/

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		tokenHearder := ctx.Request.Header.Get("Authorization")

		if tokenHearder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK,gin.H{
				"code" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken := strings.Split(tokenHearder," ")
		if len(checkToken) == 0 {
			ctx.JSON(http.StatusOK,gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK,gin.H{
				"code" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		j := NewJWT()
		claims,err := j.ParserToken(checkToken[1])

		if err != nil {
			if err == TokenExpired {
				ctx.JSON(http.StatusOK,gin.H{
					"status":  errmsg.ERROR,
					"message": "token授权已过期,请重新登录",
					"data":    nil,
				})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusOK,gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			ctx.Abort()
			return
		}

		//ctx.Set("username",key.Username)
		ctx.Set("username", claims)
		ctx.Next()
	}
}

// 绑定token
func SetToken(ctx *gin.Context,user model.User)  {
	j := NewJWT()
	claims := MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix()-100,
			ExpiresAt: time.Now().Unix()+604800,
			Issuer: "blog",
		},
	}

	token,err := j.CreateToken(claims)

	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	ctx.JSON(http.StatusOK,gin.H{
		"status":  200,
		"username": user.Username,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
}