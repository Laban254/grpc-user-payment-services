# gRPC User Payment Services 🚀

This project demonstrates a simple gRPC-based microservices architecture with two services: **User Service** and **Payment Service**. The **User Service** manages user profiles, while the **Payment Service** handles payment transactions and interacts with the **User Service** to verify user details.

## 📑 Table of Contents
-   [🗂 Project Structure](#project-structure)
-   [🔍 Services Overview](#services-overview)
-   [📋 Requirements](#requirements)
-   [⚙️ Setup](#setup)
-   [🛠 Usage](#usage)
-   [📜 Protobuf](#protobuf)
-   [💾 Database](#database)
-   [🏃‍♂️ Running the Services](#running-the-services)

##  🔍Services Overview

### 👤User Service

This service handles user-related operations such as:

-   CreateUser
-   GetUser
-   UpdateUser
-   DeleteUser

The gRPC definitions for this service can be found in `proto/user.proto`, and the implementation is in `user-service/user.go`.

### 💸Payment Service

This service handles payment-related operations such as:

-   ProcessPayment
-   RefundTransaction
-   CheckBalance

The gRPC definitions for this service can be found in `proto/payment.proto`, and the implementation is in `payment-service/payment.go`.

## 🗂Project Structure

```bash

`grpc-user-payment-services/
├── database                # Database models and connection setup
│   ├── db.go
│   └── models.go
├── docs                    # Documentation files
│   └── PROTOBUF_PATH_SETUP.md
├── gen                     # Generated protobuf files
│   ├── payment/
│   └── user/
├── payment-service         # Payment service implementation
│   ├── main.go
│   └── payment.go
├── proto                   # Protobuf definition files
│   ├── payment.proto
│   └── user.proto
├── user-service            # User service implementation
│   ├── main.go
│   └── user.go
└── README.md               # Project documentation` 
```
## 📋Requirements

-   [Go 1.16+](https://golang.org/dl/)
-   gRPC
-   Protobuf Compiler
-   [GORM](https://gorm.io/)

## ⚙️Setup

1.  Clone the repository:
    
   ```bash

    git clone https://github.com/your-repo/grpc-user-payment-services.git
    cd grpc-user-payment-services` 
  ```  
2.  Install the required Go modules:
    
   ``` bash
    go mod tidy
   ```
3.  Generate the protobuf files:
    
    ```bash
    `protoc --go_out=. --go-grpc_out=. proto/*.proto` 
    ```
4.  Configure your database connection in `database/db.go`.

## 🏃‍♂️Running the Services

### 🔨Build the Services

To run both services, you first need to **build** them. Navigate to the project root and build the services:

```bash
go build -o tmp/user-service-executable user-service/main.go
go build -o tmp/payment-service-executable payment-service/main.go` 
```
This will create the executables for both services in the `tmp/` directory.

### 🚀Start the Services

You can now start both services simultaneously by running the executables:

```bash
./tmp/user-service-executable &
./tmp/payment-service-executable &` 
```
Alternatively, you can run them directly from their respective directories without building:

```bash
go run user-service/main.go
go run payment-service/main.go` 
```
## 🛠Usage

Once the services are running, they will listen on different ports as specified in their respective `main.go` files.

### 👤Running User Service

To manually run the **User Service**, navigate to the `user-service` directory:

```bash
cd user-service
go run main.go` 
```
### 💸 Running Payment Service

To manually run the **Payment Service**, navigate to the `payment-service` directory:

```bash
cd payment-service
go run main.go` 
```
## 📜Protobuf

The service contracts are defined using protobuf files located in the `proto/` directory.

To update or modify protobuf definitions, edit the `.proto` files and regenerate the `.pb.go` files using the following command:

```bash
`protoc --go_out=. --go-grpc_out=. proto/*.proto` 
```
Refer to [PROTOBUF_PATH_SETUP.md](./docs/PROTOBUF_PATH_SETUP.md)  for more information on setting up the Protobuf path.

## 💾 Database

-   The project uses **GORM** for database interactions.
-   Models are defined in the `database/models.go` file.
-   Database initialization and configuration are handled in `database/db.go`.

Make sure to set the correct database connection details in `db.go`.

## 📝 TODO

- [ ] Refactor responses
- [ ] Implement OAuth2/JWT & TLS
- [ ] Add service discovery
- [ ] Integrate message broker
- [ ] Implement health checks
- [ ] Use Redis for caching
- [ ] Write tests
- [ ] Document API
- [ ] Optimize database
- [ ] Set up monitoring


## Contact

For any queries, feel free to reach out at [labanrotich6544@gmail.com.com](labanrotich6544@gmail.com.com).