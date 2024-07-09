package browser

import (
	"fmt"
	"time"
)

func timeDiff(futureTime time.Time) string {
	diff := futureTime.Sub(time.Now())
	totalSeconds := int(diff.Seconds())

	// Calculate days, hours, minutes, and seconds
	days := totalSeconds / (24 * 3600)
	totalSeconds %= 24 * 3600
	hours := totalSeconds / 3600
	totalSeconds %= 3600
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60

	timeDiffStr := fmt.Sprintf("%d days, %d hours and %02d:%02d", days, hours, minutes, seconds)
	return timeDiffStr
}

func checkTime() {
	time.Sleep(time.Second * 1)
	//t := time.Now().Add(time.Hour * 24 * 7)
	hardcodedTime := time.Date(2024, time.July, 15, 20, 54, 0, 0, time.UTC)

	for {
		newText := fmt.Sprintf("%s", timeDiff(hardcodedTime))
		Document.Id("voting-ends").Set("innerHTML", newText)
		hardcodedTime = hardcodedTime.Add(time.Second * -1)

		time.Sleep(time.Millisecond * 100)
	}
}
