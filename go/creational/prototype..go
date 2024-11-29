package main

import "fmt"

// prototype interface
type DocumentPrototype interface {
	Clone() DocumentPrototype
}

// concerete prototype implementation
type Invoice struct {
	item  string
	price float64
	amout int
}

func (i *Invoice) Clone() DocumentPrototype {
	return &Invoice{
		item:  i.item,
		price: i.price,
		amout: i.amout,
	}
}

func (i *Invoice) String() string {
	return fmt.Sprintf("Invoice - Item: %s, Price: $%.2f, Amount: %d", i.price, i.price, i.amout)
}

type OrderReport struct {
	id          uint64
	totalAmount int
}

func (r *OrderReport) Clone() DocumentPrototype {
	return &OrderReport{
		id:          r.id,
		totalAmount: r.totalAmount,
	}
}

func (i *OrderReport) String() string {
	return fmt.Sprintf("Report - Id: %d, Total Amount: %d", i.id, i.totalAmount)
}

func main() {

	invoice := Invoice{}

	newInvoice := invoice.Clone().(*Invoice)
	newInvoice.item = "P02"
	newInvoice.amout = 5
	newInvoice.price = 2000.0

	fmt.Println(newInvoice)

	report := OrderReport{}

	newReport := report.Clone().(*OrderReport)
	newReport.id = 5
	newReport.totalAmount = 260

	fmt.Println(newReport)
}
