# Swetrix Source Plugin Configuration Reference

## Example

The following example connects a single Swetrix project to a Postgres destination. The (top level) source spec section is described in the [Source Spec Reference](https://docs.cloudquery.io/docs/reference/source-spec). 

```yaml
kind: source
spec:
  name: "swetrix"
  path: "swetrix/swetrix"
  tables: ["*"]
  version: "v1.0.0"
  destinations:
    - "postgresql"
  spec:
    project_ids: ["<YOUR_PROECT_ID>"]
    # Get the API key from 'SWETRIX_API_KEY' environment variable
    api_key: "${SWETRIX_API_KEY}"
    time_bucket: "day"
    period: "4w"
```

## Configuration Reference

This is the (nested) spec used by the Swetrix source plugin.

- `project_ids` ([]string, required)

  Swetrix Project IDs to sync from.

- `api_key` (string, required)
  
  Swetrix API key. You can get the API key from the Swetrix portal.

- `time_bucket` (string, required)

  The time bucket to use for the `swetrix_log` and `swetrix_performance` charts. e.g. `day`. See all possible values here: https://docs.swetrix.com/statistics-api#time-buckets.

- `period` (string, optional)
  
  How long in the past to sync data (e.g. `7d`). Either `period` or `from` and `to` must be specified. See all possible values here: https://docs.swetrix.com/statistics-api#periods.

- `from` (string, optional)
  
  A date in `YYYY-MM-DD` format. If `from` is provided, `to` must also be provided.

- `to` (string, optional)
  
  A date in `YYYY-MM-DD` format. If `to` is provided, `from` must also be provided.

- `timezone` (string, optional)

  The timezone to use for the time range. The default is Etc/GMT. You can use any timezone supported by [day.js](https://day.js.org/docs/en/timezone/timezone/) library.

- `concurrrency` (int, optional, default: `1000`)

  Best effort maximum number of Go routines to use. Lower this number to reduce memory usage.