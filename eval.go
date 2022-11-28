package main

import (
	"fmt"

	"github.com/Pramod-Devireddy/go-exprtk"
)

func main() {
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	exprtkObj.SetExpression("(x^2+1)/2")

	exprtkObj.AddDoubleVariable("x")
	//exprtkObj.AddDoubleVariable("y")

	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	exprtkObj.SetDoubleVariableValue("x", 5)
	//exprtkObj.SetDoubleVariableValue("y", 32)

	fmt.Println(exprtkObj.GetEvaluatedValue())
}
