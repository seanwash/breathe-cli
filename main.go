package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func clearPreviousLine() {
	// Use ANSI escape codes to control the terminal's cursor. Note that
	// this won't work when running via Idea's Run pane.
	//
	// Reference: https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797
	fmt.Print("\033[A\033[K")
}

func printMessage(message string, duration int) {
	clearPreviousLine()

	for i := 0; i < duration; i++ {
		periods := strings.Repeat(".", i)
		fmt.Printf("\r%s%s", message, periods)
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
}

func main() {
	durationInMinutes := flag.Int("duration", 1, "Duration (in minutes) of the exercise")
	flag.Parse()

	beginDurationInSeconds := 5
	cycleDurationInSeconds := 10
	exerciseDurationInSeconds := (*durationInMinutes * 60) / cycleDurationInSeconds

	printMessage("First, be still and bring your attention to your breath", beginDurationInSeconds)

	for i := 0; i < exerciseDurationInSeconds; i++ {
		printMessage("Breath in", cycleDurationInSeconds/2)
		printMessage("And breath out", cycleDurationInSeconds/2)
	}

	printMessage("Take a moment to enjoy this sense of calm. Remember, your mind is a tool; regularly tuning it makes every part of your day better. See you next time!", 1)
}
