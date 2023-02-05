// Code generated by hertz generator.

package main

import (
	handler "gateway/biz/handler"
	"gateway/internal/videos"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	r.StaticFS("/media", &app.FS{
		Root:            videos.Dir(),
		AcceptByteRange: true,
		PathRewrite: func(c *app.RequestContext) []byte {
			return c.Path()[len("/media"):]
		},
	})
}
