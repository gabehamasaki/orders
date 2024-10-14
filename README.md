# Order Management System with Microservices

## Overview

Welcome to the Order Management System! This project employs a microservices architecture with dedicated services for `auth`, `products`, and `gateway`. Each service handles specific functionalities, ensuring a modular and scalable system. This README offers comprehensive instructions for setting up, building, and managing these services.

## Architecture

- **Auth Service**: Responsible for user authentication and authorization.
- **Products Service**: Manages product listings and inventory.
- **Gateway Service**: Acts as the entry point for client requests, routing them to the appropriate services.

## Getting Started

### Prerequisites

Ensure you have the following installed before you begin:

- **Go**: Version 1.16 or later
- **Docker**: Required for running the PostgreSQL database
- **Node.js**: Needed for the `concurrently` package
- **SQLC**: For generating SQL code
- **Tern**: For managing database migrations
- **Buf**: For generating GRPC boilerplate code

### Installation Steps

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/gabehamasaki/orders.git
   cd orders
   ```

2. **Build the Services**:
   ```bash
   make build
   ```

3. **Initialize the Database**:
   This command starts the PostgreSQL database using Docker and applies migrations:
   ```bash
   make init-db
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

### Health Check

- **Ping**
  - **Endpoint**: `/ping`
  - **Method**: `GET`
  - **Description**: A simple endpoint to check the service status.

## Makefile Targets

The project includes a Makefile to streamline common tasks. Here are the available targets:

- **buf-generate**: Regenerate GRPC boilerplate.
- **init-db**: Start Docker Compose to create the database instance.
- **build**: Compile the services and set up the environment.
- **run**: Start the services concurrently.
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

Contributions are welcome! Please fork the repository and submit a pull request with your changes. For significant modifications, discuss them via an issue first.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Conclusion

The Order Management System offers a scalable and efficient solution for managing orders through microservices. Follow the setup instructions and utilize the Makefile for a smooth development experience. If you have any questions or encounter issues, please open an issue in the repository. Happy coding!
