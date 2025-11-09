# Pack Calculator

Calculate optimal package distribution for any order size. Built with Go (GraphQL) and Svelte 5.

**Live Demo:** http://packager-frontend-demo-694226983221.s3-website.eu-central-1.amazonaws.com

## Quick Start

**Local:**
```bash
# Backend (port 8080)
go run ./cmd/api/main.go

# Frontend (port 5173)
cd web && bun install && bun run dev
```

**Docker:**
```bash
make dc
# or: docker-compose up
```

Visit:
- Local: http://localhost:5173
- Docker: http://localhost:3000
- GraphQL API: `/query`
- Playground: `/playground`

## How It Works

Enter pack sizes (e.g., 250, 500, 1000), then calculate any order. The algorithm finds the optimal combination that:
1. Uses only whole packs
2. Minimizes total items sent
3. Minimizes number of packs

**Example:** Order 251 items with packs [250, 500, 1000, 2000, 5000]
→ Result: 1×500 (not 2×250) because 500 total items < 750 total items

## Structure

```
cmd/
├── api/                # Local HTTP server
└── lambda/             # AWS Lambda handler
internal/
├── service/packer/     # packaging algorithm and types management
├── repository/mem/     # In-memory storage
└── transport/graphql/  # GraphQL API + schema
web/                    # Svelte 5 frontend
├── Dockerfile          # Multi-stage build
└── nginx.conf          # Production config
terraform/              # AWS infrastructure
Dockerfile              # Backend container (distroless)
docker-compose.yml      # Local development
```

## GraphQL API

All operations are documented in the schema. Use the Playground to explore.

**Common queries:**
```graphql
# Get all pack sizes
query { packGetAll(input: {}) { packs { UID Size } } }

# Calculate optimal packs
query { packCalculate(input: { items: 251 }) { calculations { PackSize Quantity } } }

# Create new pack size
mutation { packCreate(input: { size: 750 }) { pack { UID Size } } }
```

## Stack

- **Backend:** Go 1.23+, GraphQL (gqlgen), AWS Lambda
- **Frontend:** Svelte 5, TailwindCSS 4, Bun/TypeScript  
- **Containers:** Docker multi-stage, distroless (backend), nginx (frontend)
- **Infrastructure:** Terraform, S3, Lambda Function URLs
- **Algorithm:** Dynamic programming

## Development

```bash
make test          # Run tests
make bench         # Run benchmarks
make lint          # Lint Go code
make dc            # Run docker-compose

cd web
bun run lint       # Lint frontend
bun run codegen    # Regenerate GraphQL types
```

## License

MIT
