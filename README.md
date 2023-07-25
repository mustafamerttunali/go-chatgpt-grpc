# go-chatgpt-grpc

## Overview

This repository contains a simple, yet efficient, gRPC server-client communication setup implemented in Go. It's intended to help developers establish a reliable, quick-response interaction between a server and a client. The server is uses ```gpt-3.5-turbo``` model using OpenAI API and Go OpenAI Wrapper. 

## Motivation
This project originated from the need to seamlessly use OpenAI's API with custom features on the server-side and effortlessly stream this data to the frontend. It was designed to incorporate aspects like prompt engineering and rule-based systems. I believe this tool, born out of practical necessity, can assist others in similar projects and scenarios.

## Linux Containers
Before you begin, make sure you have Docker installed on your machine. If you don't, you can download it [here](https://docs.docker.com/get-docker/).

Create a local copy of this repository. Add your OpenAI API key to docker-compose.yml file:

```bash
environment:
      - OPENAI_KEY=your-openai-key-is-here
```

Then, run the following command in the root directory of the project:

```bash
docker-compose build
```

This sets compose and builds a local development environment.

After the build is complete (this may take a few minutes), run:

```bash
docker-compose up -d && docker attach go-grpc-client
```


## TODO

- [x] Dockerize the project.

## Contributing

Any contributions are welcome. If you have any questions, feel free to contact me at ```tunalim@mef.edu.tr```.
