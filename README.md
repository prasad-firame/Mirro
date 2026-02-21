# Mirro
<p align="center">
<img src="https://visitor-badge.laobi.icu/badge?page_id=prasad-firame.mirro" alt="Visitors" /> 
<img src="https://img.shields.io/github/actions/workflow/status/prasad-firame/mirro/go.yml" alt="Build" /> 
<img src="https://img.shields.io/codecov/c/github/prasad-firame/mirro" alt="Coverage" /> 
<a href="https://goreportcard.com/report/github.com/neooxdev/mirro">
<img src="https://goreportcard.com/badge/github.com/neooxdev/mirro" alt="Go Report Card" />
</a>
<img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License" />
</p>

<p align="center">
<img src="https://img.shields.io/github/v/release/prasad-firame/mirro" alt="Release" /> 
<img src="https://img.shields.io/badge/go-1.20-blue" alt="Go Version" /> 
<img src="https://img.shields.io/github/stars/prasad-firame/mirro?style=social" alt="Stars" /> 
<img src="https://img.shields.io/github/forks/prasad-firame/mirro?style=social" alt="Forks" />
</p>

**Mirro** is a lightweight, high-performance, open-source reverse proxy built in Go.  
Efficiently route HTTP, TCP, and WebSocket traffic, handle load balancing, and manage requests seamlessly in modern web infrastructure.

---

## Features

- ✅ Lightweight and high-performance  
- ✅ Open-source and free to use  
- ✅ Reverse proxy for HTTP, TCP, and WebSocket traffic  
- ✅ Load balancing and request routing with minimal configuration  
- ✅ Written in Go for concurrency and speed  
- ✅ Simple CLI for configuration and deployment  
- ✅ Configurable via YAML, JSON, or CLI flags  

---

## Folder Structure
```bash
mirro/
├── README.md            # Project documentation
├── LICENSE              # MIT License
├── go.mod               # Go module file
├── go.sum               # Go dependencies
├── cmd/                 # CLI entrypoints
│   └── mirro/           # Main CLI for running Mirro proxy
├── internal/            # Internal packages (not for external use)
│   ├── config/          # Config parsing and defaults
│   ├── logger/          # Logging utilities
│   ├── proxy/           # Core proxy engine
│   ├── loadbalancer/    # Load balancing algorithms
│   └── utils/           # General helpers
├── pkg/                 # Optional reusable packages
├── examples/            # Sample configuration and usage
│   └── simple/          # Minimal proxy example
├── deployments/         # Docker, Kubernetes, systemd configs
├── scripts/             # Dev scripts and setup helpers
├── tests/               # Unit and integration tests
│   ├── proxy_test.go
│   └── loadbalancer_test.go
└── docs/                # Extended documentation
    └── architecture.md
```
---

## Installation

### Prerequisites

- Go >= 1.20: [https://golang.org/dl/](https://golang.org/dl/)  
- Git: [https://git-scm.com/](https://git-scm.com/)  

---

### macOS / Linux
``` bash
## 1. Clone the repository
git clone https://github.com/neooxdev/mirro.git
cd mirro

## 2. Build the CLI
go build ./cmd/mirro

## 3. Run Mirro locally
./mirro serve

## 4. Test with sample configuration
curl http://localhost:8080
```
---

### Windows (PowerShell)
```bash
### 1. Clone the repository
git clone https://github.com/neooxdev/mirro.git
cd mirro

### 2. Build the CLI
go build ./cmd/mirro

### 3. Run Mirro locally
.\mirro.exe serve

### 4. Test with sample configuration
Invoke-WebRequest -Uri http://localhost:8080
```
---

## Quickstart Example

Create `config.yaml`:

# config.yaml
```bash
server:
  address: ":8080"

routes:
  - path: "/"
    backend: "http://localhost:9000"
    loadBalance: "roundrobin"

  - path: "/api"
    backend: "http://localhost:9001"
    loadBalance: "leastconn"
```

Run Mirro with config:
```bash
./mirro serve --config config.yaml
```

Test routing:
``` bash
curl http://localhost:8080/
curl http://localhost:8080/api
```
---

## Usage

- Configure routes in `config.yaml` or via CLI flags  
- Supports multiple backends, load balancing strategies, and protocol routing  
- Integrate with Docker or Kubernetes for dynamic service discovery  
- Monitor logs and metrics through CLI or optional HTTP dashboard  

---

## Contributing

We welcome contributions!  

1. Fork the repository  
2. Create a feature branch: `git checkout -b feature/your-feature`  
3. Commit your changes: `git commit -am 'Add new feature'`  
4. Push to the branch: `git push origin feature/your-feature`  
5. Open a Pull Request  

---

## License

MIT License

---



