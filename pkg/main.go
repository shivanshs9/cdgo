package pkg

import "github.com/shivanshs9/cdgo/pkg/daemon"

func StartServer() (app *daemon.App, err error) {
	cfg, err := daemon.InitConfig()
	if err != nil {
		return nil, err
	}
	app, err = daemon.InitApp(cfg)
	if err != nil {
		return nil, err
	}

	// Work with App

	return app, nil
}
