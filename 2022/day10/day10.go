package day10

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day10 struct {
	Input string
	Dir   string
}

func New() Day10 {
	return Day10{
		Input: "2022/day10/input.txt",
		Dir:   "2022/day10/",
	}
}

type CRT struct {
	screen [240]string
}

func (c *CRT) update(cycle, registryX int) {
	cycle -= 1
	y := cycle / 40
	tmp := cycle - (40 * y)
	if math.Abs(float64(tmp-registryX)) <= 1 {
		c.screen[cycle] = "#"
	}
}

func (c *CRT) draw() {
	fmt.Println(c.screen[:40])
	fmt.Println(c.screen[40:80])
	fmt.Println(c.screen[80:120])
	fmt.Println(c.screen[120:160])
	fmt.Println(c.screen[160:200])
	fmt.Println(c.screen[200:240])
}

func newCRT() CRT {
	var screen [240]string
	for i := 0; i < 240; i++ {
		screen[i] = "."
	}
	return CRT{screen: screen}
}

type CPU struct {
	cycles     int
	x          int
	nextUpdate int
	sum        int
	crt        CRT
}

func (c *CPU) processNoop() {
	c.updateCycle()
}

func (c *CPU) processAddrx(value int) {
	c.updateCycle()
	c.updateCycle()
	c.x += value
}

func (c *CPU) updateCycle() {
	c.cycles++
	c.crt.update(c.cycles, c.x)
	if c.cycles == c.nextUpdate {
		c.sum += c.cycles * c.x
		c.nextUpdate += 40
	}
}

func (d Day10) ProcessPuzzle1(lines []string) (string, error) {
	cpu := CPU{
		cycles:     0,
		x:          1,
		nextUpdate: 20,
		sum:        0,
	}

	for _, line := range lines {
		if line == "noop" {
			cpu.processNoop()
			continue
		}

		arr := strings.Split(line, " ")
		value, _ := strconv.Atoi(arr[1])
		cpu.processAddrx(value)
	}

	return fmt.Sprintf("%d", cpu.sum), nil
}

func (d Day10) ProcessPuzzle2(lines []string) (string, error) {
	cpu := CPU{
		cycles:     0,
		x:          1,
		nextUpdate: 40,
		sum:        0,
		crt:        newCRT(),
	}

	for _, line := range lines {
		if line == "noop" {
			cpu.processNoop()
			continue
		}

		arr := strings.Split(line, " ")
		value, _ := strconv.Atoi(arr[1])
		cpu.processAddrx(value)
	}

	cpu.crt.draw()

	return "not implemented", nil
}
