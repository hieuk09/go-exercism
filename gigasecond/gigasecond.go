package gigasecond

import (
	"time"
)

const TestVersion = 2
const GigaSecond = time.Second * 1000000000

var Birthday, _ = time.Parse("09/01/2006", "09/06/1991")

func AddGigasecond(birthday time.Time) time.Time {
	return birthday.Add(GigaSecond)
}
