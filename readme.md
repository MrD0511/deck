# 🧱 Deck CLI

Deck is a blazing-fast, cross-platform CLI tool written in Go that **automatically generates optimized Dockerfiles** for your project. It detects your tech stack (like FastAPI, Flask, React, etc.), suggests best configurations, and outputs production-ready Dockerfiles — all with a single command.

---

## ⚙️ Features

- 🔍 **Smart Framework Detection**  
  Detects your project's framework: FastAPI, Flask, Django, Node.js, React, Angular, etc.

- ✏️ **Customizable Setup**  
  Automatically detected values (like base image, entrypoint, workdir) can be reviewed and edited via TUI.

- 🐳 **Dockerfile Generator**  
  Creates clean, efficient Dockerfiles with framework-specific optimizations.

- ⚗️ **Dev & Prod Support**  
  Generates both:
  - `Dockerfile` for production
  - `Dockerfile.dev` for development (volumes, live reload, etc.)

- 📦 **.dockerignore File Generator**  
  Generates `.dockerignore` with standard ignore patterns to reduce image bloat.

- ⚡ **Force Mode**  
  `--force` flag generates all files using default/detected config with zero prompts.

---

## 🪄 Result: Ultra-Light Docker Images

Deck minimizes image size using best practices like slim base images, clean layers, and no dev dependencies in production builds.

📉 Typical image sizes:
- 🐍 **Python (FastAPI/Flask)** → ~199 MB
- ⚛️ **JS (React/Node/Angular)** → ~54 MB

This means **faster CI/CD**, quicker deployments, and lower storage costs.

---

## 🚀 Quick Start

### 📥 Install

```bash
go install github.com/MrD0511/deck@latest
```
| Or use single-command installer (coming soon)

---

## 🧪 Usage
``` bash
deck generate (-p or -d) (folder path)
```

Deck will:

- Detect your framework & language

- Let you edit or confirm configuration

- Generate:
 - Dockerfile
 - .dockerignore

---

## ⚙️ Flags

``` bash
--force        # Skip prompts, use defaults
--prod         # Generate only production Dockerfile
--dev          # Generate only development 
```
---

## setup
### 📁 Example outputs
``` bash
.
├── Dockerfile
├── Dockerfile.dev
├── .dockerignore
└── your_project/
```
---

## 💡 Why Use Deck?

Without Deck, devs waste time googling Dockerfile snippets, tweaking commands, or copy-pasting bloated templates.

### Deck gives you:

🔒 Clean, opinionated Docker setups

🎯 Slim images with zero fluff

🧠 Smart detection + TUI editing

💥 Instant setup in any project

---

## 📬 Contributions & Feedback

Want to support more frameworks or package managers? Found a bug?
Open a PR or create an issue — contributions are welcome!

---

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

---

## 🙋‍♂️ Author

### Dhruv Sharma
🔗 [LinkedIn](https://www.linkedin.com/in/dhruvsharma005/)
🐙 [GitHub](https://github.com/MrD0511)
✉️ [Email](mailto:sharmadhruv00005@gmail.com)