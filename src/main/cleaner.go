package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Helper function to clean and split the text into words
func tokenize(text string) []string {
	// Convert text to lowercase and remove non-alphabetical characters
	re := regexp.MustCompile(`[^a-zA-Z\s]`)
	text = re.ReplaceAllString(text, "")
	words := strings.Fields(strings.ToLower(text))
	return words
}

func parseDate(dateString string) Date {
	parsedTime, err := time.Parse("2006-01-02", dateString[:10])

	date := Date{
		Year:  parsedTime.Year(),
		Month: int(parsedTime.Month()),
		Day:   parsedTime.Day(),
	}

	if err != nil {
		fmt.Println(err)
	}
	return date
}

func stopWords() string {
	stop_words := "\b(i|me|my|myself|we|our|ours|ourselves|you|your|yours|yourself|yourselves|he|him|his|himself|she|her|hers|herself|it|its|itself|they|them|their|theirs|themselves|what|which|who|whom|this|that|these|those|am|is|are|was|were|be|been|being|have|has|had|having|do|does|did|doing|a|an|the|and|but|if|or|because|as|until|while|of|at|by|for|with|about|against|between|into|through|during|before|after|above|below|to|from|up|down|in|out|on|off|over|under|again|further|then|once|here|there|when|where|why|how|all|any|both|each|few|more|most|other|some|such|no|nor|not|only|own|same|so|than|too|very|s|t|can|will|just|don|should|now)\b"
	return stop_words
}

func GetWordFrequency(tweets []Tweet) map[string]int {
	wordmap := make(map[string]int)
	
	for _, currTweet := range tweets {
		for _, currToken := range currTweet.Tokens {
			if _, found := wordmap[currToken]; found {
				wordmap[currToken] += 1
			} else {
				wordmap[currToken] = 1
			}
		}
	}

	return wordmap
}

func ParseRecord(record []string) Tweet {
	tokens := tokenize(record[3])

	newTweet := Tweet{
		Date_created: 	parseDate(record[2]),
		Text:         	record[3],
		Tokens:			tokens,
		Word_count:		len(tokens),
	}

	return newTweet
}
