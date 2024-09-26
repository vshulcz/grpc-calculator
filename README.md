# gRPC Calculator

**gRPC Calculator** – just a simple example of a gRPC service written in Go. You'll find a few standard math operations implemented (Add, Subtract, Multiply, Divide), and it's all wrapped up with some middleware for logging and recovery

## Protobuf contract
https://github.com/vshulcz/grpc-protos

## Getting Started

1. **Clone the repo**:
    ```bash
    git clone https://github.com/vshulcz/grpc-calculator.git
    cd grpc-calculator
    ```
2. **Install dependencies**:
    ```bash
    go mod tidy
    ```
3. **Compile the project**:
    ```bash
    go build -o calculator ./cmd/calculator
    ```
4. **Run the gRPC server**:
    ```bash
    ./calculator --config="config/dev.yaml"
    ```
5. **Test it out with grpcurl**:
    Install grpcurl:
    ```bash
    brew install grpcurl
    ```

    To see the available services:
    ```bash
    grpcurl -plaintext localhost:5105 list
    ```

    To call the `Add` method:
    ```bash
    grpcurl -plaintext -d '{"number1": 5, "number2": 3}' localhost:5105 calculator.Calculator.Add
    ```
6. **Run the tests**:
    ```bash
    go test ./... -v
    ```

## Configuration

The configuration is stored in YAML files (e.g., `config/dev.yaml`), and we load it using `cleanenv`.