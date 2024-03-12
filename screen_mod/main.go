package main

import (
    "fmt"
    "time"

    "github.com/inancgumus/screen"
)

func main() {
    screen.Clear()

    var currentTime time.Time
    var lastTime time.Time = time.Now()
    for {
        currentTime = time.Now()
        if currentTime.Second() != lastTime.Second() {
            screen.MoveTopLeft()

            fmt.Println(currentTime)
            lastTime = currentTime
        }
        time.Sleep(time.Millisecond * 20)
    }
}
