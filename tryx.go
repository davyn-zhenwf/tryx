package tryx

type panicHandler struct {
	PanicValue interface{}
}

func Try(tryFn func()) (ph *panicHandler) {
	ph = &panicHandler{}
	defer func() {
		ph.PanicValue = recover()
	}()

	tryFn()
	return
}

func (this *panicHandler) Catch(catchFn func(ex interface{})) (ph *panicHandler) {
	ph = &panicHandler{}
	defer func() {
		ph.PanicValue = recover()
	}()

	if this.PanicValue != nil {
		catchFn(this.PanicValue)
	}

	return
}

func (this *panicHandler) Finally(finallyFn func()) {
	finallyFn()
}
