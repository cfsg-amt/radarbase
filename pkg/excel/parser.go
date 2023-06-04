package excel
import (
  "fmt"
  "crypto/sha256"
  "encoding/hex"
	"strconv"
  "strings"
  "github.com/xuri/excelize/v2"
)

func RowParse(file string, sheet string) ([]map[string]interface{}, []string, error) {
	var data []map[string]interface{}

	xlFile, err := excelize.OpenFile(file)
	if err != nil {
		return nil, nil, err
	}

	rows, err := xlFile.GetRows(sheet)
	if err != nil {
		return nil, nil, err
	}

	headers := rows[0]

	if len(headers) == 0 {
		return nil, nil, fmt.Errorf("no headers found")
	}

	for _, row := range rows[1:] {
		// We'll try to parse each row, but if we encounter an error, we'll skip it
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from error while parsing row:", r)
				}
			}()

			if len(row) == 0 {
				return
			}

			// Generate SHA256 hash of the first cell of the row
			hash := sha256.New()
			hash.Write([]byte(row[0]))
			hashStr := hex.EncodeToString(hash.Sum(nil))

			record := map[string]interface{}{
				"_id": hashStr, // Add the hash string to the record
			}

			valid := 1

			for i, cell := range row {
				if cell == "" || strings.ToLower(cell) == "n/a" {
					valid = 0
					cell = "0"
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

      record["valid"] = valid

			data = append(data, record)
		}()
	}

	return data, headers, nil
}

func ColParse(file string, sheet string) (map[string][]interface{}, []string, error) {
	xlFile, err := excelize.OpenFile(file)
	if err != nil {
		return nil, nil, err
	}

	rows, err := xlFile.GetRows(sheet)
	if err != nil {
		return nil, nil, err
	}

	headers := rows[0]
	if len(headers) == 0 {
		return nil, nil, fmt.Errorf("no headers found")
	}

	colData := make(map[string][]interface{})
	for _, header := range headers {
		colData[header] = make([]interface{}, 0)
	}

	for _, row := range rows[1:] {
		// We'll try to parse each row, but if we encounter an error, we'll skip it
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from error while parsing row:", r)
				}
			}()

			if len(row) == 0 {
				return
			}

			// Generate SHA256 hash of the first cell of the row
			hash := sha256.New()
			hash.Write([]byte(row[0]))
			hashStr := hex.EncodeToString(hash.Sum(nil))

			valid := 1
			for i, cell := range row {
				// If cell is empty or "n/a", mark record as invalid and fill the cell as "n/a"
				if cell == "" || strings.ToLower(cell) == "n/a" {
					cell = "0"
					valid = 0
				}

				if i < len(headers) {
					value, err := strconv.ParseFloat(cell, 64)
					if err != nil {
						colData[headers[i]] = append(colData[headers[i]], cell)
					} else {
						colData[headers[i]] = append(colData[headers[i]], value)
					}
				}
			}

			// Append the validity flag to the "valid" column
			colData["valid"] = append(colData["valid"], valid)
			
			// Use the value of the first column, remove spaces and add to the "_id" column.
			colData["_id"] = append(colData["_id"], hashStr)
			
		}()
	}

	return colData, headers, nil
}
