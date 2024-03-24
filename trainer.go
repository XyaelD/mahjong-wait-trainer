package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func (app *application) setup_hand() *HandToGuess {
	hand, err := app.create_hand()
	if err != nil {
		app.logger.Error(err.Error())
	}
	hand_count, _ := app.tileCount(hand)
	if err != nil {
		app.logger.Error(err.Error())
	}
	winning_tiles, err := app.findWinningTiles(hand_count)
	if err != nil {
		app.logger.Error(err.Error())
	}

	hand_to_guess := &HandToGuess{
		hand:          hand,
		winning_tiles: winning_tiles,
	}
	return hand_to_guess
}

func (app *application) get_user_guesses(hand_to_guess *HandToGuess) bool {

	user_guess := []Tile{}
	guessing := true
	var confirmation string
	for guessing {
		tile_to_add := app.get_user_guess(hand_to_guess)
		user_guess = append(user_guess, tile_to_add)
		fmt.Println("Input another tile? (y/n)")
		fmt.Scanln(&confirmation)
		switch strings.ToLower(confirmation) {
		case "y":
			continue
		case "n":
			guessing = false
		default:
			fmt.Println("Please input 'y' or 'n'")
		}
	}
	sort.Slice(user_guess, func(i, j int) bool {
		return user_guess[i].Rank < user_guess[j].Rank
	})
	return slices.Equal(user_guess, hand_to_guess.winning_tiles)
}

func (app *application) get_user_guess(hand_to_guess *HandToGuess) Tile {
	for {
		var current_guess int
		fmt.Printf("Which is one of the ranks that win for this hand? %v\n", hand_to_guess.hand)
		fmt.Scanln(&current_guess)
		if 0 < current_guess && current_guess < 10 {
			return Tile{Suit: app.suit, Rank: current_guess}
		}
		fmt.Println("Please enter a valid tile from 1-9")
	}
}
