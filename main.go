
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("no input.txt available")
	}
	counter := 0
	lines := strings.Split(string(content), "\n")
	target := string("XMAS")
	inverse_target := string("SAMX")
	Vbuffer := ""
	diagBuffer := ""
	diagBufferInv := ""
	Vcounter := 0
	DiagCount := 0
	invDiagCounter := 0
	everything := string(content)
	counter += strings.Count(everything, "XMAS")
	counter += strings.Count(everything, "SAMX")
	patterns := []string{"MMASS", "SSAMM", "MSAMS", "SMASM"}
	xmasBuffer := ""
	xmasCounter := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if (i + 3) < len(lines) {
				for k := 0; k < 4; k++ {
					Vbuffer = Vbuffer + string(lines[i+k][j])
				}
				if Vbuffer == target || Vbuffer == inverse_target {
					counter++
					Vcounter++
				}
				Vbuffer = ""
			}

			if i < len(lines)-3 {
				if j < len(lines[i])-3 {
					for k := 0; k < 4; k++ {
						diagBuffer = diagBuffer + string(lines[i+k][j+k])
					}
					if diagBuffer == target || diagBuffer == inverse_target {
						counter++
						DiagCount++
					}
				}
			}
			diagBuffer = ""
			if i > 2 {
				if j < len(lines[i])-3 {
					for k := 0; k < 4; k++ {
						diagBufferInv = diagBufferInv + string(lines[i-k][j+k])
					}
					if diagBufferInv == target || diagBufferInv == inverse_target {
						counter++
						invDiagCounter++
					}
					diagBufferInv = ""
				}
			}
			if i < len(lines)-2 {
				if j < len(lines[i])-2 {
					xmasBuffer += string(lines[i][j])
					xmasBuffer += string(lines[i][j+2])
					xmasBuffer += string(lines[i+1][j+1])
					xmasBuffer += string(lines[i+2][j])
					xmasBuffer += string(lines[i+2][j+2])
					switch xmasBuffer {
					case patterns[0]:
						xmasCounter++
					case patterns[1]:
						xmasCounter++
					case patterns[2]:
						xmasCounter++
					case patterns[3]:
						xmasCounter++
					}
					xmasBuffer = ""
				}
			}
		}
	}

	fmt.Println("First part:", counter)
	fmt.Println("Second part:", xmasCounter)

}
