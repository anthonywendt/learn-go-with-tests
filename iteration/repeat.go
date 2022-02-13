package iteration

func Repeat(letter string, numberOfTimes int) string {

	var repeated string
	for i := 0; i < numberOfTimes; i++ {
		repeated += letter
	}
	return repeated
}
