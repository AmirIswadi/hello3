package main
import "fmt"

func main() {
    var celsius float64
    fmt.Printf("Enter the temperature in Celsius: ")
    fmt.Scanf("%g", &celsius)

    var fahrenheit float64
    fahrenheit = celsius * 9/5 + 32
    fmt.Printf("Fahrenheit value is: %g\n", fahrenheit)
}