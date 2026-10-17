package lista

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

// Primitivas lista

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func crearNodo[T any](elemento T) *nodoLista[T] {
	return &nodoLista[T]{
		dato: elemento,
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := crearNodo(elemento)

	nuevo.siguiente = lista.primero

	if lista.EstaVacia() {
		lista.ultimo = nuevo
	}

	lista.primero = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := crearNodo(elemento)

	if lista.EstaVacia() {
		lista.primero = nuevo
		lista.ultimo = nuevo
	} else {
		lista.ultimo.siguiente = nuevo
		lista.ultimo = nuevo
	}

	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	dato := lista.primero.dato

	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}

	lista.largo--

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for nodo := lista.primero; nodo != nil; nodo = nodo.siguiente {
		if !visitar(nodo.dato) {
			return
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{lista: lista, actual: lista.primero}
}

// Primitivas iterador

func (iter *iteradorLista[T]) VerActual() T {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) HayAlgoMas() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Borrar() T {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}

	dato := iter.actual.dato

	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente

		if iter.actual.siguiente == nil {
			iter.lista.ultimo = nil
		}
	} else {
		iter.anterior.siguiente = iter.actual.siguiente

		if iter.actual.siguiente == nil {
			iter.lista.ultimo = iter.anterior
		}
	}

	iter.actual = iter.actual.siguiente

	iter.lista.largo--
	return dato
}

func (iter *iteradorLista[T]) Insertar(elemento T) {
	nuevo := crearNodo(elemento)

	if iter.anterior == nil {
		nuevo.siguiente = iter.lista.primero
		iter.lista.primero = nuevo
		if iter.lista.ultimo == nil {
			iter.lista.ultimo = nuevo
		}
	} else {
		nuevo.siguiente = iter.actual
		iter.anterior.siguiente = nuevo
		if iter.actual == nil {
			iter.lista.ultimo = nuevo
		}
	}

	iter.actual = nuevo
	iter.lista.largo++
}
