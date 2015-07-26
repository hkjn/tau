// Package tau deals with units of time.
package tau

import (
	"time"

	"hkjn.me/timeutils"
)

type (
	// Tau represents τ, a duration in number of seconds.
	Tau int64
	// MegaTau represents Mτ, a duration in millions of τ.
	MegaTau int64
	// GigaTau represents Gτ, a duration in billions of τ.
	GigaTau int64
	// TeraTau represents Tτ, a duration in trillions of τ.
	TeraTau int64
	// Time represents an instant in time.
	Time interface {
		// Since returns the Tau that's passed since the instant.
		Since() Tau
	}
	// ClockTime implements Time using time.Time.
	ClockTime struct {
		// Note: This limits the largest time span to 290 years.
		time.Time
	}
)

// Since returns the Tau that's passed since the instant.
func (t ClockTime) Since() Tau {
	return Tau(time.Since(t.Time) / 1e9)
}

// newClockTime returns a new clock time from given value.
func newClockTime(value string) ClockTime {
	return ClockTime{timeutils.Must(timeutils.ParseStd(value))}
}

// Mega returns the MegaTau.
func (t Tau) Mega() MegaTau {
	return MegaTau(t / 1e6)
}

// Giga returns the GigaTau.
func (t Tau) Giga() GigaTau {
	return GigaTau(t / 1e9)
}

// Tera returns the TeraTau.
func (t Tau) Tera() TeraTau {
	return TeraTau(t / 1e12)
}

// TauSince returns the Tau since given time.
func TauSince(t Time) Tau {
	return t.Since()
}
