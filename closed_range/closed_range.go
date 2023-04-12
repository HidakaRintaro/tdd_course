package closed_range

import (
	"errors"
	"fmt"
)

type (
	ClosedRange struct {
		lower int
		upper int
	}
)

func NewClosedRange(lower, upper int) (*ClosedRange, error) {
	if lower >= upper {
		return nil, errors.New("invalid range")
	}
	return &ClosedRange{
		lower: lower,
		upper: upper,
	}, nil
}

func (cr *ClosedRange) String() string {
	return fmt.Sprintf("[%d,%d]", cr.lower, cr.upper)
}

func (cr *ClosedRange) Contains(target int) bool {
	return cr.lower <= target && target <= cr.upper
}

func (cr *ClosedRange) Equals(target ClosedRange) bool {
	return cr.lower == target.lower && cr.upper == target.upper
}

func (cr *ClosedRange) RangeContains(target ClosedRange) bool {
	return cr.lower <= target.lower && target.upper <= cr.upper
}
