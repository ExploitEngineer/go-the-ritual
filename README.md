# the-go-ritual

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)

> *A discipline. A descent. A record of understanding earned the hard way.*

---

This repository is **not** a tutorial.
It is **not** a playground.
It is **not** a collection of projects.

`the-go-ritual` exists as a **structured record of Go mastery**, built slowly, deliberately, and without shortcuts.

Every directory represents a concept.
Every file exists to remove ambiguity.
Nothing is here by accident.

Go is a language that strips comfort away.
No hidden control flow. No decorative abstractions. No illusions.

This repository follows the same rules:

* Explicit over clever
* Boring over impressive
* Correct over convenient
* Understanding over speed

If something feels "too easy", it is examined until it no longer does.

---

## Repository Structure

This repository is divided into **two main sections**:

1. `go-by-example` – contains all topics from Go by Example exercises, each with its own folder.
2. `roadmap` – contains topics from roadmap.sh, each as its own folder.

All folders are **placed at the top level** to give the repository **size and presence from the start**, while maintaining a clean and organized hierarchy.

```
.
├── go-by-example/
│   ├── Hello-World
│   ├── Values
│   ├── Variables
│   ├── Constants
│   ├── For
│   ├── If-Else
│   ├── Switch
│   ├── Arrays
│   ├── Slices
│   ├── Maps
│   ├── Functions
│   ├── Multiple-Return-Values
│   ├── Variadic-Functions
│   ├── Closures
│   ├── Recursion
│   ├── Range-over-Built-in-Types
│   ├── Pointers
│   ├── Strings-and-Runes
│   ├── Structs
│   ├── Methods
│   ├── Interfaces
│   ├── Enums
│   ├── Struct-Embedding
│   ├── Generics
│   ├── Range-over-Iterators
│   ├── Errors
│   ├── Custom-Errors
│   ├── Goroutines
│   ├── Channels
│   ├── Channel-Buffering
│   ├── Channel-Synchronization
│   ├── Channel-Directions
│   ├── Select
│   ├── Timeouts
│   ├── Non-Blocking-Channel-Operations
│   ├── Closing-Channels
│   ├── Range-over-Channels
│   ├── Timers
│   ├── Tickers
│   ├── Worker-Pools
│   ├── WaitGroups
│   ├── Rate-Limiting
│   ├── Atomic-Counters
│   ├── Mutexes
│   ├── Stateful-Goroutines
│   ├── Sorting
│   ├── Sorting-by-Functions
│   ├── Panic
│   ├── Defer
│   ├── Recover
│   ├── String-Functions
│   ├── String-Formatting
│   ├── Text-Templates
│   ├── Regular-Expressions
│   ├── JSON
│   ├── XML
│   ├── Time
│   ├── Epoch
│   ├── Time-Formatting-Parsing
│   ├── Random-Numbers
│   ├── Number-Parsing
│   ├── URL-Parsing
│   ├── SHA256-Hashes
│   ├── Base64-Encoding
│   ├── Reading-Files
│   ├── Writing-Files
│   ├── Line-Filters
│   ├── File-Paths
│   ├── Directories
│   ├── Temporary-Files-and-Directories
│   ├── Embed-Directive
│   ├── Testing-and-Benchmarking
│   ├── Command-Line-Arguments
│   ├── Command-Line-Flags
│   ├── Command-Line-Subcommands
│   ├── Environment-Variables
│   ├── Logging
│   ├── HTTP-Client
│   ├── HTTP-Server
│   ├── TCP-Server
│   ├── Context
│   ├── Spawning-Processes
│   ├── Execing-Processes
│   ├── Signals
│   └── Exit
├── roadmap/
│   ├── Introduction-to-Go
│   ├── Why-Use-Go
│   ├── History-of-Go
│   ├── Setting-up-the-Environment
│   ├── Hello-World-in-Go
│   ├── Go-Command
│   ├── Variables-Constants
│   ├── Var-vs-Short-Declaration
│   ├── Zero-Values
│   ├── Const-and-Iota
│   ├── Scope-and-Shadowing
│   ├── Data-Types
│   ├── Numeric-Types
│   ├── Integers-Signed-Unsigned
│   ├── Floating-Points
│   ├── Complex-Numbers
│   ├── Boolean
│   ├── Runes
│   ├── Strings
│   ├── Raw-String-Literals
│   ├── Interpreted-String-Literals
│   ├── Type-Conversion
│   ├── Commands-and-Docs
│   ├── Composite-Types
│   ├── Arrays
│   ├── Slices
│   ├── Capacity-and-Growth
│   ├── Make
│   ├── Slice-to-Array-Conversion
│   ├── Array-to-Slice-Conversion
│   ├── Comma-Ok-Idiom
│   ├── Structs
│   ├── Struct-Tags-JSON
│   ├── Embedding-Structs
│   ├── Conditionals
│   ├── If
│   ├── If-Else
│   ├── Switch
│   ├── Loops
│   ├── For-Loop
│   ├── For-Range
│   ├── Iterating-Maps
│   ├── Iterating-Strings
│   ├── Break
│   ├── Continue
│   ├── Goto-Discouraged
│   ├── Functions-Basics
│   ├── Variadic-Functions
│   ├── Multiple-Return-Values
│   ├── Anonymous-Functions
│   ├── Named-Return-Values
│   ├── Closures
│   ├── Call-by-Value
│   ├── Pointers-Basics
│   ├── Pointers-with-Structs
│   ├── Pointers-with-Maps-Slices
│   ├── Memory-Management
│   ├── Garbage-Collection
│   ├── Methods-vs-Functions
│   ├── Pointer-Receivers
│   ├── Value-Receivers
│   ├── Interfaces-Basics
│   ├── Empty-Interfaces
│   ├── Embedding-Interfaces
│   ├── Type-Assertions
│   ├── Type-Switch
│   ├── Generics
│   ├── Generic-Functions
│   ├── Generic-Types-Interfaces
│   ├── Type-Constraints
│   ├── Type-Inference
│   ├── Error-Handling-Basics
│   ├── Error-Interface
│   ├── Errors-New
│   ├── Fmt-Errorf
│   ├── Wrapping-Unwrapping-Errors
│   ├── Sentinel-Errors
│   ├── Panic-and-Recover
│   ├── Stack-Traces-Debugging
│   ├── Modules-Dependencies
│   ├── Go-Mod-Init
│   ├── Go-Mod-Tidy
│   ├── Go-Mod-Vendor
│   ├── Package-Import-Rules
│   ├── Using-3rd-Party-Packages
│   ├── Publishing-Modules
│   ├── Goroutines
│   ├── Channels
│   ├── Select-Statement
│   ├── Buffered-vs-Unbuffered
│   ├── Worker-Pools
│   ├── Sync-Package
│   ├── Mutexes
│   ├── WaitGroups
│   ├── Context-Package
│   ├── Deadlines-Cancellations
│   ├── Concurrency-Patterns
│   ├── Fan-in
│   ├── Fan-out
│   ├── Pipeline
│   ├── Race-Detection
│   ├── Standard-Library-I-O-File-Handling
│   ├── Flag
│   ├── Time
│   ├── Encoding-JSON
│   ├── OS
│   ├── Bufio
│   ├── Slog
│   ├── Regexp
│   ├── Go-Embed
│   ├── Testing-Package-Basics
│   ├── Table-Driven-Tests
│   ├── Mocks-Stubs
│   ├── HttpTest
│   ├── Benchmarks
│   ├── Coverage
│   ├── Ecosystem-Popular-Libraries
│   ├── Building-CLIs
│   ├── Bubbletea
│   ├── Cobra
│   ├── Urfave-CLI
│   ├── Web-Development
│   ├── Net-HTTP
│   ├── Frameworks-Optional
│   ├── Gin
│   ├── Echo
│   ├── Fiber
│   ├── Beego
│   ├── GRPC-Protocol-Buffers
│   ├── ORMs-DB-Access
│   ├── PGX
│   ├── GORM
│   ├── Logging
│   ├── Zerolog
│   ├── Zap
│   ├── Realtime-Communication
│   ├── Melody
│   ├── Centrifugo
│   ├── Go-Toolchain-Tools
│   ├── Core-Go-Commands
│   ├── Go-Run
│   ├── Go-Build
│   ├── Go-Install
│   ├── Go-Fmt
│   ├── Go-Mod
│   ├── Go-Test
│   ├── Go-Clean
│   ├── Go-Doc
│   ├── Go-Version
│   ├── Code-Quality-Analysis
│   ├── Go-Vet
│   ├── GoImports
│   ├── Linters
│   ├── Revive
│   ├── StaticCheck
│   ├── GolangCI-Lint
│   ├── Security
│   ├── Govulncheck
│   ├── Code-Generation-Build-Tags
│   ├── Go-Generate
│   ├── Build-Tags
│   ├── Performance-Debugging
│   ├── PProf
│   ├── Trace
│   ├── Race-Detector
│   ├── Deployment-Tooling
│   ├── Building-Executables
│   ├── Cross-Compilation
│   ├── Memory-Management-in-Depth
│   ├── Escape-Analysis
│   ├── Reflection
│   ├── Unsafe-Package
│   ├── Build-Constraints-Tags
│   ├── CGO-Basics
│   ├── Compiler-Linker-Flags
│   └── Plugins-Dynamic-Loading
```

This structure allows **each topic its own folder** for clarity and future expansion, maintaining a **ritualistic and professional presence**.

---

Every folder, every file, every line is a deliberate step into understanding Go.
There is no shortcut, no facade. Only mastery.

![Dark Gopher](https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png)
![Gopher Night](https://blog.golang.org/gopher/gopher.png)

---

**Author:** ExploitEngineer
**Repository:** the-go-ritual
*Enter deliberately. Progress relentlessly.*
