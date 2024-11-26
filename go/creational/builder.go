package main

import "fmt"

// product
type Report struct {
	title   string
	header  string
	content string
	footer  string
}

func (r *Report) Display() {
	fmt.Println("Report:")
	fmt.Println("Header: ", r.header)
	fmt.Println("Title: ", r.title)
	fmt.Println("Content: ", r.content)
	fmt.Println("Footer: ", r.footer)
}

// builder interface
type ReportBuilder interface {
	SetHeader(header string) ReportBuilder
	SetTitle(title string) ReportBuilder
	SetContent(content string) ReportBuilder
	SetFooter(footer string) ReportBuilder
	Build() *Report
}

// concerete builders
type PDFReportBuilder struct {
	report *Report
}

func NewPDFReportBuilder() *PDFReportBuilder {
	return &PDFReportBuilder{
		report: &Report{},
	}
}

func (pr *PDFReportBuilder) SetTitle(title string) ReportBuilder {
	pr.report.title = title
	return pr
}

func (pr *PDFReportBuilder) SetHeader(header string) ReportBuilder {
	pr.report.header = header
	return pr
}

func (pr *PDFReportBuilder) SetContent(content string) ReportBuilder {
	pr.report.content = content
	return pr
}

func (pr *PDFReportBuilder) SetFooter(footer string) ReportBuilder {
	pr.report.footer = footer
	return pr
}

func (pr *PDFReportBuilder) Build() *Report {
	return pr.report
}

type ExceleportBuilder struct {
	report *Report
}

func NewExcelReportBuilder() *ExceleportBuilder {
	return &ExceleportBuilder{
		report: &Report{},
	}
}

func (pr *ExceleportBuilder) SetTitle(title string) ReportBuilder {
	pr.report.title = title
	return pr
}

func (pr *ExceleportBuilder) SetHeader(header string) ReportBuilder {
	pr.report.header = header
	return pr
}

func (pr *ExceleportBuilder) SetContent(content string) ReportBuilder {
	pr.report.content = content
	return pr
}

func (pr *ExceleportBuilder) SetFooter(footer string) ReportBuilder {
	pr.report.footer = footer
	return pr
}

func (pr *ExceleportBuilder) Build() *Report {
	return pr.report
}

func main() {

	pdfBuilder := NewPDFReportBuilder()
	pdfReport := pdfBuilder.
		SetTitle("learning desing pattern").
		SetHeader("Builder pattern").
		SetContent("lorem ipsum..").
		SetFooter("desing pattern cookbook").
		Build()
	pdfReport.Display()

	excelBuilder := NewExcelReportBuilder()
	excelReport := excelBuilder.
		SetTitle("learning desing pattern").
		SetHeader("Builder pattern").
		SetContent("lorem ipsum..").
		SetFooter("desing pattern cookbook").
		Build()
	excelReport.Display()
}
