// mandelbrot unit tests by Erik Wrenholt 2015-11-28

package mandelbrot

import "testing"

func TestMandelbrot1(t *testing.T) {
	iterations := Calc(0.0, 0.0)
	if iterations != 1000 {
		t.Errorf("TestMandelbrot1: Calc(0.0, 0.0) expected %d, but got %d", 1000, iterations)
	}
}

func TestMandelbrot2(t *testing.T) {
	iterations := Calc(0.0, 1.0)
	if iterations != 6 {
		t.Errorf("TestMandelbrot2: Calc(0.0, 1.0) expected %d, but got %d", 6, iterations)
	}
}

func TestMandelbrotRender(t *testing.T) {
	set := RenderSet(2,2)
	expected := "  # \n  # \n ###\n####\n"
	if set != expected {
		t.Errorf("TestMandelbrotRender: Calc(2,2) expected %s, but got \n%s", expected, set)
	}
}

func TestMandelbrotRenderParallel(t *testing.T) {
	set := RenderSetParallel(2,2)
	expected := "  # \n  # \n ###\n####\n"
	if set != expected {
		t.Errorf("TestMandelbrotRenderParallel: Calc(2,2) expected %s, but got \n%s", expected, set)
	}
}

func TestMandelbrotRenderCompareParallel39(t *testing.T) {
	set := RenderSetParallel(39,39)
	expected := RenderSet(39,39)
	if set != expected {
		t.Errorf("TestMandelbrotRenderCompareParallel39: RenderSet__(39,39) expected %s, but got \n%s", expected, set)
	}
}

func TestMandelbrotRenderCompareParallel40(t *testing.T) {
	set := RenderSetParallel(40,40)
	expected := RenderSet(40,40)
	if set != expected {
		t.Errorf("TestMandelbrotRenderCompareParallel40: RenderSet__(40,40) expected %s, but got \n%s", expected, set)
	}
}

// go test -bench=".*"

func BenchmarkMandelbrotRender(b *testing.B) {
	RenderSet(400,400)
}

func BenchmarkMandelbrotRenderParallel(b *testing.B) {
	RenderSetParallel(400,400)
}

