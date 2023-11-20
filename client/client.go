package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func New(logger zerolog.Logger, pluginSpec Spec) (schema.ClientMeta, error) {
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
