package main

import (
	"maps"
	"slices"
	"testing"
)

// Test the tileCount function
func TestCount(t *testing.T) {
	app := application{}
	data := []Tile{
		{Suit: "Characters", Rank: 8},
		{Suit: "Bamboo", Rank: 2},
		{Suit: "Dots", Rank: 7},
		{Suit: "Bamboo", Rank: 3},
		{Suit: "Bamboo", Rank: 4},
		{Suit: "Dots", Rank: 8},
		{Suit: "Dots", Rank: 9},
		{Suit: "Bamboo", Rank: 4},
		{Suit: "Characters", Rank: 8},
		{Suit: "Characters", Rank: 8},
	}

	expected := map[Tile]int{
		{Suit: "Bamboo", Rank: 2}:     1,
		{Suit: "Bamboo", Rank: 3}:     1,
		{Suit: "Bamboo", Rank: 4}:     2,
		{Suit: "Characters", Rank: 8}: 3,
		{Suit: "Dots", Rank: 7}:       1,
		{Suit: "Dots", Rank: 8}:       1,
		{Suit: "Dots", Rank: 9}:       1,
	}

	result, err := app.tileCount(data)

	if !maps.Equal(result, expected) || err != nil {
		t.Fatalf("test failed for counting tiles %v\n the result was: %v", err, result)
	}
}

func TestCountTooMany(t *testing.T) {
	app := application{}
	data := []Tile{
		{Suit: "Characters", Rank: 8},
		{Suit: "Characters", Rank: 8},
		{Suit: "Characters", Rank: 8},
		{Suit: "Characters", Rank: 8},
		{Suit: "Characters", Rank: 8},
	}

	_, err := app.tileCount(data)

	// Check that the error is being hit in the tileCount function
	if err == nil {
		t.Fatalf("wrong error for this test: %v", err)
	}
}

// Test for triplets and sequences in the checkSeqAndTrips function;
// Also test the helper functions checkTripsFirst and checkSeqFirst

func TestCheckSeqAndTrips(t *testing.T) {
	app := application{}

	// Two sequences (out of order, but the function handles that) and a triplet, so it is valid
	data := map[Tile]int{
		{Suit: "Dots", Rank: 8}:       1,
		{Suit: "Bamboo", Rank: 2}:     1,
		{Suit: "Bamboo", Rank: 3}:     1,
		{Suit: "Bamboo", Rank: 1}:     1,
		{Suit: "Characters", Rank: 8}: 3,
		{Suit: "Dots", Rank: 7}:       1,
		{Suit: "Dots", Rank: 9}:       1,
	}

	result, err := app.checkSeqAndTrips(data)
	if !result || err != nil {
		t.Fatalf("sequence and triplets test failed: %v\n the result was: %v", err, result)
	}
}

func TestCheckSeqAndTripsFail(t *testing.T) {
	app := application{}

	// Two sequences (out of order, but the function handles that) and a triplet, so it is valid
	data := map[Tile]int{
		{Suit: "Bamboo", Rank: 2}:     1,
		{Suit: "Bamboo", Rank: 3}:     1,
		{Suit: "Bamboo", Rank: 1}:     1,
		{Suit: "Characters", Rank: 8}: 3,
		{Suit: "Dots", Rank: 7}:       1,
		{Suit: "Dots", Rank: 9}:       1,
		{Suit: "Dots", Rank: 9}:       1,
	}

	result, _ := app.checkSeqAndTrips(data)
	if result {
		t.Fatalf("returned true when it should have returned false")
	}
}

