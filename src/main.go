// app's main

package main

import (
	"strconv"

	"github.com/astaxie/beego"
)

func main() {
	/* Routes for the app, such as:
	/sum/9/4
	/product/5/6
	*/
	beego.URLFor("./static")
	beego.Router("/:operation/:x:int/:y:int", &mainController{})
	beego.Router("/", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	// Get values of route parameters from the given route
	operation := c.Ctx.Input.Param(":operation")
	x, _ := strconv.Atoi(c.Ctx.Input.Param(":x"))
	y, _ := strconv.Atoi(c.Ctx.Input.Param(":y"))

	// Set the values to use in the template
	c.Data["operation"] = operation
	c.Data["x"] = x
	c.Data["y"] = y
	c.TplName = "result.html"

	// Perform a computation based on an 'operation' router parameter
	switch operation {
	case "sum":
		c.Data["result"] = add(x, y)
	case "product":
		c.Data["result"] = multiply(x, y)
	default:
		c.TplName = "404.html"
	}
}

func add(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}
