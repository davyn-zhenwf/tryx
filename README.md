# tryx ![](https://github.com/fatih/color/workflows/build/badge.svg)

A lightweight, dead-simple Try–Catch–Finally exception handling pattern for Go. 

Supports <b><u>chained catches</b></u> and <b><u>standalone try</b></u>.

# Installation
```
go get github.com/davyn-zhenwf/tryx
```

# Examples
### Chained Catches
```go
package main

import (
	"fmt"
	"github.com/davyn-zhenwf/tryx"
)

func main() {
	tryx.Try(func() {
		fmt.Println("☆ Medal Panic Chain Demo")
		panic("🥇")

	}).Catch(func(ex any) {
		fmt.Println("├─ 1st panic caught:", ex)
		panic("🥈")

	}).Catch(func(ex any) {
		fmt.Println("├─ 2nd panic caught:", ex)
		panic("🥉")

	}).Catch(func(ex any) {
		fmt.Println("├─ 3rd panic caught:", ex)

	}).Finally(func() {
		fmt.Println("└─ Completed — All panics caught! 🥇🥈🥉")
	})
}
```
```
☆ Medal Panic Chain Demo
├─ 1st panic caught: 🥇
├─ 2nd panic caught: 🥈
├─ 3rd panic caught: 🥉
└─ Completed — All panics caught! 🥇🥈🥉
```
### Standalone Try
```go
package main

import (
	"fmt"
	"github.com/davyn-zhenwf/tryx"
)

func main() {
	tryx.Try(func() {
		fmt.Println("☆ Standalone Try Demo")
		panic("🥇")
	})

  fmt.Println("└─ Recovered — Program continues normally!")
}
```
```
☆ Standalone Try Demo
└─ Recovered — Program continues normally!
```