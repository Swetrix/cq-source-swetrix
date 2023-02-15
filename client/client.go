package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/swetrix/cq-source-swetrix/rest"
)

type Client struct {
	Spec Spec

	SwetrixClient *rest.SwetrixRestClient

	// We multiplex based on projectId - this is the projectId this client is responsible for.
	ProjectId string

	Logger zerolog.Logger
}

func (c *Client) ID() string {
	return "swetrix"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	if err := pluginSpec.Validate(); err != nil {
		return nil, err
	}

	swetrixClient := rest.NewSetrixClient(pluginSpec.ApiKey)

	return &Client{
		Logger:        logger,
		Spec:          pluginSpec,
		SwetrixClient: &swetrixClient,
	}, nil
}

func (c *Client) withProjectId(projectId string) *Client {
	newClient := *c
	newClient.Logger = c.Logger.With().Str("project_id", projectId).Logger()
	newClient.ProjectId = projectId

	return &newClient
}
