package main

import (
	"context"
	"embed"
	"log"
	"olydown/logic"
	"olydown/types"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	app := &App{}

	err := wails.Run(&options.App{
		Title:  "Basic Demo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

type App struct {
	ctx context.Context
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *App) shutdown(ctx context.Context) {}

func (b *App) PopulateImages() []string {
	response, err := logic.GetImageList()
	if err != nil {
		return response
	}
	return response
}

func (b *App) GetImageScreennail(filename string) (types.ImageResponse, error) {
	response, err := logic.GetImageScreennail(filename)
	if err != nil {
		return types.ImageResponse{}, err
	}
	return response, nil
}
