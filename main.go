package main

import (
	"fmt"
	"os"
)

func writeCircle(buffer *[]uint, width, height, foreground, background, radius int) {
	centerX := width
	centerY := height
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := centerX - x*2 - 1
			dy := centerY - y*2 - 1
			r := radius * radius
			if (dx*dx)+(dy*dy) <= (r) {
				(*buffer)[y*width+x] = uint(foreground)
			} else {
				(*buffer)[y*width+x] = uint(background)
			}
		}
	}
}

func dumpToPpm(buffer *[]uint, filename string, width, height int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString("P3 \n")
	file.WriteString(fmt.Sprintf("%d %d 255\n", width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := (*buffer)[y*width+x]
			r := pixel >> 8 * 0 & 0xff
			g := pixel >> 8 * 1 & 0xff
			b := pixel >> 8 * 2 & 0xff
			file.WriteString(fmt.Sprintf("%d %d %d \n", r, g, b))
		}
	}
	return nil
}

func main() {
	WIDTH := 800
	HEIGHT := 600
	FOREGROUND := 0xFFFFFF
	BACKGROUND := 0x000000
	var buffer = make([]uint, WIDTH*HEIGHT)
	writeCircle(&buffer, WIDTH, HEIGHT, FOREGROUND, BACKGROUND, HEIGHT/2)
	if err := dumpToPpm(&buffer, "circle.ppm", WIDTH, HEIGHT); err != nil {
		panic(err)
	}
}
