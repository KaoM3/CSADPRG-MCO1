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
)

type Date struct {
	Year int
	Month int
	Day int
}

type Tweet struct {
	Date_created Date
	Text string
	Tokens []string
	Word_count int
}

type corpusAnalysis struct {
	word_count int
	vocabulary_size int
	word_frequency map[string]int
	character_frequency map[rune]int
	most_frequent_words map[string]int
	common_stop_words map[string]int
}

func getFilePath() string {
	var filepath string
	fmt.Println("Enter file path of CSV file")
	fmt.Scanln(&filepath)
	return filepath
}

func main() {
	filename := "C:\\Users\\Rafael\\Documents\\GitHub\\CSADPRG-MCO2\\fake_tweets.csv"
	records, err := ReadCSV(filename)
	if(err != nil) {
		log.Fatal("Cannot open and read CSV file")
	}
	
	wordCount := 0
	var tweetSlice []Tweet

	for _, eachrecord := range records {
		currTweet := ParseRecord(eachrecord)
		fmt.Println(currTweet)
		wordCount += currTweet.Word_count
		tweetSlice = append(tweetSlice, currTweet)
	}
	fmt.Println(wordCount)
	GetWordFrequency(tweetSlice)
	GetCharacterFrequency(tweetSlice)
	fmt.Println("STOP WORDS")
	fmt.Println(GetCountStopWords(tweetSlice))

	// Handle the route for displaying the bar chart
	RenderBarChart()
}
