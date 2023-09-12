package task2

import (
	"fmt"
	"log"
	"strconv"
)

type (
	EuropeanConverter interface {
		MphToKmh(mph float64)
		MsToKmh(ms float64)
	}

	EuropeanVelocity float64
)

type (
	AmericanConverter interface {
		KmhToMph(kmh float64)
		MsToMph(ms float64)
	}

	AmericanVelocity float64
)

func (a *AmericanVelocity) KmhToMph(kmh float64) {
	mph := kmh / 1.609
	*a = AmericanVelocity(mph)
}

func (a *AmericanVelocity) MsToMph(ms float64) {
	kmh := ms * 3.6
	a.KmhToMph(kmh)
}

func (e *EuropeanVelocity) MphToKmh(mph float64) {
	kmh := mph * 1.609
	*e = EuropeanVelocity(kmh)
}

func (e *EuropeanVelocity) MsToKmh(ms float64) {
	kmh := ms * 3.6
	*e = EuropeanVelocity(kmh)
}

func Task() {
	printTaskNum()
	stringVar := "104"
	intVar := 35
	strToIVar, err := strconv.Atoi(stringVar)
	if err != nil {
		log.Fatal(err)
	}
	intToStrVar := strconv.Itoa(intVar)
	fmt.Printf("stringVar \"%s\", intVar %d, strToIVar %d, intToStrVar \"%s\"\n", stringVar, intVar, strToIVar, intToStrVar)

	var kmh EuropeanVelocity = 0
	var mph AmericanVelocity = 0
	fmt.Printf("kmh=%f, mph=%f\n", kmh, mph)

	kmh.MsToKmh(20)
	mph.KmhToMph(100)
	fmt.Printf("kmh=%f, mph=%f\n", kmh, mph)

	kmh.MphToKmh(88)
	mph.MsToMph(26.8224)
	fmt.Printf("kmh=%f, mph=%f\n", kmh, mph)

}

func printTaskNum() {
	fmt.Println()
	fmt.Println()
	fmt.Println("====================Task 2====================")
	fmt.Println()
	fmt.Println()
}
