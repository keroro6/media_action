package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan //close 获取其余error会发这个channel
	Data       dataChan
	dataSize   int
	longLived  bool //是否长期存活的一个runner
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longLived bool, d fn, e fn) *Runner {
	return &Runner{Controller: make(controlChan, 1),
		Error:      make(controlChan, 1),
		Data:       make(dataChan, size),
		longLived:  longLived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,}
}

func (r *Runner) startDispatch() {

	//如果不是long lived 关闭channl
	defer func() {
		if ! r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXCUTE
				}
			}
			if c == READY_TO_EXCUTE { //如果准备执行就执行函数，等完成后就告诉可以调度了
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}
func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
