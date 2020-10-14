package main

func Burbuja(s []int64) {
	// Total elements
	total := len(s)
	var temp int64

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			if s[i] < s[j] {
				temp = s[i]
				s[i] = s[j]
				s[j] = temp
			}
		}
	}
}

func main() {

}
