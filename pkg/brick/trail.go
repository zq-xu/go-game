package brick

import (
	"math"
	"math/rand"
)

func GenerateStableTrailFromAdd(baseNum, minNum, maxNum, speedFactor float64, startToAdd bool) func() float64 {
	num := baseNum
	return func() float64 {
		if num >= maxNum {
			startToAdd = false
		} else if num <= minNum {
			startToAdd = true
		}

		if startToAdd {
			num = num + speedFactor
			return num
		}

		num = num - speedFactor
		return num
	}
}

func GenerateRandomStableTrail(baseNum, minNum, maxNum, speedFactor float64) func() float64 {
	startToAdd := rand.Intn(2) == 0
	return GenerateStableTrailFromAdd(baseNum, minNum, maxNum, speedFactor, startToAdd)
}

func GenerateSinTrail(maxNum float64, timesToTop int) func() float64 {
	var i float64

	unit := math.Pi / float64(2*timesToTop)

	return func() float64 {
		i++
		return maxNum * math.Sin(i*unit)
	}
}

func GenerateRandomSinTrail(maxNum float64, timesToTop int) func() float64 {
	fn := GenerateSinTrail(maxNum, timesToTop)
	r := rand.Intn(2) == 0

	return func() float64 {
		if r {
			return fn()
		}

		return -fn()
	}
}

func GenerateRandomSinTrailWithBase(baseNum, maxNum float64, timesToTop int) func() float64 {
	fn := GenerateRandomSinTrail(maxNum, timesToTop)

	return func() float64 { return baseNum + fn() }
}

func GenerateSinTrailWithBase(baseNum, maxNum float64, timesToTop int) func() float64 {
	fn := GenerateSinTrail(maxNum, timesToTop)

	return func() float64 { return baseNum + fn() }
}

func GenerateReverseSinTrailWithBase(baseNum, maxNum float64, timesToTop int) func() float64 {
	fn := GenerateSinTrail(maxNum, timesToTop)

	return func() float64 { return baseNum - fn() }
}
