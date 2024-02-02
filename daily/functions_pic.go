package main
 
import "golang.org/x/tour/pic"

func calc_pow() {
pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
}


	
func Pic(dx, dy int) [][]uint8 {
	myFirstSlice := make([][]uint8, dy)
	
	for i := range myFirstSlice {
		myFirstSlice[i] = make([]uint8, dx)
	}
	
	for i := range myFirstSlice {
		for j := range myFirstSlice[i] {
			myFirstSlice[i][j] = uint8(min(i,j))
		}
	}
	
	return myFirstSlice
}


func Pic2(dx, dy int) [][]uint8 {
	myFirstSlice := make([][]uint8, dy)
	
	for i := range myFirstSlice {
		myFirstSlice[i] = make([]uint8, dx)
	}
	
	for i := range myFirstSlice {
		for j := range myFirstSlice[i] {
			myFirstSlice[i][j] = uint8(i*j)
		}
	}
	
	return myFirstSlice
}

func Pic3(dx, dy int) [][]uint8 {
	myFirstSlice := make([][]uint8, dy)
	
	for i := range myFirstSlice {
		myFirstSlice[i] = make([]uint8, dx)
	}
	
	for i := range myFirstSlice {
		for j := range myFirstSlice[i] {
			myFirstSlice[i][j] = uint8((i*i)+(2*i*j)+(j*j))
		}
	}
	
	return myFirstSlice
}


func Pic4(dx, dy int) [][]uint8 {
	myFirstSlice := make([][]uint8, dy)
	
	for i := range myFirstSlice {
		myFirstSlice[i] = make([]uint8, dx)
	}
	
	for i := range myFirstSlice {
		for j := range myFirstSlice[i] {
			myFirstSlice[i][j] = uint8((i*i)+(j*j))
		}
	}
	
	return myFirstSlice
}

func Pic5(dx, dy int) [][]uint8 {
	myFirstSlice := make([][]uint8, dy)
	
	for i := range myFirstSlice {
		myFirstSlice[i] = make([]uint8, dx)
	}
	
	for i := range myFirstSlice {
		for j := range myFirstSlice[i] {
			myFirstSlice[i][j] = uint8(+(i*i)+(2*i*j)-(j*j))
		}
	}
	
	return myFirstSlice
}

func main() {
	pic.Show(Pic5)
}
