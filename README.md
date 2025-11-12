![Logo](https://fuchsia-defiant-cicada-56.mypinata.cloud/ipfs/bafkreibbvgqo7qyy5f7g4nzbiaaiecxnq3nn7uko7sgljhqfg3ti3dw54y)

# cobra-gin-starter

A modern Go starter template that unifies a **Cobra CLI**, **Gin-powered REST API**, and **Templ + HTMX + TailwindCSS** (GoTTH) web frontend — all served together with a single command.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/6missedcalls/cobra-gin-starter)
![Go Report Card](https://goreportcard.com/badge/github.com/6missedcalls/cobra-gin-starter)(https://goreportcard.com/report/github.com/6missedcalls/cobra-gin-starter)
![GitHub License](https://img.shields.io/github/license/6missedcalls/cobra-gin-starter)

## Features

- **Cobra CLI** – easy command creation and management.
- **Gin** – high-performance HTTP web framework.
- **Swagger (gin-swagger)** – live, auto-generated API documentation.
- **Templ** – type-safe HTML templating in Go.
- **HTMX** – minimal-JS interactivity using server responses.
- **TailwindCSS v4** – modern utility-first styling.
- **Unified Serve Command** – run API + web UI with `cobra-gin-starter serve`.

## Run Locally

Clone the project

```bash
  git clone https://github.com/6missedcalls/cobra-gin-starter.git
```

Go to the project directory

```bash
  cd cobra-gin-starter
```

Install dependencies

```bash
  go mod tidy
  bun install
```

Generate Templ Components & Build TailwindCSS

```bash
  templ generate
  bun run build:css
```

Start the server (development)

```bash
  go run main.go serve
```

Start the server (production)

```bash
  ./cobra-gin-starter serve
```

## Tech Stack

**Client:** HTMX, Templ, TailwindCSS v4

**Server:** Go, Gin, gin-swagger

**CLI:** Cobra

## Contributing

Contributions are always welcome!

See [docs/contributing.md](docs/contributing.md) for ways to get started,

## License

[MIT](https://choosealicense.com/licenses/mit/)
