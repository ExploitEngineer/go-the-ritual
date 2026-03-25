/*
Packages
A package is a folder of .go files that share the same `package` declaration.
This is Go's way of organizing and reusing code across files.

JS/TS → Go equivalents:
  a JS file with exports     → a .go file inside a package
  a JS module/folder         → a Go package (folder)
  export const x = ...       → func X() or var X (capital letter = exported)
  export default             → does not exist in Go
  module.exports             → does not exist in Go
*/

package main

import "fmt"

/*
 1. What is a package?

 Every .go file starts with a package declaration:
   package main       → special: this is the entry point of your program
   package utils      → a reusable package named "utils"
   package models     → a reusable package named "models"

 All files in the same folder MUST have the same package name.
 The folder name and package name are usually the same (convention, not rule).

 JS:  every file is its own module, you pick what to export
 Go:  all files in a folder are ONE package and share everything
*/

/*
 2. Exported vs unexported

 JS/TS:
   export const Add = (a, b) => a + b    exported
   const helper = () => {}               not exported (no export keyword)

 Go:
   func Add(a, b int) int { ... }        exported   (capital A)
   func helper() { ... }                unexported (lowercase h)

 That's the entire rule. Capital first letter = exported. No keyword needed.
 This applies to: functions, types, variables, constants, struct fields.
*/

/*
 3. package main is special

 JS:  index.js or whatever you point node at
 Go:  any file with `package main` and a `func main()` is the entry point

 You can only have ONE func main() in your whole program.
 All other packages are libraries — they don't have main().
*/

/*
 4. Folder structure example

 myproject/
 ├── go.mod                 → module github.com/you/myproject
 ├── main.go                → package main
 ├── utils/
 │   └── math.go            → package utils
 └── models/
     └── user.go            → package models

 To use utils from main.go:
   import "github.com/you/myproject/utils"
   utils.Add(1, 2)

 JS equivalent:
   import { Add } from './utils/math'
   Add(1, 2)

 Key difference: in Go you always call the package name first (utils.Add),
 you cannot destructure imports like JS ({ Add }).
*/

/*
 5. init() function

 Each package can have an init() function that runs automatically
 before main() — used for setup like registering drivers, config, etc.

 JS:  no direct equivalent (closest is top-level code in a module)
 Go:  func init() { ... } runs once when the package is first imported
*/

func init() {
	// this runs before main() automatically
	// useful for setup tasks
}

func Packages() {
	fmt.Println("Capital letter   = exported  (public in JS terms)")
	fmt.Println("lowercase letter = unexported (private in JS terms)")
}
