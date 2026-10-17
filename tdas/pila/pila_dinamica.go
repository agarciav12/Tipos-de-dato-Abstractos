package pila

const (
	TAM_INICIAL               = 16
	FACTOR_REDIMENSIONAMIENTO = 2
	FACTOR_REDUCCION          = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	p := &pilaDinamica[T]{}

	p.cantidad = 0
	p.datos = make([]T, TAM_INICIAL)

	return p
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

// ...
func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * FACTOR_REDIMENSIONAMIENTO)
	}

	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	tope := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad*FACTOR_REDUCCION <= len(p.datos) && len(p.datos)/FACTOR_REDIMENSIONAMIENTO >= TAM_INICIAL {
		p.redimensionar(len(p.datos) / FACTOR_REDIMENSIONAMIENTO)
	}

	return tope
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	nuevosDatos := make([]T, tam)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}
