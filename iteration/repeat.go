package iteration

const repeatedCount int = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += "a"
	}
	return repeated
}

