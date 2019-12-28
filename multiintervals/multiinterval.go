package multiintervals

import "container/list"

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

func newInternal(parts []*IntervalPart) *MultiInterval {
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
	if other == nil || i.GetNumParts() == 0 || other.GetNumParts() == 0 {
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

func (i *MultiInterval) Join(other *MultiInterval) *MultiInterval {
	if other == nil || other.GetNumParts() == 0 {
		return i
	}

	if i.GetNumParts() == 0 {
		return other
	}

	currentPartI := 0
	currentPartOther := 0
	partLst := list.New()
	var lastJoinedPart *IntervalPart

	if i.GetPart(currentPartI).start < other.GetPart(currentPartOther).start {
		lastJoinedPart = i.GetPart(currentPartI)
		currentPartI++
	} else {
		lastJoinedPart = other.GetPart(currentPartOther)
		currentPartOther++
	}

	for {
		if currentPartI >= i.GetNumParts() {
			lastJoinedPart = getLastJoined(other, lastJoinedPart, currentPartOther, partLst)
			partLst.PushBack(lastJoinedPart)
			for j := currentPartOther + 1; j < other.GetNumParts(); j++ {
				partLst.PushBack(other.GetPart(j))
			}
			break
		}

		if currentPartOther >= other.GetNumParts() {
			lastJoinedPart = getLastJoined(i, lastJoinedPart, currentPartI, partLst)
			partLst.PushBack(lastJoinedPart)
			for j := currentPartI + 1; j < i.GetNumParts(); j++ {
				partLst.PushBack(i.GetPart(j))
			}
			break
		}

		if i.GetPart(currentPartI).start < other.GetPart(currentPartOther).start {
			lastJoinedPart = getLastJoined(i, lastJoinedPart, currentPartI, partLst)
			currentPartI++
		} else {
			lastJoinedPart = getLastJoined(other, lastJoinedPart, currentPartOther, partLst)
			currentPartOther++
		}
	}

	index := 0
	intervalParts := make([]*IntervalPart, partLst.Len())
	for element := partLst.Front(); element != nil; element = element.Next() {
		intervalParts[index] = element.Value.(*IntervalPart)
		index++
	}

	return newInternal(intervalParts)
}

func getLastJoined(mi *MultiInterval, joinedPart *IntervalPart, currentPart int, partLst *list.List) *IntervalPart {
	if mi.GetPart(currentPart).Intersects(joinedPart) || mi.GetPart(currentPart).touches(joinedPart) {
		return mergeParts(joinedPart, mi.GetPart(currentPart))
	} else {
		partLst.PushBack(joinedPart)
		return mi.GetPart(currentPart)
	}
}

func (p *IntervalPart) touches(other *IntervalPart) bool {
	return p.end == other.start || p.start == other.end
}

func mergeParts(part1 *IntervalPart, part2 *IntervalPart) *IntervalPart {
	s1 := part1.start
	s2 := part2.start
	e1 := part1.end
	e2 := part2.end

	if s1 < s2 {
		if e1 > e2 {
			return newPart(s1, e1)
		} else {
			return newPart(s1, e2)
		}
	} else {
		if e1 > e2 {
			return newPart(s2, e1)
		} else {
			return newPart(s2, e2)
		}
	}
}
