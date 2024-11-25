package main

import "fmt"

// abstract factory
type PaymentFactory interface {
	CreateCreditCardProcessor() PaymentProcessor
	CreateBankTransferProcessor() PaymentProcessor
}

// abstract product
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// concrete factories
type StripeCreditCardProcessor struct{}

func (s *StripeCreditCardProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Stripe: Processing credit card payment of $%.2f\n", amount)
}

type StripeBankTransferProcessor struct{}

func (s *StripeBankTransferProcessor) ProcessPayment(amount float64) {
	fmt.Printf("Stripe: Processing bank transfer of $%.2f\n", amount)
}

type PayPalCreditCardProcessor struct{}

func (p *PayPalCreditCardProcessor) ProcessPayment(amount float64) {
	fmt.Printf("PayPal: Processing credit card payment of $%.2f\n", amount)
}

type PayPalBankTransferProcessor struct{}

func (p *PayPalBankTransferProcessor) ProcessPayment(amount float64) {
	fmt.Printf("PayPal: Processing bank transfer of $%.2f\n", amount)
}

// concrete products
type StripeFactory struct{}

func (s *StripeFactory) CreateCreditCardProcessor() PaymentProcessor {
	return &StripeCreditCardProcessor{}
}

func (s *StripeFactory) CreateBankTransferProcessor() PaymentProcessor {
	return &StripeBankTransferProcessor{}
}

type PayPalFactory struct{}

func (p *PayPalFactory) CreateCreditCardProcessor() PaymentProcessor {
	return &PayPalCreditCardProcessor{}
}

func (p *PayPalFactory) CreateBankTransferProcessor() PaymentProcessor {
	return &PayPalBankTransferProcessor{}
}

func ProcessPayments(factory PaymentFactory, amount float64) {
	creditCardProcessor := factory.CreateCreditCardProcessor()
	bankTransferProcessor := factory.CreateBankTransferProcessor()

	creditCardProcessor.ProcessPayment(amount)
	bankTransferProcessor.ProcessPayment(amount * 2)
}

func main() {
	var factory PaymentFactory

	gateway := "stripe"

	if gateway == "stripe" {
		factory = &StripeFactory{}
	} else if gateway == "paypal" {
		factory = &PayPalFactory{}
	} else {
		fmt.Println("Unsupported payment gateway")
		return
	}

	ProcessPayments(factory, 100.00)
}
