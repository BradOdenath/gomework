package main

import (
	"log"
	"math"
	"os"
	"bufio"
	"unicode"
	"fmt"
	"strings"
)

const (
	default_name = "City"
	default_x = 0
	default_y = 0
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	log.Println("Gomework!")
	readCityDocument("dj38.tsp.txt")
}

/* ===== CITY ===== */

//Go is not OOP, so accessors are a default thing.
type city struct {
	x    float64
	y    float64
	name string
}

func defaultCity() city {
	/* SideNote: Since Go is not technically OOP, you can't declare default values
				within the struct.  This feels super salty, but hey whatever.*/
	return city{0,0,"Default_City"}
}

func distance(c1 city, c2 city) int64 {
	/* SideNote: I'm used to putting + on the pre statement, not post.
				Go's structure does not allow pre statement operators.	*/
	return int64(
		math.Round(
			math.Sqrt(math.Pow(c1.x-c2.y, 2) +
				math.Pow(c1.y-c2.y, 2))))
}

//Figured for convention of the assignment I'd throw this in there.
func toString(c city) string {
	return c.name
}

/* ===== TOUR ===== */
/*

type tour struct {
	tourAL[] Channel	//getTour()
	distance in64	//getistance
}
func createTour() tour {
	return tour{context.TODO(),0}
}
*/

/* ===== FILE READER AND CITY LIST ===== */

type cityList struct {
	cityX[] string
	cityY[] string
}

func readCityDocument(fileName string) (string) {
	//file, err := ioutil.ReadFile(fileName)

	cities := cityList{}

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	outStr := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanStr := scanner.Text()
		//log.Println(scanStr)

		tempRune := []rune(scanStr)

		if unicode.IsDigit(tempRune[0]) {
			//Break them coords apart without breaking my neck
			tempCoords:= strings.Fields(scanStr)
			cities.cityX = append(cities.cityX, tempCoords[1])
			cities.cityY = append(cities.cityY, tempCoords[2])
		}

		fmt.Println(outStr)

	}

	return outStr

	}