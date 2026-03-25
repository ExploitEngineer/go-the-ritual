/*
Publishing Modules
Publishing a Go module means pushing your code to a public git repo (usually GitHub).
There is no `npm publish` step — the URL of your repo IS the package.

JS/TS → Go equivalents:
  npm publish                  → git push (that's it)
  npmjs.com registry           → GitHub (or any git host)
  npm version 1.2.3            → git tag v1.2.3
  .npmignore                   → no equivalent needed
*/

package main

import "fmt"

/*
1. How publishing works

 JS:  you build → npm publish → package appears on npmjs.com
 Go:  you push to GitHub → package is immediately importable by anyone

 No account needed on a registry. No publish command. Just git push.
 pkg.go.dev automatically indexes your module once someone first fetches it.
*/

/*
2. Step by step

 1. Initialize your module with a path matching your GitHub repo:
      go mod init github.com/yourname/mypackage

 2. Write your exported functions (capital letters):
      package mypackage
      func Add(a, b int) int { return a + b }

 3. Push to GitHub:
      git push origin main

 4. Tag a version (required for versioned imports):
      git tag v1.0.0
      git push origin v1.0.0

 5. Anyone can now use it:
      go get github.com/yourname/mypackage@v1.0.0
      import "github.com/yourname/mypackage"
*/

/*
3. Versioning rules

 Go follows semantic versioning strictly:
   v1.0.0  → initial stable release
   v1.0.1  → bug fix (no API changes)
   v1.1.0  → new features (backwards compatible)
   v2.0.0  → breaking changes → import path MUST change to /v2

 JS:  breaking change → bump version in package.json, same import path
 Go:  breaking change → bump version AND update the module path in go.mod

 go.mod for v2:
   module github.com/yourname/mypackage/v2

 Users import v2 as:
   import "github.com/yourname/mypackage/v2"

 This lets users stay on v1 and v2 in the same project with no conflict.
*/

/*
4. What to export

 JS:  you decide per-file what to export with the export keyword
 Go:  capital letter = exported automatically, everywhere in the package

 Good practice — only export what users of your package need:
   func Add(a, b int) int { ... }     exported — part of your API
   func validate(n int) bool { ... }  unexported — internal detail
*/

/*
5. Documentation

 JS:  JSDoc comments, separate README, or tools like Storybook
 Go:  comments directly above exported functions become the docs

 pkg.go.dev reads these comments automatically:

    Add returns the sum of a and b.
   func Add(a, b int) int {
       return a + b
   }

 The comment above Add() will appear word-for-word on pkg.go.dev.
 No separate docs tool needed.
*/

/*
6. go.sum and reproducibility

 When someone imports your module, Go verifies it against a global
 checksum database (sum.golang.org) — similar to npm's package integrity
 checks but enforced by default for all packages.

 This means: you cannot silently change code on a tagged version.
 Once v1.0.0 is tagged and fetched by someone, its hash is recorded.
 Changing the code without a new tag will cause checksum failures for users.
*/

/*
7. GOPATH vs modules (historical context)

 Before Go 1.11, all Go code had to live in a specific folder called GOPATH.
 This confused everyone coming from other languages.

 Modules (go.mod) replaced this. You can now put your project anywhere.
 If you ever see old tutorials mentioning GOPATH for code organization,
 that approach is outdated — use modules.
*/

func PublishingModules() {
	fmt.Println("Publishing steps:")
	fmt.Println("  1. go mod init github.com/you/pkg")
	fmt.Println("  2. write exported functions (capital letters)")
	fmt.Println("  3. git push origin main")
	fmt.Println("  4. git tag v1.0.0 && git push origin v1.0.0")
	fmt.Println("  5. importable immediately — no npm publish needed")
}
