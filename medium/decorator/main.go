package main

import (
	"fmt"
	"log"
	"strings"
)

type DataSource interface {
	fetchData() string
}

type FileDataSource struct{}

func (*FileDataSource) fetchData() string {
	return "Hello, World!"
}

type DataDourceDecorator struct {
	wrappee DataSource
}

func NewDataSourceDecorator(wrappee DataSource) *DataDourceDecorator {
	return &DataDourceDecorator{wrappee: wrappee}
}
func (dd *DataDourceDecorator) fetchData() string {
	return dd.wrappee.fetchData()
}

type LoggingDecorator struct {
	wrappee DataSource
}

func NewLoggingDecorator(wrappee DataSource) *LoggingDecorator {
	return &LoggingDecorator{wrappee: wrappee}
}
func (ld *LoggingDecorator) fetchData() string {
	fmt.Println("[Logging] About to fetch data...")
	data := ld.wrappee.fetchData()
	fmt.Printf("[Logging] Fetched data: %q\n", data)
	return data
}

type UppercaseDecorator struct {
	wrappee DataSource
}

func NewUppercaseDecorator(wrappee DataSource) *UppercaseDecorator {
	return &UppercaseDecorator{wrappee: wrappee}
}
func (ud *UppercaseDecorator) fetchData() string {
	data := ud.wrappee.fetchData()
	return strings.ToUpper(data)
}

func main() {
	fmt.Println("=== DECORATOR ===")

	dataSource := NewDataSourceDecorator(&FileDataSource{})
	log.Printf("[+] Raw output: %q", dataSource.fetchData())

	fmt.Println("..................................................")
	loggingSource := NewLoggingDecorator(dataSource)
	log.Printf("[+] Logged output: %q", loggingSource.fetchData())

	fmt.Println("..................................................")
	uppercaseSource := NewUppercaseDecorator(loggingSource)
	log.Printf("[+] Transformed output: %q", uppercaseSource.fetchData())
}
