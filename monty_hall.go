package main

import (
	"fmt"
	"math/rand"
)

// import "time"

func main() {
	tries := 100
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
	if doors[pick] == true {
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
