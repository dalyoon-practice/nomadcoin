# Nomadcoin

## Description

- specification: Golang
- purpose: Lecture Practice Following [Nomadcoders](https://nomadcoders.co/)
- createAt: 2021-08

---

## Notes

### - Section 4 -

<!-- 4.1 -->
<details><summary>#4.1 Our First Block</summary>

- What is one-way-function
- What is hash
- Generate genesis block with SHA256

</details>

<!-- 4.2 -->
<details><summary>#4.2 Our First Blockchain</summary>
  
- Make blocks as chain
- Genarate blocks through functions

</details>

<!-- 4.3 -->
<details><summary>#4.3 Singleton Pattern</summary>

- Refactoring to separated module
- What is singleton pattern

</details>

<!-- 4.4 -->
<details><summary>#4.4 Refactoring part One</summary>

- What is [sync.Once.Do](https://pkg.go.dev/sync@go1.16.7#Once.Do)
- Reactor codes to apply singleton pattern

</details>

<!-- 4.5 -->
<details><summary>#4.5 Refactoring part Two</summary>

- Continuing refactoring

</details>

### - Section 5 -

<!-- 5.0 -->
<details><summary>#5.0 Setup</summary>

- Basic Setting for web server

```go
// Simple Example
package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Server listening on port http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

```

</details>

<!-- 5.1 -->
<details><summary>#5.1 Rendering Templates</summary>

- How to use HTML template on go server

</details>

<!-- Section 6 -->
<details><summary><b>Section 6: RestAPI</b></summary>

- [Json](https://pkg.go.dev/encoding/json) in GO
- What is [Stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
- [Json - Marshal](https://pkg.go.dev/encoding/json#Marshal)
- [Encoding - TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler)
- [http-mux](https://pkg.go.dev/net/http#NewServeMux)

</details>
