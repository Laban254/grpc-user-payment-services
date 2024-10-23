## 🚀 Setting Up the Protobuf Path

To set up Protobuf for your Go gRPC services, follow the steps below:

### 1. 📦 Install Protobuf and Go Plugins

Begin by installing the Protobuf compiler and the necessary Go plugins:

```bash
# Install the Protobuf compiler
# For macOS
brew install protobuf  
# For Ubuntu
sudo apt install protobuf-compiler  
```
# Install the Go plugins for Protobuf and gRPC
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` 

### 2. 📁 Create a Generated Code Directory

Create a `gen` directory in your project to store the generated code:

```bash
mkdir gen
```
### 3. 🛠️ Generate Go Code

Use the `protoc` command to generate Go code for your `.proto` files. Run the following commands for each file:

```bash
protoc --go_out=gen --go-grpc_out=gen proto/user.proto
protoc --go_out=gen --go-grpc_out=gen proto/payment.proto` 
```
### 4. ✏️ Update Imports

In your `main.go` files for both the User and Payment services, update the import paths to point to the generated code:

-   **User Service** (`user-service/main.go`):
    
    ```go
	pb "path/to/your/project/gen/user"` 
    ```
-   **Payment Service** (`payment-service/main.go`):
    
    ```go
    `pb "path/to/your/project/gen/payment"` 
	```
### 5. ⚙️ Build the Services

After generating the code and updating the imports, build the services using the following commands:

```bash
# Build the Payment Service
go build -o ../tmp/payment-service-executable
```
# Build the User Service
`go build -o ../tmp/user-service-executable` 

### 🔍 Additional Notes

-   Ensure that you have the correct path to your project in the import statements.
-   The generated code will be placed in the `gen` directory. Verify that the files are created successfully.
-   Consider adding the `gen` directory to your `.gitignore` file to avoid version control issues with generated code.