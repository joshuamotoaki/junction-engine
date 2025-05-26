# TigerJunction Engine

WIP TigerJunction 2.0. For now, see the old [TigerJunction 1.0](https://github.com/TigerAppsOrg/tiger-junction) repository.

## 🏗️ Architecture

- **Web App** (`apps/web`): React Router v7 application
- **API** (`apps/api`): Elysia backend running on Bun
- **Shared** (`packages/shared`): Common utilities and types

## 🚀 Quick Start

```bash
# Install dependencies
npm install

# Start development servers
npm run dev

# Build all apps
npm run build

# Run linting and formatting
npm run lint:fix
npm run format
```

## 📁 Project Structure

```
junction-engine/
├── apps/
│   ├── web/          # React Router frontend (port 5173)
│   └── api/          # Elysia backend (port 3000)
├── packages/
│   └── shared/       # Shared utilities
└── turbo.json        # Turborepo configuration
```

## 📦 Available Scripts

- `npm run dev` - Start all development servers
- `npm run build` - Build all applications
- `npm run lint` - Run ESLint across all packages
- `npm run format` - Format code with Prettier
- `npm run typecheck` - Run TypeScript type checking

Code is automatically formatted and linted on commit via Husky pre-commit hooks.

## 🤝 Contributing

TODO: Add contribution guidelines

## ⚖️ License

This project is licensed under the BSD-3-Clause License. See the [LICENSE](LICENSE) file for details.
