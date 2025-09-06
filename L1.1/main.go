package main

import "fmt"

// моя реализация структур
type Action struct {
    Human
    Name   string
    Target string
}

type Human struct {
    Name   string
    Age    int
    Gender string
}

func (h *Human) Greet() string {
    return "Hello, my name is " + h.Name
}

func (h *Human) IsAdult() bool {
    return h.Age >= 18
}

func (a *Action) Describe() string {
    return "Action: " + a.Name + ", Target: " + a.Target + ", Human: " + a.Greet()
}

func main() {
    a := &Action{
        Human:  Human{Name: "Ivan", Age: 30},
        Name:   "Speak",
        Target: "Audience",
    }

    fmt.Println(a.Greet())    // Hello, my name is Ivan
    fmt.Println(a.Describe()) // Action: Speak, Target: Audience, Human: Hello, my name is Ivan
    fmt.Println("Is Adult:", a.IsAdult())
}

