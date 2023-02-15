package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"swetrix",
		Version,
		schema.Tables{
			resources.Log(),
			resources.Performance(),
			resources.LiveVisitors(),
		},
		client.New,
	)
}
