# Go SSE

Server Sent Events (SSE) using Go-chi.

> [!NOTE]
> This repo was mainly created to benchmark SSE performance in Go, so as to implement it as a `push-only` messaging service (Server-side Push Notifications). The `make benchmark` command will help you test your target loads, to see if this implementation fits your needs.

## Getting Started

### Environment Setup

For `NixOS` users, you can use the `shell.nix` file to setup your environment with all the required dependencies.

```bash
nix-shell
```

Refer to the `shell.nix` file to see the required dependencies.

### Running the Service

- Ensure that an `.env` file is created in the root directory of the project. You can use the `.env.sample` file as a template.

- Run `make` to see the available commands.

## Usage

Visit `http://localhost:6000` to see the SSE stream.

or

```bash
curl http://localhost:6000/stocks/stream/Apple
```
