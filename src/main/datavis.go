/*
********************
Last name: Bernandino
Language: Go
Paradigm(s): Procedural, Multi-paradigm
********************
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func RenderWordCloud(wordFrequencyMap map[string]int, filename string) {
	var cloudData []opts.WordCloudData

	for key, value := range wordFrequencyMap {
		cloudData = append(cloudData, opts.WordCloudData{
			Name: key,
			Value: float64(value),
		})
	}

	// Create new word cloud
	wordCloud := charts.NewWordCloud()
	wordCloud.AddSeries("wordcloud", cloudData)

	// Create the file to save the chart as HTML
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Render the chart into the HTML file
	if err := wordCloud.Render(f); err != nil {
		log.Fatal(err)
	}

	log.Println("Chart saved as", filename)
}

func RenderTweetFrequency(tweetFrequencyMap map[int]map[string]int, filename string) {
	monthNames := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var seriesNames []string
	var barData [][]opts.BarData // 2D Array: Rows as years, Columns as Months

	// Iterate over the tweetFrequencyMap to prepare the data
	for year, monthData := range tweetFrequencyMap {
		var monthlyCounts []opts.BarData
		for _, month := range monthNames {
			// If there's no tweet data for this month, use 0
			count := float64(monthData[month]) // Convert to float64
			monthlyCounts = append(monthlyCounts, opts.BarData{
				Value: count,
			})
		}
		barData = append(barData, monthlyCounts)
		seriesNames = append(seriesNames, fmt.Sprintf("%d", year))
	}

	// Create new bar
	bar := charts.NewBar()
	bar.SetXAxis(monthNames[:])

	for i, data := range barData {
		bar.AddSeries(seriesNames[i], data)
	}

	// Create the file to save the chart as HTML
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Render the chart into the HTML file
	if err := bar.Render(f); err != nil {
		log.Fatal(err)
	}

	log.Println("Chart saved as", filename)
}
