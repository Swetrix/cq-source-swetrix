package resources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/rest"
)

func fetchLog(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cqClient := meta.(*client.Client)

	getLogRequest := rest.GetLogRequest{
		ProjectId:  cqClient.ProjectId,
		TimeBucket: cqClient.Spec.TimeBucket,
		Period:     cqClient.Spec.Period,
		From:       cqClient.Spec.From,
		To:         cqClient.Spec.To,
		Timezone:   cqClient.Spec.Timezone,
	}

	log, err := cqClient.SwetrixClient.GetLog(ctx, &getLogRequest)
	if err != nil {
		return err
	}

	res <- log

	return nil
}
