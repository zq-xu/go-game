package brick

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// go test -v trail.go trail_test.go -test.run TestGenerateinTrail -count=1
func TestGenerateinTrail(t *testing.T) {
	Convey("TestGenerateinTrail", t, func() {
		maxNum, timeToTop := 500, 5
		numList := make([]float64, timeToTop)

		fn := GenerateSinTrail(float64(maxNum), timeToTop)
		for i := 0; i < timeToTop; i++ {
			numList[i] = fn()
		}

		Printf("%+v", numList)
	})
}
