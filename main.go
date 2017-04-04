package main

import ()
import "go-plugin-example/models"

const (
	PluginsFolder   = "plugins-build"
	PluginExtension = ".so"
)

func main() {
	var (
		plugins     = initPlugins()
		initialData = models.Data{
			A: 1,
		}
	)

	println(plugins.processPipeline(initialData).A)
}
