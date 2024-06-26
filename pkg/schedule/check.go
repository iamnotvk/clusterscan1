package schedule

import "time"

func TimeUntilSchedule(schedule string) (time.Duration, error) {
	now := time.Now().UTC()
	layout := "2024-05-14T15:04:05Z"
	scheduledTime, err := time.Parse(layout, schedule)
	if err != nil {
		return time.Duration(0), err
	}

	return scheduledTime.Sub(now), nil
}
