.PHONY:build_plugin

PLUGIN_DIR := plugins
PLUGINS := $(foreach dir,$(PLUGIN_DIR),$(wildcard $(dir)/*.go))

build_plugins: $(PLUGINS)
	for file in $(PLUGINS); do \
		go build -buildmode=plugin -o plugins-build/$$(basename $$file).so $$file; \
	done

