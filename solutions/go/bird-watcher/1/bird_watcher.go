package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birdCount int
	for i := 0; i < len(birdsPerDay); i++ {
		birdCount += birdsPerDay[i]
	}
	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdCount int
	weekStarts := (week - 1) * 7
	weekFinishes := weekStarts + 7

	// if starts in out of boundaries
	if weekStarts >= len(birdsPerDay) {
		return 0
	}
	// if finishes in out of boundaries, set it the last item of the week
	if weekFinishes > len(birdsPerDay) {
		weekFinishes = len(birdsPerDay)
	}

	weekBirds := birdsPerDay[weekStarts:weekFinishes]

	for i := 0; i < len(weekBirds); i++ {
		birdCount += weekBirds[i]
	}

	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
