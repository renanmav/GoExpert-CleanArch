# GoExpert-CleanArch

This project is a simple example of a clean architecture in Go. It is a basic Create and Read application that allows you to manage orders respecting the business rules.

## About Clean Architecture

Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. For more information, check out [Uncle Bob's Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

<div align="center">
  <img src="./assets/CleanArchitecture.jpg" alt="Clean Architecture" width="500">
</div>

## Business Rules / Requirements

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
1. When an order is created, a message is sent to the RabbitMQ exchange
1. All orders can be read
1. An order can be read by ID

## Running the application

1. Clone the repository
1. Run `make docker-up` to start the infrastructure
1. Run `make run` to start the application

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

## Accessing the web server API

Once the application is running, you can access the web server API at `http://localhost:8000` (check `.env` file).

Example request to create an order:

```bash
curl -X POST http://localhost:8000/order \
-H "Content-Type: application/json" \
-d '{
    "price": 100.50,
    "tax": 10.05
  }'
```

Example request to read all orders:

```bash
curl -X GET http://localhost:8000/orders
```

Example request to read an order by ID:

```bash
curl -X GET http://localhost:8000/order?id=123
```

## Generating gRPC code

Make sure you have the protocol buffer compiler and the Go plugin installed. You can install the protocol buffer compiler [here](https://grpc.io/docs/protoc-installation/).

After installing the protocol buffer compiler, you can generate the gRPC code by running the following command:

```bash
make proto
```

## Accessing the gRPC server

Once the application is running, you can access the gRPC server at `http://localhost:50051` (check `.env` file).

You can use a GRPC client to send requests to the server, like [Evans](https://github.com/ktr0731/evans).

Example request to create an order:

```bash
evans -r repl

# From the repl, run the following commands
package proto
service OrderService
call CreateOrder
```

Example request to read all orders:

```bash
evans -r repl

# From the repl, run the following commands
package proto
service OrderService
call RealAllOrders
```

Example request to read an order by ID:

```bash
evans -r repl

# From the repl, run the following commands
package proto
service OrderService
call ReadOrderById
```

## Generating GraphQL code

Make sure you have gqlgen installed. You can install gqlgen [here](https://github.com/99designs/gqlgen).

After installing gqlgen, you can generate the GraphQL code by running the following command:

```bash
make graphql
```

## Accessing the GraphQL server

Once the application is running, you can access the GraphQL server at `http://localhost:8080` (check `.env` file).

You can use a GraphQL client to send requests to the server, like [GraphiQL](https://github.com/graphql/graphiql).

Example request to create an order:

```graphql
mutation CreateOrder {
  createOrder(input:{
    price: 100.2,
    tax: 0.1
  }) {
    id
    price
    tax
    finalPrice
  }
}
```

Example request to read all orders:

```graphql
query ReadAllOrders {
  readAllOrders {
    id
    price
    tax
    finalPrice
  }
}
```

Example request to read an order by ID:

```graphql
query ReadOrderById {
  readOrderById(id: "123") {
    id
    price
    tax
    finalPrice
  }
}
```

## RabbitMQ

Once an order is created, a message is sent to the RabbitMQ exchange. You can access the RabbitMQ management dashboard at `http://localhost:15672`. Username is `user` and password is `password`. The queue is automatically created by the application.

To check messages in the queue:

1. Go to the "Queues" tab
1. Select the "orders" queue
1. Click on "Get messages"
1. You should see the message you created

## Using Wire to inject dependencies

This project uses [Wire](https://github.com/google/wire) to inject dependencies. To generate the dependency injection container, run the following command:

```bash
make wire
```

This will generate a `wire_gen.go` file with the dependency injection container.