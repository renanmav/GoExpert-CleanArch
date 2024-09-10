# GoExpert-CleanArch

This project is a simple example of a clean architecture in Go. It is a basic Create and Read application that allows you to manage orders respecting the business rules.

## About Clean Architecture

Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. For more information, check out [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

<div align="center">
  <img src="./assets/CleanArchitecture.jpg" alt="Clean Architecture" width="500">
</div>

## Business Rules

1. The order can be created
    1. The order must contain an ID of type string (UUID v4)
    1. The order must contain a price of type float64
    1. The order must contain a tax of type float64
    1. The order can contain a final price of type float64
1. The order can calculate its final price
    1. FinalPrice = Price + Tax
1. The order can be validated
    1. The order cannot have an empty ID
    1. The order cannot have a price equal or less than 0
    1. The order cannot have a tax equal or less than 0
1. The order can be read by ID

## Running the application

1. Clone the repository
1. Run `docker-compose up -d` to start the infrastructure
1. Run `go run cmd/main.go` to start the application

## Checking the data

1. Run `docker exec -it mysql bash` to access the MySQL container

```bash
mysql -u user -p orders # Enter password: password
```

1. Check the data

```bash
DESCRIBE orders;
SELECT * FROM orders;
```
