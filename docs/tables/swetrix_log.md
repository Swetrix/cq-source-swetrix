# Table: swetrix_log

This table shows data for Swetrix Log.

The primary key for this table is **project_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|params|`json`|
|chart|`json`|
|avg_sdur|`float64`|
|customs|`json`|
|applied_filters|`json`|