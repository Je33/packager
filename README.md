# Pack Calculator - Quick Start Guide

## Overview

A full-stack application for calculating optimal package distribution for orders, built with Go backend (GraphQL) and Svelte 5 frontend.

## Prerequisites

- **Go 1.21+** for backend
- **Node.js 18+ or Bun** for frontend
- Make sure ports 8080 (backend) and 5173 (frontend) are available

## Starting the Application

### 1. Start the Backend

```bash
# From project root
go run ./cmd/main.go
```

The GraphQL API will be available at:
- API endpoint: `http://localhost:8080/query`
- GraphQL Playground: `http://localhost:8080/playground`

### 2. Start the Frontend

```bash
# Navigate to web directory
cd web

# Install dependencies (first time only)
bun install
# or: npm install

# Start dev server
bun run dev
# or: npm run dev
```

The frontend will be available at: `http://localhost:5173`

## Using the Application

### Step 1: Configure Pack Sizes

1. Enter pack sizes in the input fields (e.g., 250, 500, 1000, 2000, 5000)
2. Click "Add Pack Size" to add more inputs if needed
3. Click "Submit pack sizes change" to save

### Step 2: Calculate Order

1. Enter the number of items to order (e.g., 263)
2. Click "Calculate"
3. View the optimal pack distribution in the results table

### Example Scenarios

**Example 1: Standard packs**
- Pack sizes: 250, 500, 1000, 2000, 5000
- Order: 251 items
- Result: 1 × 500 = 500 items (minimizes total items)

**Example 2: Edge case**
- Pack sizes: 23, 31, 53
- Order: 500,000 items
- Result: 2×23 + 7×31 + 9,429×53 = 500,000 items exactly

## Algorithm Rules

The calculator follows these rules in order of priority:

1. **Only whole packs** can be sent (packs cannot be broken)
2. **Minimize total items** sent (primary objective)
3. **Minimize number of packs** (secondary objective)

## Architecture

```
packager/
├── cmd/main.go                     # Backend entry point
├── internal/
│   ├── service/packer/             # Core calculation logic
│   │   ├── packer_calculate.go    # Optimized DP algorithm
│   │   └── packer_calculate_test.go
│   ├── repository/mem/             # In-memory storage
│   └── transport/graphql/          # GraphQL API
└── web/                            # Svelte frontend
    └── src/routes/+page.svelte     # Main UI
```

## API Endpoints (GraphQL)

### Query: Get All Packs
```graphql
query {
  packGetAll(input: {}) {
    packs {
      UID
      Size
    }
  }
}
```

### Mutation: Create Pack
```graphql
mutation {
  packCreate(input: { size: 250 }) {
    pack {
      UID
      Size
    }
  }
}
```

### Query: Calculate Packs
```graphql
query {
  packCalculate(input: { items: 251 }) {
    calculations {
      PackSize
      Quantity
    }
  }
}
```

## Performance

The optimized algorithm handles large orders efficiently:
- **Standard orders (12K items)**: ~82 μs
- **Edge case (500K items)**: ~12 ms
- **Memory efficient**: Uses backpointers instead of copying

## Development

### Running Tests

```bash
# Backend tests
make test

# With benchmarks
make bench
```

### Code Quality

```bash
# Frontend linting
cd web
bun run lint

# Format code
bun run format
```

## Technologies Used

**Backend:**
- Go 1.25+
- GraphQL (gqlgen)
- Dynamic Programming algorithm

**Frontend:**
- Svelte 5 (with runes)
- SvelteKit
- Tailwind CSS 4
- TypeScript

## License

MIT
