package main

import "fmt"

func main() {
	var a [5]int
	b := [3]string{"e", "hi", "hihihi"}

	d := make([]int, 5)

	d[0] = 10
	d[1] = 20
	d[2] = 30
	d[3] = 40
	d = append(d, 90)
	fmt.Println(d)

	/////////////////////////////////////
	p := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(p)
	fmt.Println(p[0:4])
	fmt.Println(p[:4])
	fmt.Println(p[4:])

	///////////////////////////////
	var z [2][3]int

	a[0] = 10
	a[2] = 30
	a[3] = 40
	fmt.Println(a)
	fmt.Println("สมาชิก Array มีทั้งหมด => ", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println("สมาชิกใน Array ตัวที่ ", i+1, " = ", a[i])
	}

	for i := range a {
		fmt.Println("for 2 =>", a[i])
	}

	for i, v := range a {
		fmt.Println("for 3=>", i, " -- ", v)
	}

	for _, v := range b {
		fmt.Println(v)
	}

	for i := 0; i < len(z); i++ {
		for j := 0; j < len(z[i]); j++ {
			z[i][j] = j
		}
	}

	////////////////////////////

	r := make(map[string]string)
	r["hello"] = "gopher"
	r["name"] = "acosoft"

	fmt.Println(r)

	/////////////////////////////
	//slide
	name := []string{}

	name = append(name, "Parnaat")
	name = append(name, "Aticom")
	name = append(name, "Chaiyanint")

	fmt.Println(name)

	/////////////////////////////
	//map

	me := make(map[string]int)

	me["chayanint"] = 1
	me["theppayut"] = 2

	fmt.Println(me)

	delete(me, "chayanint")

	fmt.Println(me)

	me["nenenen"] = 3

	fmt.Println(me)

	for key, value := range me {
		fmt.Println(key, ": ", value)
	}

	//////
	//กำหนดค่าเริ่มต้นให้ map
	my := map[string]string{
		"1": "helloworld",
		"2": "222",
	}

	for key, value := range my {
		fmt.Println(key, " : ", value)
	}
}
