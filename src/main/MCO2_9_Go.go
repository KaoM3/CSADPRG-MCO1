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
	"strings"
	"encoding/csv"
	"sort"
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
	filepath = strings.ReplaceAll(filepath, `"`, "")
	return filepath
}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	reader.Read()	// Skips the header
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func printWordFrequency(wordmap map[string]int) {
	var keys []string

	for key := range wordmap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordmap[keys[i]] < wordmap[keys[j]]
	})

	for _, key := range keys {
		fmt.Println("Count:", wordmap[key], "\t", key)
	}
}

func printCharFrequency(runemap map[rune]int) {
	var keys []rune

	for key := range runemap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return runemap[keys[i]] < runemap[keys[j]]
	})

	for _, key := range keys {
		fmt.Printf("Count: %d \t '%c'\n", runemap[key], key)
	}
}


func main() {
	filename := getFilePath()
	records, err := readCSV(filename)
	if(err != nil) {
		log.Fatal("Cannot open and read CSV file")
	}
	
	// Data Analysis
	wordCount := 0
	var tweetSlice []Tweet

	for _, eachrecord := range records {
		currTweet := ParseRecord(eachrecord)
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

	fmt.Println("Word count:", dataAnalysis.word_count)
	fmt.Println("Vocabulary size:", dataAnalysis.vocabulary_size)
	fmt.Printf("\nWord frequency in ascending order:\n")
	printWordFrequency(dataAnalysis.word_frequency)
	fmt.Printf("\nCharacter frequency in ascending order:\n")
	printCharFrequency(dataAnalysis.character_frequency)
	fmt.Printf("\nTop 20 Most Frequent Words:\n")
	printWordFrequency(GetMostFrequentWords(dataAnalysis.word_frequency, 20))
	fmt.Printf("\nStop Words:\n")
	printWordFrequency(dataAnalysis.stop_word_count)


	RenderWordCloud(GetMostFrequentWords(dataAnalysis.word_frequency, 20), "word-cloud.html")
	RenderTweetFrequency(dataAnalysis.tweet_frequency, "tweet-frequency-chart.html")
	RenderSymbolPieChart(GetSymbols(dataAnalysis.character_frequency), "symbol-frequency-chart.html")
}