// Test findWinningTiles and its helper function alreadyWinningTile
func TestWinningTiles(t *testing.T) {
	app := application{
		suit: "Dots",
	}
	// Test a three-sided wait
	dataThree := map[Tile]int{
		{Suit: app.suit, Rank: 2}: 2,
		{Suit: app.suit, Rank: 4}: 1,
		{Suit: app.suit, Rank: 5}: 1,
		{Suit: app.suit, Rank: 6}: 1,
		{Suit: app.suit, Rank: 7}: 1,
		{Suit: app.suit, Rank: 8}: 1,
	}

	expectedThree := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 6},
		{Suit: app.suit, Rank: 9},
	}

	// Test a pair wait on a sequence of four
	dataFour := map[Tile]int{
		{Suit: app.suit, Rank: 4}: 1,
		{Suit: app.suit, Rank: 5}: 1,
		{Suit: app.suit, Rank: 6}: 1,
		{Suit: app.suit, Rank: 7}: 1,
	}

	expectedFour := []Tile{
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 7},
	}

	// Test a five-sided wait
	dataFive := map[Tile]int{
		{Suit: app.suit, Rank: 3}: 1,
		{Suit: app.suit, Rank: 4}: 1,
		{Suit: app.suit, Rank: 5}: 2,
		{Suit: app.suit, Rank: 6}: 1,
		{Suit: app.suit, Rank: 7}: 3,
		{Suit: app.suit, Rank: 8}: 2,
		{Suit: app.suit, Rank: 9}: 3,
	}

	expectedFive := []Tile{
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 6},
		{Suit: app.suit, Rank: 7},
		{Suit: app.suit, Rank: 8},
		{Suit: app.suit, Rank: 9},
	}

	// Test a result the would add the same tile twice if the helper function does not work
	dataMulti := map[Tile]int{
		{Suit: app.suit, Rank: 3}: 2,
		{Suit: app.suit, Rank: 6}: 1,
		{Suit: app.suit, Rank: 5}: 2,
		{Suit: app.suit, Rank: 4}: 2,
	}

	expectedMulti := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 6},
	}

	// Test a result the would run with five of the same tile if not caught and handled properly
	dataTooMany := map[Tile]int{
		{Suit: app.suit, Rank: 9}: 1,
		{Suit: app.suit, Rank: 6}: 4,
		{Suit: app.suit, Rank: 8}: 1,
		{Suit: app.suit, Rank: 7}: 1,
	}

	expectedTooMany := []Tile{
		{Suit: app.suit, Rank: 9},
	}

	result, err := app.findWinningTiles(dataThree)
	if !slices.Equal(result, expectedThree) || err != nil {
		t.Fatalf("there was an error finding the winning tiles: %v\n the result was: %v", err, result)
	}
	result, err = app.findWinningTiles(dataFour)
	if !slices.Equal(result, expectedFour) || err != nil {
		t.Fatalf("there was an error finding the winning tiles: %v\n the result was: %v", err, result)
	}

	result, err = app.findWinningTiles(dataFive)
	if !slices.Equal(result, expectedFive) || err != nil {
		t.Fatalf("there was an error finding the winning tiles: %v\n the result was: %v", err, result)
	}

	result, err = app.findWinningTiles(dataMulti)
	if !slices.Equal(result, expectedMulti) || err != nil {
		t.Fatalf("there was an error finding the winning tiles: %v\n the result was: %v", err, result)
	}

	result, err = app.findWinningTiles(dataTooMany)
	if !slices.Equal(result, expectedTooMany) || err != nil {
		t.Fatalf("there was an error finding the winning tiles: %v\n the result was: %v", err, result)
	}
}

// Test createSet
func TestCreateSet(t *testing.T) {

	suit := "Bamboo"

	expected := []Tile{
		{Suit: suit, Rank: 1},
		{Suit: suit, Rank: 2},
		{Suit: suit, Rank: 3},
		{Suit: suit, Rank: 4},
		{Suit: suit, Rank: 5},
		{Suit: suit, Rank: 6},
		{Suit: suit, Rank: 7},
		{Suit: suit, Rank: 8},
		{Suit: suit, Rank: 9},
	}

	expected_count := map[Tile]int{
		{Suit: suit, Rank: 1}: 4,
		{Suit: suit, Rank: 2}: 4,
		{Suit: suit, Rank: 3}: 4,
		{Suit: suit, Rank: 4}: 4,
		{Suit: suit, Rank: 5}: 4,
		{Suit: suit, Rank: 6}: 4,
		{Suit: suit, Rank: 7}: 4,
		{Suit: suit, Rank: 8}: 4,
		{Suit: suit, Rank: 9}: 4,
	}

	result, result_count, err := createSet(suit)
	if !slices.Equal(result, expected) || !maps.Equal(result_count, expected_count) || err != nil {
		t.Fatalf("there was an error creating a set: %v\n the result was: %v\n the result_count was: %v", err, result, result_count)
	}
}

// Test add_pair
func TestPair(t *testing.T) {

	count_data := map[Tile]int{
		{Suit: "Bamboo", Rank: 3}: 2,
		{Suit: "Bamboo", Rank: 4}: 1,
		{Suit: "Bamboo", Rank: 5}: 1,
	}

	app := application{
		suit:              "Bamboo",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 5},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
	}

	result := []Tile{
		{Suit: app.suit, Rank: 1},
	}
	result, err := app.add_pair(data, result)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a pair: %v\n the result was: %v", err, result)
	}
}

