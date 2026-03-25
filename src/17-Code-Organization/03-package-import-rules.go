/*
Package Import Rules
Go has strict, explicit import rules. Unused imports are compile errors.
Understanding these rules removes most of the confusion coming from JS/TS.

JS/TS → Go equivalents:
  import React from 'react'              → import "fmt"
  import { useState } from 'react'       → no destructuring in Go
  import * as fs from 'fs'               → import "os"
  import myAlias from './something'      → import myAlias "some/package"
   unused import (just a warning)      → unused import = COMPILE ERROR in Go
*/

package main

import (
	"fmt" // used below — Go will refuse to compile if this is unused
	"strings"
)

/*
1. Basic import
      then use: filepath.Join(...)

 You always use the package name to call its functions.
 No destructuring. No picking individual functions.
*/

/*
2. Multiple imports

 JS:
   import fs from 'fs'
   import path from 'path'

 Go (grouped import — preferred style):
   import (
       "fmt"
       "os"
       "path/filepath"
   )

 The group style is standard Go — gofmt enforces it automatically.
*/

/*
Aliasing imports

 JS:  import * as MyFmt from 'some-fmt-lib'
 Go:  import MyFmt "some/fmt/lib"
      then use: MyFmt.Println(...)

 Useful when two packages have the same name:
   import (
       pgx "github.com/jackc/pgx/v5"
       mysql "github.com/go-sql-driver/mysql"
   )
*/

/*
3. Blank import

 Sometimes you import a package ONLY for its side effects (its init() runs).
 But Go would complain about an unused import — so you use _ as the alias.

 JS:  import './setup'   (side-effect only import)
 Go:  import _ "github.com/lib/pq"   (loads postgres driver via init())

 The _ tells Go: "I know I'm not using this directly, that's intentional."
*/

/*
4. Dot import

 Imports everything into the current namespace (avoid this in real code):
   import . "fmt"
   Println("hello")    no fmt. prefix needed

 JS:  like `import * from 'fmt'` if that existed
 Go:  legal but considered bad practice — makes code hard to read
*/

/*
5. Unused imports = compile error

 JS/TS:  unused import = warning (or ignored)
 Go:     unused import = your code WILL NOT compile

 This is intentional — Go forces clean code. If you imported something
 and aren't using it, remove it. The tool `goimports` does this automatically.
*/

/*
6. Standard library vs third party

 Standard library paths have no domain:
   "fmt", "os", "strings", "net/http", "encoding/json"

 Third party paths always start with a domain:
   "github.com/gin-gonic/gin"
   "github.com/jackc/pgx/v5"

 JS:  no visual difference between built-in and third-party in import path
 Go:  you can tell immediately just by looking at the path
*/

func PackageImportRules() {
	// using both imports so Go doesn't complain
	fmt.Println(strings.ToUpper("import rules in go"))
}
