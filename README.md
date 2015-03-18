# protobuf 3.0 alpha / grpc presentation

## Installation (OS X)

```
PREFIX=/Users/alindeman/tmp/protobuf

brew install autoconf automake
wget https://github.com/google/protobuf/archive/v3.0.0-alpha-2.tar.gz
./autogen.sh
./configure --prefix=$PREFIX
make
make install
```

## Code Generation

```
mkdir -p calculator

cd protos
$PREFIX/bin/protoc --go_out=plugins=grpc:../calculator calculator.proto
```

## Usage

### Server

```
go run main.go -listen=127.0.0.1:5001
```

### Client

```
go run main.go -connect=127.0.0.1:5001 -operand1=25 -operand2=50
```
