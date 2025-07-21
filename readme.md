# HyMatrix Container

HyMatrix Container is a Docker-based runtime environment for executing and managing AO smart contracts.

## Features

- Lua 5.3 runtime support
- Ollama runtime support
- HTTP API interface
- Contract deployment and execution support

## Quick Start from docker

### Build Image

```bash
# golua runtime
./docker_build.sh golua <GITHUB_TOKEN>

# ollama runtime 
./docker_build.sh ollama <GITHUB_TOKEN>
```

### Run Container
```bash
./docker_run.sh
```
## local run
```bash
go run -tags=lua53 main.go

# or
go build -tags=lua53 -o main
./main

# test
go test -tags=lua53 -v

```

## API Endpoints
### Health Check
```bash
POST /vmm/health
```

### Spawn ao
```bash
POST /vmm/spawn
```
Body:
```json
{
    "pid": "0x8454",
    "owner": "0x123",
    "cuAddr": "0x84534",
    "data": "",
    "tags": []
}
```

### apply
```bash
POST /vmm/apply
```
Body:
```json
{
    "action": "Info",
    "nonce": 1,
    "params": {
        "Action": "Info",
        "From": "0x123",
        "Module": "0x84534"
    }
}
```
