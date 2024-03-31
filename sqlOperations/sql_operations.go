// sqlOperations.go
package sqlOperations

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Employee represents an employee structure
type Employee struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	Postal      string `json:"postal"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}

// InitializeMySQL initializes the MySQL database connection
func InitializeMySQL() error {
	var err error
	db, err = sql.Open("mysql", "root:abc@123@tcp(localhost:3306)/assignment")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to MySQL database")
	return nil
}

// GetEmployeeByID retrieves an employee by ID from the database
func GetEmployeeByID(id int) (Employee, error) {
	var employee Employee
	query := "SELECT * FROM employee WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.CompanyName, &employee.Address, &employee.City, &employee.County, &employee.Postal, &employee.Phone, &employee.Email, &employee.Web)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

// GetAllEmployees retrieves all employees from the database
func GetAllEmployees() ([]Employee, error) {
	var employees []Employee
	query := "SELECT * FROM employee"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.CompanyName, &employee.Address, &employee.City, &employee.County, &employee.Postal, &employee.Phone, &employee.Email, &employee.Web)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

// InsertEmployee inserts a new employee into the database
func InsertEmployee(employee *Employee) error {
	query := "INSERT INTO employee (id,first_name, last_name, company_name, address, city, county, postal, phone, email, web) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, employee.ID, employee.FirstName, employee.LastName, employee.CompanyName, employee.Address, employee.City, employee.County, employee.Postal, employee.Phone, employee.Email, employee.Web)
	if err != nil {
		return err
	}
	return nil
}

// UpdateEmployee updates an existing employee in the database
func UpdateEmployee(employee *Employee) error {
	query := "UPDATE employee SET first_name=?, last_name=?, company_name=?, address=?, city=?, county=?, postal=?, phone=?, email=?, web=? WHERE id=?"
	_, err := db.Exec(query, employee.FirstName, employee.LastName, employee.CompanyName, employee.Address, employee.City, employee.County, employee.Postal, employee.Phone, employee.Email, employee.Web, employee.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteEmployee deletes an employee from the database
func DeleteEmployee(id int) error {
	query := "DELETE FROM employee WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
