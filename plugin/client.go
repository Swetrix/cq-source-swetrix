package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/rs/zerolog"
	"github.com/swetrix/cq-source-swetrix/client"
	"github.com/swetrix/cq-source-swetrix/resources"
)

type Client struct {
	logger    zerolog.Logger
	config    client.Spec
	options   plugin.NewClientOptions
	scheduler *scheduler.Scheduler
	tables    schema.Tables
	plugin.UnimplementedDestination
}

func (c *Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}

	tt, err := c.tables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	schedulerClient, err := client.New(c.logger, c.config)
	if err != nil {
		return fmt.Errorf("failed to create scheduler client: %w", err)
	}

	return c.scheduler.Sync(ctx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
}

func Configure(_ context.Context, logger zerolog.Logger, specBytes []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	tables, err := getTables()
	if err != nil {
		return nil, fmt.Errorf("failed to get tables: %w", err)
	}

	if opts.NoConnection {
		return &Client{
			options: opts,
			logger:  logger,
			tables:  tables,
		}, nil
	}

	var config client.Spec
	if err := json.Unmarshal(specBytes, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	return &Client{
		options: opts,
		logger:  logger,
		config:  config,
		tables:  tables,
		scheduler: scheduler.NewScheduler(
			scheduler.WithLogger(logger),
			scheduler.WithConcurrency(config.Concurrency),
		),
	}, nil
}

func getTables() (schema.Tables, error) {
	tables := schema.Tables{
		resources.Log(),
		resources.Performance(),
		resources.LiveVisitors(),
	}
	if err := transformers.TransformTables(tables); err != nil {
		return nil, err
	}
	for _, table := range tables {
		schema.AddCqIDs(table)
	}
	return tables, nil
}
