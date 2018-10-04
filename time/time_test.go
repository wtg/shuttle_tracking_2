package time

import (
	"testing"
	"time"
)

func TestCreateTime(t *testing.T) {
	timeTest, _ := time.Parse("15:04:05", "12:05:06")
	var t1 Time
	t1.FromTime(timeTest)
	if !(t1.GetTimeString() == "12:05:06") {
		t.Errorf("Time string did not match anticipated string")
	}
}

func TestTimeAfter(t *testing.T) {
	timeTest, _ := time.Parse("15:04:05", "12:05:06")
	timeTestTwo := timeTest.Add(1 * time.Minute)

	var t1 Time
	t1.FromTime(timeTest)
	var t2 Time
	t2.FromTime(timeTestTwo)
	var t3 Time
	t3.FromTime(timeTestTwo.Add(time.Hour))
	var t4 Time
	t4.FromTime(timeTest.Add(time.Second))
	if t1.After(t2) {
		t.Errorf("t2 should be after t1")
	}
	if !t2.After(t1) {
		t.Errorf("t1 should not be after t2")
	}
	if !t3.After(t1) {
		t.Errorf("t1 should not be after t3")
	}
	if !t4.After(t1) {
		t.Errorf("t1 should not be after t4")
	}
	if t1.After(t1) {
		t.Errorf("a time cannot be after itself")
	}

}

func TestTimeAfterWithDay(t *testing.T) {
	timeTest, _ := time.Parse("15:04:05", "12:05:06")
	timeTestTwo := timeTest.Add(1 * time.Minute)
	var t1 Time
	t1.FromTime(timeTest)
	var t2 Time
	t2.FromTime(timeTestTwo)
	t2.Day = 1
	t1.Day = 2
	if !t1.After(t2) {
		t.Errorf("t1 should be after t2")
	}
	if t2.After(t1) {
		t.Errorf("t2 should not be after t1")
	}
	if t1.After(t1) {
		t.Errorf("a time cannot be after itself")
	}

}

func TestSorting(t *testing.T) {
	timeTest, _ := time.Parse("15:04:05", "12:05:06")
	timeTestTwo := timeTest.Add(1 * time.Minute)
	timeTestThree := timeTest.Add(-1 * time.Minute)
	var t1 Time
	t1.FromTime(timeTest)
	var t2 Time
	t2.FromTime(timeTestTwo)
	var t3 Time
	t3.FromTime(timeTestThree)
	t2.Day = 1
	t1.Day = 2
	t3.Day = 1
	var times []Time

	times = append(times, t1, t2, t3)
	Sort(times)
	if times[0] != t3 {
		t.Error()
	}
	if times[1] != t2 {
		t.Error()
	}
	if times[2] != t1 {
		t.Error()
	}

	times = append(times, t3, t2, t1)
	Sort(times)
	if times[0] != t3 || times[1] != t3 {
		t.Error()
	}
	if times[2] != t2 || times[3] != t2 {
		t.Error()
	}
	if times[4] != t1 || times[5] != t1 {
		t.Error()
	}

	times = append(times, t3, t1, t2)
	Sort(times)
	if times[0] != t3 || times[1] != t3 || times[2] != t3 {
		t.Error()
	}
	if times[3] != t2 || times[4] != t2 || times[5] != t2 {
		t.Error()
	}
	if times[6] != t1 || times[7] != t1 || times[8] != t1 {
		t.Error()
	}
}
