# Commands & Documentation in Go

Go treats **documentation as part of the language**, not as an afterthought. The Go toolchain ships with built-in commands that allow you to **read, explore, and write documentation directly from source code**.

This makes Go one of the few ecosystems where:

- Docs live next to code
- Docs are always in sync
- No external tools are required

---

## Philosophy: Documentation as Code

In Go:

- Documentation is written using **comments**
- Comments are parsed by Go tools
- Well-written comments become **official documentation**

If your code compiles but your comments are bad, your Go code is considered incomplete.

---

## `go help`

The `go help` command provides built-in documentation for the Go toolchain itself.

### Basic usage

```bash
go help
```

Lists all available `go` subcommands.

```bash
go help build

go help test
```

Shows detailed help for a specific command.

### Why this matters

- No need to Google basic usage
- Always matches your installed Go version
- Works offline

---

## `go doc`

`go doc` allows you to read **documentation directly in the terminal**.

### Examples

```bash
go doc fmt

go doc fmt.Println

go doc strings.Split
```

This pulls documentation from:

- Package comments
- Type comments
- Function comments

### Key point

`go doc` reads **source code comments**, not external docs.

---

## `godoc`

`godoc` provides a **web-based documentation interface** generated from Go source code.

### Start a local documentation server

```bash
godoc -http=:6060
```

Then open:

```
http://localhost:6060
```

You can:

- Browse the standard library
- Search packages
- Read formatted documentation

---

## How Go Documentation Comments Work

Go uses **structured comments** placed immediately before declarations.

### Package comments

```go
// Package mathutils provides basic mathematical helpers.
package mathutils
```

### Function comments

```go
// Add returns the sum of a and b.
func Add(a, b int) int {
 return a + b
}
```

Rules:

- Comment must start with the **name of the item**
- Must be directly above the declaration
- Written in complete sentences

---

## Why This System Is Powerful

Go documentation is:

- **Always close to code**
- **Automatically generated**
- **Tooling-friendly**
- **Version-aware**

This prevents:

- Outdated docs
- Copy-paste documentation rot
- Dependency on third-party doc generators

---

## Exploring the Standard Library (Recommended Practice)

Instead of searching online, get used to:

```bash
go doc net/http

go doc os.File

go doc time.Time
```

Reading standard library docs this way:

- Improves Go fluency
- Exposes idiomatic patterns
- Teaches you how Go authors write comments

---

## Writing Good Documentation (Rule of Thumb)

Ask yourself:

- What does this do?
- Why does it exist?
- When should it be used?
- What are the edge cases?

If the comment answers those questions, it’s good Go documentation.

---

## Mental Model (Lock This In)

- Comments → Documentation
- `go doc` → Terminal reader
- `godoc` → Web reader
- `go help` → Toolchain manual

Documentation in Go is **not optional** — it is part of writing correct code.

---

## TL;DR

- Go ships with built-in documentation tools
- Docs are written as comments
- `go help` explains commands
- `go doc` reads code documentation
- `godoc` provides a web interface

If you can’t explain your code with comments, you don’t understand it yet.
