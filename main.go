package main

import "fmt"


func getBiggerNum(nums ...int) int {
  res := nums[0]
  for _, n := range nums {
    if n > res {
      res = n
    }
  }
  return res
}

func fibonacci(n int) int {
  if n == 1 {
    return 0
  } else if n == 2 {
    return 1
  } else {
    return fibonacci(n - 1) + fibonacci(n - 2) 
  }
}

func oddGenerator() func() uint {
	i := uint(0)
	return func() uint {
    for isOdd := false; isOdd == false; {
      if i%2 == 0 {
        isOdd = true
      }
      i++
    }
		var odd = i
		return odd
	}
}

func exchange(a *int, b *int) {
  *a, *b = *b, *a
}

func main() {
  fmt.Println("FIBONNACI")
  fmt.Println(fibonacci(10))

  fmt.Println("BIGGER NUMBER")
  fmt.Println(getBiggerNum(-123,-2,-43,4,-5,-6))

  fmt.Println("ODD NUMBERS")
  nextOdd := oddGenerator()
  for i := 0; i < 20 ; i++ {
    fmt.Println(nextOdd())
  }

  a := 10
  b := -235
  exchange(&a, &b)
  fmt.Printf("a: %d \n", a)
  fmt.Printf("b: %d \n", b)

}