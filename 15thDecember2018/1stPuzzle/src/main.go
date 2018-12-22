package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type character struct {
	position coordinate
	hitpoints int
	attackPower int
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getCharactersAndLand(lines []string) (elves []character, goblins []character, walkableLand []coordinate, maxX int, maxY int) {
	maxX = len(lines[0]) - 1
	maxY = len(lines) - 1
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 'G' {
				goblins = append(goblins, character{coordinate{x, y}, 200, 3})
				walkableLand = append(walkableLand, coordinate{x,y})
			} else if lines[y][x] == 'E' {
				elves = append(elves, character{coordinate{x, y}, 200, 3})
				walkableLand = append(walkableLand, coordinate{x,y})
			} else if lines[y][x] == '.' {
				walkableLand = append(walkableLand, coordinate{x,y})
			}
		}
	}
	return
}

// Returns a list with coordinates corresponding with the elf/goblin who's turn it is
func getTurns(elves []character, goblins []character, sizeX int, sizeY int) (turns []coordinate) {
	for y := 0; y <= sizeY; y++ {
		for x := 0; x <= sizeX; x++ {
			point := coordinate{x,y}
			for _, elf := range elves {
				if elf.position == point {
					turns = append(turns, elf.position)
				}
			}
			for _, goblin := range goblins {
				if goblin.position == point {
					turns = append(turns, goblin.position)
				}
			}
		}
	}
	return
}

func canAttack(characterPoint coordinate, enemies []character) bool {
	possiblePoints := []coordinate{ {characterPoint.x, characterPoint.y+1},
									{characterPoint.x, characterPoint.y-1},
									{characterPoint.x+1, characterPoint.y},
									{characterPoint.x-1, characterPoint.y}}
	for _, enemy := range enemies {
		for _, point := range possiblePoints {
			if enemy.position == point {
				return true
			}
		}
	}

	return false
}

func attack(attacker character, enemies []character) []character {
	// In reading order
	possiblePoints := []coordinate{ {attacker.position.x, attacker.position.y-1},
									{attacker.position.x-1, attacker.position.y},
									{attacker.position.x+1, attacker.position.y},
									{attacker.position.x, attacker.position.y+1}}
	for _, point := range possiblePoints {
		for i := range enemies {
			if enemies[i].position == point {
				enemies[i].hitpoints -= attacker.attackPower
				if enemies[i].hitpoints <= 0 {
					enemies = remove(enemies, i)
				}
				return enemies
			}
		}
	}
	return enemies
}

func remove(list []character, index int) []character {
	if len(list) == 1 {
		return []character{}
	} else if index == len(list) -1 {
		return list[:index]
	}
	return append(list[:index], list[index+1:]...)
}

func identifyTargets(enemies []character, walkableLand []coordinate) []character {
	return []character{}
}

func isReachable(position coordinate, goal coordinate) bool {
	return false
}

func move() {

}

func breadthFirstSearch(position coordinate, endPoint coordinate, walkableLand []coordinate) (path []coordinate) {
	
	return nil
}

func main() {
	file, err := ioutil.ReadFile("../testinput.txt")
	checkError(err)
	elves, goblins, walkableLand, sizeX, sizeY := getCharactersAndLand(strings.Split(string(file), "\n"))
	fmt.Println("Elves:", elves)
	fmt.Println("Goblins:", goblins)
	fmt.Println("Walkable land:", walkableLand)

	turnOrder := getTurns(elves, goblins, sizeX, sizeY)
	for turn := 1; turn < 1000; turn++ {
		fmt.Println("Turn:", turn)
		for _, point := range turnOrder {
			for _, elf := range elves {
				if elf.position == point {
					if canAttack(elf.position, goblins) {
						goblins = attack(elf, goblins)
					} else {
						targets := identifyTargets(goblins, walkableLand)
						reachableTargets := []character{}
						for _, target := range targets {
							if !isReachable(elf.position, target.position) {
								reachableTargets = append(reachableTargets, target)
							}
							if len(reachableTargets) > 0 {
								// move
							}
						}
					}
				}
			}
			for _, goblin := range goblins {
				if goblin.position == point {
					if canAttack(goblin.position, elves) {
						elves = attack(goblin, elves)
					} else {
						targets := identifyTargets(elves, walkableLand)
						reachableTargets := []character{}
						for _, target := range targets {
							if !isReachable(goblin.position, target.position) {
								reachableTargets = append(reachableTargets, target)
							}
							if len(reachableTargets) > 0 {
								// move
							}
						}
					}
				}
			}
		}
	}
}
