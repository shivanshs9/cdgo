package app

import (
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/skademlia"
	"github.com/shivanshs9/cdgo/pkg"
)

type App struct {
	WorkingDir string
	CurrNode   *pkg.Node
}

func InitApp(cfg *Config) (app *App, err error) {
	params := noise.DefaultParams()
	params.Host = cfg.Host
	params.Port = uint16(cfg.Port)
	params.Keys = skademlia.RandomKeys()

	node, err := noise.NewNode(params)
	if err != nil {
		return nil, err
	}

	app = &App{
		CurrNode: pkg.NewNode(node, cfg.Nick),
	}
	return app, nil
}
