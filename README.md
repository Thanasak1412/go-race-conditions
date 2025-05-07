# go-race-conditions

A repository dedicated to demonstrating, identifying, and solving race conditions in Go applications. This project serves as both a learning resource and a reference guide for dealing with concurrency issues in Go.

## Introduction

Race conditions occur when multiple goroutines access shared data concurrently, leading to unpredictable and erroneous behavior. This repository includes:

- Examples of common race conditions
- Tools and techniques to identify race conditions
- Patterns and best practices to prevent race conditions
- Solutions for fixing race conditions

## Installation

To get started with this repository:

```bash
# Clone the repository
git clone https://github.com/Thanasak1412/go-race-conditions.git

# Navigate to the project directory
cd go-race-conditions

# Install dependencies (if any)
go mod tidy
```

## Usage

Each example in this repository demonstrates a different aspect of race conditions:

```bash
# Run a specific example
go run examples/basic_race/main.go

# Run with Go's race detector
go run -race examples/basic_race/main.go
```

## Examples

This repository contains several examples including:

1. **Basic Race Conditions** - Simple demonstrations of race conditions
2. **Mutex Solutions** - Using mutexes to prevent race conditions
3. **Channel-based Solutions** - Using channels for safe concurrency
4. **Atomic Operations** - Using atomic operations for simple shared state
5. **Wait Groups** - Proper synchronization using sync.WaitGroup

Each example directory includes:
- Code that demonstrates the race condition
- Tests that can reproduce the issue
- Solutions that fix the race condition

## Best Practices

- Always use Go's race detector during development
- Prefer using channels for communication between goroutines
- Use appropriate synchronization primitives (Mutex, RWMutex)
- Minimize shared state between goroutines
- Consider using atomic operations for simple counters

## Contributing

Contributions are welcome! If you'd like to contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-fix`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-fix`)
5. Open a Pull Request

Please ensure your code examples include tests and documentation.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
