package cola

const mensajeColaVacia = "La cola esta vacia"

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(mensajeColaVacia)
	}

	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elemento T) {
	nuevo := &nodoCola[T]{
		dato: elemento,
	}

	if c.EstaVacia() {
		c.primero = nuevo
	} else {
		c.ultimo.prox = nuevo
	}

	c.ultimo = nuevo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(mensajeColaVacia)
	}
	dato := c.primero.dato
	c.primero = c.primero.prox

	if c.primero == nil {
		c.ultimo = nil
	}

	return dato
}
