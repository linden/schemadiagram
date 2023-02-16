# Schema Diagram
## Getting Started
1. download `schemadiagram`
```sh
go install github.com/linden/schemadiagram@latest
```
2. dump your schema to `JSON`.
```sh
psql postgres://postgres:postgres@localhost:5432/postgres \
	--command="SELECT row_to_json(result) FROM (SELECT table_name, table_schema, column_name, data_type FROM information_schema.columns) AS result;" \
	--output="dump.json" \
	--quiet \
	--no-align \
	--tuples-only
```
3. create your `SVG`s
```sh
schemadiagram
```
4. drag and drop them to figma.
