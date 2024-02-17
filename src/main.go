package main

import (
	"strconv"

	"github.com/astaxie/beego"
)

func main(){
    beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {

	// Get values from route params
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	// Set template values
	c.Data["operation"] = operation
	c.Data["num1"] = num1
	c.Data["2"] = num2
	c.TplName = "result.html"

	// Perform operation
	switch(operation){
	case "sum":
		c.Data["result"] = add(num1, num2)
	case "product":
		c.Data["result"] = multiply(num1, num2)
	}
}

func add(num1, num2 int) int {
	return num1 + num2
}

func multiply(num1, num2 int) int {
	return num1 * num2
}