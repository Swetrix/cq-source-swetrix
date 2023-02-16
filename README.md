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

### Release a new version

1. Follow [this link](https://github.com/Swetrix/cq-source-swetrix/releases/new) to draft a new release.
2. Click `Choose a tag` and enter the new version number:
   ![image](https://user-images.githubusercontent.com/26760571/219360662-0ad1f83d-84c9-47c8-afb9-fe774ce03dcc.png)
3. Click `Create new tag: <version> on publish` assuming it's a new tag.
4. Click `Generate release notes` to automatically generate release notes.
5. Click `Publish release` to publish the release.

> Once the tag is pushed, a new GitHub Actions workflow will be triggered to build and upload the release binaries to the release
