package tau

import (
	"log"
	"testing"
	"time"
)

type TestTime struct {
	time.Time
	current time.Time
}

func (t TestTime) Since() Tau {
	return Tau(t.current.Sub(t.Time) / 1e9)
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
