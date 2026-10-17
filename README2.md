# TDAs en Go

Colección de **Tipos de Dato Abstractos** implementados desde cero en Go, usando genéricos. Cada TDA expone una interfaz clara y oculta su implementación interna.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)

## Contenido

| TDA | Paquete | Implementación | Descripción |
|-----|---------|----------------|-------------|
| Pila | `pila` | Dinámica (arreglo redimensionable) | LIFO: último en entrar, primero en salir |
| Cola | `cola` | Enlazada | FIFO: primero en entrar, primero en salir |
| Lista | `lista` | Enlazada | Secuencia con inserción/borrado en cualquier posición e iterador |
| Diccionario | `diccionario` | Hash abierto (listas enlazadas) | Pares clave–valor con iterador |
| Diccionario ordenado | `abb` | Árbol binario de búsqueda | Pares clave–valor ordenados por clave, con iterador en orden |

## Estructura del repositorio

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
└── README.md
```

## Instalación

```bash
go get github.com/<usuario>/<repo>
```

Y luego se importa el TDA que se necesite:

```go
import (
    "github.com/<usuario>/<repo>/pila"
    "github.com/<usuario>/<repo>/cola"
    "github.com/<usuario>/<repo>/lista"
    "github.com/<usuario>/<repo>/diccionario"
    "github.com/<usuario>/<repo>/abb"
)
```

## Interfaces y complejidad

> Las firmas son orientativas: ajustalas a los nombres exactos de tu código.

### Pila

```go
type Pila[T any] interface {
    EstaVacia() bool
    VerTope() T          // panic si está vacía
    Apilar(T)
    Desapilar() T        // panic si está vacía
}
```

| Operación | Complejidad |
|-----------|-------------|
| `EstaVacia`, `VerTope` | O(1) |
| `Apilar`, `Desapilar` | O(1) amortizado |

### Cola

```go
type Cola[T any] interface {
    EstaVacia() bool
    VerPrimero() T       // panic si está vacía
    Encolar(T)
    Desencolar() T       // panic si está vacía
}
```

| Operación | Complejidad |
|-----------|-------------|
| Todas | O(1) |

### Lista

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

| Operación | Complejidad |
|-----------|-------------|
| `InsertarPrimero`, `InsertarUltimo`, `BorrarPrimero`, `VerPrimero`, `VerUltimo`, `Largo` | O(1) |
| `Iterar`, recorrido completo | O(n) |
| Operaciones del iterador (`Insertar`, `Borrar`, `Siguiente`) | O(1) |

### Diccionario (hash)

```go
type Diccionario[K comparable, V any] interface {
    Guardar(clave K, dato V)
    Pertenece(clave K) bool
    Obtener(clave K) V   // panic si no existe
    Borrar(clave K) V    // panic si no existe
    Cantidad() int
    Iterar(visitar func(clave K, dato V) bool)
    Iterador() IterDiccionario[K, V]
}
```

| Operación | Promedio | Peor caso |
|-----------|----------|-----------|
| `Guardar`, `Pertenece`, `Obtener`, `Borrar` | O(1) | O(n) |
| `Iterar` | O(n + m)* | O(n + m)* |

\* n = elementos, m = tamaño de la tabla.

El hash usa **direccionamiento abierto con listas enlazadas** (encadenamiento) y redimensiona la tabla según el factor de carga.

### Diccionario ordenado (ABB)

```go
type DiccionarioOrdenado[K any, V any] interface {
    Diccionario[K, V]
    IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)
    IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}
```

Se construye con una función de comparación:

```go
dic := abb.CrearABB[string, int](strings.Compare)
```

| Operación | Promedio | Peor caso (árbol degenerado) |
|-----------|----------|------------------------------|
| `Guardar`, `Pertenece`, `Obtener`, `Borrar` | O(log n) | O(n) |
| `Iterar` (in-order) | O(n) | O(n) |

## Ejemplos de uso

```go
package main

import (
    "fmt"

    "github.com/<usuario>/<repo>/pila"
    "github.com/<usuario>/<repo>/diccionario"
)

func main() {
    // Pila
    p := pila.CrearPilaDinamica[int]()
    p.Apilar(1)
    p.Apilar(2)
    fmt.Println(p.Desapilar()) // 2

    // Diccionario
    d := diccionario.CrearHash[string, int]()
    d.Guardar("a", 10)
    d.Guardar("b", 20)
    fmt.Println(d.Obtener("a")) // 10

    d.Iterar(func(clave string, dato int) bool {
        fmt.Println(clave, dato)
        return true // false corta la iteración
    })
}
```

## Manejo de errores

Las operaciones inválidas (por ejemplo `Desapilar` sobre una pila vacía u `Obtener` de una clave inexistente) provocan un `panic` con un mensaje descriptivo, siguiendo la convención de los TDAs de la materia:

| Situación | Mensaje |
|-----------|---------|
| Pila/cola vacía | `"La pila esta vacia"` / `"La cola esta vacia"` |
| Lista vacía | `"La lista esta vacia"` |
| Clave inexistente | `"La clave no pertenece al diccionario"` |
| Iterador finalizado | `"El iterador termino de iterar"` |

Cada paquete incluye sus pruebas unitarias.

## Notas de diseño

- **Genéricos**: los TDAs funcionan con cualquier tipo sin conversiones ni `interface{}`.
- **Encapsulamiento**: solo se exponen las interfaces y los constructores; las estructuras internas son privadas.
- **Iteradores**: se ofrecen dos formas de recorrer, interna (`Iterar` con callback) y externa (`Iterador`).
- **Sin dependencias externas**: solo la librería estándar de Go.

## Licencia

Distribuido bajo la licencia MIT. Ver [`LICENSE`](LICENSE) para más información.
