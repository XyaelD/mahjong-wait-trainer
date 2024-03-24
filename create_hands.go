package main

import (
	"math/rand"
)

func createSet(suit string) ([]Tile, map[Tile]int, error) {
	mahjongSet := []Tile{}
	mahjongSetCount := make(map[Tile]int, 36)
	for i := 1; i < 10; i++ {
		mahjongSet = append(mahjongSet, Tile{Suit: suit, Rank: i})
		mahjongSetCount[Tile{Suit: suit, Rank: i}] = 4
	}
	return mahjongSet, mahjongSetCount, nil
}

func (app *application) add_pair(mahjongSet []Tile, current_hand []Tile) ([]Tile, error) {

	for {
		pairIdx := rand.Intn(len(mahjongSet))
		chosen_tile := mahjongSet[pairIdx]

		if app.mahjong_set_count[chosen_tile] > 1 {
			current_hand = append(current_hand, chosen_tile)
			current_hand = append(current_hand, chosen_tile)
			app.mahjong_set_count[chosen_tile] -= 2
			app.tiles_remaining -= 2
			break
		}
	}
	return current_hand, nil
}

func (app *application) add_triplet(mahjongSet []Tile, current_hand []Tile) ([]Tile, error) {
	for {
		tripleIdx := rand.Intn(len(mahjongSet))
		chosen_tile := mahjongSet[tripleIdx]

		if app.mahjong_set_count[chosen_tile] > 2 {
			current_hand = append(current_hand, chosen_tile)
			current_hand = append(current_hand, chosen_tile)
			current_hand = append(current_hand, chosen_tile)
			app.mahjong_set_count[chosen_tile] -= 3
			app.tiles_remaining -= 3
			break
		}
	}
	return current_hand, nil
}

func (app *application) add_unfinished_sequence(mahjongSet []Tile, current_hand []Tile) ([]Tile, error) {
	for {
		seqIdx := rand.Intn(len(mahjongSet))
		chosen_tile := mahjongSet[seqIdx]
		found_valid := false

		switch seqValue := chosen_tile.Rank; seqValue {
		case 6, 7, 8, 9:
			secondary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank - 1,
			}
			if app.mahjong_set_count[chosen_tile] != 0 && app.mahjong_set_count[secondary_tile] != 0 {
				current_hand = append(current_hand, chosen_tile)
				current_hand = append(current_hand, secondary_tile)
				app.mahjong_set_count[chosen_tile] -= 1
				app.mahjong_set_count[secondary_tile] -= 1
				app.tiles_remaining -= 2
				found_valid = true
			}
		default:
			secondary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank + 1,
			}
			if app.mahjong_set_count[chosen_tile] != 0 && app.mahjong_set_count[secondary_tile] != 0 {
				current_hand = append(current_hand, chosen_tile)
				current_hand = append(current_hand, secondary_tile)
				app.mahjong_set_count[chosen_tile] -= 1
				app.mahjong_set_count[secondary_tile] -= 1
				app.tiles_remaining -= 2
				found_valid = true
			}
		}
		if found_valid {
			break
		}
	}
	return current_hand, nil
}

func (app *application) add_finished_sequence(mahjongSet []Tile, current_hand []Tile) ([]Tile, error) {
	for {
		seqIdx := rand.Intn(len(mahjongSet))
		chosen_tile := mahjongSet[seqIdx]
		found_valid := false

		switch seqValue := chosen_tile.Rank; seqValue {
		case 6, 7, 8, 9:
			secondary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank - 1,
			}
			tertiary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank - 2,
			}
			if app.mahjong_set_count[chosen_tile] != 0 &&
				app.mahjong_set_count[secondary_tile] != 0 &&
				app.mahjong_set_count[tertiary_tile] != 0 {
				current_hand = append(current_hand, chosen_tile)
				current_hand = append(current_hand, secondary_tile)
				current_hand = append(current_hand, tertiary_tile)
				app.mahjong_set_count[chosen_tile] -= 1
				app.mahjong_set_count[secondary_tile] -= 1
				app.mahjong_set_count[tertiary_tile] -= 1
				app.tiles_remaining -= 3
				found_valid = true
			}
		default:
			secondary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank + 1,
			}
			tertiary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank + 2,
			}
			if app.mahjong_set_count[chosen_tile] != 0 &&
				app.mahjong_set_count[secondary_tile] != 0 &&
				app.mahjong_set_count[tertiary_tile] != 0 {
				current_hand = append(current_hand, chosen_tile)
				current_hand = append(current_hand, secondary_tile)
				current_hand = append(current_hand, tertiary_tile)
				app.mahjong_set_count[chosen_tile] -= 1
				app.mahjong_set_count[secondary_tile] -= 1
				app.mahjong_set_count[tertiary_tile] -= 1
				app.tiles_remaining -= 3
				found_valid = true
			}
		}
		if found_valid {
			break
		}
	}
	return current_hand, nil
}

