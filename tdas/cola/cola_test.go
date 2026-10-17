package cola_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
)

const _CANTIDAD_ELEMENTOS_VOLUMEN = 10000

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	require.True(t, cola.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})

	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
}

func TestVaciarCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i < 10; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	require.False(t, cola.EstaVacia())

	for i := 0; i < 10; i++ {
		require.Equal(t, i, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})

	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})

	cola.Encolar(10)
	cola.Encolar(20)

	require.False(t, cola.EstaVacia())
	require.Equal(t, 10, cola.VerPrimero())
	require.Equal(t, 10, cola.Desencolar())
	require.Equal(t, 20, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestEncolarYDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()

	elementos := []float64{10.23, 20.001, 30.1234, 12.09, 12.06}

	for _, elemento := range elementos {
		cola.Encolar(elemento)
		require.Equal(t, elementos[0], cola.VerPrimero())
	}

	for i := 0; i < len(elementos); i++ {
		require.Equal(t, elementos[i], cola.Desencolar())
		if i < len(elementos)-1 {
			require.Equal(t, elementos[i+1], cola.VerPrimero())
		}
	}

	require.True(t, cola.EstaVacia())
}

func TestEncolarYDesencolarVarios(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	elementos := []int{1, 2, 3, 4, 5}

	for _, elemento := range elementos {
		cola.Encolar(elemento)
	}

	require.False(t, cola.EstaVacia())

	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.Desencolar())

	cola.Encolar(6)
	cola.Encolar(7)

	elementosEsperados := []int{3, 4, 5, 6, 7}

	for _, elemento := range elementosEsperados {
		require.Equal(t, elemento, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())
}

func TestColaVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i < _CANTIDAD_ELEMENTOS_VOLUMEN; i++ {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	for i := 0; i < _CANTIDAD_ELEMENTOS_VOLUMEN; i++ {
		require.Equal(t, i, cola.Desencolar())
		if i < _CANTIDAD_ELEMENTOS_VOLUMEN-1 {
			require.Equal(t, i+1, cola.VerPrimero())
		}
	}
	require.True(t, cola.EstaVacia())
}
