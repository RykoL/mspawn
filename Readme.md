# mspawn
One command to run them all. Orchestrate your entire development stack with a single entry point and graceful shutdown.

## 🚀 Introduction
Modern development stacks are complex. To start working, you often find yourself juggling three or four terminal tabs: one for the Go backend, one for the Tailwind watcher, another for the React frontend, and maybe a database proxy.

Managing these is a hassle. Worse, when you are done, you have to visit every tab to `Ctrl+C` them individually, or hunt down PIDs when a process refuses to die.

**mspawn** solves this. It is a lightweight CLI tool written in Go that takes process arguments from the cli and launches all your required processes concurrently. When you are finished, a single `Ctrl+C` propagates the signal to all children, ensuring a clean, synchronized shutdown of your entire stack.

## ✨ Key Features
- **Concurrent Execution**: Launch unlimited processes simultaneously without blocking.

- **Unified Signal Handling**: Captures `SIGINT` (Ctrl+C) and `SIGTERM` at the parent level and propagates it to all child processes, preventing "zombie" processes.

- **Log Aggregation**: (Planned) Streams stdout/stderr from all services into a single unified view.

- **Zero Dependencies**: A single binary is all you need.

## 🛠 Installation
Prerequisites
- Go 1.25+ installed on your machine.

Build from Source
We use a `Makefile` to streamline dependency management and building.

1. **Clone the repository**:

```bash
git clone https://github.com/username/mspawn.git
cd mspawn
```
2. **Install dependencies**:

```bash
make deps
```
3. **Build the binary**:

```bash
make build
```
**Note**: This will compile the binary to ./bin/mspawn.

## 📖 Usage

1. Starting multiple processes

To spawn multiple processes simply pass each command via the cmd flag:

```bash
mspawn run \
  --cmd "python -m http.server" \
  --cmd "npm run start"
```

2. Stopping
When you are ready to stop working, press `Ctrl+C` once. `mspawn` will catch the signal, send a shutdown request to the backend, frontend, and watchers, wait for them to exit, and then close itself.

## 🗺 Roadmap
We are actively working on making `mspawn` the go-to tool for local process management.

- [ ] **Log Prefixing**: color-coded output prefixes (e.g., [backend] Log message...) to easily distinguish streams.

- [ ] **Environment Variables**: Support for loading .env files into the process context.

- [ ] **Loading from config file**: Instead of using cli flags configuraitons can be persisted in yaml files.

## 🤝 Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project

2. Create your Feature Branch (git checkout -b feature/AmazingFeature)

3. Commit your Changes (git commit -m 'Add some AmazingFeature')

4. Push to the Branch (git push origin feature/AmazingFeature)

5. Open a Pull Request
