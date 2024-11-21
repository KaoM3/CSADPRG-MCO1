/*
********************
Last name: Bernandino
Language: Go
Paradigm(s): Procedural, Multi-paradigm
********************
*/
package main

import (
	"log"
	"os"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func RenderBarChart() {
	barchart := charts.NewBar()
	xAxis := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// Set up the y-axis data (values for the bars)
	yAxis := []opts.BarData{
		{Value: 120}, {Value: 132}, {Value: 101}, {Value: 134}, {Value: 90},
		{Value: 115}, {Value: 160}, {Value: 110}, {Value: 130}, {Value: 140},
		{Value: 180}, {Value: 150},
	}

	// Add the bar series to the chart
	barchart.AddSeries("Sales", yAxis).SetXAxis(xAxis)
	barchart.AddSeries("OtherSales", yAxis).SetXAxis(xAxis)

	// Create the file to save the chart as HTML
	f, err := os.Create("chart.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Render the chart into the HTML file
	if err := barchart.Render(f); err != nil {
		log.Fatal(err)
	}

	log.Println("Chart saved as chart.html")
}