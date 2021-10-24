package protocol

import (
	"github.com/rreubenreyes/gzclp/internal/core"
)

type Modifier string

const (
	RepMax Modifier = "REP_MAX"
	Select Modifier = "SELECT"
	AMRAP  Modifier = "AMRAP"
)

type ExerciseProgression struct {
	Weight    *core.Progression
	Sets      *core.Progression
	Reps      *core.Progression
	Modifiers []Modifier
}

type ExerciseProgressionOptions struct {
	StartingWeight  int
	WeightIncrement int
}

func (ep *ExerciseProgression) Append(next *ExerciseProgression) {
	ep.Weight.Append(next.Weight)
	ep.Sets.Append(next.Sets)
	ep.Reps.Append(next.Reps)
}

func (ep *ExerciseProgression) NextWeightIncrement() {
	ep.Weight.NextIncrement()
}

func (ep *ExerciseProgression) NextProgression(selectedWeight int) {
	nextWeight := ep.Weight.Value()
	for _, mod := range ep.Modifiers {
		if mod == RepMax || mod == Select {
			nextWeight = selectedWeight
			break
		}
	}
	ep.Weight, _ = ep.Weight.Next(nextWeight)
	ep.Sets, _ = ep.Sets.Next(nil)
	ep.Reps, _ = ep.Reps.Next(nil)
}

func startingWeightProgression(start int, increment int) *core.Progression {
	return core.NewProgression(&core.ProgressionOptions{
		Dimension:     core.Weight,
		Value:         start,
		Increment:     increment,
		NextDirective: core.Arbitrary,
	})
}

func ongoingWeightProgression(increment int) *core.Progression {
	return core.NewProgression(&core.ProgressionOptions{
		Dimension:     core.Weight,
		Value:         0,
		Increment:     increment,
		NextDirective: core.Arbitrary,
	})
}

func setsProgression(value int) *core.Progression {
	return core.NewProgression(&core.ProgressionOptions{
		Dimension: core.Sets,
		Value:     value,
	})
}

func repsProgression(value int) *core.Progression {
	return core.NewProgression(&core.ProgressionOptions{
		Dimension: core.Reps,
		Value:     value,
	})
}

func T1(opts *ExerciseProgressionOptions) *ExerciseProgression {
	firstStage := &ExerciseProgression{
		Weight: startingWeightProgression(opts.StartingWeight, opts.WeightIncrement),
		Sets:   setsProgression(5),
		Reps:   setsProgression(3),
	}
	secondStage := &ExerciseProgression{
		Weight: ongoingWeightProgression(opts.WeightIncrement),
		Sets:   setsProgression(6),
		Reps:   repsProgression(2),
	}
	thirdStage := &ExerciseProgression{
		Weight: ongoingWeightProgression(opts.WeightIncrement),
		Sets:   setsProgression(10),
		Reps:   repsProgression(1),
	}
	retest := &ExerciseProgression{
		Weight:    ongoingWeightProgression(opts.WeightIncrement),
		Sets:      setsProgression(1),
		Reps:      repsProgression(5),
		Modifiers: []Modifier{RepMax},
	}

	firstStage.Append(secondStage)
	secondStage.Append(thirdStage)
	thirdStage.Append(retest)
	retest.Append(firstStage)

	return firstStage
}

func T2(opts *ExerciseProgressionOptions) *ExerciseProgression {
	firstStage := &ExerciseProgression{
		Weight: startingWeightProgression(opts.StartingWeight, opts.WeightIncrement),
		Sets:   setsProgression(3),
		Reps:   setsProgression(10),
	}
	secondStage := &ExerciseProgression{
		Weight: ongoingWeightProgression(opts.WeightIncrement),
		Sets:   setsProgression(3),
		Reps:   repsProgression(8),
	}
	thirdStage := &ExerciseProgression{
		Weight:    ongoingWeightProgression(opts.WeightIncrement),
		Sets:      setsProgression(3),
		Reps:      repsProgression(6),
		Modifiers: []Modifier{Select},
	}

	firstStage.Append(secondStage)
	secondStage.Append(thirdStage)
	thirdStage.Append(firstStage)

	return firstStage
}
