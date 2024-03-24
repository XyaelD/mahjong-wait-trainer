# Richii Mahjong Wait Trainer
## Description:

### What the program does:

This is a CLI program coded in Go to practice which tiles will make a winning hand in richii mahjong when one away from a valid hand. This program does not handle honor tiles nor Yakuman; it is meant to practice more usual sequences. 

### How to use:

In its current implementation, the user chooses a suit and a hand-size to consider then a one-away hand will be created of a specified size (4, 7, 10 or 13). The user then choose a rank (1-9) of tile that creates a winning hand before being prompted for another winning tile. When finished, the user will be informed if their choices are correct, and what the winning tiles are.

## Details:

### go_test.go

Contains all the tests written for the varying functions as the program was developed. 

### create_hands.go

Contains functions responsible for creating the set of tiles a one-away hand will be created from as well as the one-away hand.

### trainer.go

Contains functions responsible for setting up the one-away hand and getting the user's guess.

### validate.go

Contains functions which confirm the created one-away hand has winning tiles and does not contain illegal permutations.

### main.go 

Contains the function to get the user's options then initializes the program using them. Keeps playing until the user exits.
