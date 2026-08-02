# 📊 Organizer CLI Benchmark Report

## Project

**Organizer CLI (Go)**

---

# Purpose

The purpose of this benchmark is to compare the Go implementation of Organizer CLI with the original Python implementation. The comparison focuses on architecture, deployment, maintainability, and runtime characteristics rather than exact performance measurements.

---

# Comparison Overview
____________________________________________________________________________________
|     Feature     |     Python Organizer CLI     |        Go Organizer CLI         |
|-----------------|------------------------------|---------------------------------|
| Language        | Python                       | Go                              |
| Compilation     | Interpreted                  | Compiled                        |
| Executable      | Requires Python runtime      | Standalone executable           |
| Typing          | Dynamic                      | Static                          |
| Architecture    | Dictionary-based             | Modular package-based           |
| Deployment      | Python installation required | Single executable               |
| Maintainability | Good                         | Improved through modular design |
------------------------------------------------------------------------------------
---

# Project Structure Comparison

## Python Version

```
main.py
functions.py
file_extensions.py
```

The Python implementation organizes functionality using multiple dictionaries and helper functions.

---

## Go Version

```
main.go

organizer/
    commands.go
    extensions.go
    logger.go
    organize.go
    types.go
    utils.go
```

The Go implementation separates responsibilities into individual modules, making the project easier to understand, maintain, and extend.

---

# Startup Characteristics

### Python

- Requires Python interpreter
- Depends on installed packages
- Executes through the Python runtime

### Go

- Compiled into a native executable
- No external runtime required after compilation
- Simple execution using a single binary

---

# Deployment Comparison

## Python

Requirements

- Python installed
- Required packages installed
- Environment configuration

Execution

```bash
python main.py
```

---

## Go

Requirements

- No Go installation required after building the executable

Execution

```bash
organizer.exe all
```

This simplifies distribution because the application can be shared as a single executable.

---

# Code Organization

## Python

- Large extension dictionaries
- Function-oriented implementation
- Less separation between responsibilities

---

## Go

Responsibilities are divided into dedicated modules:

- `commands.go`   – Supported CLI commands
- `extensions.go` – File categories and extensions
- `organize.go`   – Core organization logic
- `logger.go`     – Movement logging
- `utils.go`      – Helper functions
- `types.go`      – Shared data structures

This modular design improves readability and maintainability.

---

# Functional Comparison
________________________________________
| Feature               | Python  | Go  |
|-----------------------|---------|-----|
| Directory Scanning    |   ✅    | ✅ |
| File Classification   |   ✅    | ✅ |
| Folder Creation       |   ✅    | ✅ |
| File Organization     |   ✅    | ✅ |
| Duplicate Handling    |   ✅    | ✅ |
| Confirmation Prompt   |   ✅    | ✅ |
| Logging               |   ✅    | ✅ |
| Modular Architecture  | Limited | ✅ |
| Standalone Executable |   ❌    | ✅ |
-----------------------------------------
---

# Observations

The Go implementation provides several practical advantages:

- Strong static typing
- Modular project structure
- Easier long-term maintenance
- Standalone executable after compilation
- Cleaner separation of responsibilities

Both implementations provide equivalent file organization functionality, while the Go version focuses on improved project organization and deployment.

---

# Limitations

This benchmark does **not** include measured execution time, CPU usage, or memory usage. The comparison is based on implementation characteristics and project design rather than quantitative performance testing.

Future work may include benchmarking with large datasets to compare execution time and resource usage.

---

# Conclusion

The Go implementation successfully reproduces the functionality of the original Organizer CLI while introducing a cleaner modular architecture and simplified deployment through a standalone executable. These improvements make the project easier to maintain, extend, and distribute without changing its core behavior.