package task1

import (
	"awesomeProject/utility"
	"fmt"
	"log"
	"math"
	"strconv"
)

type (
	EuropeanConverter interface {
		MphToKmh(mph float64)
		MsToKmh(ms float64)
	}

	europeanVelocity float64
)

type (
	AmericanConverter interface {
		KmhToMph(kmh float64)
		MsToMph(ms float64)
	}

	americanVelocity float64
)

func (a *americanVelocity) KmhToMph(kmh float64) {
	mph := kmh / 1.609
	*a = americanVelocity(mph)
}

func (a *americanVelocity) MsToMph(ms float64) {
	kmh := ms * 3.6
	a.KmhToMph(kmh)
}

func (e *europeanVelocity) MphToKmh(mph float64) {
	kmh := mph * 1.609
	*e = europeanVelocity(kmh)
}

func (e *europeanVelocity) MsToKmh(ms float64) {
	kmh := ms * 3.6
	*e = europeanVelocity(kmh)
}

func Task() {
	utility.AnnounceNewTask(2, 1)
	stringVar := "104"
	intVar := 35
	strToIVar, err := strconv.Atoi(stringVar)
	if err != nil {
		log.Fatal(err)
	}
	intToStrVar := strconv.Itoa(intVar)
	fmt.Printf("stringVar \"%s\", intVar %d, strToIVar %d, intToStrVar \"%s\"\n", stringVar, intVar, strToIVar, intToStrVar)

	var kmh europeanVelocity
	var mph americanVelocity
	fmt.Printf("kmh=%.2f, mph=%.2f\n", math.Round(float64(kmh*100)/100), math.Round(float64(mph*100)/100))

	kmh.MsToKmh(120.4)
	mph.KmhToMph(100)
	fmt.Printf("kmh=%.2f, mph=%.2f\n", math.Round(float64(kmh*100)/100), math.Round(float64(mph*100)/100))

	kmh.MphToKmh(88)
	mph.MsToMph(130)
	fmt.Printf("kmh=%.2f, mph=%.2f\n", math.Round(float64(kmh)*100)/100, math.Round(float64(mph*100)/100))

}
