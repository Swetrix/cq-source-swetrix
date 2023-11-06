package resources

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func LiveVisitors() *schema.Table {
	return &schema.Table{
		Name:      "swetrix_live_visitors",
		Resolver:  fetchLiveVisitors,
		Transform: transformers.TransformWithStruct(rest.LiveVisitors{}),
		Multiplex: client.ProjectMultiplex,
	}
}
