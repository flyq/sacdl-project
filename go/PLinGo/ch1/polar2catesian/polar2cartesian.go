package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	th     float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //*nix
		prompt = fmt.Sprintf(prompt, "Ctrl + D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := creatSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			th := polarCoord.th * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(th)
			y := polarCoord.radius * math.Sin(th)
			answers <- cartesian{x, y}
		}

	}()
	return answers
}

const result = "Polar radius=%.02f th=%.02f'->Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Println("Radius and angle: ")
		line, err := reader.Reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, thfloat64
		if _, err := fmt.Sscan(line, "%f %f", &radius, &th); err != nil {
			fmt.Println(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, th}
		coord := <- answers
		fmt.Printf(result, radius, th, coord.x, coord.y)
	}
	fmt.Println()
}

			
