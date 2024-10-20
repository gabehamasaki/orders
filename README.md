# Order Management System with Microservices

## Overview

Welcome to the **Order Management System**! This project utilizes a microservices architecture, featuring dedicated services for **Authentication**, **Product Management**, and a **Gateway**. Each service is designed to handle specific functionalities, ensuring a modular, scalable, and maintainable system. This README provides detailed instructions for setting up, building, and managing these services effectively.

## Architecture

The system is composed of the following services:

- **Auth Service**: Handles user authentication and authorization processes.
- **Products Service**: Manages product listings, inventory, and related operations.
- **Gateway Service**: Serves as the entry point for client requests, routing them to the appropriate services seamlessly.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following software installed:

- **Go**: Version 1.21 or later
- **Docker**: Required for running the services
- **SQLC**: For generating SQL code
- **Tern**: For managing database migrations
- **Buf**: For generating GRPC boilerplate code

### Installation Steps

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/gabehamasaki/orders.git
   cd orders
   ```

2. **Build the Services and Initialize the Database**:
   ```bash
   make run
   ```

## API Endpoints

### Auth Service

- **Register User**
  - **Endpoint**: `/auth/register`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "name": "string",
      "password": "string",
      "email": "string"
    }
    ```
  - **Description**: Creates a new user account.

- **Login User**
  - **Endpoint**: `/auth/login`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "name": "string",
      "password": "string"
    }
    ```
  - **Description**: Authenticates a user and returns a token.

### Products Service

- **Create Product**
  - **Endpoint**: `/products`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "name": "string",
      "description": "string",
      "price": 0.0,
      "image_url": "string"
    }
    ```
  - **Description**: Creates a new product.

- **List Products**
  - **Endpoint**: `/products`
  - **Method**: `GET`
  - **Query Parameters**:
    - `page`: Page number for pagination (default: 1)
    - `per_page`: Number of products per page (default: 15)
  - **Description**: Retrieves a list of products.

- **Get Product**
  - **Endpoint**: `/products/:id`
  - **Method**: `GET`
  - **Description**: Retrieves a specific product by its ID.

### Health Check

- **Ping**
  - **Endpoint**: `/ping`
  - **Method**: `GET`
  - **Description**: A simple endpoint to check the service status.

## Makefile Targets

The project includes a Makefile to streamline common tasks. Here are the available targets:

- **buf-generate**: Regenerate GRPC boilerplate.
- **init-db**: Start Docker Compose to create the database instance.
- **build**: Compile the services.
- **run**: Start the services using Docker.
- **sqlc-gen**: Generate SQL code based on schema definitions.
- **migrate**: Apply database migrations.
- **drop-db**: Drop the existing database.
- **create-db**: Create a new database.
- **fresh-db**: Drop, create, and migrate the database.

### Example Usage

- **To build the services**:
  ```bash
  make build
  ```

- **To run the services**:
  ```bash
  make run
  ```

- **To generate SQL code**:
  ```bash
  make sqlc-gen
  ```

- **To manage the database**:
  - Create Database: `make create-db`
  - Drop Database: `make drop-db`
  - Migrate Database: `make migrate`
  - Fresh Database: `make fresh-db`

## Contributing

We welcome contributions! Please fork the repository and submit a pull request with your changes. For significant modifications, we encourage discussing them via an issue first.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Conclusion

The Order Management System provides a scalable and efficient solution for managing orders through microservices. Follow the setup instructions and utilize the Makefile for a smooth development experience. If you have any questions or encounter issues, please open an issue in the repository. Happy coding!
