# Assignment
# Employee Management System

This project is a simple Employee Management System implemented in GoLang.
It allows users to import data from an Excel file, store it in MySQL, and cache it in Redis. 
It also provides CRUD (Create, Read, Update, Delete) operations to manage employee records.

## Features

- Import data from an Excel file
- Store data in MySQL database
- Cache data in Redis
- CRUD operations for managing employee records
- Separated router handler functions for modularity

## Installation

1. Clone the repository:
git clone https://github.com/akshatjain14/assignment.git
2. Install dependencies:
go mod tidy
3. Run the application:
go run main.go


The application will start running on port 8080 by default.

## API Endpoints

- `GET /employees`: Get all employees
- `GET /employee/:id`: Get an employee by ID
- `PUT /employee/:id`: Update an employee by ID
- `DELETE /employee/:id`: Delete an employee by ID
- `GET /redis/employees`: Get all employees from Redis cache
- `GET /redis/employee/:id`: Get an employee by ID from Redis cache

## Dependencies

- Gin: Web framework for building APIs in GoLang
- Go-Redis: Redis client for GoLang
- Go-MySQL-Driver: MySQL driver for GoLang

## License

This project is licensed under the [MIT License](LICENSE).





