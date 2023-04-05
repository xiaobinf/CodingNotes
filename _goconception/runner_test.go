package _goconception

import (
  "os"
  "os/signal"
  "time"
)

type Runner struct {
  interrupt chan os.Signal
  timeout   <-chan time.Time
  timeout1  time.Time
}

var (
  interrupt chan os.Signal
  timeout   <-chan int
  timeout1  time.Time
  _time     = time.After(time.Duration(1))
  _time1    = time.Time{}
  timeChan  = make(chan time.Time, 10)
)

func A() {
  signal.Notify(interrupt, os.Interrupt)

}
