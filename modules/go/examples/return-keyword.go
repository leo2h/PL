import main

// tag::return[]
func noResult() {
	return        // <1>
}

func singleResult() int {
	return 42
}

func multiResults() (int, float32) {
	return 42, 42.0
}

func namedResult() (x int, y float32) {
    x = 42
	y = 42.0
	return
}
// end::return[]

func main() {

}