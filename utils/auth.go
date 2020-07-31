package utils

import (
	"fmt"
	"time"
	constant "cabservice/constants"
	jwt "github.com/dgrijalva/jwt-go"
)

/*
Function name : CreateToken
Description   : this function creates encrypted jwt token
Params        : userid
Return        : token(string)
*/

func CreateToken(userId string) string {
	fmt.Println("fcm payload")

	expiryTimeSec := int64(time.Now().Unix()) + 31622400
	claim := TokenClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiryTimeSec,
			Issuer:    "Aatmaja",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(constant.ENCRYPTION_KEY))
	if nil != err {
		fmt.Println(err.Error())
	}
	return ss
}

/*
Function name : DecodeToken
Description   : DecodeToken decodes the JWT token
Params        : token(string)
Return        : decodedtoken(string), error
*/


func DecodeToken(tokenStr string) (string, error) {
	claim := TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.ENCRYPTION_KEY), nil
	})

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Uid, claims.StandardClaims.ExpiresAt)
		return claims.Uid, nil
	} else {
		fmt.Printf("DecodeToken 1 %+v", err)
		return "", err
	}
}

type TokenClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}
