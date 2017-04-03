package main

import "plugin"

func main() {
  p, _ := plugin.Open("./plugin.so")
  add, _ := p.Lookup("Add")
  sum := add.(func(int, int) int)(40, 2)
  println(sum)
}

