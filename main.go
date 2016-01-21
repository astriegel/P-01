package main

import ("fmt"; "io/ioutil")

func main() {
	dat, err := ioutil.ReadFile("data.txt")
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

