package main

import "fmt"

func main() {
	fmt.Println("Hello", "world!", "how", "are", "you")

	// Variable

	// menggunakan var beserta tipe data
	var firstName1 string = "john"

	fmt.Println("First name 1: " + firstName1)
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"

	fmt.Printf("firstname : %s ,\nlastname : %s", firstName, lastName)

	// constanta

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	fmt.Print("\n", satu, dua, three, four)

	garis()

	kondisi()
}

func kondisi() {
	// if else
	var point float32 = 10

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %f\n", point)
	}

	garis()
	// if else dengan temporary data
	point = 10000

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	garis()
	// switch tidak perlu break
	point = 10

	switch point := point - 2; point { // note : ini menggunakan temporary data
	case 8:
		fmt.Println("perfect")
	case 7, 10:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	fmt.Println("Ini nilait dari point", point)

	garis()
	// switch yang salah pergaulan
	point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println(point)

}

func garis() {
	fmt.Print("\n ================ \n")
}
