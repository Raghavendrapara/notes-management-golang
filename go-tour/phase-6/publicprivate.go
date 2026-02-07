package main

import (
	"context"
	"fmt"
	"time"
)

type PMS interface {
	book()
	cancel()
	change()
}

//	type User struct {
//		Name         string
//		password     string
//		items        map[string]string
//		appointments map[string]Appointment
//	}
type Appointment struct {
	DoctorName  Doctor
	BookingTime string
}

type Doctor struct {
	Name string
	Fees int
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println("helo")
	<-ctx.Done()

	user := &User{Balance: 500}

	stripe := &Stripe{}

	ProcessCheckout(user, stripe)

	paypal := &PayPal{}

	ProcessCheckout(user, paypal)

	v := 23425453434666675
	interfaceSwitch(v)
}

//Payments

type PaymentGateway interface {
	Pay(amount float64) bool
}

type Stripe struct{}

func (s *Stripe) Pay(amount float64) bool {
	fmt.Println("From Stripe")
	return true
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) bool {
	fmt.Println("From PayPal")
	return true
}

func ProcessCheckout(u *User, gateway PaymentGateway) {

	price := 500.0

	if gateway.Pay(price) {
		err := u.Charge(price)
		if err == nil {
			fmt.Println("Success")
		} else {
			fmt.Println(err)
		}
	}
}

type User struct {
	ID      int
	Balance float64
}

func (u *User) Charge(amount float64) error {
	if u.Balance < amount {
		return fmt.Errorf("Insufficient")
	}

	u.Balance = u.Balance - amount

	return nil
}

func interfaceSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer : %d \n", v*2)
	case string:
		fmt.Printf("I am a string of len %d\n", len(v))
	case User:
		fmt.Printf("I am a User object: \n", v.ID)
	default:
		fmt.Printf("I don't know what I am\n")

	}
}
