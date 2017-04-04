# Chain of responsibility based on Go plugins

This repository and particularly this document describes Chain of responsibility pattern implementation, based on Go 1.8+ Plugin capabilities.
  
#### Restrictions

At this point Go 1.8+ plugins working only in Linux environment.

#### High level overview

![](https://s3-eu-west-1.amazonaws.com/serhiiherasymov/Work/chain+of+responsibility+-+Page+1-2.png)

Current implementation includes Pipeline Manager (_main.go_), Plugin update service (_plugins.go_) and Plugin Repository (_plugins.go_).
Plugin updating service look up for changes in `PluginsFolder` and launches plugin updating process (re-reading list of plugins) when changes available. All long-term operations with Plugin repository are wrapped with mutexes.

#### Pipeline

Pipeline is building dynamically, basing on available plugins. Order of execution is based on plugin meta-information. 

#### Plugin

Each plugin has meta-information and handler. Here's basic example of plugin:

```go
package main

import (
	"context"
	"go-plugin-example/models"
	"time"
)

func Weight() int {
	return 20
}

func Handler(ctx context.Context, data models.Data) (models.Data, error) {
	time.Sleep(20 * time.Second)
	return models.Data{A: data.A + 10}, nil
}
```

Type signatures for `Weight` and `Handler` functions should be constant across all plugins.
 Notice the `Weight` function - this is meta-information, which shows place of plugin execution in Pipeline query.