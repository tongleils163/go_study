package pkg

func events() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- 1
		}
	}()
	return ch
}

func genEvents(ch chan int) {
	go func() {
		for {
			ch <- 1
		}
	}()
}

func Invoke() {

}
