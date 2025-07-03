# dictionaries
A simple in-memory dictionary implementation in Go with support for basic CRUD operations (Create, Read, Update, Delete). The package includes thorough unit tests and examples to demonstrate its behavior and correctness. Designed for eductational and demonstrational purposes, especially when learning Go, testing practices, and TDD (Test-Driven Development).

[Go Reference Documentation](https://pkg.go.dev/github.com/imunoka/dictionaries)

---

## 📦 Features

- **Search**: Look up a value by its key. Returns an error if the key does not exist.

- **Add**: Insert a new key/value pair. Returns an error if the key already exists.
 
- **Update**: Update the value of an existing key. Returns an error if the key does not exist.

- **Delete**: Remove a key/value pair. Returns an error if the key does not exist.

---

## 🧨 Error Handling
Consistent and predictable error responses for edge cases:

    ErrNotFound when a key is not found during a search.

    ErrKeyExists when attempting to add a key that already exists.

    ErrKeyDoesNotExist when updating or deleting a non-existent key.

---

## 🛠️ Test Coverage
Includes unit tests for all operations:

- Positive and negative test cases for each method.

- Example-based documentation tests for GoDoc output.

---

## 💡 Example
dictionary := Dictionary{}

```go
package main

import (
  "fmt"
  . "github.com/imunoka/wallet"
)

func main() {
	dictionary := Dictionary{}
	err := dictionary.Add("key", "value")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	value, err := dictionary.Search("key")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("found:", value)

	// Output: found: value
}
```

---

## 🧪 Testing

Run tests with the following:

```bash
go test
```

To see full test output, use the following:
```bash
go test -v
```
