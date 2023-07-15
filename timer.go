package main

import (
	"os/exec"
	"fmt"
	"time"
)



func main() {

fmt.Println("How long should the countdown be?")
var secs float64
fmt.Scanln(&secs)
duration := time.Duration(secs) * (time.Second)

for duration != 0 {

fmt.Println(duration)
time.Sleep(time.Second)
duration -= time.Second

 
  }

if duration == 0 {
  play := exec.Command("mpg123", "~/alarm.mp3")
stuff, err := play.Output()
	if err != nil {
fmt.Println(err)
    }
  
fmt.Println(string(stuff))
fmt.Println("Time is up!")
    }
  }
