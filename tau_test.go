package tau

import (
	"testing"
	"time"

	"hkjn.me/timeutils"
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
	return TestTime{
		Time:    timeutils.Must(timeutils.ParseStd(value)),
		current: timeutils.Must(timeutils.ParseStd(current)),
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
			in:   newTestTime("2016-11-26 16:47", "1985-03-20 15:00"),
			want: MegaTau(1000),
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
