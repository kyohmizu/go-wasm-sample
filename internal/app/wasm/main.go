package main

import (
	"strconv"
	"syscall/js"
)

type Operator int

var (
	cachedNum = 0
	cached    = false
	operator  = None
)

const (
	None Operator = iota
	Add
	Sub
	Mul
	Div
)

func operate(this js.Value, i []js.Value) interface{} {
	if !cached {
		if operator != None {
			cachedNum = calc(cachedNum, toInt(js.Global().Get("document").Call("getElementById", "result").Get("textContent").String()), operator)
			js.Global().Get("document").Call("getElementById", "result").Set("textContent", cachedNum)
		} else {
			cachedNum = toInt(js.Global().Get("document").Call("getElementById", "result").Get("textContent").String())
		}
		cached = true
	}

	switch i[0].String() {
	case "+":
		operator = Add
	case "-":
		operator = Sub
	case "*":
		operator = Mul
	case "/":
		operator = Div
	default:
		operator = None
	}
	return nil
}

func toInt(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}

func calc(num1, num2 int, op Operator) int {
	switch op {
	case Add:
		return num1 + num2
	case Sub:
		return num1 - num2
	case Mul:
		return num1 * num2
	case Div:
		return num1 / num2
	default:
		//log.Fatal("unexpected operator")
		return num2
	}
}

func inputNum(this js.Value, i []js.Value) interface{} {
	currentNum := js.Global().Get("document").Call("getElementById", "result").Get("textContent").String()
	if cached || currentNum == "0" {
		currentNum = i[0].String()
		cached = false
	} else {
		currentNum += i[0].String()
	}
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", currentNum)
	return nil
}

func reverseSign(this js.Value, i []js.Value) interface{} {
	if cached {
		cachedNum *= -1
	}
	currentNum := toInt(js.Global().Get("document").Call("getElementById", "result").Get("textContent").String())
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", currentNum*-1)
	return nil
}

func reset(this js.Value, i []js.Value) interface{} {
	cachedNum = 0
	cached = false
	operator = None
	js.Global().Get("document").Call("getElementById", "result").Set("textContent", "0")
	return nil
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("inputNum", js.FuncOf(inputNum))
	js.Global().Set("operate", js.FuncOf(operate))
	js.Global().Set("reverseSign", js.FuncOf(reverseSign))
	js.Global().Set("reset", js.FuncOf(reset))
	<-c
}
