package multiintervals

type MultiInterval struct {
	intervalParts []*IntervalPart
}

type IntervalPart struct {
	start int
	end   int
}

func newPart(start int, end int) *IntervalPart {
	if end <= start {
		panic("Start of interval needs to be smaller than the end")
	}

	part := IntervalPart{
		start: start,
		end:   end,
	}
	return &part
}

func New(boundaries ...int) *MultiInterval {
	if len(boundaries)%2 != 0 {
		panic("An even amount of boundaries is needed")
	}

	parts := make([]*IntervalPart, len(boundaries)/2)

	for i := 0; i < len(boundaries); i += 2 {
		parts[i/2] = newPart(boundaries[i], boundaries[i+1])
	}

	interval := MultiInterval{intervalParts: parts}

	return &interval
}

func (i *MultiInterval) GetNumParts() int {
	return len(i.intervalParts)
}

func (i *MultiInterval) GetGlobalStart() int {
	if i.GetNumParts() == 0 {
		panic("Interval not initialized!")
	}

	return i.GetStart(0)
}

func (i *MultiInterval) GetStart(part int) int {
	if i.GetNumParts() <= part {
		panic("Too few parts")
	}

	return i.intervalParts[part].start
}

func (i *MultiInterval) GetGlobalEnd() int {
	if i.GetNumParts() == 0 {
		panic("Interval not initialized!")
	}

	return i.GetEnd(i.GetNumParts() - 1)
}

func (i *MultiInterval) GetEnd(part int) int {
	if i.GetNumParts() <= part {
		panic("Too few parts")
	}

	return i.intervalParts[part].end
}

func (p *IntervalPart) Intersects(other *IntervalPart) bool {
	return p.start >= other.start && p.start < other.end ||
		p.end <= other.end && p.end > other.start
}

func (i *MultiInterval) GetPart(part int) *IntervalPart {
	if i.GetNumParts() <= part {
		panic("Too few parts")
	}

	return i.intervalParts[part]
}

func (i *MultiInterval) Intersects(other *MultiInterval) bool {
	if i.GetNumParts() == 0 || other.GetNumParts() == 0 {
		return false
	}

	currentPartI := 0
	currentPartOther := 0
	p1 := i.GetPart(currentPartI)
	p2 := other.GetPart(currentPartOther)

	for {
		if currentPartI >= i.GetNumParts() || currentPartOther >= other.GetNumParts() {
			return false
		}

		if p1.Intersects(p2) {
			return true
		}

		if p1.start < p2.start {
			currentPartI++
		} else {
			currentPartOther++
		}
	}
}
