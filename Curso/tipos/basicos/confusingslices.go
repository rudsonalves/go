package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	fmt.Println(x)
	x = append(x, 1, 2, 3, 4)
	fmt.Println(x)
	y := x[:2]
	z := x[2:]
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("cap(x) = cap(y), but cap(z) is only 3!")

	fmt.Println("\nz = append(z, 70)")
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("Append to z affect only z slice")

	fmt.Println("\ny = append(y, 10)")
	y = append(y, 10)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("But append do y change all, x and z slices")

	fmt.Println("\nAvoid complicated slice situations:")
	w := x[:2:2]
	fmt.Println("w := x[:2:2]")
	fmt.Println("w:", w, " cap(w)", cap(w))
	w = append(w, 30, 40)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("w:", w)
	fmt.Println("\nTo avoid complicated slice situations, you should either")
	fmt.Println("never use append with a subslice or make sure that append")
	fmt.Println("doesn´t cause an overwrite by using a full slice expression.")

	k := x[2:4:4]
	fmt.Println("k := x[2:4:4]")
	fmt.Println("k:", k, " cap(k)", cap(k))
	k = append(k, 50, 60)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("k:", k)
	fmt.Println("\nBoth w and k have the capacity of 2. Because we limited the")
	fmt.Println("capacity of the subslices to their legths, appending additional")
	fmt.Println("elements onto w and k create a new slices that didn't interect")
	fmt.Println("with the others slices.")

	k[0] = 111
	w[0] = 222
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("w:", w)
	fmt.Println("k:", k)
}
