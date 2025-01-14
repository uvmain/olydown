package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"olydown/logic"
	"olydown/network"
	"olydown/types"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	app := &App{}

	err := wails.Run(&options.App{
		Title:  "olydown",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// Frameless:  true,
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
	network.StartSsidTicker(b.ctx)
}

func (b *App) shutdown(ctx context.Context) {
	network.StopSsidTicker()
}

func (b *App) PopulateImages() []string {
	response, err := logic.GetImageList()
	if err != nil {
		return response
	}
	return response
}

func (b *App) GetImageThumbnail(filename string) (types.ImageResponse, error) {
	response, err := logic.GetImageThumbnail(filename)
	if err != nil {
		return types.ImageResponse{}, err
	}
	return response, nil
}

func (b *App) OpenDirectoryDialog(dialogOptions interface{}) (string, error) {
	directoryPath, err := runtime.OpenDirectoryDialog(b.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: fmt.Sprintf("%v", dialogOptions.(map[string]interface{})["defaultDirectory"]),
		Title:            fmt.Sprintf("%v", dialogOptions.(map[string]interface{})["title"]),
	})
	if err != nil {
		return "", fmt.Errorf("failed opening dialog - %s", err.Error())
	}
	return directoryPath, nil
}

func (b *App) GetSSID() (string, error) {
	response, err := network.GetSSID()
	if err != nil {
		return "", err
	}
	return response, nil
}
