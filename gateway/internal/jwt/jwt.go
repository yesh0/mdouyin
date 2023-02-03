package jwt

import (
	"common/utils"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-jwt/jwt/v4"
)

const (
	JwtIdField   = "i"
	JwtNameField = "n"
	JwtTimeField = "t"
)

var (
	hmacSecret []byte
	duration   time.Duration
)

func Init(secret string, timeSpan time.Duration) error {
	if hmacSecret != nil {
		return fmt.Errorf("jwt secret already initialized")
	}
	if timeSpan <= 0 {
		return fmt.Errorf("effective time span for a token must be positive")
	}

	bytes, err := hex.DecodeString(secret)
	if err != nil {
		return fmt.Errorf("not a hex encoded string")
	}

	hmacSecret = bytes
	duration = timeSpan

	return nil
}

func NewAuthorization(id uint64, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		JwtIdField:   strconv.FormatUint(id, 16),
		JwtNameField: name,
		JwtTimeField: strconv.FormatInt(time.Now().Unix(), 16),
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		hlog.Error("error generating jwt token", err)
		return "", err
	}

	return tokenString, nil
}

func Validate(tokenString string) (uint64, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return hmacSecret, nil
		} else {
			return nil, utils.ErrorUnauthorized
		}
	})
	if err != nil {
		return 0, "", utils.ErrorUnauthorized
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		createdAt, err := strconv.ParseInt(claims[JwtTimeField].(string), 16, 64)
		if err != nil {
			hlog.Error("jwt inconsistency detected")
			return 0, "", utils.ErrorInternalError.Wrap(err)
		}

		id, err := strconv.ParseInt(claims[JwtIdField].(string), 16, 64)
		if err != nil {
			hlog.Error("jwt inconsistency detected")
			return 0, "", utils.ErrorInternalError.Wrap(err)
		}

		if time.Since(time.Unix(createdAt, 0)) >= duration {
			return 0, "", utils.ErrorExpiredToken
		} else {
			return uint64(id), claims[JwtNameField].(string), nil
		}
	} else {
		hlog.Error("jwt inconsistency detected")
		return 0, "", utils.ErrorInternalError
	}
}
