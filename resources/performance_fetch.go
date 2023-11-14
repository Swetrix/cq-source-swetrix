package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func fetchPerformance(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cqClient := meta.(*client.Client)

	request := rest.GetPerformanceRequest{
		ProjectId:  cqClient.ProjectId,
		TimeBucket: cqClient.Spec.TimeBucket,
		Period:     cqClient.Spec.Period,
		From:       cqClient.Spec.From,
		To:         cqClient.Spec.To,
		Timezone:   cqClient.Spec.Timezone,
	}

	item, err := cqClient.SwetrixClient.GetPerformance(ctx, &request)
	if err != nil {
		return err
	}

	res <- item

	return nil
}
