package resources

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func Log() *schema.Table {
	return &schema.Table{
		Name:      "swetrix_log",
		Resolver:  fetchLog,
		Transform: transformers.TransformWithStruct(rest.Log{}),
		Multiplex: client.ProjectMultiplex,
	}
}
