package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveProjectId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cqClient := meta.(*Client)

	return r.Set(c.Name, cqClient.ProjectId)
}
