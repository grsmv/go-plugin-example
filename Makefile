.PHONY:build_plugin

build_plugin: plugin.go
	go build -buildmode=plugin $<
