package main

import "fmt"

func main() {
	var (
		a int64   = 2
		b uint16  = 8
		c float64 = 2.43342325
		d bool    = true
		e byte    = 'a'
		f string  = "albert"
	)

	fmt.Printf("a: %v %T\n", a, a)
	fmt.Printf("b: %v %T\n", b, b)
	fmt.Printf("c: %v %T\n", c, c)
	fmt.Printf("d: %v %T\n", d, d)
	fmt.Printf("e: %v %T\n", e, e)
	fmt.Printf("f: %v %T\n", f, f)

	fmt.Println()

	fmt.Printf("%%d format in a: %03d##%+3d\n", a, a)
	fmt.Printf("%%d format in e: %02d  %+2d\n", e, e)
	fmt.Printf("%%f format in c: %1.1f  %.2f\n", c, c)
	fmt.Printf("%%q format in e and f:%q %q\n", e, f)
	fmt.Printf("%%o and %%O format in d:%o %O\n", b, b)
	fmt.Printf("%%s format in f: >|%s|<    >|%8s|<\n", f, f)
}
