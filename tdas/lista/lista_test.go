package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

type alumno struct {
	nombre string
	padron int
}

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "Una lista recien creada debe reportar estar vacia")
	require.Equal(t, 0, lista.Largo(), "Una lista recien creada debe tener largo 0")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Intentar ver el primero de una lista vacia debe lanzar panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Intentar ver el ultimo de una lista vacia debe lanzar panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Intentar borrar el primero de una lista vacia debe lanzar panic")
}

func TestInsertarBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarPrimero(10)
	require.False(t, lista.EstaVacia(), "La lista no debe reportar estar vacia despues de insertar un elemento")
	require.Equal(t, 1, lista.Largo(), "El largo debe ser 1 despues de insertar un elemento")
	require.Equal(t, 10, lista.VerPrimero(), "El primer elemento debe de ser 10")
	require.Equal(t, 10, lista.VerUltimo(), "El ultimo elemento debe de ser 10 al haber un solo elemento")

	lista.InsertarPrimero(20)
	require.Equal(t, 2, lista.Largo(), "El largo debe actualizarse a 2 tras insertar un segundo elemento al principio")
	require.Equal(t, 20, lista.VerPrimero(), "El primer elemento debe ser el recien insertado (20)")
	require.Equal(t, 10, lista.VerUltimo(), "El ultimo elemento debe continuar siendo 10")

	require.Equal(t, 20, lista.BorrarPrimero(), "Borrar primero debe devolver 20")
	require.Equal(t, 1, lista.Largo(), "El largo debe decrementar a 1")
	require.Equal(t, 10, lista.VerPrimero(), "El nuevo primer elemento debe de ser 10")
	require.Equal(t, 10, lista.VerUltimo(), "El ultimo elemento debe mantenerse en 10")

	require.Equal(t, 10, lista.BorrarPrimero(), "Borrar primero del ultimo elemento debe devolver 10")
	require.True(t, lista.EstaVacia(), "La lista debe quedar vacia")
	require.Equal(t, 0, lista.Largo(), "El largo debe ser 0")
}

func TestInsertarUltimoYBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarUltimo("A")
	require.False(t, lista.EstaVacia(), "La lista no debe reportar estar vacia despues de insertar un elemento")
	require.Equal(t, 1, lista.Largo(), "El largo debe ser 1")
	require.Equal(t, "A", lista.VerPrimero(), "El primer elemento debe ser A")
	require.Equal(t, "A", lista.VerUltimo(), "El ultimo elemento debe ser A")

	lista.InsertarUltimo("B")
	lista.InsertarUltimo("C")
	require.Equal(t, 3, lista.Largo(), "El largo debe ser 3")
	require.Equal(t, "A", lista.VerPrimero(), "El primer elemento debe mantenerse en A")
	require.Equal(t, "C", lista.VerUltimo(), "El ultimo elemento debe actualizarse a C")

	require.Equal(t, "A", lista.BorrarPrimero(), "El borrado debe seguir orden FIFO (A)")
	require.Equal(t, "B", lista.VerPrimero(), "El nuevo primer elemento debe ser B")
	require.Equal(t, "C", lista.VerUltimo(), "El ultimo elemento debe mantenerse en 'C'")

	require.Equal(t, "B", lista.BorrarPrimero(), "El segundo borrado debe devolver 'B'")
	require.Equal(t, "C", lista.BorrarPrimero(), "El tercer borrado debe devolver 'C'")

	require.True(t, lista.EstaVacia(), "La lista debe reportar estar vacia tras borrar todos los elementos")
	require.Equal(t, 0, lista.Largo(), "El largo debe ser 0")
}

func TestComportamientoVaciada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()

	lista.InsertarPrimero(0.67)
	lista.InsertarUltimo(3.14)
	require.False(t, lista.EstaVacia(), "La lista no debe estar vacia")

	require.Equal(t, 0.67, lista.BorrarPrimero(), "El primer borrado debe devolver 0.67")
	require.Equal(t, 3.14, lista.BorrarPrimero(), "El segundo borrado debe devolver 3.14")
	require.True(t, lista.EstaVacia(), "La lista debe estar vacia tras borrar sus elementos")

	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Ver el primero en una lista vaciada debe de lanzar panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Ver el ultimo en una lista vaciada debe de lanzar panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Intentar borrar el primero en una lista vaciada debe de lanzar panic")
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const VOLUMEN = 10000

	for i := 0; i < VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i+1, lista.Largo(), "El largo debe actualizarse progresivamente")
		require.Equal(t, 0, lista.VerPrimero(), "El primero debe ser siempre 0 durante inserciones al final")
		require.Equal(t, i, lista.VerUltimo(), "El ultimo debe ser el elemento recien insertado")
	}

	require.False(t, lista.EstaVacia(), "La lista con 10.000 elementos no debe reportar estar vacia")

	for i := 0; i < VOLUMEN; i++ {
		require.Equal(t, i, lista.VerPrimero(), "El primero debe ser consistente con la posicion borrada")
		require.Equal(t, VOLUMEN-1, lista.VerUltimo(), "El ultimo debe mantenerse en VOLUMEN-1 hasta borrar el ultimo elemento")
		require.Equal(t, i, lista.BorrarPrimero(), "Borrar primero debe devolver el elemento en el orden correcto")
		require.Equal(t, VOLUMEN-1-i, lista.Largo(), "El largo debe de decrementarse progresivamente")
	}

	require.True(t, lista.EstaVacia(), "La lista debe estar vacia tras borrar todos los elementos")
	require.Equal(t, 0, lista.Largo(), "El largo debe ser 0")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Ver el primero tras vaciar volumen debe lanzar panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Ver el ultimo tras vaciar volumen debe lanzar panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Borrar primero tras vaciar volumen debe lanzar panic")
}

