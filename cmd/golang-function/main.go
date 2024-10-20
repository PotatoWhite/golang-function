package main

import (
	"fmt"
	"golang-function/pkg/chaining"
	"golang-function/pkg/factory"
	"golang-function/pkg/observer"
)

func consoleObserver(msg string) {
	println("console:", msg)
}

func fileObserver(msg string) {
	println("file:", msg)
}

func main() {
	// observer pattern
	emitter := &observer.EventEmitter{}
	emitter.Register(consoleObserver)
	emitter.Register(fileObserver)

	emitter.NotifyAll("Hello World!")

	// factory pattern
	tenPercentDiscount := factory.GetDiscountFunction(10)
	twentyPercentDiscount := factory.GetDiscountFunction(20)

	fmt.Printf("10%% 할인: %.2f\n", tenPercentDiscount(100.0))    // 출력: 90.00
	fmt.Printf("20%% 할인: %.2f\n", twentyPercentDiscount(100.0)) // 출력: 80.00

	// function chaining
	fmt.Println(chaining.Chain(2, chaining.Double, chaining.AddTen)) // 출력: 14
}
