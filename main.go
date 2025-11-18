/*
 * Author Miles Aube
 * Version 1.0.0
 * Date 2025-11-17
 * This is the This program shows area of a circle
 */

package main

import "fmt"

func main () {
	 // set CircleRadius and Pi
  var CircleRadius float64 = 5.6
	var Pi float64 = 3.14159

	//set answer 
	var CircleArea float64 = Pi * CircleRadius * CircleRadius

	// Display Circcle area
	fmt.Println ("The area of a circle with a radius of ",CircleRadius," cm is ",CircleArea,"cm\u00B2")
}