package jwt

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/yesh0/mdouyin/common/utils"
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

func AuthorizedUser(c *app.RequestContext) (uint64, error) {
	if c.Keys == nil {
		return 0, utils.ErrorUnauthorized
	}
	if id, ok := c.Keys[JwtIdField]; ok {
		if i, ok := id.(uint64); ok {
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