func (app *application) add_wait_sequence(mahjongSet []Tile, current_hand []Tile) ([]Tile, error) {
	for {
		seqIdx := rand.Intn(len(mahjongSet))
		chosen_tile := mahjongSet[seqIdx]
		found_valid := false

		switch seqValue := chosen_tile.Rank; seqValue {
		case 6, 7, 8, 9:
			secondary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank - 1,
			}
			tertiary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank - 2,
			}
			quaternary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank - 3,
			}
			if app.mahjong_set_count[chosen_tile] != 0 &&
				app.mahjong_set_count[secondary_tile] != 0 &&
				app.mahjong_set_count[tertiary_tile] != 0 &&
				app.mahjong_set_count[quaternary_tile] != 0 {
				current_hand = append(current_hand, chosen_tile)
				current_hand = append(current_hand, secondary_tile)
				current_hand = append(current_hand, tertiary_tile)
				current_hand = append(current_hand, quaternary_tile)
				app.mahjong_set_count[chosen_tile] -= 1
				app.mahjong_set_count[secondary_tile] -= 1
				app.mahjong_set_count[tertiary_tile] -= 1
				app.mahjong_set_count[quaternary_tile] -= 1
				app.tiles_remaining -= 4
				found_valid = true
			}
		default:
			secondary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank + 1,
			}
			tertiary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank + 2,
			}
			quaternary_tile := Tile{
				Suit: app.suit,
				Rank: chosen_tile.Rank + 3,
			}
			if app.mahjong_set_count[chosen_tile] != 0 &&
				app.mahjong_set_count[secondary_tile] != 0 &&
				app.mahjong_set_count[tertiary_tile] != 0 &&
				app.mahjong_set_count[quaternary_tile] != 0 {
				current_hand = append(current_hand, chosen_tile)
				current_hand = append(current_hand, secondary_tile)
				current_hand = append(current_hand, tertiary_tile)
				current_hand = append(current_hand, quaternary_tile)
				app.mahjong_set_count[chosen_tile] -= 1
				app.mahjong_set_count[secondary_tile] -= 1
				app.mahjong_set_count[tertiary_tile] -= 1
				app.mahjong_set_count[quaternary_tile] -= 1
				app.tiles_remaining -= 4
				found_valid = true
			}
		}
		if found_valid {
			break
		}
	}
	return current_hand, nil
}

func (app *application) handle_hand(hand []Tile) []Tile {

	handle_pair := rand.Intn(2)
	switch handle_pair {
	case 0:
		hand, _ = app.add_pair(app.mahjong_set, hand)
	case 1:
		hand, _ = app.add_wait_sequence(app.mahjong_set, hand)
	}

	for app.tiles_remaining > 0 {
		switch {
		case app.tiles_remaining >= 3:
			choose_permutation := rand.Intn(2)
			switch choose_permutation {
			case 0:
				hand, _ = app.add_triplet(app.mahjong_set, hand)
			case 1:
				hand, _ = app.add_finished_sequence(app.mahjong_set, hand)
			}
		default:
			choose_permutation := rand.Intn(2)
			switch choose_permutation {
			case 0:
				hand, _ = app.add_pair(app.mahjong_set, hand)
			case 1:
				hand, _ = app.add_unfinished_sequence(app.mahjong_set, hand)
			}
		}
	}
	return hand
}

func (app *application) create_hand() ([]Tile, error) {
	finished_hand := []Tile{}
	finished_hand = app.handle_hand(finished_hand)
	return finished_hand, nil
}
