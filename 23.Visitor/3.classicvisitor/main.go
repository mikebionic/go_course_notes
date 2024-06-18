package main

import (
	"fmt"
	"strings"
)

// Classic double dispatch visitor

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAdditionExpression(e *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (d *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(d)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(de *DoubleExpression) {
	ee.result = de.value
}
func (ee *ExpressionEvaluator) VisitAdditionExpression(de *AdditionExpression) {
	de.left.Accept(ee)
	x := ee.result
	de.right.Accept(ee)
	x += ee.result
	ee.result = x
}

func main() {
	// 1 + (2+3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	e.Accept(ep)
	fmt.Println(ep.String())

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g \n", ep.String(), ee.result)
}
