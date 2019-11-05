package structsmethodsintrerfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 37.0},
	}

	for _, it := range areaTests {
		t.Run(it.name, func(t *testing.T) {
			got := it.shape.Area()
			if got != it.want {
				t.Errorf("%#v got %g want %g", it.shape, got, it.want)
			}
		})
	}

	/*checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})*/
}

func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}
