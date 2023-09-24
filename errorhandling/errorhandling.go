package errorhandling

import (
	"fmt"
	"log"
	"math"
	"os"
)

type sqrtNegativeErr struct {
	value float64
	err   error
}

func (err *sqrtNegativeErr) Error() string {
	if err.value < 0 {
		return fmt.Sprintf("cannot use negative number: %f", err.value)
	}
    
	return fmt.Sprint(err.value)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &sqrtNegativeErr{value: x, err: nil}
	}

	return math.Sqrt(x), nil
}

func TrySqrt() {
	sqrtPositive, err := sqrt(2)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(sqrtPositive)

	sqrtNegative, err := sqrt(-2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sqrtNegative)
}

func Hello(name string) {
    message := recover();
    
    if message != nil {
        fmt.Println("error because: ", message);
    }
    fmt.Println(name)
}

func ErrorDefer(err bool) {
    defer Hello("maulana")

    if err {
        panic("Error")
    }
}
