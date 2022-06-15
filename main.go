package main

var test *Trie

//Setup routes and handlers then serve on port 8080.
func main() {
	Execute()
	// C.getGPU()
	//Init and run Imgui using a go routine.

	// imgui.InitImgui()
	// SetupRoutes()
	//Init and test Trie.
	// start := time.Now()

	AddToTrie()

	// fmt.Printf("\033[32mTime Taken: %s\033[0m\n", time.Since(start))

	// test = initializeTrie()
	// test.insertToTrie("21e8", "21e893411ac5c7f3896fe57fb7d8a8f150ee18a7256fe73990a17c47a498c8b5")
	// test.insertToTrie("21e8", "21eb2f005c551eca25903ab09dbd08f512d9cbb6af226152690583cbcac51135")
	// test.insertToTrie("21e8", "21eabf80faebc12002aec48f82ba433758130924fde0c0b03dace7b0c9c42f09")

	// test.insertToTrie("21e", "21e813411aa5c7f3896fe57fb7d8a8f150ee18a7256fe73990a17c47a498c8b5")
	//End Test.

	// InitSocket()

	// http.HandleFunc("/", wsEndpoint)

}
