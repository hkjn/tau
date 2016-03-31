// Package tau deals with units of time.
package tau

import "time"

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
		// Since returns the Tau that's currently passed since the instant.
		Since() Tau
		// AddTau returns the Time advanced by Tau.
		AddTau(Tau) Time
	}
	// ClockTime implements Time using time.Time.
	//
	// Note: This limits the largest time span to 290 years.
	ClockTime time.Time
)

// Since returns the Tau that's passed since the instant.
func (ct ClockTime) Since() Tau {
	return Tau(time.Since(time.Time(ct)) / 1e9)
}

// AddTau returns the ClockTime advanced by t Tau.
func (ct ClockTime) AddTau(t Tau) Time {
	return ClockTime(time.Time(ct).Add(time.Duration(t) * time.Second))
}

func (ct ClockTime) String() string {
	return time.Time(ct).String()
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

func (mt MegaTau) Tau() Tau { return Tau(mt * 1e6) }
func (gt GigaTau) Tau() Tau { return Tau(gt * 1e9) }
func (tt TeraTau) Tau() Tau { return Tau(tt * 1e12) }

// TauSince returns the Tau since given time.
func TauSince(t Time) Tau {
	return t.Since()
}
