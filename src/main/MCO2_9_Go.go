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
	stop_word_count map[string]int
	tweet_frequency map[int]map[string]int
}

func getFilePath() string {
	var filepath string
	fmt.Println("Enter file path of CSV file")
	fmt.Scanln(&filepath)
	return filepath
}

func main() {
	// Reading file
	// TODO: Replace with getting input of file location
	filename := "C:\\Users\\Rafael\\Documents\\GitHub\\CSADPRG-MCO2\\fake_tweets.csv"
	records, err := ReadCSV(filename)
	if(err != nil) {
		log.Fatal("Cannot open and read CSV file")
	}
	
	// Data Analysis
	wordCount := 0
	var tweetSlice []Tweet

	for _, eachrecord := range records {
		currTweet := ParseRecord(eachrecord)
		fmt.Println(currTweet)
		wordCount += currTweet.Word_count
		tweetSlice = append(tweetSlice, currTweet)
	}

	dataAnalysis := corpusAnalysis{
		word_count: 			wordCount,
		vocabulary_size: 		len(GetWordFrequency(tweetSlice)),
		word_frequency:			GetWordFrequency(tweetSlice),
		character_frequency:	GetCharacterFrequency(tweetSlice),
		stop_word_count:		GetCountStopWords(tweetSlice),
		tweet_frequency:		GetTweetFrequency(tweetSlice),
	}

	fmt.Print(dataAnalysis)
	// TODO: Order then render

	RenderWordCloud(GetMostFrequentWords(dataAnalysis.word_frequency, 20), "word-cloud.html")
	RenderTweetFrequency(dataAnalysis.tweet_frequency, "tweet-frequency-chart.html")
	RenderSymbolPieChart(GetSymbols(dataAnalysis.character_frequency), "symbol-frequency-chart.html")
}
