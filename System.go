package main

import "fmt"

const iter = 15

func logistic(r float64, x float64) float64 {
	return r*x*(1-x)
}

func getLogLine(r float64, x0 float64) [iter]float64 {
	var logLine [iter]float64
	logLine[0] = x0
	for i:=0; i<(iter-1); i++ {
		logLine[i+1] = logistic(r, logLine[i])
	}
	return logLine
}

func main() {

	var r float64 = 3.6
	for x0:=0.01; x0<1; x0+=0.048 {
		logLine := getLogLine(r, x0)
		
		for i:=0; i<iter; i++ {
			fmt.Print("(", i+1, ",", x0, ",", logLine[i], ") ")
		}
		fmt.Println()
		fmt.Println()
	}

	
}
