/*
Using Third Party Packages
Go has no central registry like npmjs.com. Packages are identified by their URL (usually GitHub). You pull them directly from source control.

JS/TS → Go equivalents:
  npmjs.com                    → no central registry (URL is the identity)
  npm install express          → go get github.com/gin-gonic/gin
  import express from 'express'→ import "github.com/gin-gonic/gin"
  node_modules/                → ~/.cache/go (lives outside your project)
  package.json dependencies    → go.mod require block
*/

package main

import "fmt"

/*
1. How to add a third party package

 Step 1: find the package (usually on GitHub, pkg.go.dev is the docs site)
 Step 2: run go get
   go get github.com/gin-gonic/gin

 This does three things:
   - downloads the package into your global cache (~/.cache/go)
   - adds it to go.mod under `require`
   - updates go.sum with the hash

 Step 3: import and use it in your code
   import "github.com/gin-gonic/gin"

 Step 4: run go mod tidy (cleans up any unused entries)
   go mod tidy
*/

/*
2. pkg.go.dev — the Go package docs site

 JS:  npmjs.com — search for packages
 Go:  pkg.go.dev — search for packages AND read their docs

 Every public Go package automatically gets a docs page at pkg.go.dev
 generated from its source code comments. No separate docs site needed.
*/

/*
3. Versioning

 JS:  "gin": "^1.2.3"   (caret = allow minor/patch updates)
 Go:  github.com/gin-gonic/gin v1.9.1   (exact version, always pinned)

 Go always pins exact versions — no ^ or ~ like npm.
 To upgrade: go get github.com/gin-gonic/gin@latest
 To pin a version: go get github.com/gin-gonic/gin@v1.9.1
*/

/*
4. Major versions (v2, v3...)

 This is unique to Go and confuses JS devs at first.

 When a package releases a breaking v2, the import path itself changes:
   v1:  import "github.com/jackc/pgx"
   v5:  import "github.com/jackc/pgx/v5"

 JS:  you just install the new version, same import path
 Go:  the /v5 is literally part of the import path

 Why? So you can use v1 and v5 in the same project without conflict.
 You'll see this on almost every popular package.
*/

/*
5. The global module cache

 JS:  node_modules/ lives inside each project (can be huge, per-project)
 Go:  ~/.cache/go is shared across ALL your Go projects

 Consequences:
   - no need to re-download the same package for every project
   - no need to gitignore node_modules (there is no local copy)
   - your project folder stays clean — only your code lives there

 To clear the cache: go clean -modcache (rarely needed)
*/

/*
6. go mod tidy

 Run this whenever you:
   - add a new import in your code
   - remove an import from your code
   - want to clean up go.mod and go.sum

 It's safe to run anytime. Think of it as the Go equivalent of
 running `npm install` after pulling someone else's repo.

 JS workflow:  git pull → npm install → run
 Go workflow:  git pull → go mod tidy → run
*/

/*
7. Example: what go.mod looks like with dependencies

 module github.com/yourname/myproject

 go 1.22

 require (
     github.com/gin-gonic/gin v1.9.1
     github.com/jackc/pgx/v5 v5.5.4
 )

 indirect dependencies (deps of your deps) are also listed:
 require (
     github.com/bytedance/sonic v1.9.1  indirect
 )

 JS equivalent in package.json:
 {
   "dependencies": {
     "express": "^4.18.2"    ← your direct deps
   }
 }
 (indirect deps hidden inside node_modules)
*/

func UsingThirdPartyPackages() {
	fmt.Println("Key steps:")
	fmt.Println("  1. go get github.com/some/package")
	fmt.Println("  2. import it in your .go file")
	fmt.Println("  3. go mod tidy to clean up")
	fmt.Println("  4. browse docs at pkg.go.dev")
}
