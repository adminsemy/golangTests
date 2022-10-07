package testAsyncMap

func main() {
	async := NewAsyncMap()
	for i := 0; i <= 100; i++ {
		go func(async *CasheMap) {

		}(&async)
	}
}
