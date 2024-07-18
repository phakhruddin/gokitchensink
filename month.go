package main

func Season(month int) string {
	switch month {
	case 1, 2, 12:
		return "Winter"
	case 3, 4, 5:
		return "Springs"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Summer"
	default:
		return "unknown"
	}
	return ""
}

func main() {
	println(Season(20))
}
