package diccionario

import (
	TDAPila "tdas/pila"
)

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func (abb *abb[K, V]) buscarEnlace(clave K) **nodoAbb[K, V] {
	nodoRef := &abb.raiz
	for *nodoRef != nil {
		resComparacion := abb.cmp(clave, (*nodoRef).clave)
		if resComparacion < 0 {
			nodoRef = &(*nodoRef).izquierdo
		} else if resComparacion > 0 {
			nodoRef = &(*nodoRef).derecho
		} else {
			break
		}
	}
	return nodoRef
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	enlace := abb.buscarEnlace(clave)
	if *enlace == nil {
		*enlace = &nodoAbb[K, V]{clave: clave, dato: dato}
		abb.cantidad++
	} else {
		(*enlace).dato = dato
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return *abb.buscarEnlace(clave) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := *abb.buscarEnlace(clave)
	if nodo == nil {
		panic(_PANIC_NO_EXISTE)
	}
	return nodo.dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Borrar(clave K) V {
	enlace := abb.buscarEnlace(clave)
	nodo := *enlace
	if nodo == nil {
		panic(_PANIC_NO_EXISTE)
	}
	dato := nodo.dato

	switch {
	case nodo.izquierdo == nil && nodo.derecho == nil:
		*enlace = nil
		abb.cantidad--
	case nodo.izquierdo == nil:
		*enlace = nodo.derecho
		abb.cantidad--
	case nodo.derecho == nil:
		*enlace = nodo.izquierdo
		abb.cantidad--
	default:
		reemplazante := nodo.derecho
		for reemplazante.izquierdo != nil {
			reemplazante = reemplazante.izquierdo
		}
		claveReemplazante := reemplazante.clave
		datoReemplazante := abb.Borrar(claveReemplazante)
		nodo.clave = claveReemplazante
		nodo.dato = datoReemplazante
	}
	return dato
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || abb.cmp(*desde, nodo.clave) < 0 {
		if !abb.iterarRango(nodo.izquierdo, desde, hasta, visitar) {
			return false
		}
	}

	if (desde == nil || abb.cmp(*desde, nodo.clave) <= 0) &&
		(hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if hasta == nil || abb.cmp(nodo.clave, *hasta) < 0 {
		if !abb.iterarRango(nodo.derecho, desde, hasta, visitar) {
			return false
		}
	}

	return true
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		abb:   abb,
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
	}
	iter.apilarDesde(abb.raiz)
	return iter
}

func (iter *iterAbb[K, V]) apilarDesde(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.abb.cmp(*iter.desde, nodo.clave) > 0 {
			nodo = nodo.derecho
		} else if iter.hasta != nil && iter.abb.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izquierdo
		} else {
			iter.pila.Apilar(nodo)
			nodo = nodo.izquierdo
		}
	}
}

func (iter *iterAbb[K, V]) HayAlgoMas() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic(_PANIC_ITER)
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iterAbb[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic(_PANIC_ITER)
	}
	nodo := iter.pila.Desapilar()
	iter.apilarDesde(nodo.derecho)
}
