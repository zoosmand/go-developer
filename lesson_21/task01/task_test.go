package task01

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

type argsT struct {
	x int16
	y uint8
	z float32
}

type testingDataT struct {
	name string
	args argsT
	want float32
}

func TestTypesCastingRandoms(t *testing.T) {
	singleTestingData := testingDataT{
		"The result is type of float32",
		argsT{int16(rand.Intn(math.MaxInt32)), uint8(rand.Uint32()), rand.Float32()},
		0.0,
	}
	result := TypesCasting(singleTestingData.args.x, singleTestingData.args.y, singleTestingData.args.z)

	if reflect.TypeOf(singleTestingData.args.z).Kind() != reflect.TypeOf(result).Kind() {
		t.Errorf("%s - %f", t.Name(), result)
		t.Errorf("x: %d, y: %d, z: %f", singleTestingData.args.x, singleTestingData.args.y, singleTestingData.args.z)
	}
}

func TestTypesCasting(t *testing.T) {
	testingData := []testingDataT{
		{
			"Max values",
			argsT{math.MaxInt16, math.MaxUint8, math.MaxFloat32},
			251.000000,
		},
		{
			"Min values",
			argsT{math.MinInt16, 0, (math.MaxFloat32 * -1)},
			2.000000,
		},
		{
			"Sample values",
			argsT{32000, 250, 32000.001},
			-1288.000122,
		},
	}
	for _, test := range testingData {
		t.Run(test.name, func(t *testing.T) {
			result := TypesCasting(test.args.x, test.args.y, test.args.z)
			if result != test.want {
				t.Errorf("Error in %s function, expected value is %f, got %f", "TypesCasting", result, test.want)
			}
		})
	}
}
