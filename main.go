package main

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/ajstarks/svgo"
)

type RawColumn struct {
	Table  string `json:"table_name"`
	Schema string `json:"table_schema"`

	Name string `json:"column_name"`
	Type string `json:"data_type"`
}

type Column struct {
	Name string
	Type string
}

var (
	DefaultTextStyle = "font-family: Inter; fill: black;"
)

func main() {
	file, err := os.Open("./dump.json")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	database := make(map[string][]Column)

	for scanner.Scan() {
		line := scanner.Bytes()

		var raw RawColumn

		err = json.Unmarshal(line, &raw)

		if err != nil {
			panic(err)
		}

		name := raw.Schema + "." + raw.Table

		table, exists := database[name]

		if exists == false {
			database[name] = []Column{}
		}

		table = append(table, Column{
			Name: raw.Name,
			Type: raw.Type,
		})

		database[name] = table
	}

	for name, table := range database {
		output, err := os.Create("./output/" + name + ".svg")

		if err != nil {
			panic(err)
		}

		canvas := svg.New(output)
		defer canvas.End()

		rowHeight := 40
		width := 300
		height := rowHeight + ((rowHeight + rowHeight) * len(table))

		canvas.Start(width, height)

		canvas.Rect(1, 1, width-2, height-2, "fill: white; stroke: #8A8A8A; stroke-width: 2px; rx: 4px;")
		canvas.Text(width/2, 25, name, "text-anchor: middle; font-size: 18px; font-weight: bolder;"+DefaultTextStyle)

		for index, column := range table {
			offset := ((rowHeight + rowHeight) * index) + 40

			canvas.Group()
			canvas.Rect(1, offset, width-2, 40, "fill: #D9D9D9; stroke: #8A8A8A; stroke-width: 2px;")
			canvas.Text(10, offset+25, column.Name, "font-weight: bold; font-size: 16px;"+DefaultTextStyle)
			canvas.Gend()

			canvas.Text(10, offset+rowHeight+25, column.Type, "font-size: 14px;"+DefaultTextStyle)
		}
	}
}
