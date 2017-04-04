package main

import (
	"context"
	"go-plugin-example/models"
	"io/ioutil"
	"path/filepath"
	"plugin"
	"sort"
)

type internalPlugin struct {
	name   string
	weight int
}

// callHandler calls handler function, stored in given plugin
func (pl internalPlugin) callHandler(ctx context.Context, data models.Data) models.Data {
	fn := getFunction(pl.name, "Handler")
	updatedData := fn.(func(context.Context, models.Data)models.Data)(ctx, data)
	return updatedData
}

type internalPlugins []internalPlugin

// initPlugins returns sorted slice with detected plugins
func initPlugins() (pls internalPlugins) {

	// finding plugins
	files, _ := ioutil.ReadDir(PluginsFolder)
	for _, f := range files {
		if filepath.Ext(f.Name()) == PluginExtension {
			pls = append(pls, internalPlugin{
				name:   f.Name(),
				weight: getPluginWeight(f.Name()),
			})
		}
	}

	// sorting plugins slice by weight
	if len(pls) > 0 {
		sort.Slice(pls, func(a, b int) bool {
			return pls[a].weight >= pls[b].weight
		})
	}
	return
}

// processPipeline
func (pls internalPlugins) processPipeline(data models.Data) models.Data {
	var updatedData = data
	for _, pl := range pls {
		updatedData = pl.callHandler(context.Background(), updatedData)
	}
	return updatedData
}

// pluginWeight extracts weight info from plugin
func getPluginWeight(pluginName string) int {
	fn := getFunction(pluginName, "Weight")
	weight := fn.(func()int)()
	return weight
}

// getFunction
func getFunction(pluginName, functionName string) plugin.Symbol {
	p, _ := plugin.Open(filepath.Join(PluginsFolder, pluginName))
	function, _ := p.Lookup(functionName)
	return function
}
