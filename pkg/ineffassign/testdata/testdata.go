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
	x = 0 // want "ineffectual assignment to x"
}

func _() {
	var x int
	x = 0 // want "ineffectual assignment to x"
	x = 0
	_ = x
}

func _() {
	x := 0
	x = 0
	_ = x
}

func _() {
	x := false
	x = true
	_ = x
}

func _() {
	x := true // want "ineffectual assignment to x"
	x = false
	_ = x
}

func _() {
	false := "not the real false"
	x := false // want "ineffectual assignment to x"
	x = "also not false"
	_ = x
}

func _() {
	x := int(0)
	x = 0
	_ = x
}

func _() {
	x := ""
	x = "abc"
	_ = x
}

func _() {
	x := "abc" // want "ineffectual assignment to x"
	x = "def"
	_ = x
}

func _() {
	x := 1 // want "ineffectual assignment to x"
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
	x := 1 // want "ineffectual assignment to x"
	if b {
		x = 0 // want "ineffectual assignment to x"
	}
	x = 0
	_ = x
}

func _() {
	x := 1 // want "ineffectual assignment to x"
	for b {
		x = 0 // want "ineffectual assignment to x"
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
	x := 1 // want "ineffectual assignment to x"
	if b {
		x = 0 // want "ineffectual assignment to x"
		x = 0 // want "ineffectual assignment to x"
	}
	if b {
		x = 0 // want "ineffectual assignment to x"
	}
	x = 0
	_ = x
}

func _() {
	x := 0
	if b {
		x = 0 // want "ineffectual assignment to x"
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
		x = 0 // want "ineffectual assignment to x"
		x = 0
	}
}

func _() {
	x := 0
	for {
		x += 0 // want "ineffectual assignment to x"
		x = 0
	}
}

func _() {
	x := 0
	for {
		x++ // want "ineffectual assignment to x"
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
	type T struct{ f int }
	x := T{}
	_ = x.f
	x = T{}
}

func _() {
	x := []int{}
	_ = &x[0]
	x = []int{}
}

func _() {
	x := []int{}
	_ = x[:]
	x = []int{}
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

func _() {
	x := 0
	x = 0 // want "ineffectual assignment to x"
	x = 0
	_ = func() {
		x = 0

		_ = x
	}
	_ = x
}

func _() {
	x := 0
	x = 0 // want "ineffectual assignment to x"
	x = 0
	_ = func() {
		_ = x
	}
	x = 0
	_ = x
}

func _() {
	x := 0
	x = 0 // want "ineffectual assignment to x"
	x = 0
	_ = &x
}

func _() {
	x := []int{} // want "ineffectual assignment to x"
	x = []int{}
	_ = &x[0]
}

func _() {
	x := []int{} // want "ineffectual assignment to x"
	x = []int{}
	_ = x[:]
}

func _() (x int) {
	x = 0
	return
}

func _() (x int) {
	x = 0 // want "ineffectual assignment to x"
	x = 0
	return
}

func _() (x int) {
	x = 0 // want "ineffectual assignment to x"
	return 0
}

func _() (x int) {
	x = 0
	return x
}

func _(anyFunctionMightPanic func()) (x int) {
	x = 1
	defer func() {
		recover()
	}()
	anyFunctionMightPanic()
	return 2
}

func _(anyFunctionMightPanic func()) (x int) {
	x = 1 // want "ineffectual assignment to x"
	anyFunctionMightPanic()
	return 2
}

func _(a []int) (x int) {
	x = 1
	defer func() {
		recover()
	}()
	_ = a[1]
	return 2
}

func _(a []int) (x int) {
	x = 1
	defer func() {
		recover()
	}()
	_ = a[2:4]
	return 2
}

func _(a, b interface{}) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a == b
	return 2
}

func _(a, b int) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a / b
	return 2
}

func _(a, b int) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a / b
	return 2
}

func _(a, b int) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a % b
	return 2
}

func _(a, b int) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a % b
	return 2
}

func _(a *struct{ b int }) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a.b
	return 2
}

func _(a *int) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = *a
	return 2
}

func _(a interface{}) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	_ = a.(int)
	return 2
}

func _(a chan int) (x int) {
	defer func() {
		recover()
	}()
	x = 1
	a <- 1
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
		x = 0 // want "ineffectual assignment to x"
		fallthrough
	case b:
	}
	x = 0
	_ = x
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
		x = 0 // want "ineffectual assignment to x"
	case <-ch:
		x = 0 // want "ineffectual assignment to x"
	default:
		_ = x
	}
}

func _() {
	x := 1 // want "ineffectual assignment to x"
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
		x = 0 // want "ineffectual assignment to x"
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
		x = 0 // want "ineffectual assignment to x"
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
			x = 0 // want "ineffectual assignment to x"
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
	x := 0
	if b {
		x = 1 // want "ineffectual assignment to x"
		return
	}
	_ = x
}

func returnErr() error {
	return nil
}

func _() (err error) {
	if 1 > 2 {
		err = returnErr() // want "ineffectual assignment to err"
	}
	err = returnErr()
	return
}
