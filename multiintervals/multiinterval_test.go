package multiintervals

import "testing"

func TestIntersects(t *testing.T) {
	interval1 := New(10, 20, 30, 40, 50, 60)
	interval2 := New(20, 30, 40, 50, 60, 70)
	interval3 := New(20, 30, 40, 55, 60, 70)
	interval4 := New(100, 200, 300, 400, 500, 600)
	interval5 := New(55, 200, 300, 400, 500, 600)

	inter12 := interval1.Intersects(interval2) // false
	inter13 := interval1.Intersects(interval3) // true
	inter14 := interval1.Intersects(interval4) // false
	inter15 := interval1.Intersects(interval5) // true
	inter51 := interval1.Intersects(interval5) // true

	if !inter12 && inter13 && !inter14 && inter15 && inter51 {
		t.Error("Intersection not working")
	}
}
