package main

import (
	//"io/ioutil"
	//"path/filepath"
)

const (
	PluginsFolder   = "plugins-build"
	PluginExtension = ".so"
)

func main() {
	plugins := initPlugins()
	for _, pl := range plugins {
		println(pl.weight, pl.name)
	}
}
