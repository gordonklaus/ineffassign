package p

var b bool

func _() {
	var x int
	x = 0
	_ = x
}

func _() {
	var x int
	x = 0
	if b {
		_ = x
	}
}

func _() {
	var x int
	_ = x
	x = 0 //x
}

func _() {
	var x int
	x = 0 //x
	x = 0
	_ = x
}

func _() {
	x := 0
	x = 0
	_ = x
}

func _() {
	x := T(0)
	x = 0
	_ = x
}

func _() {
	x := ""
	x = "abc"
	_ = x
}

func _() {
	x := "abc" //x
	x = "def"
	_ = x
}

func _() {
	x := 1 //x
	x = 0
	_ = x
}

func _() {
	x := 0
	x = x + 0
	_ = x
}

func _() {
	x := 0
	x += 0
	_ = x
}

func _() {
	x := 0
	x++
	_ = x
}

func _() {
	x := 0
	if b {
		x = 0
	}
	_ = x
}

func _() {
	x := 1 //x
	if b {
		x = 0 //x
	}
	x = 0
	_ = x
}

func _() {
	x := 1 //x
	for b {
		x = 0 //x
	}
	x = 0
	_ = x
}

func _() {
	x := 0
	if b {
		x = 0
	}
	if b {
		x = 0
	}
	_ = x
}

func _() {
	x := 1 //x
	if b {
		x = 0 //x
		x = 0 //x
	}
	if b {
		x = 0 //x
	}
	x = 0
	_ = x
}

func _() {
	x := 0
	if b {
		x = 0 //x
		x = 0
	}
	if b {
		x = 0
	}
	_ = x
}

func _() {
	x := 0
	for {
		_ = x
		x = 0
	}
}

func _() {
	x := 0
	for {
		_ = x
		x = 0 //x
		x = 0
	}
}

func _() {
	x := 0
	for {
		x += 0 //x
		x = 0
	}
}

func _() {
	x := 0
	for {
		x++ //x
		x = 0
	}
}

func _() {
	x := 0
	for {
		x++
	}
}

func _() {
	x := 0
	_ = &x
	x = 0
}

func _() {
	x := 0
	_ = x.f
	x = 0
}

func _() {
	x := 0
	_ = &x[0]
	x = 0
}

func _() {
	x := 0
	_ = x[:]
	x = 0
}

func _() {
	x := 0
	func() {
		_ = x
	}()
	x = 0
}

func _() {
	x := 0
	func() {
		x++
	}()
	x = 0
}

func _() {
	x := 0
	func() {
		x += 0
	}()
	x = 0
}

func _() {
	x := 0
	func() {
		x = 0
	}()
	_ = x
}

func _() {
	x := 0
	_ = x
	func() {
		x = 0
	}()
}

func _() (x int) {
	x = 0
	return
}

func _() (x int) {
	x = 0 //x
	x = 0
	return
}

func _() (x int) {
	x = 0 //x
	return 0
}

func _() (x int) {
	x = 0
	return x
}

func _() (x int) {
	x = 1
	anyFunctionMightPanic()
	return 2
}

func _() (x int) {
	x = 1
	_ = a[i]
	return 2
}

func _() (x int) {
	x = 1
	_ = a[i:j]
	return 2
}

func _() (x int) {
	x = 1
	_ = a == b
	return 2
}

func _() (x int) {
	x = 1
	_ = a / b
	return 2
}

func _() (x int) {
	x = 1
	a /= b
	return 2
}

func _() (x int) {
	x = 1
	_ = a % b
	return 2
}

func _() (x int) {
	x = 1
	a %= b
	return 2
}

func _() (x int) {
	x = 1
	_ = a.b
	return 2
}

func _() (x int) {
	x = 1
	_ = *a
	return 2
}

func _() (x int) {
	x = 1
	_ = a.(int)
	return 2
}

func _() (x int) {
	x = 1
	a <- b
	return 2
}

func _() {
	global = 0
}

var global int

func _() {
	global = 0
	global = 0
}

func _() {
	var x int
	if b {
		x = 0
	} else {
		x = 0
	}
	_ = x
}

func _() {
	var x int
	switch b {
	case true:
		x = 0
	case false:
		x = 0
	}
	_ = x
}

func _() {
	var x int
	switch b {
	default:
		x = 0 //x
		fallthrough
	case b:
	}
}

func _() {
	var x int
	switch b {
	default:
		x = 0
		fallthrough
	case b:
		_ = x
	}
}

func _() {
	var x int
	switch interface{}(b).(type) {
	case bool:
		x = 0
	case int:
		x = 0
	}
	_ = x
}

func _() {
	var x int
	var ch chan int
	select {
	case ch <- 0:
		x = 0
	case <-ch:
		x = 0
	}
	_ = x
}

func _() {
	var x int
	var ch chan int
	select {
	case ch <- 0:
		x = 0 //x
	case <-ch:
		x = 0 //x
	default:
		_ = x
	}
}

func _() {
	x := 1 //x
	var ch chan int
	select {
	case ch <- 0:
		x = 0
	case <-ch:
		x = 0
	default:
		x = 0
	}
	_ = x
}

func _() {
	x := 0
	var ch chan int
	select {
	case ch <- 0:
		x = 0
	case <-ch:
		x = 0
	}
	_ = x
}

func _() (x int) {
	if b {
		x = 0
		return
	}
	x = 0
	return
}

func _() {
	var x int
	if b {
		x = 0
	} else if x = 0; b {

	}
	_ = x
}

func _() {
	var x int
	if b {
		x = 0 //x
	}
	if x = 0; b {

	}
	_ = x
}

func _() {
	var x int
	if b {
		x = 0
	} else if b {
		x = 0
	}
	_ = x
}

func _() {
	var x int
	if b {
		x = 0
	} else {
		x = 0
	}
	_ = x
}

func _() {
	x := 0
	for b {
		_ = x
		x = 0
	}
	x = 0
	_ = x
}

func _() {
	x := 0
	for {
		_ = x
		x = 0
		if b {
			break
		}
		x = 0
	}
	_ = x
}

func _() {
	x := 0
	for x < 0 {
		x = 0 //x
		if b {
			break
		}
		x = 0
	}
}

func _() {
	x := 0
	for {
		_ = x
		x = 0
		if b {
			continue
		}
		x = 0
	}
}

func _() {
	var x int
	for x = range []int{} {
		_ = x
		x = 0
		if b {
			continue
		} else {
			break
		}
	}
	_ = x
}

func _() {
	var x int
	for {
		if b {
			x = 0 //x
			break
		}
		_ = x
	}
}

func _() {
	var x int
loop:
	for b {
		_ = x
		for b {
			x = 0
			break loop
		}
		x = 0
	}
	_ = x
}

func _() {
	var x int
	for range []int{} {
		if b {
			x = 0
			continue
		}
		x = 0
	}
	_ = x
}

func _() {

	x := 1
	y := 2

	if y != 0 {
		x = y // x
		return
	}

	_ = x
}
