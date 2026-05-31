# tryx ![](https://github.com/fatih/color/workflows/build/badge.svg)

A lightweight, dead-simple Try–Catch–Finally exception handling pattern — in Go.

# Installation
```
go get github.com/davyn-zhenwf/tryx
```

# Example
```
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