// Code generated by hertz generator.

package main

import (
	"common/snowy"
	"common/utils"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"gateway/internal/db"
	"gateway/internal/jwt"
	"gateway/internal/services"
	"gateway/internal/videos"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/logger/zap"
	"github.com/urfave/cli/v2"
)

const (
	cli_storage = "storage"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     cli_storage,
				Required: true,
				Usage:    "the video storage path",
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		hlog.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	if err := initialize(ctx); err != nil {
		return err
	}

	h := server.Default(
		server.WithHostPorts(":8000"),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
		server.WithMaxRequestBodySize(1*1024*1024*1024),
	)

	register(h)
	h.Spin()
	return nil
}

// TODO: Read from config files or command line
func initialize(ctx *cli.Context) error {
	hlog.SetLogger(getLogger())
	utils.InitEnvVars()

	if utils.Env.Secret == "" {
		secret := make([]byte, 256/8)
		_, err := rand.Read(secret)
		if err != nil {
			return fmt.Errorf("unable to generate secrets")
		} else {
			return fmt.Errorf("generated secret: %s", hex.EncodeToString(secret))
		}
	}

	if err := db.Init(utils.GormDialector()); err != nil {
		return err
	}

	if err := jwt.Init(utils.Env.Secret, time.Hour*24*7); err != nil {
		return err
	}

	if err := videos.Init(ctx.Path(cli_storage), utils.Env.Base); err != nil {
		return err
	}

	if err := services.Init(); err != nil {
		return err
	}

	if err := snowy.Init(); err != nil {
		return err
	}

	return nil
}

func getLogger() *zap.Logger {
	logger := zap.NewLogger(zap.WithCoreEnc(utils.GetZapEncoder()))
	logger.SetLevel(hlog.LevelTrace)
	return logger
}
