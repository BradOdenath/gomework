package main

import (
	"log"
	"math"
	"os"
	"bufio"
	"unicode"
	"strings"
	"fmt"
	"strconv"
)

const default_name = "City"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//Pride
	fmt.Println("Gomework!")

	//Parse the document for coordiantes
	cityCoords := readCityDocument("dj38.tsp.txt")

	//Put them things into a thing.
	cityTour := populateTour(cityCoords)

	//Feels good mate!
	log.Println(getDistance(cityTour))
}

/* ===== CITY ===== */

//Go is not OOP, so accessors are a default thing.
type city struct {
	x    float64
	y    float64
	name string
}

func distance(c1 city, c2 city) int64 {
	/* SideNote: I'm used to putting + on the pre statement, not post.
				Go's structure does not allow pre statement operators.	*/
	return int64(
		math.Round(
			math.Sqrt(math.Pow(c1.x-c2.x, 2) +
				math.Pow(c1.y-c2.y, 2))))
}

//Figured for convention of the assignment I'd throw this in there.
func toString(c city) string {
	return c.name
}

/* ===== TOUR ===== */

//Note: the city needed can be obtained through accessing the path
//		directly rather than making a method.
type tour struct {
	path [] city
}

func addCity(t *tour, c city) {
	t.path = append(t.path, c)
	log.Println("Succesfully added city to tour.")
}

func numberOfCities(t tour) int64 {
	return int64(len(t.path))
}

func populateTour(coords cityCoordinates) tour {
	t := tour{}
	for i, cityCoordX := range coords.cityX { //X same length as Y
		tempX, err := strconv.ParseFloat(cityCoordX, 64)
		check(err)
		tempY, err := strconv.ParseFloat(coords.cityY[i], 64)
		check(err)
		t.path = append(t.path, city{tempX, tempY, default_name})
	}
	return t
}

func getDistance(t tour) int64 {
	totalDistance := int64(0)
	for i, c := range t.path {
		destination := city{}

		if int64(i+1) < numberOfCities(t) {
			destination = t.path[i+1]
		} else {
			destination = t.path[0]
		}

		totalDistance += distance(c, destination)
	}
	return totalDistance
}

/* ===== COORDINATE FILE GRABBER STUFF ===== */

type cityCoordinates struct {
	cityX [] string
	cityY [] string
}

//Heart palpitations ensued as I wrote this.
func readCityDocument(fileName string) cityCoordinates {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	cities := cityCoordinates{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanStr := scanner.Text()
		tempRune := []rune(scanStr)

		if unicode.IsDigit(tempRune[0]) {
			//Break them coords apart without breaking my neck
			tempCoords := strings.Fields(scanStr)
			cities.cityX = append(cities.cityX, tempCoords[1])
			cities.cityY = append(cities.cityY, tempCoords[2])
		}
	}
	return cities
}