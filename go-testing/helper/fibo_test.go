package helper

import "testing"

// Run benchmark with go test -v -bench=BenchmarkName (use . for running all benchmark)
// go test -v -run=NoMatch -bench=BenchmarkName for running benchmark without unit test
func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibo(5)
	}
}

// Sub benchmark
func BenchmarkFiboSub(b *testing.B) {
	b.Run("Fibo 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fibo(10)
		}
	})

	b.Run("Fibo 15", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fibo(15)
		}
	})
}

// Table benchmark
func BenchmarkFiboTable(b *testing.B) {
	test := []struct {
		Name  string
		Value int
	}{
		{
			Name:  "Fibo 3",
			Value: 3,
		},
		{
			Name:  "Fibo 3",
			Value: 4,
		},
		{
			Name:  "Fibo 3",
			Value: 5,
		},
		{
			Name:  "Fibo 3",
			Value: 6,
		},
		{
			Name:  "Fibo 3",
			Value: 7,
		},
	}

	for _, test := range test {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fibo(test.Value)
			}
		})
	}
}
