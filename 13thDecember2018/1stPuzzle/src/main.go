package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cart struct {
	positionX int
	positionY int
	direction string
	directionNextIntersection string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getCarts(tracks [][]string) ([]cart, [][]string) {
	carts := []cart{}
	for y := range tracks {
		for x := range tracks[y] {
			if tracks[y][x] == ">" {
				carts = append(carts, cart{x, y, "right", "left"})
				tracks[y][x] = "-"
			} else if tracks[y][x] == "^" {
				carts = append(carts, cart{x, y, "up", "left"})
				tracks[y][x] = "|"
			} else if tracks[y][x] == "<" {
				carts = append(carts, cart{x, y, "left", "left"})
				tracks[y][x] = "-"
			} else if tracks[y][x] == "v" {
				carts = append(carts, cart{x, y, "down", "left"})
				tracks[y][x] = "|"
			}
		}
	}
	return carts, tracks
}

func getCart(xCoordinate int, yCoordinate int, carts []cart) (bool, int) {
	for i := range carts {
		if carts[i].positionX == xCoordinate && carts[i].positionY == yCoordinate {
			return true, i
		}
	}
	return false, 0
}

func updateCartDirection(car cart, tracks [][]string) cart {
	DIRECTION_UP := "up"
	DIRECTION_DOWN := "down"
	DIRECTION_LEFT := "left"
	DIRECTION_RIGHT := "right"
	DIRECTION_STRAIGHT := "straight"

	INTERSECTION := "+"
	BEND1 := "/"
	BEND2 := "\\"

	pieceOfTrack := tracks[car.positionY][car.positionX]

	switch pieceOfTrack {
	case INTERSECTION:
		if car.directionNextIntersection == DIRECTION_STRAIGHT {
			car.directionNextIntersection = DIRECTION_RIGHT
			// in this case, do not change the cart's direction
		} else {
			if car.direction == DIRECTION_UP {
				car.direction = car.directionNextIntersection
				if car.direction == DIRECTION_LEFT {
					car.directionNextIntersection = DIRECTION_STRAIGHT
				} else if car.direction == DIRECTION_RIGHT {
					car.directionNextIntersection = DIRECTION_LEFT
				}
			} else if car.direction == DIRECTION_DOWN {
				if car.directionNextIntersection == DIRECTION_LEFT {
					car.direction = DIRECTION_RIGHT
					car.directionNextIntersection = DIRECTION_STRAIGHT
				} else if car.directionNextIntersection == DIRECTION_RIGHT {
					car.direction = DIRECTION_LEFT
					car.directionNextIntersection = DIRECTION_LEFT
				}
			} else if car.direction == DIRECTION_RIGHT {
				if car.directionNextIntersection == DIRECTION_LEFT {
					car.direction = DIRECTION_UP
					car.directionNextIntersection = DIRECTION_STRAIGHT
				} else if car.directionNextIntersection == DIRECTION_RIGHT {
					car.direction = DIRECTION_DOWN
					car.directionNextIntersection = DIRECTION_LEFT
				}
			} else if car.direction == DIRECTION_LEFT {
				if car.directionNextIntersection == DIRECTION_LEFT {
					car.direction = DIRECTION_DOWN
					car.directionNextIntersection = DIRECTION_STRAIGHT
				} else if car.directionNextIntersection == DIRECTION_RIGHT {
					car.direction = DIRECTION_UP
					car.directionNextIntersection = DIRECTION_LEFT
				}
			}
		}
	case BEND1:
		if car.direction == DIRECTION_UP {
			car.direction = DIRECTION_RIGHT
		} else if car.direction == DIRECTION_DOWN {
			car.direction = DIRECTION_LEFT
		} else if car.direction == DIRECTION_LEFT {
			car.direction = DIRECTION_DOWN
		} else if car.direction == DIRECTION_RIGHT {
			car.direction = DIRECTION_UP
		}
	case BEND2:
		if car.direction == DIRECTION_UP {
			car.direction = DIRECTION_LEFT
		} else if car.direction == DIRECTION_DOWN {
			car.direction = DIRECTION_RIGHT
		} else if car.direction == DIRECTION_RIGHT {
			car.direction = DIRECTION_DOWN
		} else if car.direction == DIRECTION_LEFT {
			car.direction = DIRECTION_UP
		}
	}
	return car
}

func crashOccured(cartToCheck cart, carts []cart) bool {
	for i := range carts {
		if carts[i].positionX == cartToCheck.positionX && carts[i].positionY == cartToCheck.positionY {
			return true
		}
	}
	return false
}

// returns true if a cart is crashed and the position
func updateCarts(carts []cart, tracks [][]string) (bool, int, int, []cart) {
	updatedCarts := []cart{}
	for y := range tracks {
		for x := range tracks[y] {
			cartFound, index := getCart(x, y, carts)
			if cartFound {
				curCart := carts[index]
				carts[index] = cart{}
				if curCart.direction == "left" {
					curCart.positionX -= 1
				} else if curCart.direction == "up" {
					curCart.positionY -= 1
				} else if curCart.direction == "right" {
					curCart.positionX += 1
				} else if curCart.direction == "down" {
					curCart.positionY += 1
				}

				checkForCrash := append(carts, updatedCarts...)
				if crashOccured(curCart, checkForCrash) {
					return true, curCart.positionX, curCart.positionY, []cart{}
				}
				curCart = updateCartDirection(curCart, tracks)
				updatedCarts = append(updatedCarts, curCart)
			}
		}
	}
	return false, 0, 0, updatedCarts
}

func main() {
	file, err := ioutil.ReadFile("../input.txt")
	checkError(err)
	tmpTracks := strings.Split(string(file), "\n")

	tracks := [][]string{}
	for i := range tmpTracks {
		tracks = append(tracks, strings.Split(tmpTracks[i], ""))
	}
	carts, tracks := getCarts(tracks)
	crashed := false
	positionXCrash := 0
	positionYCrash := 0
	for !crashed {
		crashed, positionXCrash, positionYCrash, carts = updateCarts(carts, tracks)
	}
	fmt.Println("The carts crashed at position:", positionXCrash, ",", positionYCrash)
}