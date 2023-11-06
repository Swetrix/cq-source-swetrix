package client

import "github.com/cloudquery/plugin-sdk/v4/schema"

func ProjectMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	cqClient := meta.(*Client)

	l := make([]schema.ClientMeta, 0)

	for _, projectId := range cqClient.Spec.ProjectIds {
		l = append(l, cqClient.withProjectId(projectId))
	}

	return l
}
