package main

import (
	"io/ioutil"
	"path/filepath"
	"plugin"
)

const (
	PluginsFolder   = "plugins-build"
	PluginExtension = ".so"
)

func main() {
	files, _ := ioutil.ReadDir(PluginsFolder)
	for _, f := range files {
		if filepath.Ext(f.Name()) == PluginExtension {
			println(pluginWeight(f.Name()))
		}
	}
}

func pluginWeight(pluginName string) int {
	fn := getFunction(pluginName, "Weight")
	weight := fn.(func() int)()
	return weight
}

func getFunction(pluginName, functionName string) plugin.Symbol {
	p, _ := plugin.Open(filepath.Join(PluginsFolder, pluginName))
	function, _ := p.Lookup(functionName)
	return function
}
