package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/home/cmd/home/api"
	"github.com/lishimeng/home/internal/etc"
	"github.com/lishimeng/home/internal/setup"
	"github.com/lishimeng/home/static"
	"net/http"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}

		builder.
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			EnableWeb(":80", api.Application).
			ComponentBefore(setup.Setup).
			PrintVersion()
		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
