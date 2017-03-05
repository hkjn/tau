package tau

import (
	"fmt"
	"log"
	"testing"
	"time"
)

type TestTime struct {
	time.Time
	current time.Time
}

func (tt TestTime) Since() Tau {
	return Tau(tt.current.Sub(tt.Time) / 1e9)
}

func (tt TestTime) AddTau(t Tau) Time {
	return TestTime{
		Time:    tt.Time.Add(time.Duration(t) * time.Second),
		current: tt.current,
	}
}

// newTestTime creates a deterministic TestTime for given value and current time.
func newTestTime(current, value string) TestTime {
	layout := "2006-01-02 15:04" // a simple time layout
	parse := func(s string) time.Time {
		t, err := time.Parse(layout, s)
		if err != nil {
			log.Fatalf("got err: %v\n", err)
		}
		return t
	}
	return TestTime{
		Time:    parse(value),
		current: parse(current),
	}
}

func TestTau(t *testing.T) {
	cases := []struct {
		in   Time
		want Tau
	}{
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-29 12:00"),
			want: Tau(0),
		},
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-17 22:14"),
			want: Tau(999960),
		},
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-17 22:13"),
			want: Tau(1000020),
		},
		{
			in:   newTestTime("2016-11-26 16:47", "1985-03-20 15:00"),
			want: Tau(1000000020),
		},
		{
			in:   newTestTime("2019-03-10 18:34", "1955-10-24 15:00"),
			want: Tau(2000000040),
		},
	}

	for i, tt := range cases {
		out := TauSince(tt.in)
		if tt.want != out {
			t.Errorf("[%d] TauSince(%v) => %v; want %v\n", i, tt.in, out, tt.want)
		}
	}
}

func TestMegaTau(t *testing.T) {
	cases := []struct {
		in   Time
		want MegaTau
	}{
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-29 12:00"),
			want: MegaTau(0),
		},
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-17 22:14"),
			want: MegaTau(0),
		},
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-17 22:13"),
			want: MegaTau(1),
		},
		{
			in:   newTestTime("2016-11-26 16:46", "1985-03-20 15:00"),
			want: MegaTau(999),
		},
		{
			in:   newTestTime("2015-05-09 09:00", "1985-03-20 15:00"),
			want: MegaTau(950),
		},
		{
			in:   newTestTime("2016-11-26 16:47", "1985-03-20 15:00"),
			want: MegaTau(1000),
		},
		{
			in:   newTestTime("2016-06-18 22:00", "1955-10-24 15:00"),
			want: MegaTau(1914),
		},
		{
			in:   newTestTime("2016-06-18 22:00", "1987-07-19 03:00"),
			want: MegaTau(912),
		},
		{
			in:   newTestTime("2016-06-18 22:00", "1990-01-04 21:00"),
			want: MegaTau(834),
		},
		{
			in:   newTestTime("2016-06-18 22:00", "1992-03-14 15:00"),
			want: MegaTau(765),
		},
		{
			in:   newTestTime("2016-06-18 22:00", "2014-06-18 22:00"),
			want: MegaTau(63),
		},
		{
			in:   newTestTime("2018-06-30 12:00", "1985-03-20 15:00"),
			want: MegaTau(1050),
		},
		{
			in:   newTestTime("2014-12-03 17:47", "1983-03-27 15:00"),
			want: MegaTau(1000),
		},
		{
			in:   newTestTime("2016-07-04 12:00", "1983-03-27 15:00"),
			want: MegaTau(1050),
		},
		{
			in:   newTestTime("2018-02-03 12:00", "1983-03-27 15:00"),
			want: MegaTau(1100),
		},
		{
			in:   newTestTime("2015-06-26 15:00", "1985-03-20 15:00"),
			want: MegaTau(955),
		},
		{
			// TODO(hkjn): This is the largest time span that can be
			// represented with time.Duration. Add more test cases once this
			// limitation is removed.
			in:   newTestTime("2277-06-25 07:26", "1985-03-20 15:00"),
			want: MegaTau(9222),
		},
	}
	for i, tt := range cases {
		out := TauSince(tt.in).Mega()
		if tt.want != out {
			t.Errorf("[%d] TauSince(%v).Mega() => %v; want %v\n", i, tt.in, out, tt.want)
		}
	}
}

func TestGigaTau(t *testing.T) {
	cases := []struct {
		in   Time
		want GigaTau
	}{
		{
			in:   newTestTime("2015-06-29 12:00", "2015-06-29 12:00"),
			want: GigaTau(0),
		},
		{
			in:   newTestTime("2016-11-26 16:46", "1985-03-20 15:00"),
			want: GigaTau(0),
		},
		{
			in:   newTestTime("2016-11-26 16:47", "1985-03-20 15:00"),
			want: GigaTau(1),
		},
		{
			in:   newTestTime("2019-03-10 19:00", "1955-10-24 15:00"),
			want: GigaTau(2),
		},
		{
			in:   newTestTime("2019-03-10 18:34", "1955-10-24 15:00"),
			want: GigaTau(2),
		},
		{
			in:   newTestTime("2277-06-25 07:26", "1985-03-20 15:00"),
			want: GigaTau(9),
		},
	}
	for i, tt := range cases {
		out := TauSince(tt.in).Giga()
		if tt.want != out {
			t.Errorf("[%d] TauSince(%v).Giga() => %v; want %v\n", i, tt.in, out, tt.want)
		}
	}
}

