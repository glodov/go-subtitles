package main

import (
	"fmt"

	"github.com/glodov/go-subtitles/greet"
)

func main() {
	fmt.Println(greet.MORNING)
	// lines := loadSubtitles("captions.ru.sbv")
	// for _, line := range lines[:5] {
	// 	fmt.Println(line.format())
	// }

	// fmt.Println(lines[:5])
	// b, err := ioutil.ReadAll(file)
	// fmt.Println(b)

	// fmt.Println(Subtitle{TimeFrom: "00:00:01", TimeTo: "00:05:20"})
}
