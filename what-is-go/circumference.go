package main
import "fmt"

func main() {
	var pi float32 = 3.14;
	var radius float32 = 3.45;
	var cf float32 = 2.00;

	fmt.Println(cf * radius * pi);
	hoistingMaybe(); // hoisting is real
}


func hoistingMaybe() {
	pi := 3.14;
	radius := 5.0;

	circumference := 2 * pi * float64(radius);

	fmt.Println(circumference);
	fmt.Printf("For a radius of %v, the circle circumference is %.2f", radius, circumference);
}