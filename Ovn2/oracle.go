package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	// Answer questions.
	go answerQuestions(questions, answers)

	// Make prophecies.
	go makeRandomPredictions(answers)

	// Print answers.
	go printAnswer(answers)

	return questions
}

// Ansers questions by making a new prophecy
func answerQuestions(question chan string, answer chan string) {
	for message := range question { // for every question on the channel,
		go prophecy(message, answer) // make a new prophecy
	}
}

// Makes a prediction at random time.
func makeRandomPredictions(answer chan string) {
	for {
		time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)
		prophecy("", answer)
	}
}

// Prints the answer from channel where answers are sent.
func printAnswer(answer chan string) {
	for message := range answer {
		for _, c := range message {
			time.Sleep(time.Duration(180+rand.Intn(10)) * time.Millisecond) // Prints one character at the time.
			fmt.Print(string(c))
		}
		fmt.Print("\n" + prompt)
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.

	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	word := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.

	fortuneCookieMessages := []string{
		"run.",
		"I see money in your future... It's not yours though.",
		"You laugh now, wait till you get home.",
		"Yor pet is planing to eat you.",
		"You think it’s a secret, but they know.",
	}

	buddhistSayings := []string{
		"Speak only when you feel that your words are better than silence.",
		"Never stop learning, because life never stops teaching.",
		"Peace comes from within. Do not seek it without",
		"I'm not what you think I am. You are what you think I am.",
		"Be where you are; otherwise you will miss it",
		"To conquer oneself is a greater task than conquering others.",
		"The root of suffering is attachment",
		"You only lose what you cling to.",
		"Nothing rests in your mind",
		"Purity or impurity depends on oneself. No one can purify another",
	}

	buddhistAnswers := []string{
		"You learn nothing from life if you are right all the time.",
		"You are the problem, only YOU are the solution.",
		"Be with someone who is proud to have you",
		"Don't get upset with people and situations, because both are powerless without your reaction.",
		"Wear your ego like a loose fitting garment.",
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"The horse may converge",
		"One man's nonsense is another man's sense.",
		"To appreciate nonsense requires a serious interest in life.",
		"Nonsense, seems to sum up everything.",
	}

	// Questions keys
	questionKeys := map[string]struct{}{
		"why":   {},
		"which": {},
		"who":   {},
		"what":  {},
		"where": {},
		"when":  {},
		"how":   {},
		"whose": {},
	}

	// Answer algorithm
	for i, w := range words {

		w = strings.ToLower(w)		// 
		_, qok := questionKeys[w]	// Checks if w is in questionKeys map

		if qok {
			answer <- w + "?... " + buddhistAnswers[rand.Intn(len(buddhistAnswers))]
			break
		}
		if strings.ContainsRune(w, 36) { // Looks for '?' = 36
			select {
			case answer <- word + "... " + fortuneCookieMessages[rand.Intn(len(fortuneCookieMessages))]:
			case answer <- word + "... " + buddhistAnswers[rand.Intn(len(buddhistAnswers))]:
			}
		}
		if i == len(words)-1 {
			select {
			case answer <- word + "... " + fortuneCookieMessages[rand.Intn(len(fortuneCookieMessages))]:
			case answer <- word + "... " + buddhistSayings[rand.Intn(len(buddhistSayings))]:
			}
		}
	}

	if word == "" {
		select {
		case answer <- word + "... " + nonsense[rand.Intn(len(nonsense))]:
		case answer <- word + "... " + fortuneCookieMessages[rand.Intn(len(fortuneCookieMessages))]:
		case answer <- word + "... " + buddhistAnswers[rand.Intn(len(buddhistAnswers))]:
		case answer <- word + "... " + buddhistSayings[rand.Intn(len(buddhistSayings))]:
		}
	}

}
func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
