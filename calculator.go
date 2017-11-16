package main

import (
	"fmt"  //Package fmt implements formatted I/O with functions analogous to C's printf
	"math"
	"reflect"  //Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.


)

// Stores the result of calculations
var result struct {
	last     float32 // stores last result
	operator string  // the next operator to use
	active   bool    // whether calculation has started
}

