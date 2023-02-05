package jwt

import (
	"common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// Attaches authorization info to the current request context.
//
// Use `AuthorizedUser` or `IsAuthorized` to check.
func Attach(c *app.RequestContext, token string) error {
	id, name, err := Validate(token)
	if err != nil {
		return err
	}

	if c.Keys == nil {
		c.Keys = make(map[string]interface{}, 2)
	}
	c.Keys[JwtIdField] = id
	c.Keys[JwtNameField] = name
	return nil
}

func AuthorizedUser(c *app.RequestContext) (int64, error) {
	if c.Keys == nil {
		return 0, utils.ErrorUnauthorized
	}
	if id, ok := c.Keys[JwtIdField]; ok {
		if i, ok := id.(int64); ok {
			return i, nil
		}
		return 0, utils.ErrorInternalError
	}
	return 0, utils.ErrorUnauthorized
}

func IsAuthorized(c *app.RequestContext) bool {
	_, err := AuthorizedUser(c)
	return err == nil
}
