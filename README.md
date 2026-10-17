# Data Structures in Go

🇦🇷 [Leer en español](README2.md)

A collection of **Abstract Data Types (ADTs)** implemented from scratch in Go using generics. Each ADT exposes a clear interface and hides its internal implementation.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)

## Contents

| ADT | Package | Implementation | Description |
|-----|---------|----------------|-------------|
| Stack | `pila` | Dynamic (resizable array) | LIFO: last in, first out |
| Queue | `cola` | Linked | FIFO: first in, first out |
| List | `lista` | Linked | Sequence with insertion/removal at any position and an iterator |
| Dictionary | `diccionario` | Hash table (separate chaining) | Key–value pairs with an iterator |
| Ordered dictionary | `abb` | Binary search tree (BST) | Key–value pairs sorted by key, with an in-order iterator |

## Repository structure

```
.
├── pila/
│   ├── pila.go
│   └── pila_test.go
├── cola/
│   ├── cola.go
│   └── cola_test.go
├── lista/
│   ├── lista.go
│   └── lista_test.go
├── diccionario/
│   ├── hash.go
│   └── hash_test.go
├── abb/
│   ├── abb.go
│   └── abb_test.go
├── go.mod
├── README.md
└── README.en.md
```

## Installation

```bash
go get github.com/<user>/<repo>
```

Then import whichever ADT you need:

```go
import (
    "github.com/<user>/<repo>/pila"
    "github.com/<user>/<repo>/cola"
    "github.com/<user>/<repo>/lista"
    "github.com/<user>/<repo>/diccionario"
    "github.com/<user>/<repo>/abb"
)
```

## Interfaces and complexity

> Signatures are indicative: adjust them to match the exact names in the code.

### Stack

```go
type Pila[T any] interface {
    EstaVacia() bool
    VerTope() T          // panics if empty
    Apilar(T)
    Desapilar() T        // panics if empty
}
```

| Operation | Complexity |
|-----------|------------|
| `EstaVacia`, `VerTope` | O(1) |
| `Apilar`, `Desapilar` | O(1) amortized |

### Queue

```go
type Cola[T any] interface {
    EstaVacia() bool
    VerPrimero() T       // panics if empty
    Encolar(T)
    Desencolar() T       // panics if empty
}
```

| Operation | Complexity |
|-----------|------------|
| All | O(1) |

### List

```go
type Lista[T any] interface {
    EstaVacia() bool
    InsertarPrimero(T)
    InsertarUltimo(T)
    BorrarPrimero() T
    VerPrimero() T
    VerUltimo() T
    Largo() int
    Iterar(visitar func(T) bool)
    Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
    VerActual() T
    HaySiguiente() bool
    Siguiente()
    Insertar(T)
    Borrar() T
}
```

| Operation | Complexity |
|-----------|------------|
| `InsertarPrimero`, `InsertarUltimo`, `BorrarPrimero`, `VerPrimero`, `VerUltimo`, `Largo` | O(1) |
| `Iterar`, full traversal | O(n) |
| Iterator operations (`Insertar`, `Borrar`, `Siguiente`) | O(1) |

### Dictionary (hash table)

```go
type Diccionario[K comparable, V any] interface {
    Guardar(clave K, dato V)
    Pertenece(clave K) bool
    Obtener(clave K) V   // panics if the key does not exist
    Borrar(clave K) V    // panics if the key does not exist
    Cantidad() int
    Iterar(visitar func(clave K, dato V) bool)
    Iterador() IterDiccionario[K, V]
}
```

| Operation | Average | Worst case |
|-----------|---------|------------|
| `Guardar`, `Pertenece`, `Obtener`, `Borrar` | O(1) | O(n) |
| `Iterar` | O(n + m)* | O(n + m)* |

\* n = number of elements, m = table size.

The hash table uses **separate chaining with linked lists** and resizes according to its load factor.

### Ordered dictionary (BST)

```go
type DiccionarioOrdenado[K any, V any] interface {
    Diccionario[K, V]
    IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)
    IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}
```

It is created with a comparison function:

```go
dic := abb.CrearABB[string, int](strings.Compare)
```

| Operation | Average | Worst case (degenerate tree) |
|-----------|---------|------------------------------|
| `Guardar`, `Pertenece`, `Obtener`, `Borrar` | O(log n) | O(n) |
| `Iterar` (in-order) | O(n) | O(n) |

## Usage examples

```go
package main

import (
    "fmt"

    "github.com/<user>/<repo>/pila"
    "github.com/<user>/<repo>/diccionario"
)

func main() {
    // Stack
    p := pila.CrearPilaDinamica[int]()
    p.Apilar(1)
    p.Apilar(2)
    fmt.Println(p.Desapilar()) // 2

    // Dictionary
    d := diccionario.CrearHash[string, int]()
    d.Guardar("a", 10)
    d.Guardar("b", 20)
    fmt.Println(d.Obtener("a")) // 10

    d.Iterar(func(clave string, dato int) bool {
        fmt.Println(clave, dato)
        return true // returning false stops the iteration
    })
}
```

## Error handling

Invalid operations (for example `Desapilar` on an empty stack, or `Obtener` with a missing key) trigger a `panic` with a descriptive message:

| Situation | Message |
|-----------|---------|
| Empty stack/queue | `"La pila esta vacia"` / `"La cola esta vacia"` |
| Empty list | `"La lista esta vacia"` |
| Missing key | `"La clave no pertenece al diccionario"` |
| Exhausted iterator | `"El iterador termino de iterar"` |

Each package ships with its own unit tests.

## Design notes

- **Generics**: the ADTs work with any type, with no conversions or `interface{}`.
- **Encapsulation**: only interfaces and constructors are exposed; internal structs are private.
- **Iterators**: two ways to traverse are provided, internal (`Iterar` with a callback) and external (`Iterador`).
- **No external dependencies**: standard library only.

## Contributing

1. Fork the repository.
2. Create a branch: `git checkout -b feature/new-adt`
3. Add tests for any change.
4. Make sure `go test ./...` and `go vet ./...` pass.
5. Open a Pull Request.

## License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for details.
