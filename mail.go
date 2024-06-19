package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintUniverse(univ [15][35]int) {
	for i := 0; i < 15; i++ {
		for j := 0; j < 35; j++ {
			if univ[i][j] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func RandomUniverse() [15][35]int {
	var a, b int
	var univ [15][35]int
	//Заполнение массива нулями
	for i := 0; i < 15; i++ {
		for j := 0; j < 35; j++ {
			univ[i][j] = 0
		}
	}
	//Заполнение массива единицми в случайных местах.
	for i := 0; i < 10; i++ {
		a = rand.Intn(15)
		b = rand.Intn(35)
		univ[a][b] = 1
	}
	return univ
}

func CountNeighbours(i, j int, univ [15][35]int) int {
	var count int = 0
	if univ[i][j] == 1 {
		count--
	}
	if i == 0 {
		i = 15
	}
	if j == 0 {
		j = 35
	}
	i--
	j--
	jj := j
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if univ[i][j] == 1 {
				count++
			}
			j++
			if j > 34 {
				j = 0
			}
		}
		i++
		if i > 14 {
			i = 0
		}
		j = jj
	}
	return count
}

// Один жизненный цикл
func Step(univ [15][35]int) [15][35]int {
	newuniv := univ
	var count int
	for i := 0; i < 15; i++ {
		for j := 0; j < 35; j++ {
			count = CountNeighbours(i, j, univ)
			if count == 2 || count == 3 {
				newuniv[i][j] = 1
			} else {
				newuniv[i][j] = 0
			}
		}
	}
	return newuniv
}

func main() {
	fmt.Print("\033[H")
	universe := RandomUniverse()
	PrintUniverse(universe)
	for i := 0; i < 300; i++ {
		time.Sleep(time.Second / 1)
		fmt.Print("\033[H")
		universe = Step(universe)
		PrintUniverse(universe)
	}
}
