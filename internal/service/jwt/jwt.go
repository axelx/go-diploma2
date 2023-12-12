package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/axelx/go-diploma2/internal/models"
	"github.com/axelx/go-diploma2/internal/utils"
)

var jwtKey = []byte("my_secret_key")

func CreateJWT(user models.User) models.JWT {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := jwt.StandardClaims{
		Id:        utils.IntToStr(user.ID),
		Subject:   user.Login,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Println("JWT: error creteJWT", err)
		return models.JWT("")
	}
	fmt.Println(tokenString)
	return models.JWT(tokenString)
}

func CheckJWT(strJWT models.JWT) bool {
	token, err := jwt.ParseWithClaims(string(strJWT), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(string(jwtKey)), nil
	})
	fmt.Println("CheckJWT:1: ", token, err)
	if err != nil {
		//!!!!!!!
		//return false
		return true
	}

	return true
}

type PayloadJwtStruct struct {
	ID  string `json:"jti"`
	Exp int    `json:"exp"`
}

func UserIDFromJwt(token string) int {
	s := strings.Split(token, ".")
	fmt.Println("UserIDFromJwt: ", s[1])

	data, err := base64.StdEncoding.DecodeString(s[1])
	fmt.Println("UserIDFromJwt: data:", string(data))
	if err != nil {
		log.Println("error, base64.StdEncoding.DecodeString  ", err, string(data), "костыль, ", token)
		fmt.Println("00")
	}
	// костыль. Непонятно почему криво расшифровывает...
	data = []byte(string(data) + "\"}")
	fmt.Printf("data:: %q\n", data)

	p := PayloadJwtStruct{}
	err3 := json.Unmarshal(data, &p)
	if err3 != nil {
		log.Println(err3)
	}
	fmt.Println("UserIDFromJwt: DAta JWT:", p)

	return utils.StrToInt(p.ID)
}
