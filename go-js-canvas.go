package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
)

func main() {

	//js.Global.Get("document").Call("write", "Test Canvas")

	document := js.Global.Get("document")
	canvas := document.Call("getElementById", "cnvs")

	DrawLine(canvas, 80, 22, 110, 22)

	DrawNode(canvas, 10, 10, 1)
	DrawNode(canvas, 40, 10, 2)
	DrawNode(canvas, 70, 10, 3)
	DrawNode(canvas, 100, 10, 4)
}

func DrawNode(canvas *js.Object, x, y, id int) {
	ctx := canvas.Call("getContext", "2d")
	// TODO get this from node color
	ctx.Set("fillStyle", "#AA00AA")
	ctx.Call("fillRect", x, y, 20, 25)
	ctx.Set("font", "18px Arial")
	ctx.Set("fillStyle", "#FFFFFF")
	// TODO set to node id
	str := fmt.Sprint(id)
	ctx.Call("fillText", str, x+5, y+20)
}

func DrawLine(canvas *js.Object, x1, y1, x2, y2 int) {
	ctx := canvas.Call("getContext", "2d")
	ctx.Set("strokeStyle", "#000000")
	ctx.Call("moveTo", x1, y1)
	ctx.Call("lineTo", x2, y2)
	ctx.Call("stroke")
}
