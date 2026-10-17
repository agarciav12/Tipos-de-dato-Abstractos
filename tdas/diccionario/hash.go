package diccionario

import (
	"fmt"
	"hash/fnv" //Fuente: http://www.isthe.com/chongo/tech/comp/fnv/
	TDALista "tdas/lista"
)

const (
	_TAM_INICIAL        = 16
	_FACTOR_CARGA       = 3
	_FACTOR_REDIMENSION = 2
	_PANIC_NO_EXISTE    = "La clave no pertenece al diccionario"
	_PANIC_ITER         = "El iterador termino de iterar"
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterDicc[K comparable, V any] struct {
	hash      *hashAbierto[K, V]
	pos       int
	iterLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashear[K comparable](clave K, tam int) int {
	h := fnv.New64a()
	h.Write(convertirABytes(clave))
	return int(h.Sum64() % uint64(tam))
}

func crearTabla[K comparable, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		tabla:    crearTabla[K, V](_TAM_INICIAL),
		tam:      _TAM_INICIAL,
		cantidad: 0,
	}
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, existe := hash.buscar(clave)
	return existe
}

func (hash *hashAbierto[K, V]) buscar(clave K) (TDALista.IteradorLista[parClaveValor[K, V]], bool) {
	iter := hash.tabla[hashear(clave, hash.tam)].Iterador()

	for iter.HayAlgoMas() {
		if iter.VerActual().clave == clave {
			return iter, true
		}
		iter.Avanzar()
	}

	return nil, false
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	iter, existe := hash.buscar(clave)

	nuevoPar := parClaveValor[K, V]{
		clave: clave,
		dato:  dato,
	}

	if existe {
		iter.Borrar()
		iter.Insertar(nuevoPar)
		return
	}

	hash.tabla[hashear(clave, hash.tam)].InsertarUltimo(nuevoPar)
	hash.cantidad++

	if hash.cantidad >= hash.tam*_FACTOR_CARGA {
		hash.redimensionar(hash.tam * _FACTOR_REDIMENSION)
	}
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	iter, existe := hash.buscar(clave)
	if !existe {
		panic(_PANIC_NO_EXISTE)
	}
	return iter.VerActual().dato
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	iter, existe := hash.buscar(clave)
	if !existe {
		panic(_PANIC_NO_EXISTE)
	}

	dato := iter.Borrar().dato
	hash.cantidad--

	if hash.tam > _TAM_INICIAL && hash.cantidad*_FACTOR_CARGA*_FACTOR_REDIMENSION < hash.tam {
		hash.redimensionar(hash.tam / _FACTOR_REDIMENSION)
	}
	return dato
}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	nuevo := &hashAbierto[K, V]{
		tabla: crearTabla[K, V](nuevoTam),
		tam:   nuevoTam,
	}
	hash.Iterar(func(clave K, dato V) bool {
		nuevo.Guardar(clave, dato)
		return true
	})
	hash.tabla = nuevo.tabla
	hash.tam = nuevo.tam
}

func (hash *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i, seguir := 0, true; i < hash.tam && seguir; i++ {
		hash.tabla[i].Iterar(func(par parClaveValor[K, V]) bool {
			seguir = visitar(par.clave, par.dato)
			return seguir
		})
	}
}
func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iterHash := &iterDicc[K, V]{
		hash: hash,
		pos:  0,
	}
	iterHash.encontrarListaConElementos()
	return iterHash
}

func (iterHash *iterDicc[K, V]) encontrarListaConElementos() {
	for iterHash.pos < iterHash.hash.tam && iterHash.hash.tabla[iterHash.pos].EstaVacia() {
		iterHash.pos++
	}
	if iterHash.pos < iterHash.hash.tam {
		iterHash.iterLista = iterHash.hash.tabla[iterHash.pos].Iterador()
	} else {
		iterHash.iterLista = nil
	}
}

func (iterHash *iterDicc[K, V]) HayAlgoMas() bool {
	return iterHash.pos < iterHash.hash.tam
}

func (iterHash *iterDicc[K, V]) VerActual() (K, V) {
	if !iterHash.HayAlgoMas() {
		panic(_PANIC_ITER)
	}
	par := iterHash.iterLista.VerActual()
	return par.clave, par.dato
}

func (iterHash *iterDicc[K, V]) Avanzar() {
	if !iterHash.HayAlgoMas() {
		panic(_PANIC_ITER)
	}
	iterHash.iterLista.Avanzar()
	if !iterHash.iterLista.HayAlgoMas() {
		iterHash.pos++
		iterHash.encontrarListaConElementos()
	}
}
