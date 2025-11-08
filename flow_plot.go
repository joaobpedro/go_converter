package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
	"strconv"
	s "strings"
)

// check the return errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// check if a string is numeric
func isFloat(a string) bool {
	_, err := strconv.ParseFloat(a, 64)
	return err == nil
}

func change_units(s string) string {
	// change the units in the density word
	val, err := strconv.ParseFloat(s, 64)
	check(err)
	val = val / 1000.0
	s2 := strconv.FormatFloat(val, 'f', -1, 64)
	return s2
}

func main() {
	args := os.Args
	fmt.Print("Loading the file: ", args[1], "\n")
	// open file in go, this is just a dump of the info into memory
	f, err := os.Open(args[1])
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// declare the columns that I want to take from the file
	// let's keep it simple for now and take time, length and density
	var arclength []string
	var time []string
	var density []string

	// this just reads data form the file
	for scanner.Scan() {
		line := scanner.Text()

		time = append(time, s.Split(line, " ")[0])
		arclength = append(arclength, s.Split(line, " ")[1])
		density = append(density, s.Split(line, " ")[2])
	}

	const tdp float64 = 1023.576628 // this values needs to be changes
	// now loop through the arclength data take the TDP location

	var time_plot []float64
	var densityplot []float64

	for i, val := range arclength {
		// transform val into a float
		val_float, err := strconv.ParseFloat(val, 64)
		check(err)
		if val_float == tdp {
			time_val, err := strconv.ParseFloat(time[i], 64)
			check(err)
			time_plot = append(time_plot, time_val)

			density_val, err := strconv.ParseFloat(density[i], 64)
			check(err)
			densityplot = append(densityplot, density_val)
		}
	}

	// fmt.Println(time_plot)
	// fmt.Println(densityplot)

	pts := make(plotter.XYs, len(time_plot))
	for i, _ := range time_plot {
		pts[i].X = time_plot[i]
		pts[i].Y = densityplot[i]
	}

	plt := plot.New()
	sct, err := plotter.NewScatter(pts)
	check(err)

	plt.Add(sct)
	err2 := plt.Save(7*vg.Inch, 4*vg.Inch, "test_plot.svg")
	check(err2)
}
