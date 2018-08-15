package main

import (
	"fmt"
	"runtime"

	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

type Order struct {
	ID uint
	transition.Transition
}

func main() {
	var order Order
	var OrderStateMachine = transition.New(&Order{})
	// Define initial state
	OrderStateMachine.Initial("draft")
	// Define States
	OrderStateMachine.State("validated").Enter(enterState).Exit(exitState)
	OrderStateMachine.State("checkedout").Enter(enterState).Exit(exitState)
	OrderStateMachine.State("paid").Enter(enterState).Exit(exitState)
	OrderStateMachine.State("cancelled").Enter(enterState).Exit(exitState)
	OrderStateMachine.State("surveyed").Enter(enterState).Exit(exitState)

	// Define Events
	OrderStateMachine.Event("start").To("validated").From("draft").Before(beforeTransition).After(afterTransition)
	OrderStateMachine.Event("buy").To("checkedout").From("draft", "validated")
	OrderStateMachine.Event("pay").To("paid").From("checkedout")
	OrderStateMachine.Event("cancel").To("cancelled").From("checkedout", "draft", "paid").Before(beforeTransition).After(afterTransition)

	OrderStateMachine.Trigger("start", &order, nil)
	OrderStateMachine.Trigger("buy", &order, nil)
	OrderStateMachine.Trigger("pay", &order, nil)
	OrderStateMachine.Trigger("start", &order, nil)
	OrderStateMachine.Trigger("cancel", &order, nil)

	// fmt.Printf("The value of order = %s \n", order.State)

}

func enterState(intf interface{}, tx *gorm.DB) error {
	fmt.Printf("enter state \t")
	order := intf.(*Order)
	fmt.Printf("%s\tState=%s\n", trace(), order.GetState())
	return nil
}

func exitState(intf interface{}, tx *gorm.DB) error {
	fmt.Printf("exit state \t")
	order := intf.(*Order)
	fmt.Printf("%s\tState=%s\n", trace(), order.GetState())
	return nil
}

func beforeTransition(intf interface{}, tx *gorm.DB) error {
	fmt.Printf("before transition \t")
	order := intf.(*Order)
	fmt.Printf("%s\tState=%s\n", trace(), order.GetState())
	return nil
}

func afterTransition(intf interface{}, tx *gorm.DB) error {
	fmt.Printf("after transition \t")
	order := intf.(*Order)
	fmt.Printf("%s\tState=%s\n", trace(), order.GetState())
	return nil
}

func trace() string {
	return ""
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	// return fmt.Sprintf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
	return fmt.Sprintf(":%d %s\n", frame.Line, frame.Function)
}
