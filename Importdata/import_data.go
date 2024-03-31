// import_data.go
package importdata

import (
	"assignment/redisOperations"
	"assignment/sqlOperations"
	"fmt"
	"log"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// ImportDataFromExcel reads data from an Excel file and stores it in MySQL and Redis
func ImportDataFromExcel(filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	xlsx, err := excelize.OpenFile(absPath)
	if err != nil {
		return err
	}

	rows, _ := xlsx.GetRows("uk-500")
	for id, row := range rows {
		if id == 0 {
			continue
		}
		employee := sqlOperations.Employee{
			ID:          id,
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			County:      row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		}

		// Insert data into MySQL
		err := sqlOperations.InsertEmployee(&employee)
		if err != nil {
			log.Printf("Error inserting employee: %v", err)
			continue
		}

		// Store data in Redis cache
		key := fmt.Sprintf("employee:%d", id)
		err = redisOperations.SetEmployee(key, employee)
		if err != nil {
			log.Printf("Error setting employee in Redis: %v", err)
			continue
		}
	}
	return nil
}
