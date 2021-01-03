package visitor

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Visitor interface {
	VisitCabinet(Equipment)
	VisitMotherBoard(Equipment)
	VisitCPU(Equipment)
	VisitRAM(Equipment)
	VisitVideoCard(Equipment)
	VisitHDD(Equipment)
}

type VisitorGPower struct {
	total Watt
}

func NewVisitorGPower() *VisitorGPower {
	return &VisitorGPower{}
}

func (v *VisitorGPower) VisitCabinet(e Equipment) {
	v.total += e.GPower()
}

func (v *VisitorGPower) VisitMotherBoard(e Equipment) {
	v.total += e.GPower()
}

func (v *VisitorGPower) VisitCPU(e Equipment) {
	v.total += e.GPower()
}

func (v *VisitorGPower) VisitRAM(e Equipment) {
	v.total += e.GPower()
}

func (v *VisitorGPower) VisitVideoCard(e Equipment) {
	v.total += e.GPower()
}

func (v *VisitorGPower) VisitHDD(e Equipment) {
	v.total += e.GPower()
}

func (v *VisitorGPower) GPower() int {
	return int(v.total)
}

type VisitorBill struct {
	tw            *tabwriter.Writer
	totalPrice    Price
	totalDiscount Price
}

func NewVisitorBill() *VisitorBill {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "Name\tPrice")
	return &VisitorBill{tw, NewPrice(0), NewPrice(0)}
}

func (v *VisitorBill) println(label, name string, price Price) {
	fmt.Fprintf(v.tw, "%s: %s\t%s\n", label, name, price.String())
}

func (v *VisitorBill) accumulate(price, discount Price) {
	v.totalPrice = v.totalPrice.Add(price)
	v.totalDiscount = v.totalDiscount.Add(discount)
}

func (v *VisitorBill) VisitCabinet(e Equipment) {
	v.println("Cabinet", e.PrintName(), e.NetPrice())
	v.accumulate(e.NetPrice(), e.DiscountPrice())
}

func (v *VisitorBill) VisitMotherBoard(e Equipment) {
	v.println("Mother Board", e.PrintName(), e.NetPrice())
	v.accumulate(e.NetPrice(), e.DiscountPrice())
}

func (v *VisitorBill) VisitCPU(e Equipment) {
	v.println("CPU", e.PrintName(), e.NetPrice())
	v.accumulate(e.NetPrice(), e.DiscountPrice())
}

func (v *VisitorBill) VisitRAM(e Equipment) {
	v.println("RAM", e.PrintName(), e.NetPrice())
	v.accumulate(e.NetPrice(), e.DiscountPrice())
}

func (v *VisitorBill) VisitVideoCard(e Equipment) {
	v.println("Video Card", e.PrintName(), e.NetPrice())
	v.accumulate(e.NetPrice(), e.DiscountPrice())
}

func (v *VisitorBill) VisitHDD(e Equipment) {
	v.println("HDD", e.PrintName(), e.NetPrice())
	v.accumulate(e.NetPrice(), e.DiscountPrice())
}

func (v *VisitorBill) PrintBill() {
	fmt.Fprintln(v.tw, "")
	fmt.Fprintf(v.tw, "Discount:\t%s\n", v.totalDiscount.String())
	fmt.Fprintf(v.tw, "Total:\t%s\n", v.totalPrice.String())
	fmt.Fprintf(v.tw, "Amount Due:\t%s\n", v.totalPrice.Sub(v.totalDiscount))
	v.tw.Flush()
}
