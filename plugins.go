package main

import (
  "plugin"
  "path/filepath"
)

// pluginWeight
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
