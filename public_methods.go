package figure

import (
	"fmt"
	"io"
	"strings"
	"time"
)

//stdout
func (fig Figure) Print() {
	for _, printRow := range fig.Slicify() {
		fmt.Println(printRow)
	}
}

func (fig Figure) String() string {
	s := ""
	for _, printRow := range fig.Slicify() {
		s += fmt.Sprintf("%s\n", printRow)
	}
	return s
}

func (fig Figure) Scroll(duration, stillness int, direction string) {
	endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
	fig.phrase = fig.phrase + "   "
	clearScreen()
	for time.Now().Before(endTime) {
		var shiftedPhrase string
		chars := []byte(fig.phrase)
		if strings.HasPrefix(strings.ToLower(direction), "r") {
			shiftedPhrase = string(append(chars[len(chars)-1:], chars[0:len(chars)-1]...))
		} else {
			shiftedPhrase = string(append(chars[1:len(chars)], chars[0]))
		}
		fig.phrase = shiftedPhrase
		fig.Print()
		sleep(stillness)
		clearScreen()
	}
}

func (fig Figure) Blink(duration, timeOn, timeOff int) {
	if timeOff < 0 {
		timeOff = timeOn
	}
	endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
	clearScreen()
	for time.Now().Before(endTime) {
		fig.Print()
		sleep(timeOn)
		clearScreen()
		sleep(timeOff)
	}
}

func (fig Figure) Dance(duration, freeze int) {
	endTime := time.Now().Add(time.Duration(duration) * time.Millisecond)
	font := fig.font //TODO: change to deep copy
	font.evenLetters()
	figures := []Figure{Figure{font: font}, Figure{font: font}}
	clearScreen()
	for i, c := range fig.phrase {
		appenders := []string{" ", " "}
		appenders[i%2] = string(c)
		for f, _ := range figures {
			figures[f].phrase = figures[f].phrase + appenders[f]
		}
	}
	for p := 0; time.Now().Before(endTime); p ^= 1 {
		figures[p].Print()
		figures[1-p].Print()
		sleep(freeze)
		clearScreen()
	}
}

//writers
func Write(w io.Writer, fig Figure) {
	for _, printRow := range fig.Slicify() {
		fmt.Fprintf(w, "%v\n", printRow)
	}
}

//helpers
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
