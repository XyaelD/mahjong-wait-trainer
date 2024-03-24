package main

import (
	"errors"
	"maps"
	"sort"
)

func (app *application) tileCount(hand []Tile) (map[Tile]int, error) {
	counts := make(map[Tile]int)

	for _, tile := range hand {
		counts[tile]++
		if counts[tile] > 4 {
			return nil, errors.New("cannot have more than 4 of the same tile")
		}
	}
	return counts, nil
}

func (app *application) checkTripsFirst(curr_tiles []Tile, tilesMinusPair map[Tile]int) int {
	for _, tile := range curr_tiles {
		// Check for a valid triplet first
		if tilesMinusPair[tile] > 2 {
			tilesMinusPair[tile] = tilesMinusPair[tile] - 3
		}

		second_sequence_tile := Tile{}
		third_sequence_tile := Tile{}
		// Check for valid sequence
		if tile.Rank <= 7 && tilesMinusPair[tile] > 0 {
			second_sequence_tile = Tile{tile.Suit, tile.Rank + 1}
			third_sequence_tile = Tile{tile.Suit, tile.Rank + 2}
		}

		if tilesMinusPair[second_sequence_tile] > 0 && tilesMinusPair[third_sequence_tile] > 0 {
			tilesMinusPair[tile]--
			tilesMinusPair[second_sequence_tile]--
			tilesMinusPair[third_sequence_tile]--
		}
	}

	tile_count := 0
	for _, v := range tilesMinusPair {
		tile_count += v
	}
	return tile_count
}

func (app *application) checkSeqFirst(curr_tiles []Tile, tilesMinusPair map[Tile]int) int {
	for _, tile := range curr_tiles {
		second_sequence_tile := Tile{}
		third_sequence_tile := Tile{}
		// Check for valid sequence first
		if tile.Rank <= 7 && tilesMinusPair[tile] > 0 {
			second_sequence_tile = Tile{tile.Suit, tile.Rank + 1}
			third_sequence_tile = Tile{tile.Suit, tile.Rank + 2}
		}

		if tilesMinusPair[second_sequence_tile] > 0 && tilesMinusPair[third_sequence_tile] > 0 {
			tilesMinusPair[tile]--
			tilesMinusPair[second_sequence_tile]--
			tilesMinusPair[third_sequence_tile]--
		}

		// Check for a valid triplet
		if tilesMinusPair[tile] > 2 {
			tilesMinusPair[tile] = tilesMinusPair[tile] - 3
		}
	}

	tile_count := 0
	for _, v := range tilesMinusPair {
		tile_count += v
	}
	return tile_count
}

func (app *application) checkSeqAndTrips(tilesMinusPair map[Tile]int) (bool, error) {

	var curr_tiles []Tile
	for key := range tilesMinusPair {
		curr_tiles = append(curr_tiles, key)
	}

	// Sort the current tiles so the sequence logic will work properly
	sort.Slice(curr_tiles, func(i, j int) bool {
		if curr_tiles[i].Suit != curr_tiles[j].Suit {
			return curr_tiles[i].Suit < curr_tiles[j].Suit
		}
		return curr_tiles[i].Rank < curr_tiles[j].Rank
	})

	tripFirstCount := app.checkTripsFirst(curr_tiles, tilesMinusPair)
	seqFirstCount := app.checkSeqFirst(curr_tiles, tilesMinusPair)
	// If no tiles remain, everything had a valid sequence or triplet
	return tripFirstCount == 0 || seqFirstCount == 0, nil
}

func alreadyWinningTile(winning_tiles []Tile, newTile Tile) bool {
	for _, winning_tile := range winning_tiles {
		if winning_tile == newTile {
			return true
		}
	}
	return false
}

func (app *application) findWinningTiles(tileCount map[Tile]int) ([]Tile, error) {
	winningTiles := []Tile{}

	for i := 1; i < 10; i++ {
		newTile := Tile{
			Suit: app.suit,
			Rank: i,
		}

		// Create a copy then add a tile to the hand and then check for a valid win
		copiedTiles := map[Tile]int{}
		maps.Copy(copiedTiles, tileCount)

		copiedTiles[newTile]++
		// Do not test if five of the same tile now exist in hand
		if copiedTiles[newTile] > 4 {
			continue
		}
		pairs := []Tile{}

		for key, value := range copiedTiles {
			if value > 1 {
				pairs = append(pairs, key)
			}
		}
		for _, pair := range pairs {
			copiedTiles[pair] = copiedTiles[pair] - 2
			// Copy the modified map so it runs correctly and not with more pairs than the original slice contained
			tilesMinusPair := map[Tile]int{}
			maps.Copy(tilesMinusPair, copiedTiles)
			result, err := app.checkSeqAndTrips(tilesMinusPair)
			if err != nil {
				app.logger.Error(err.Error())
			}

			// If the check returned true, and the tile does not already exist in winning tiles, then the added tile formed a winning hand
			if result && !alreadyWinningTile(winningTiles, newTile) {
				winningTiles = append(winningTiles, newTile)
			}
			copiedTiles[pair] = copiedTiles[pair] + 2
		}
	}
	return winningTiles, nil
}
