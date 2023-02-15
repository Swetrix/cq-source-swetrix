package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func fetchLiveVisitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cqClient := meta.(*client.Client)

	request := rest.LiveVisitorsRequest{
		ProjectId: cqClient.ProjectId,
	}

	item, err := cqClient.SwetrixClient.GetLiveVisitors(ctx, &request)
	if err != nil {
		return err
	}

	res <- item

	return nil
}
