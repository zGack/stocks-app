# AGENTS.md

## Build, Lint, and Test Commands

**Frontend (web/)**
- Build: `npm run build`
- Lint: `npm run lint`
- Format: `npm run format`
- Type-check: `npm run type-check`
- Unit tests: `npm run test:unit`
- Run a single test: `npx vitest run src/path/to/testfile` or `npx vitest --testNamePattern 'name'`

**Backend (Go)**
- Build: `go build ./...`
- Test: `go test ./...`
- Run a single test: `go test -run TestName ./...`
- Migrations: `make migrate-up`, `make migrate-down`

## Code Style Guidelines

- **Frontend:** Use Prettier (no semicolons, single quotes, print width 100). ESLint enforces Vue/TypeScript recommended rules. Use camelCase for variables, PascalCase for types. Prefer absolute imports.
- **Backend:** Use Go conventions (PascalCase for exported, camelCase for local). Handle errors with explicit error returns. Use standard Go imports.
- **General:** Keep code modular, readable, and well-documented. Avoid unused variables and imports. Use type safety and clear error handling.
