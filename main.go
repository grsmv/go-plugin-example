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
	var (
		plugins       = initPlugins()
		initialData   = models.Data{A: 1}
		processedData = plugins.processPipeline(context.Background(), initialData)
	)

	println(processedData.A)
}
