package convert

var entry = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}

func ConvertNumberToWords(number int) string {
	if number < 0 {
		return `error`
	}
	if number <= 20 {
		return entry[number]
	}
	return "error"
}
