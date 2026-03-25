/*
Modules & Dependencies
A module is the root of your Go project — like package.json in JS/TS.
It defines the module name and tracks all dependencies.

JS/TS → Go equivalents:
  npm init              → go mod init
  package.json          → go.mod
  package-lock.json     → go.sum
  node_modules/         → ~/.cache/go (global cache, not in your project)
  npm install           → go mod tidy
  npm install <pkg>     → go get <pkg>
  npm install --save    → go get <pkg> (always saved to go.mod)
*/

package main

import "fmt"

/*
1. go mod init

Run once when starting a project:
go mod init github.com/yourname/projectname

This creates go.mod — the module name is important because other files
import your packages using this path as the root.

JS:  no equivalent — in JS you just use relative paths or package names
Go:  every import is an absolute path from the module root

go.mod looks like this:

module github.com/yourname/projectname

go 1.22

require (
       github.com/some/package v1.2.3   // like "some-package": "^1.2.3" in package.json
)
*/

/*
2. go mod tidy

Run after adding or removing imports in your code:
go mod tidy

JS:  like running `npm install` after editing package.json manually
Go:  scans all .go files, adds missing deps, removes unused ones also updates go.sum (the lockfile)
*/

/*
3. go get

Add a specific package:
go get github.com/some/package
go get github.com/some/package@v1.2.3   // specific version

JS:  npm install some-package
Go:  go get github.com/some/package

Key difference: Go packages are identified by their URL (usually GitHub), not a short name like in npm. There is no central registry like npmjs.com — the URL IS the package identifier.
*/

/*
4. go mod vendor

Copies all dependencies into a /vendor folder inside your project:
go mod vendor

JS:  similar to committing node_modules (rare, usually avoided)
Go:  more common in Go, useful for:
      - offline builds
      - Docker builds without internet access
      - ensuring reproducible builds in CI

Once you have a vendor folder, build with:
go build -mod=vendor
*/

/*
5. go.sum

Auto-generated lockfile — never edit this manually.

JS:  package-lock.json or yarn.lock
Go:  go.sum

It stores cryptographic hashes of every dependency so Go can verify nothing has been tampered with. Always commit this file to git.
*/

func ModulesAndDependencies() {
	fmt.Println("Module commands (run in terminal, not in Go code):")
	fmt.Println("  go mod init github.com/yourname/project")
	fmt.Println("  go get github.com/some/package")
	fmt.Println("  go mod tidy")
	fmt.Println("  go mod vendor")
}