func TestAddTau(t *testing.T) {
	type in struct {
		time Time
		tau  Tau
	}
	cases := []struct {
		in   in
		want Tau
	}{
		{
			in: in{
				time: newTestTime("2015-06-29 12:00", "2015-06-29 12:00"),
				tau:  Tau(0),
			},
			want: Tau(0),
		},
		{
			in: in{
				time: newTestTime("2015-06-29 12:00", "2015-06-29 11:59"),
				tau:  Tau(0),
			},
			want: Tau(60),
		},
		{
			in: in{
				time: newTestTime("2015-06-29 12:00", "2015-06-29 12:01"),
				tau:  Tau(0),
			},
			want: Tau(-60),
		},
		{
			in: in{
				time: newTestTime("2016-11-26 16:47", "1985-03-20 15:00"),
				tau:  Tau(0),
			},
			want: Tau(1000000020),
		},
		{
			in: in{
				time: newTestTime("2016-11-26 16:47", "1985-03-20 15:00"),
				tau:  Tau(1000),
			},
			want: Tau(999999020),
		},
		{
			in: in{
				time: newTestTime("2016-11-26 16:47", "2017-01-01 12:00"),
				tau:  Tau(0),
			},
			want: Tau(-3093180),
		},
		{
			in: in{
				time: newTestTime("2016-03-30 06:37", "2016-03-30 06:37"),
				tau:  Tau(1000000),
			},
			want: Tau(-1000000),
		},
	}

	for i, tt := range cases {
		got := tt.in.time.AddTau(tt.in.tau).Since()
		if tt.want != got {
			t.Errorf("[%d] TauSince(%v).AddTau(%v).Since() => %v; want %v\n", i, tt.in.time, tt.in.tau, got, tt.want)
		}
	}
}

func ExampleClockTime_AddTau() {
	layout := "2006-01-02 15:04" // a simple time layout
	parse := func(s string) time.Time {
		t, err := time.Parse(layout, s)
		if err != nil {
			log.Fatalf("got err: %v\n", err)
		}
		return t
	}
	t0 := ClockTime(parse("1985-03-20 15:00"))
	fmt.Printf("Time t0 is %v\n", t0)
	for mt := 900; mt <= 1200; mt += 50 {
		t1 := t0.AddTau(MegaTau(mt).Tau())
		fmt.Printf("After %d mt: %v\n", mt, t1)
	}

	start := 1000
	t1 := t0.AddTau(MegaTau(start).Tau())
	fmt.Printf("Time t1 is %v\n", t1)
	for mt := 1; mt <= 15; mt += 1 {
		t2 := t1.AddTau(MegaTau(mt).Tau())
		fmt.Printf("[%d] After %d mt: %v\n", start+mt, mt, t2)
	}
	// Output: Time t0 is 1985-03-20 15:00:00 +0000 UTC
	// After 900 mt: 2013-09-26 07:00:00 +0000 UTC
	// After 950 mt: 2015-04-27 23:53:20 +0000 UTC
	// After 1000 mt: 2016-11-26 16:46:40 +0000 UTC
	// After 1050 mt: 2018-06-28 09:40:00 +0000 UTC
	// After 1100 mt: 2020-01-28 02:33:20 +0000 UTC
	// After 1150 mt: 2021-08-28 19:26:40 +0000 UTC
	// After 1200 mt: 2023-03-30 12:20:00 +0000 UTC
	// Time t1 is 2016-11-26 16:46:40 +0000 UTC
	// [1001] After 1 mt: 2016-12-08 06:33:20 +0000 UTC
	// [1002] After 2 mt: 2016-12-19 20:20:00 +0000 UTC
	// [1003] After 3 mt: 2016-12-31 10:06:40 +0000 UTC
	// [1004] After 4 mt: 2017-01-11 23:53:20 +0000 UTC
	// [1005] After 5 mt: 2017-01-23 13:40:00 +0000 UTC
	// [1006] After 6 mt: 2017-02-04 03:26:40 +0000 UTC
	// [1007] After 7 mt: 2017-02-15 17:13:20 +0000 UTC
	// [1008] After 8 mt: 2017-02-27 07:00:00 +0000 UTC
	// [1009] After 9 mt: 2017-03-10 20:46:40 +0000 UTC
	// [1010] After 10 mt: 2017-03-22 10:33:20 +0000 UTC
	// [1011] After 11 mt: 2017-04-03 00:20:00 +0000 UTC
	// [1012] After 12 mt: 2017-04-14 14:06:40 +0000 UTC
	// [1013] After 13 mt: 2017-04-26 03:53:20 +0000 UTC
	// [1014] After 14 mt: 2017-05-07 17:40:00 +0000 UTC
	// [1015] After 15 mt: 2017-05-19 07:26:40 +0000 UTC
}