func TestDiferentesTiposDeDatos(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaInt.InsertarPrimero(42)
	require.Equal(t, 42, listaInt.BorrarPrimero(), "La lista generica debe permitir almacenar enteros")

	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarPrimero("Genio")
	require.Equal(t, "Genio", listaString.BorrarPrimero(), "La lista generica debe permitir almacenar strings")

	listaStruct := TDALista.CrearListaEnlazada[alumno]()
	al := alumno{nombre: "John Pork", padron: 100000}
	listaStruct.InsertarPrimero(al)
	require.Equal(t, al, listaStruct.BorrarPrimero(), "La lista generica debe permitir almacenar estructuras")

	listaPunt := TDALista.CrearListaEnlazada[*int]()
	valor := 123
	listaPunt.InsertarPrimero(&valor)
	require.Equal(t, &valor, listaPunt.BorrarPrimero(), "La lista generica debe permitir almacenar punteros")
}

func TestIteradorExternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HayAlgoMas(), "HayAlgoMas debe de ser falso en un iterador de lista vacia")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() }, "VerActual en iterador terminado debe lanzar panic")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() }, "Avanzar en iterador terminado debe lanzar panic")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() }, "Borrar en iterador terminado debe lanzar panic")
}

func TestIteradorExternoInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Insertar(1)

	require.Equal(t, 1, iter.VerActual(), "VerActual debe devolver el nuevo elemento insertado (1)")
	require.Equal(t, 1, lista.VerPrimero(), "VerPrimero en la lista debe reflejar que el elemento 1 esta al principio")
	require.Equal(t, 3, lista.Largo(), "El largo de la lista debe incrementarse a 3")

	iter.Avanzar()
	require.Equal(t, 2, iter.VerActual(), "Al avanzar, se debe acceder al elemento previo (2)")
	iter.Avanzar()
	require.Equal(t, 3, iter.VerActual(), "Al avanzar nuevamente, se debe acceder a (3)")
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas(), "El iterador debe reportar que no hay siguiente tras recorrer la lista")
}

func TestIteradorExternoInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iter := lista.Iterador()
	for iter.HayAlgoMas() {
		iter.Avanzar()
	}

	require.False(t, iter.HayAlgoMas(), "El iterador debe estar al final")
	iter.Insertar(3)

	require.True(t, iter.HayAlgoMas(), "Tras insertar al final, HayAlgoMas debe ser verdadero")
	require.Equal(t, 3, iter.VerActual(), "VerActual debe devolver 3")
	require.Equal(t, 3, lista.VerUltimo(), "VerUltimo en la lista debe actualizarse a 3")
	require.Equal(t, 3, lista.Largo(), "El largo de la lista debe incrementarse a 3")

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas(), "Al avanzar tras el ultimo elemento insertado, debe terminar la iteracion")
}

func TestIteradorExternoInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Avanzar()

	iter.Insertar(2)

	require.Equal(t, 2, iter.VerActual(), "VerActual debe devolver el elemento insertado en el medio (2)")
	require.Equal(t, 3, lista.Largo(), "El largo debe actualizarse a 3")

	iter.Avanzar()
	require.Equal(t, 3, iter.VerActual(), "Al avanzar debe acceder al elemento posterior (3)")

	require.Equal(t, 1, lista.BorrarPrimero(), "El primer elemento de la lista debe ser 1")
	require.Equal(t, 2, lista.BorrarPrimero(), "El segundo elemento de la lista debe ser 2")
	require.Equal(t, 3, lista.BorrarPrimero(), "El tercer elemento de la lista debe ser 3")
}

func TestIteradorExternoRemoverAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	borrado := iter.Borrar()

	require.Equal(t, 1, borrado, "El elemento borrado debe ser 1")
	require.Equal(t, 2, lista.VerPrimero(), "VerPrimero en la lista debe cambiar a 2")
	require.Equal(t, 2, iter.VerActual(), "VerActual en el iterador debe pasar a ser 2")
	require.Equal(t, 2, lista.Largo(), "El largo debe decrementarse a 2")
}

func TestIteradorExternoRemoverAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Avanzar()
	iter.Avanzar()

	borrado := iter.Borrar()

	require.Equal(t, 3, borrado, "El elemento borrado debe ser 3")
	require.Equal(t, 2, lista.VerUltimo(), "VerUltimo en la lista debe actualizarse a 2")
	require.False(t, iter.HayAlgoMas(), "Tras borrar el ultimo elemento, HayAlgoMas debe ser falso")
	require.Equal(t, 2, lista.Largo(), "El largo debe decrementarse a 2")
}

func TestIteradorExternoRemoverEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Avanzar()

	borrado := iter.Borrar()

	require.Equal(t, 2, borrado, "El elemento borrado del medio debe ser 2")
	require.Equal(t, 3, iter.VerActual(), "VerActual debe pasar a apuntar al siguiente elemento (3)")
	require.Equal(t, 2, lista.Largo(), "El largo de la lista debe ser 2")

	require.Equal(t, 1, lista.BorrarPrimero(), "La lista reducida debe contener 1 primero")
	require.Equal(t, 3, lista.BorrarPrimero(), "La lista reducida debe contener 3 al final")
}

func TestIteradorExternoCasosBorde(t *testing.T) {
	l1 := TDALista.CrearListaEnlazada[int]()
	iter1 := l1.Iterador()
	iter1.Insertar(100)
	require.True(t, iter1.HayAlgoMas(), "HayAlgoMas debe ser true tras insertar en lista vacia")
	require.Equal(t, 100, iter1.VerActual(), "VerActual debe ser 100")
	require.Equal(t, 100, l1.VerPrimero(), "VerPrimero en lista debe ser 100")
	require.Equal(t, 100, l1.VerUltimo(), "VerUltimo en lista debe ser 100")

	l2 := TDALista.CrearListaEnlazada[int]()
	l2.InsertarUltimo(10)
	l2.InsertarUltimo(20)

	iter2 := l2.Iterador()
	require.Equal(t, 10, iter2.Borrar(), "Primer borrado con iterador debe devolver 10")
	require.Equal(t, 20, iter2.Borrar(), "Segundo borrado con iterador debe devolver 20")
	require.False(t, iter2.HayAlgoMas(), "El iterador debe quedar al final")
	require.True(t, l2.EstaVacia(), "La lista debe haber quedado vacia")
	require.Equal(t, 0, l2.Largo(), "El largo de la lista debe ser 0")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Borrar() }, "Borrar en iterador terminado debe lanzar panic")

	l3 := TDALista.CrearListaEnlazada[int]()
	iter3 := l3.Iterador()
	iter3.Insertar(3)
	iter3.Insertar(2)
	iter3.Insertar(1)

	require.Equal(t, 1, l3.BorrarPrimero(), "Inserciones consecutivas al inicio deben mantener orden [1, 2, 3]")
	require.Equal(t, 2, l3.BorrarPrimero(), "El segundo elemento debe ser 2")
	require.Equal(t, 3, l3.BorrarPrimero(), "El tercer elemento debe ser 3")
}

func TestIteradorInternoSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	elementos := []int{10, 20, 30, 40, 50}
	for _, elem := range elementos {
		lista.InsertarUltimo(elem)
	}
	suma := 0
	visitados := 0

	lista.Iterar(func(v int) bool {
		suma += v
		visitados++
		return true
	})
	require.Equal(t, 150, suma, "La suma total acumulada por el iterador interno debe ser 150")
	require.Equal(t, len(elementos), visitados, "Se deben haber visitado todos los elementos de la lista")
}

func TestIteradorInternoConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	elementos := []int{1, 2, 3, 4, 5, 6, 7}

	for _, elem := range elementos {
		lista.InsertarUltimo(elem)
	}

	suma := 0
	visitados := 0

	lista.Iterar(func(v int) bool {
		suma += v
		visitados++
		if v == 4 {
			return false
		}
		return true
	})

	require.Equal(t, 10, suma, "La suma acumulada hasta cortar en 4 debe ser 10 (1+2+3+4)")
	require.Equal(t, 4, visitados, "Se deben haber visitado exactamente 4 elementos")
}

func TestIteradorInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	visitados := 0
	lista.Iterar(func(v int) bool {
		visitados++
		return true
	})
	require.Equal(t, 0, visitados, "En una lista vacia la funcion visitar del iterador interno no debe ejecutarse")
}
