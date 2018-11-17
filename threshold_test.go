package threshold

import "testing"

func TestAlertWhenReachedThreshold(t *testing.T) {
	alerted := false
	alerter := NewThreshold([]int{10}, func(threshold int) {
		alerted = true
		if threshold != 10 {
			t.Error("Expected alert threshold is", 10, "but was", threshold)
		}
	})
	alerter.Check(11)
	if !alerted {
		t.Error("Expected an alert but there wasn't any")
	}
}

func TestAlertOnceWhenPassedSeveralThresholds(t *testing.T) {
	alertedCount := 0
	alerter := NewThreshold([]int{10, 20}, func(threshold int) {
		alertedCount++
		if threshold != 20 {
			t.Error("Expected alert threshold is", 20, "but was", threshold)
		}
	})
	alerter.Check(30)
	if alertedCount != 1 {
		t.Error("Expected 1 alert but was", alertedCount)
	}
}
func TestNoAlertWhenNoReachedThreshold(t *testing.T) {
	alerted := false
	alerter := NewThreshold([]int{100}, func(threshold int) { alerted = true })
	alerter.Check(0)
	alerter.Check(50)
	alerter.Check(99)
	if alerted {
		t.Error("Expected no alert but it was")
	}
}
func TestAlertOnceOnReachedThreshold(t *testing.T) {
	alertedCount := 0
	alerter := NewThreshold([]int{10, 20}, func(threshold int) {
		alertedCount++
		if threshold != 10 {
			t.Error("Expected alert threshold is", 10, "but was", threshold)
		}
	})
	alerter.Check(11)
	alerter.Check(17)
	if alertedCount != 1 {
		t.Error("Expected 1 alert but was", alertedCount)
	}
}
func TestAlertWhenValueGoesDownAndUpThreshold(t *testing.T) {
	alertedCount := 0
	actualThreshold := 0
	alerter := NewThreshold([]int{10, 20}, func(threshold int) {
		alertedCount++
		actualThreshold = threshold
	})
	alerter.Check(30)
	if alertedCount != 1 {
		t.Error("Expected 1 alert but was", alertedCount)
	}
	if actualThreshold != 20 {
		t.Error("Expected alert threshold is", 20, "but was", actualThreshold)
	}
	alerter.Check(1)
	if alertedCount != 1 {
		t.Error("Expected 1 alert but was", alertedCount)
	}
	alerter.Check(15)
	if alertedCount != 2 {
		t.Error("Expected 2 alerts but was", alertedCount)
	}
	if actualThreshold != 10 {
		t.Error("Expected alert threshold is", 10, "but was", actualThreshold)
	}
}
func TestNoThresholds(t *testing.T) {
	t.Skip("FIX")
	alertedCount := 0
	alerter := NewThreshold([]int{}, func(threshold int) {
		alertedCount++
	})
	alerter.Check(10)
	if alertedCount != 0 {
		t.Error("Expected 0 alerts but was", alertedCount)
	}
}

func TestNoHandler(t *testing.T) {
	alerter := NewThreshold([]int{10}, nil)
	// No panic
	alerter.Check(10)
}

func TestNilThresholds(t *testing.T) {
	t.Skip("FIX")
	alerter := NewThreshold(nil, nil)
	alerter.Check(0)
}

func TestNoCallToNewThreshold(t *testing.T) {
	alerter := &Threshold{}
	alerter.Check(0)
}

func TestUnsortedThresholds(t *testing.T) {
	alerted := false
	alerter := NewThreshold([]int{20, 10}, func(threshold int) {
		alerted = true
		if threshold != 10 {
			t.Error("Expected threshold", 10, "but got", threshold)
		}
	})
	alerter.Check(10)
	if !alerted {
		t.Error("Handler expected to be calld but wasn't")
	}
}

func TestSortAndDistinct(t *testing.T) {
	in := []int{30, 20, 10, 40, 30, 20, 10, 40, 10}
	in = sortAndDistinct(in)
	if !sliceEq(in, []int{10, 20, 30, 40}) {
		t.Error("Expected slice value is", []int{10, 20, 30, 40}, "but was", in)
	}
}

func sliceEq(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
