package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type application struct {
	logger            *slog.Logger
	suit              string
	tiles_remaining   int
	mahjong_set       []Tile
	mahjong_set_count map[Tile]int
}

type Tile struct {
	Suit string
	Rank int
}

type HandToGuess struct {
	winning_tiles []Tile
	hand          []Tile
}

// Get the suit and hand size the user wants to use from a valid list
func user_options() (string, int, error) {
	suits := []string{"Bamboo", "Characters", "Dots"}
	hand_sizes := []int{4, 7, 10, 13}

	var user_suit string
	var user_hand_size int
	for {
		valid_suit := false
		fmt.Println("Which suit do you want to practice with?")
		fmt.Scanln(&user_suit)
		for _, suit := range suits {
			if user_suit == suit {
				valid_suit = true
				break
			}
		}
		if valid_suit {
			break
		}
		fmt.Println("The valid suits are: Bamboo, Characters or Dots. Please try again.")
	}
	for {
		valid_size := false
		fmt.Println("What hand size to you want to practice with?")
		fmt.Scanln(&user_hand_size)
		for _, hand_size := range hand_sizes {
			if user_hand_size == hand_size {
				valid_size = true
				break
			}
		}
		if valid_size {
			break
		}
		fmt.Println("The valid hand sizes for the trainer are: 4, 7, 10 and 13. Please try again.")
	}
	return user_suit, user_hand_size, nil
}

// Initialize the program
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	suit, hand_size, err := user_options()
	if err != nil {
		logger.Error(err.Error())
	}
	mahjong_set, mahjong_set_count, err := createSet(suit)
	if err != nil {
		logger.Error(err.Error())
	}
	app := &application{
		suit:              suit,
		logger:            logger,
		tiles_remaining:   hand_size,
		mahjong_set:       mahjong_set,
		mahjong_set_count: mahjong_set_count,
	}

	// Keep playing until user chooses to exit program
	still_playing := true
	for still_playing {
		hand_to_guess := app.setup_hand()
		correct := app.get_user_guesses(hand_to_guess)

		if correct {
			fmt.Printf("Correct! The winning tiles were: %v\n", hand_to_guess.winning_tiles)
		} else {
			fmt.Printf("Incorrect. The winning tiles were: %v\n", hand_to_guess.winning_tiles)
		}
		var keep_playing_response string
		fmt.Println("Try another hand? (y/n)")
		fmt.Scanln(&keep_playing_response)
		switch strings.ToLower(keep_playing_response) {
		//Reset the app's values if an additional hand is to be played
		case "y":
			app.tiles_remaining = hand_size
			app.mahjong_set = mahjong_set
			app.mahjong_set_count = mahjong_set_count
		case "n":
			still_playing = false
			fmt.Println("Thanks for playing!")
		default:
			fmt.Println("Please input 'y' or 'n'")
		}
	}
	os.Exit(0)
}
