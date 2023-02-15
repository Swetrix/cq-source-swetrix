# CloudQuery Swetrix Source Plugin

A Swetrix source plugin for CloudQuery that loads data from Swetrix to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Configuration](docs/configuration.md)
 - [Supported Tables](docs/tables/README.md)

## Authentication

The Swetrix source plugin uses the [Swetrix API](https://docs.swetrix.com/statistics-api) to fetch data. You will need to create an API key in the Swetrix portal. See the [Configuration](docs/configuration.md) section for more information on running the Swetrix plugin.

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```