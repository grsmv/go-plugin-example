package main

import (
	"path/filepath"
	"plugin"
	"io/ioutil"
	"sort"
)

type internalPlugin struct {
	name   string
	weight int
}

type internalPlugins []internalPlugin

// initPlugins returns sorted slice with detected plugins
func initPlugins() (pls internalPlugins) {

	// finding plugins
	files, _ := ioutil.ReadDir(PluginsFolder)
	for _, f := range files {
		if filepath.Ext(f.Name()) == PluginExtension {
			pls = append(pls, internalPlugin{
				name: f.Name(),
				weight: pluginWeight(f.Name()),
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

// pluginWeight extracts weight info from plugin
func pluginWeight(pluginName string) int {
	fn := getFunction(pluginName, "Weight")
	weight := fn.(func() int)()
	return weight
}

// getFunction
func getFunction(pluginName, functionName string) plugin.Symbol {
	p, _ := plugin.Open(filepath.Join(PluginsFolder, pluginName))
	function, _ := p.Lookup(functionName)
	return function
}
