package main

import (
	"fmt"
	"os"
)

const (
	pi         = 3.14159265
	e  float32 = 2.718281828459
)

type weekday int
type status int

// The weekday consts, starts from 0
const (
	Sunday   weekday = iota // 0
	Monday                  // 1
	Tuesday                 // 2
	Wensday                 // 3
	Thursday                // 4
	Friday                  // 5
	Saturday                // 6
)

// Weekdays consts, starts from 1
const (
	MON weekday = 1 + iota // 1
	TUE                    // 2
	WEN                    // 3
	THU                    // 4
	FRI                    // 5
	SAT                    // 6
	SUN                    // 7
)

// Status consts, initialize every element by move 1 bit to the left (times by 2)
const (
	Initialized status = 1 << iota // 1
	Processing                     // 2
	Failed                         // 4
	Succeeded                      // 8
)

// data size unit
const (
	_   = 1 << (10 * iota) // Skip the first one
	KiB                    // You can use _ to skip the second const as well
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Fprintf(os.Stdout, "const pi is %v (%T)\n", pi, pi)

	fmt.Fprintf(os.Stdout, "const e (%v) is explicted specified with type (%T)\n", e, e)

	weekdays := []weekday{Sunday, Monday, Tuesday, Wensday, Thursday, Friday, Saturday}

	for _, day := range weekdays {
		fmt.Fprintf(os.Stdout, "Weekday: %v\n", day)
	}

	fmt.Fprintln(os.Stdout)

	weekdaysAlternative := []weekday{MON, TUE, WEN, THU, FRI, SAT, SUN}

	for _, day := range weekdaysAlternative {
		fmt.Fprintf(os.Stdout, "Weekday: %v\n", day)
	}

	fmt.Fprintln(os.Stdout)

	statusList := []status{Initialized, Processing, Failed, Succeeded}

	for _, statusValue := range statusList {
		fmt.Fprintf(os.Stdout, "Status: %v\n", statusValue)
	}

	fmt.Fprintln(os.Stdout)

	dataSizeUnits := []int64{MiB, GiB, TiB, PiB, EiB} // ZiB, YiB overflow int64

	for _, dataSizeUnit := range dataSizeUnits {
		fmt.Fprintf(os.Stdout, "Data Size Unit: %v\n", dataSizeUnit)
	}
}
