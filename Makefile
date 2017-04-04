PLUGIN_DIR := plugins
PLUGINS := $(foreach dir,$(PLUGIN_DIR),$(wildcard $(dir)/*.go))

.PHONY:build_plugins
build_plugins: $(PLUGINS)
	for file in $(PLUGINS); do \
		go build -buildmode=plugin -o plugins-build/$$(basename $$file).so $$file; \
	done