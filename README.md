# Order System with Microservices

## Overview

Welcome to the Order Management System! This project implements a microservices architecture, featuring multiple services: `auth`, `products`, and `gateway`. Each service is responsible for specific functionalities within the system. This README provides detailed instructions on setting up, building, and managing these services.

## Architecture

- **Auth Service**: Manages user authentication and authorization.
- **Products Service**: Handles product listings and inventory management.
- **Gateway Service**: Serves as the entry point for client requests, routing them to the appropriate services.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.16 or later
- **Docker**: Required for running the PostgreSQL database
- **Node.js**: Necessary for the `concurrently` package
- **SQLC**: For generating SQL code
- **Tern**: For managing database migrations

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
   This step will start the PostgreSQL database using Docker and apply migrations:
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

The project includes a Makefile to streamline common tasks:

### Available Targets

- **init-db**: Run Docker Compose to create the database instance.
- **build**: Compile the services and prepare the environment.
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

Contributions are welcome! Please fork the repository and submit a pull request with your changes. For significant modifications, please discuss them via an issue first.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Conclusion

The Order Management System provides a scalable and efficient way to handle orders using microservices. Follow the setup instructions and utilize the Makefile for a smooth development experience. If you have any questions or encounter issues, feel free to open an issue in the repository. Happy coding!
