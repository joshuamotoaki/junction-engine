# Junction Engine Backend

The Go-based backend for TigerJunction Engine.

## Getting Started

Clone the repository. Make sure you have Go installed. On MacOS, you can use Homebrew:

```bash
brew install go
```

If you don't have Go installed, you can download it from the [official Go website](https://golang.org/dl/). You will also want to install [Air](https://github.com/air-verse/air) for hot-reloading during development.

```bash
# Install dependencies
go mod tidy

# Run a hot-reloading development server
make dev
```

## Overview

TODO

### API

TODO

### Course Service

TODO

### Degree Service

TODO

### Eval Service

TODO

### Prereq Service

TODO

### Seat Service

TODO

## Folder Structure

```
backend/
├── cmd/                         # API/CLI entry points (minimal main packages)
├── ├── api/
│   └── cli/
├── internal/                    # Internal packages (not meant for public use)
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── api.go
│   │   └── router.go
│   ├── database/
│   │   ├── neo4j/
│   │   └── redis/
│   ├── services/
│   │   ├── course-service/
│   │   ├── degree-service/
│   │   ├── eval-service/
│   │   ├── prereq-service/
│   │   └── seat-service/
│   └── shared/
│       ├── auth/
│       ├── config/
│       ├── helpers/
│       └── models/
├── pkg/                         # Public packages
│   └── oit-client/
├── go.mod
└── go.sum
```
