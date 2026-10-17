package pila_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
)

const _CANTIDAD_ELEMENTOS_VOLUMEN = 10000

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	require.True(t, pila.EstaVacia()) // Verifica que "EstaVacia()" funciona correctamente

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})

	// Ambos tests verifican que tanto "VerTope()" como "Desapilar()" son primitivas invalidas para una pila vacia

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

func TestVaciarPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}

	require.False(t, pila.EstaVacia()) // Pureba que luego de apilar elementos la pila ya no esta vacia

	for i := 9; i >= 0; i-- {
		pila.Desapilar()
	}

	require.True(t, pila.EstaVacia()) // Prueba que al vaciar una pila "EstaVacia()" devuelve true

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})

	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
}

func TestApilarYDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()

	elementos := []float64{10.23, 20.001, 30.1234}

	for _, elemento := range elementos {
		pila.Apilar(elemento)
		require.Equal(t, elemento, pila.VerTope())
	}

	for i := len(elementos) - 1; i >= 0; i-- {
		require.Equal(t, elementos[i], pila.Desapilar()) // Este test prueba que "Desapilar()" devuelve correctamente el ultimo elemento apilado (LIFO)
	}

	require.True(t, pila.EstaVacia())
}

func TestPilaEnVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= _CANTIDAD_ELEMENTOS_VOLUMEN; i++ { // Prueba el correcto funcionamiento de una pila en volumen (redimensionamiento)
		pila.Apilar(i)

		require.Equal(t, i, pila.VerTope()) // Prueba que el elemento tope siempre sea correcto (i)
	}

	for i := _CANTIDAD_ELEMENTOS_VOLUMEN; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar()) // Prueba que la pila cumpla con el LIFO (al desapilar, devuelve el ultimo elemento ingresado)

		if i != 0 {
			require.Equal(t, i-1, pila.VerTope()) // Lo mismo pero al desapilar el ultimo (i - 1)
		} else {
			require.True(t, pila.EstaVacia())
		}
	}
}

func TestCapacidadMinima(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < 8; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia())

	pila.Apilar(50)
	require.Equal(t, 50, pila.VerTope())
}
