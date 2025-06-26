# TigerJunction Engine

WIP TigerJunction 2.0. For the current production version, see the old [TigerJunction 1.0](https://github.com/TigerAppsOrg/tiger-junction) repository.

## 🏗️ Architecture

- **Web App** (`apps/web`): React Router v7 application
- **API** (`apps/api`): Go backend with the [Gin](https://gin-gonic.com/) framework
    - **Database**: Neo4J
    - **Authentication**: JWT-based with Princeton CAS
- **Shared** (`packages/shared`): Common utilities and types

## 📁 Project Structure

```
junction-engine/
├── apps/
│   ├── web/          # React Router frontend (port 5173)
│   └── api/          # Go backend (port 3000)
├── packages/
│   └── shared/       # Shared utilities
└── <root>            # Config files
```

Code is automatically formatted and linted on commit via Husky pre-commit hooks. Use the command line to run git commands, the VSCode GUI may cause problems.

## 🤝 Contributing

Reach out to the TigerApps team at it.admin@tigerapps.org if you're interested in contributing! If it's something small, feel free to open an issue or pull request directly.

## ⚖️ License

This project is licensed under the BSD-3-Clause License. See the [LICENSE](LICENSE) file for details.
