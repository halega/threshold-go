package threshold

import (
	"sort"
)

// Threshold implements Go version of ThresholdAlerter.
type Threshold struct {
	thresholds            []int
	alertHandler          func(int)
	reachedThresholdIndex int
}

// Check validates if a given value reached a new threshold and
// calls handler if it is.
func (t *Threshold) Check(value int) {
	index := t.indexOfMaxReachedThreshold(value)
	if index > t.reachedThresholdIndex && t.alertHandler != nil {
		t.alertHandler(t.thresholds[index])
	}
	t.reachedThresholdIndex = index
}

func (t *Threshold) indexOfMaxReachedThreshold(value int) int {
	i := sort.SearchInts(t.thresholds, value)
	if i == len(t.thresholds) || t.thresholds[i] != value {
		i--
	}
	return i
}

// NewThreshold returns a new instance of Threshold.
func NewThreshold(thresholds []int, alertHandler func(int)) *Threshold {
	if thresholds == nil {
		// TODO: ???
	}
	t := make([]int, len(thresholds))
	copy(t, thresholds)
	t = sortAndDistinct(t)
	return &Threshold{t, alertHandler, -1}
}

func sortAndDistinct(t []int) []int {
	sort.Ints(t)
	j := 0
	for i := 1; i < len(t); i++ {
		if t[j] == t[i] {
			continue
		}
		j++
		t[i], t[j] = t[j], t[i]
	}
	return t[:j+1]
}
