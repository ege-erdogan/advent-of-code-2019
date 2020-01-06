package main

import (
  "os"
  "bufio"
  "fmt"
  "strconv"
)

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()

  scanner := bufio.NewScanner(file)
  sum := 0
  for scanner.Scan() {
    mass, _ := strconv.Atoi(scanner.Text())
    fuel := calculateFuel(mass)
    sum += fuel
  }

  fmt.Println(sum)
}

func calculateFuel(mass int) int {
  fuel := (mass / 3) - 2
  if fuel < 0 {
    return 0
  }
  return fuel + calculateFuel(fuel)
}
