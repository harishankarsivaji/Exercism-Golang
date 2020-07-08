package gigasecond

import "time"

// GIGASECOND is 10^9 (1,000,000,000) seconds
const GIGASECOND = time.Duration(1e9) * time.Second

// AddGigasecond adds 1 gigasecond to the date.
func AddGigasecond(t time.Time) time.Time {

	return t.Add(GIGASECOND)
}
