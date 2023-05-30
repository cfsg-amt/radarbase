package excel

import (
  "fmt"
	"strconv"
  "strings"
  "github.com/xuri/excelize/v2"
)

func Parse(file string, sheet string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	xlFile, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}

	rows, err := xlFile.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	headers := rows[0]

	if len(headers) == 0 {
		return nil, fmt.Errorf("no headers found")
	}

	for _, row := range rows[1:] {

		// We'll try to parse each row, but if we encounter an error, we'll skip it
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from error while parsing row:", r)
				}
			}()

			record := make(map[string]interface{})
			valid := 1

			for i, cell := range row {
				// If cell is empty or "n/a", mark record as invalid and fill the cell as "n/a"
				if cell == "" || strings.ToLower(cell) == "n/a" {
					cell = "n/a"
					valid = 0
				}

				if i < len(headers) {
					value, err := strconv.ParseFloat(cell, 64)
					if err != nil {
						record[headers[i]] = cell
					} else {
						record[headers[i]] = value
					}
				}
			}

			// Use the value of the first column, remove spaces and add a "stockid" column.
			if len(row) > 0 {
				record["stockid"] = strings.ReplaceAll(row[0], " ", "")
			}

			// Append the validity flag to the record
			record["valid"] = valid
			data = append(data, record)
		}()
	}

	return data, nil
}
