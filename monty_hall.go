package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	displayOptions()

	//Getting input parameters from user:
	tries, err := receiveInput()
	if err != nil {
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	switchTotal := 0
	for i := 0; i < tries; i++ {
		success, err := simulateGame(true)
		if err != nil {
			fmt.Printf("simulation failure: %v\n", err)
		}
		if success {
			switchTotal++
		}
	}

	stayTotal := 0
	for i := 0; i < tries; i++ {
		success, err := simulateGame(false)
		if err != nil {
			fmt.Printf("simulation failure: %v\n", err)
		}
		if success {
			stayTotal++
		}
	}
	fmt.Printf("Switch success ratio: %v / %v\n", switchTotal, tries)
	fmt.Printf("Stay success ratio: %v / %v\n", stayTotal, tries)
}

func displayOptions() {
	fmt.Print("\nWelcome to the Monty Hall simulator!\nThis simulator takes in a single integer (N) and tries the monty hall problem with switching between doors and not switching between doors N times.\nIt then will return the result.\n\nPlease insert an integer and press [Enter]. (Please do not provide any additional whitespace.)\n->")
}

func receiveInput() (int, error) {
	scanner := bufio.NewReader(os.Stdin)
	returnLine, err := scanner.ReadString('\n')
	returnLine = strings.TrimSuffix(returnLine, "\r\n")

	if err != nil {
		fmt.Printf("receiving input error: %v", err)
	}

	tries, err := strconv.Atoi(returnLine)

	if err != nil {
		fmt.Printf("conversion from user input failed: %v\n", err)
	}
	return tries, err
}

func simulateGame(playerSwitch bool) (bool, error) {
	doors, err := createDoors()
	if err != nil {
		fmt.Printf("error creating doors: %v\n", err)
	}
	playerDoor, err := playerPickFirstDoor()
	if err != nil {
		fmt.Printf("error player picking door: %v\n", err)
	}
	revealDoor, err := montyRemovesDoor(doors, playerDoor)
	if err != nil {
		fmt.Printf("error removing door: %v\n", err)
	}

	//Does Player switch?
	if playerSwitch {
		for i := 0; i < len(doors); i++ {
			if i != revealDoor && i != playerDoor {
				playerDoor = i
				break
			}
		}
	}

	success := doors[playerDoor]

	return success, nil
}

func createDoors() ([3]bool, error) {
	var doors [3]bool
	doors[(rand.Int() % 3)] = true
	return doors, nil
}

func playerPickFirstDoor() (int, error) {
	return rand.Int() % 3, nil
}

func montyRemovesDoor(doors [3]bool, pick int) (int, error) {
	revealDoor := 0
	if doors[pick] {
		//pick a random door to reveal
		for {
			revealDoor = rand.Int() % 3
			if revealDoor != pick {
				break
			}
		}
	} else {
		//pick the only other door to reveal
		for i := 0; i < len(doors); i++ {
			if !doors[i] && pick != i {
				revealDoor = i
				break
			}
		}
	}
	return revealDoor, nil
}
