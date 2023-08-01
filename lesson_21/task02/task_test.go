package task02

import (
	"testing"
)

type testingDataT struct {
	name string
	f    func(int, int) int
	a    int
	b    int
	want int
}

func TestAnonymousFunctions(t *testing.T) {
	testingData := []testingDataT{
		{
			"Multiplying",
			func(a int, b int) int { return a * b },
			7,
			8,
			56,
		},
		{
			"Subtracting",
			func(a int, b int) int { return a - b },
			7,
			8,
			-1,
		},
		{
			"Addition",
			func(a int, b int) int { return a + b },
			7,
			8,
			15,
		},
	}

	for _, test := range testingData {
		t.Run(test.name, func(t *testing.T) {
			result := CallbackExecutor(test.f, test.a, test.b)
			if result != test.want {
				t.Errorf("Error in %s function, expected value is %d, got %d", "CallbackExecutor", result, test.want)
			}
		})
	}

}
