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

  var headers []string
  consecutiveEmptyRows := 0
  start := 0
  for _, row := range rows {
    if len(row) < 4 {
      consecutiveEmptyRows++
      if consecutiveEmptyRows >= 4 {
        break
      }
      continue
    } else {
      headers = row
      start = consecutiveEmptyRows + 1
      consecutiveEmptyRows = 0
      break
    }
  }

	if len(headers) == 0 {
		return nil, nil, fmt.Errorf("no headers found")
	}

	for _, row := range rows[start:] {
    if len(row) < 4 {
      consecutiveEmptyRows++
      if consecutiveEmptyRows >= 4 {
        break
      }
      continue
    }

    consecutiveEmptyRows = 0 // Reset the count

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

      name := strings.ReplaceAll(row[0], " ", "")

			// Generate SHA256 hash of the first cell of the row
			hash := sha256.New()
			hash.Write([]byte(name))
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
            // Format float with 2 decimal places
            formattedValue, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", value), 64)
            record[headers[i]] = formattedValue
					}
				}
			}

      record["valid"] = valid

      record["name"] = name

			data = append(data, record)
		}()
	}

  headers = append([]string{"name"}, headers...)
  headers = append([]string{"valid"}, headers...)

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

  var headers []string
  consecutiveEmptyRows := 0
  start := 0

  for _, row := range rows {
    if len(row) < 4 {
      consecutiveEmptyRows++
      if consecutiveEmptyRows >= 4 {
        break
      }
      continue
    } else {
      start = consecutiveEmptyRows + 1
      headers = row
      consecutiveEmptyRows = 0
      break
    }
  }

	if len(headers) == 0 {
		return nil, nil, fmt.Errorf("no headers found")
	}

	colData := make(map[string][]interface{})
	for _, header := range headers {
		colData[header] = make([]interface{}, 0)
	}

	for _, row := range rows[start:] {
    if len(row) < 4 {
      consecutiveEmptyRows++
      if consecutiveEmptyRows >= 4 {
        break
      }
      continue
    }

    consecutiveEmptyRows = 0 // Reset the count

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

      name := strings.ReplaceAll(row[0], " ", "")

			// Generate SHA256 hash of the first cell of the row
			hash := sha256.New()
			hash.Write([]byte(name))
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
            // Format float with 2 decimal places
            formattedValue, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", value), 64)
						colData[headers[i]] = append(colData[headers[i]], formattedValue)
					}
				}
			}

			// Append the validity flag to the "valid" column
			colData["valid"] = append(colData["valid"], valid)

      colData["name"] = append(colData["name"], name)
			
			// Use the value of the first column, remove spaces and add to the "_id" column.
			colData["_id"] = append(colData["_id"], hashStr)
			
		}()
	}

  headers = append([]string{"name"}, headers...)
  headers = append([]string{"valid"}, headers...)

	return colData, headers, nil
}
