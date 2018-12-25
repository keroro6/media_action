package taskrunner

import (
	"testing"
	"time"
	"fmt"
	"github.com/pkg/errors"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			fmt.Printf("Dispatcher send:%d\n", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				fmt.Printf("Executor received: %v\n", d)
			default:
				break forloop
			}
		}
		return errors.New("error")
	}

	runner := NewRunner(30, false, d, e)
	go runner.StartAll()

	time.Sleep(time.Second * 3)

}
