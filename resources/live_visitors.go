package resources

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func LiveVisitors() *schema.Table {
	return &schema.Table{
		Name:      "swetrix_live_visitors",
		Resolver:  fetchLiveVisitors,
		Transform: transformers.TransformWithStruct(rest.LiveVisitors{}),
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProjectId,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
