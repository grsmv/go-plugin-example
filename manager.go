package main

import (
	"io/ioutil"
	"path/filepath"
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

