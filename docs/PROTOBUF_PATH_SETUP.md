### Setting Up the Protobuf Path

1. **Install Protobuf and Go Plugins**:
    ```bash
    # Install protobuf compiler
    brew install protobuf  # macOS
    sudo apt install protobuf-compiler  # Ubuntu
    
    # Install Go plugins
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

2. **Create a Generated Code Directory**: Create a `gen` directory in your project.

3. **Generate Go Code**: Run the following for each `.proto` file:
    ```bash
    protoc --go_out=gen --go-grpc_out=gen proto/user.proto
    protoc --go_out=gen --go-grpc_out=gen proto/payment.proto
    ```

4. **Update Imports**: In your `main.go` files, update the import paths:

    - **User Service**:
      ```go
      pb "path/to/your/project/gen/user"
      ```
      
    - **Payment Service**:
      ```go
      pb "path/to/your/project/gen/payment"
      ```

**go build**
`go build -o ../tmp/payment-service-executable`
`go build -o ../tmp/user-service-executable`
