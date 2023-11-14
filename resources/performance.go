package resources

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func Performance() *schema.Table {
	return &schema.Table{
		Name:      "swetrix_performance",
		Resolver:  fetchPerformance,
		Transform: transformers.TransformWithStruct(rest.Performance{}),
		Multiplex: client.ProjectMultiplex,
	}
}
