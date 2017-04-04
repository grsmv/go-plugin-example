package main

import (
	"context"
	"go-plugin-example/models"
)

const (
	PluginsFolder   = "plugins-build"
	PluginExtension = ".so"
)

func main() {

	var plugins = initPlugins()

	// running plugin updater
	go initPluginUpdater()

	var initialData = models.Data{A: 1}
	var processedData = plugins.processPipeline(context.Background(), initialData)

	println(processedData.A)
}
