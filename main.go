package main

func main() {
	saved := Save[ServiceInterface](&Service{"Putri"})
	println(saved)
	updated := Update[ServiceInterface](&Service{"Putri"})
	println(updated)
}
