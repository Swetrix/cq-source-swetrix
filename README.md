# CloudQuery Swetrix Source Plugin

A Swetrix source plugin for CloudQuery that loads data from Swetrix to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

- [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
- [Configuration](docs/overview.md)
- [Supported Tables](docs/tables/README.md)

## Authentication

The Swetrix source plugin uses the [Swetrix API](https://docs.swetrix.com/statistics-api) to fetch data. You will need to create an API key in the Swetrix portal. See the [Configuration](docs/overview.md) section for more information on running the Swetrix plugin.

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

### Compile
```bash
make build
```

### Release
```bash
VERSION=1.0.0 make dist
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub  

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