// Test add_triplet
func TestTriplet(t *testing.T) {

	count_data := map[Tile]int{
		{Suit: "Characters", Rank: 5}: 3,
		{Suit: "Characters", Rank: 2}: 2,
		{Suit: "Characters", Rank: 7}: 1,
	}

	app := application{
		suit:              "Characters",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 7},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
	}
	result, err := app.add_triplet(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

// Test add_unfinished_sequence for both switch cases
func TestUnfinishedSeqHigh(t *testing.T) {

	count_data := map[Tile]int{
		{Suit: "Characters", Rank: 9}: 3,
		{Suit: "Characters", Rank: 8}: 2,
	}

	app := application{
		suit:              "Characters",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 9},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 9},
		{Suit: app.suit, Rank: 8},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
	}
	result, err := app.add_unfinished_sequence(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

func TestUnfinishedSeqLow(t *testing.T) {

	count_data := map[Tile]int{
		{Suit: "Dots", Rank: 4}: 4,
		{Suit: "Dots", Rank: 5}: 1,
	}

	app := application{
		suit:              "Dots",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 4},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 5},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 5},
	}
	result, err := app.add_unfinished_sequence(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

// Test add_finished_sequence for both switch cases
func TestFinishedSeqLow(t *testing.T) {

	count_data := map[Tile]int{
		{Suit: "Dots", Rank: 1}: 2,
		{Suit: "Dots", Rank: 2}: 3,
		{Suit: "Dots", Rank: 3}: 1,
	}

	app := application{
		suit:              "Dots",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 1},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
	}
	result, err := app.add_finished_sequence(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

func TestFinishedSeqHigh(t *testing.T) {

	count_data := map[Tile]int{
		{Suit: "Bamboo", Rank: 8}: 2,
		{Suit: "Bamboo", Rank: 7}: 1,
		{Suit: "Bamboo", Rank: 9}: 4,
	}

	app := application{
		suit:              "Bamboo",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 9},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 9},
		{Suit: app.suit, Rank: 8},
		{Suit: app.suit, Rank: 7},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
	}
	result, err := app.add_finished_sequence(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

// Test add_wait_sequence for both switch cases
func TestWaitSequenceHigh(t *testing.T) {
	count_data := map[Tile]int{
		{Suit: "Bamboo", Rank: 6}: 2,
		{Suit: "Bamboo", Rank: 7}: 1,
		{Suit: "Bamboo", Rank: 5}: 3,
		{Suit: "Bamboo", Rank: 4}: 4,
	}

	app := application{
		suit:              "Bamboo",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 7},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 7},
		{Suit: app.suit, Rank: 6},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 4},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
	}
	result, err := app.add_wait_sequence(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

func TestWaitSequenceLow(t *testing.T) {
	count_data := map[Tile]int{
		{Suit: "Characters", Rank: 3}: 2,
		{Suit: "Characters", Rank: 5}: 1,
		{Suit: "Characters", Rank: 6}: 3,
		{Suit: "Characters", Rank: 4}: 4,
	}

	app := application{
		suit:              "Characters",
		mahjong_set_count: count_data,
	}

	data := []Tile{
		{Suit: app.suit, Rank: 3},
	}

	expected := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 3},
		{Suit: app.suit, Rank: 4},
		{Suit: app.suit, Rank: 5},
		{Suit: app.suit, Rank: 6},
	}

	current_hand := []Tile{
		{Suit: app.suit, Rank: 1},
		{Suit: app.suit, Rank: 2},
		{Suit: app.suit, Rank: 3},
	}
	result, err := app.add_wait_sequence(data, current_hand)
	if !slices.Equal(result, expected) || err != nil {
		t.Fatalf("there was an error adding a triplet: %v\n the result was: %v", err, result)
	}
}

// Test create_hand and the helper handle_hand
func TestCreateHand(t *testing.T) {

	suit := "Dots"

	mahjong_set, mahjong_set_count, _ := createSet(suit)

	app := application{
		suit:              suit,
		tiles_remaining:   13,
		mahjong_set:       mahjong_set,
		mahjong_set_count: mahjong_set_count,
	}

	_, err := app.create_hand()
	if err != nil || app.tiles_remaining != 0 {
		t.Fatalf("there was an error creating a valid hand: %v\n the remaining tiles: %v", err, app.tiles_remaining)
	}
}

// Test everything together at maximum tile count
func TestApp(t *testing.T) {
	suit := "Bamboo"

	mahjong_set, mahjong_set_count, _ := createSet(suit)

	app := application{
		suit:              suit,
		tiles_remaining:   13,
		mahjong_set:       mahjong_set,
		mahjong_set_count: mahjong_set_count,
	}

	hand, _ := app.create_hand()
	hand_count, _ := app.tileCount(hand)
	winning_tiles, err := app.findWinningTiles(hand_count)
	if err != nil || winning_tiles == nil {
		t.Fatalf("a hand was generated with no winning tiles: %v\n the remaining tiles: %v\n the hand: %v\n the hand_count: %v\n winning tiles: %v\n", err, app.tiles_remaining, hand, hand_count, winning_tiles)
	}
}
