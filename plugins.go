package main

import (
	"context"
	"github.com/howeyc/fsnotify"
	"go-plugin-example/models"
	"io/ioutil"
	"log"
	"path/filepath"
	"plugin"
	"sort"
	"strings"
	"sync"
)

type internalPlugin struct {
	name   string
	weight int
}

// callHandler calls handler function, stored in given plugin
func (pl internalPlugin) callHandler(ctx context.Context, data models.Data) models.Data {
	fn := getFunction(pl.name, "Handler")
	updatedData, err := fn.(func(context.Context, models.Data) (models.Data, error))(ctx, data)
	if err != nil {
		return data
	}
	return updatedData
}

type pluginRepository []internalPlugin

var pluginsLock = &sync.Mutex{}

// initPlugins returns sorted slice with detected plugins
func initPlugins() (pls pluginRepository) {

	pluginsLock.Lock()

	log.Println("(re)building plugin list")

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

	defer pluginsLock.Unlock()

	return
}

// processPipeline
func (pls pluginRepository) processPipeline(ctx context.Context, data models.Data) models.Data {

	pluginsLock.Lock()

	var updatedData = data
	for _, pl := range pls {
		updatedData = pl.callHandler(ctx, updatedData)
	}

	defer pluginsLock.Unlock()

	return updatedData
}

// pluginWeight extracts weight info from plugin
func getPluginWeight(pluginName string) int {
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

// initPluginUpdater
func initPluginUpdater() {
	log.Println("launching plugin watcher")

	watcher, _ := fsnotify.NewWatcher()
	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("some updates detected")
				if strings.HasSuffix(ev.Name, PluginExtension) {
					log.Println("plugins should be updated")
					initPlugins()
				}
				//case err := <- watcher.Error:
				//	log.Println("error:", err)
			}
		}
	}()

	watcher.Watch(PluginsFolder)
	<-done // hanging

	watcher.Close()
}
