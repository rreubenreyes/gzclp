package core

import (
	"errors"
	"fmt"
)

type (
	NextDirective string
	Dimension     string
)

const (
	Weight    Dimension     = "WEIGHT"
	Sets      Dimension     = "SETS"
	Reps      Dimension     = "REPS"
	Arbitrary NextDirective = "ARBITRARY"
)

type Progression struct {
	dimension     Dimension
	value         int
	increment     int
	nextDirective NextDirective
	next          *Progression
}

// ProtocolOptions defines the necessary input to create a new Progression.
type ProgressionOptions struct {
	Dimension     Dimension     `json:"dimension"`
	Value         int           `json:"value"`
	Increment     int           `json:"increment_value"`
	NextDirective NextDirective `json:"next_directive"`
}

func NewProgression(opts *ProgressionOptions) *Progression {
	p := &Progression{
		dimension:     opts.Dimension,
		value:         opts.Value,
		increment:     opts.Increment,
		nextDirective: opts.NextDirective,
	}
	return p
}

func (p *Progression) Dimension() Dimension {
	return p.dimension
}

func (p *Progression) Value() int {
	return p.value
}

func (p *Progression) NextIncrement() {
  p.value += p.increment
}

func (p *Progression) Next(i interface{}) (*Progression, error) {
	if p.next == nil {
		return nil, nil
	}

	var msg string
	next := p.next
	if p.nextDirective == Arbitrary {
		switch v := i.(type) {
		case int:
			next.value = v
			return next, nil
		default:
			msg = fmt.Sprintf("%s directive expects an int", p.nextDirective)
			fmt.Println(msg)
			return nil, errors.New(msg)
		}
	}
	next.value = p.value + p.increment
	return next, nil
}

func (p *Progression) Append(next *Progression) {
  p.next = next
}
