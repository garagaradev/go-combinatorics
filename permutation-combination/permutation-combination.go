// Permutation and Combination Code with Input from Terminal
package main
import "fmt"


// function to calculate factorial 
func factorial(n int) int {
  if n == 0 {
    return 1
  }

  result := 1
  for i := 2; i <= n; i++ {
    result *= i
  }
  return result
}

// function to calculate permutation 
func permutation(n, r int) int {
  return factorial(n) / factorial(n-r)
}

// function to calculate combination
func combination(n, r int) int {
  return factorial(n) / (factorial(r) * factorial(n-r))
}

func main(){
  var n, r int

  // Input values for n and r from terminal 
  fmt.Printf("[INFO] n refers to the total number of items or elements in a set. \n r refers to the number of items or elements to be chosen from that set.\n")
  fmt.Println()
  
  fmt.Println("Enter the value of n:")
  fmt.Scan(&n)
  fmt.Println("Enter the value of r:")
  fmt.Scan(&r)

  // calculating permutation and combination
  perm := permutation(n, r)
  comb := combination(n, r)

  //output of the results
  fmt.Printf("Permutation P(%d, %d) = %d\n",n, r, perm)
  fmt.Printf("Combination C(%d, %d) = %d\n", n, r, comb)
}
