package main

import (
	"fmt"
	"math"
)

func boolfunc1(y1 int) bool {
	var x1 = 2
	if x1 == y1 {
		return true
	} else {
		return false
	}
}

func boolfunc2(x2, y2 bool) bool {
	if x2 || y2 == true {
		return true
	} else {
		return false
	}
}

func boollfunc4(x, y, R float64) bool {
	if (math.Sqrt(x-x*0))+(math.Sqrt(y-y*0)) <= math.Sqrt(R) {
		return true
	} else {
		return false
	}
}

func frombooltostrlfunc(x3, y3, R float64) string {
	if (math.Sqrt(x3-x3*0) + math.Sqrt(y3-y3*0)) <= math.Sqrt(R) {
		return "true"
	} else {
		return "false"
	}
}

func comparefunc1(c1 int) bool {
	if c1 >= 0 {
		return true
	} else {
		return false
	}
}

func comparefunc2(c2 int) string {
	if (c2 >= 80) && (c2 <= 100) {
		return "Great"
	} else if (c2 >= 40) && (c2 <= 79) {
		return "Good"
	} else if (c2 >= 20) && (c2 <= 39) {
		return "Not bad"
	} else {
		return "Bad"
	}
}

func comparefunc3(c3 int) string {
	if (c3 > 70_000_000) && (c3 <= 120_000_000) {
		return "The nearest plamet is Mercury"
	} else if (c3 > 120000000) && (c3 <= 180000000) {
		return "The nearest plamet is Venus"
	} else if (c3 > 180000000) && (c3 <= 500000000) {
		return "The nearest plamet is Earth"
	} else if (c3 > 500000000) && (c3 <= 1100000000) {
		return "The nearest plamet is Mars"
	} else if (c3 > 1100000000) && (c3 <= 2200000000) {
		return "The nearest plamet is Jupiter"
	} else if (c3 > 2200000000) && (c3 <= 3500000000) {
		return "The nearest plamet is Saturn"
	} else if (c3 > 3500000000) && (c3 <= 4300000000) {
		return "The nearest plamet is Uranus"
	} else if (c3 > 43_000_000_000) && (c3 <= 80_000_000_000) {
		return "The nearest plamet is Neptune"
	} else if c3 > 80000000000 {
		return "It is not a solar system"
	} else {
		return ""
	}
}

func comparefunc4(c4 int, c44 int) string {
	if (c4 >= 0) && (c4 <= 20) && (c44 >= 1500) && (c44 <= 3000) {
		return "The first gear"
	} else if (c4 >= 21) && (c4 <= 35) && (c44 >= 1500) && (c44 <= 3000) {
		return "The second gear"
	} else if (c4 >= 36) && (c4 <= 50) && (c44 >= 1500) && (c44 <= 3000) {
		return "The third gear"
	} else if (c4 >= 51) && (c4 <= 65) && (c44 >= 1500) && (c44 <= 3000) {
		return "The fourth gear"
	} else if (c4 >= 66) && (c4 <= 80) && (c44 >= 1500) && (c44 <= 3000) {
		return "The fifth gear"
	} else if c44 <= 1499 {
		return "You need to shift into a lower gear"
	} else if c44 >= 3001 {
		return "You need to shift into a higher gear"
	} else {
		return ""
	}
}

func signalStrenght(distance int) int {
	var signalOriginPerCent = 100
	var km = 100
	var flowRatePerCent = 2
	return signalOriginPerCent - distance/km*flowRatePerCent
}

func downloadTime(gig float64, intSpeed float64) float64 {
	return (gig * 1000) / (intSpeed / 8)
}

func main() {
	fmt.Println(boolfunc1(2))
	fmt.Println(boolfunc2(false, false))
	fmt.Println(frombooltostrlfunc(4, 4, 2))
	fmt.Println(boollfunc4(0, 0, 0))

	fmt.Println(comparefunc1(1))
	fmt.Println(comparefunc2(39))
	fmt.Println(comparefunc3(10000000000))
	fmt.Println(comparefunc4(40, 3050))

	fmt.Println(signalStrenght(600))
	fmt.Println(downloadTime(1, 100))
}
