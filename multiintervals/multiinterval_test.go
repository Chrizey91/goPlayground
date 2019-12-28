package multiintervals

import (
	"strconv"
	"testing"
)

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

func TestJoin(t *testing.T) {
	interval1 := New(10, 20, 30, 40)
	interval2 := New(20, 30, 40, 50)
	interval3 := New(41, 49)
	interval4 := New(35, 55, 100, 200, 300, 400)

	interval12 := interval1.Join(interval2)
	interval21 := interval2.Join(interval1)
	interval11 := interval1.Join(interval1)
	interval13 := interval1.Join(interval3)
	interval14 := interval1.Join(interval4)

	if interval12.GetNumParts() != 1 || interval21.GetNumParts() != 1 {
		t.Error("Interval should be one big interval after join. NumParts: " + strconv.Itoa(interval12.GetNumParts()))
	}

	if interval11.GetNumParts() != 2 {
		t.Error("Should still be the same")
	}

	if interval13.GetNumParts() != 3 {
		t.Error("Parts should be distinct")
	}

	if interval14.GetNumParts() != 4 {
		t.Error("Only last part should join")
	}
}
