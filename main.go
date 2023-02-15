package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/swetrix/cq-source-swetrix/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
