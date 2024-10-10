# Order System with Microservices

This project implements an Order System using microservices architecture. The system consists of multiple services that handle different aspects of order management, including order processing, notifications, authentication, and storage.

## Overview

The Order System consists of the following microservices:

- **Order Service**: Manages order creation, retrieval, and updates.
- **Notification Service**: Handles notifications related to order status changes.
- **Auth Service**: Manages user authentication, including login and registration.
- **Storage Service**: Manages data persistence and retrieval for orders.

### Upcoming Services

- **Orders Service**: A dedicated service for managing orders, including new features for order tracking and history.
- **Notification Service**: A new service to send notifications via various channels (e.g., email, SMS) based on order events.
- **Storage Service**: Enhancements to the existing storage solution for better data handling and retrieval.

## Requirements

Before you begin, ensure you have the following installed:

- **Go (version 1.16 or later)**: The programming language used for this project. [Download Go](https://golang.org/dl/)
- **Protocol Buffers (protoc)**: The protocol used for serializing structured data. [Install Protocol Buffers](https://protobuf.dev/docs/proto3/quickstart/go/)
- **Buf**: A tool for managing Protocol Buffers. [Install Buf](https://docs.buf.build/installation)
- **Tern**: A migration tool for managing database migrations. [Install Tern](https://github.com/tern-tools/tern)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/gabehamasaki/orders.git
   cd orders
   ```

2. **Install the required dependencies for each service:**

   Navigate to each service directory and run:

   ```bash
   go mod tidy
   ```

3. **Generate the gRPC code for the services (if applicable):**

   Each service should have its own proto files, and you can generate code using:

   ```bash
   buf generate
   ```

4. **Run Database Migrations:**

   To set up your database, run the following migration command:

   ```bash
   tern migrate --migrations auth/db/migrations --config auth/db/migrations/tern.conf
   ```
## Future Enhancements

- **Implement Orders Service**: Expand the functionality of the current order management service.
- **Add Notification Service**: Introduce a robust notification system for user updates.
- **Improve Storage Service**: Optimize data management and implement advanced storage solutions.
- **Enhance Auth Service**: Add features such as password reset and user role management.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
