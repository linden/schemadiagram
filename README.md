```sh
psql postgres://postgres:postgres@localhost:5432/postgres \
	--command="SELECT row_to_json(result) FROM (SELECT table_name, table_schema, column_name, data_type FROM information_schema.columns) AS result;" \
	--output="dump.json" \
	--quiet \
	--no-align \
	--tuples-only
```