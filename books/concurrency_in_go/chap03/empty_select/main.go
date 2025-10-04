package main

func main() {
	// deadlock!
	select {}
}
