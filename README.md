# go-chatgpt-grpc
## Overview
This repository contains a simple, yet efficient, gRPC server-client communication setup implemented in Go. It's intended to help developers establish a reliable, quick-response interaction between a server and a client. The server is uses ```gpt-3.5-turbo``` model using OpenAI API and Go OpenAI Wrapper.
## Installation
1. Install Go from https://golang.org/dl/
2. Install dependencies:
```bash
cd go-chatgpt-grpc/
go mod tidy
```
3. Open two terminals and run the following commands in each:
```bash
cd server/
go run main.go
```
```bash
cd client/
go run main.go
```

## TODO
- [ ] Next.js implementation for the client.
- [ ] Dockerize the project.

## Contributing
Any contributions are welcome. If you have any questions, feel free to contact me at ```tunalim@mef.edu.tr```.