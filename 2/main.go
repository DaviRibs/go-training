package main

const a = "Hello, World"

// Escopo Global
var (
	b bool    = true
	c int     = 10
	d string  = "Davi"
	e float64 = 1.2
)

// escopo  local
func main() {
	a := "X" // string feito na primeira vez que a variavel é criada

	println(a)
}

func x() {

}