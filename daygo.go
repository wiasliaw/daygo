package daygo

import "time"

// Daygo is the main time/date utilities.
type Daygo struct {
	Unix int64
}

// NewByDefault return a pointer of Daygo with time.Now().
func (dg *Daygo) NewByDefault() *Daygo {
	t := time.Now()
	dg = &Daygo{
		t.Unix(),
	}
	return dg
}
