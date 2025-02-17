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
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
)

// Helper function to clean and split the text into words
func tokenize(text string) []string {
	// Convert text to lowercase and remove non-alphabetical characters
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	text = re.ReplaceAllString(text, "")
	words := strings.Fields(strings.ToLower(text))
	return words
}

func parseDate(dateString string, format string) (Date, error) {
	parsedTime, err := time.Parse(format, dateString[:10])
	
	if err != nil {
		fmt.Println(err)
	}

	date := Date{
		Year:  parsedTime.Year(),
		Month: int(parsedTime.Month()),
		Day:   parsedTime.Day(),
	}

	return date, err
}

func stopWords() string {
	stop_words := `\b(i|me|my|myself|we|our|ours|ourselves|you|your|yours|yourself|yourselves|he|him|his|himself|she|her|hers|herself|it|its|itself|they|them|their|theirs|themselves|what|which|who|whom|this|that|these|those|am|is|are|was|were|be|been|being|have|has|had|having|do|does|did|doing|a|an|the|and|but|if|or|because|as|until|while|of|at|by|for|with|about|against|between|into|through|during|before|after|above|below|to|from|up|down|in|out|on|off|over|under|again|further|then|once|here|there|when|where|why|how|all|any|both|each|few|more|most|other|some|such|no|nor|not|only|own|same|so|than|too|very|s|t|can|will|just|don|should|now)\b`
	return stop_words
}

func GetMostFrequentWords(mapping map[string]int, size int) map[string]int {
	var keys []string
	wordmap := make(map[string]int)

	for key := range mapping {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return mapping[keys[i]] > mapping[keys[j]]
	})

	for i, key := range keys {
		wordmap[key] = mapping[key]
		if i >= size-1 {
			break
		}
	}
	return wordmap
}

func GetSymbols(mapping map[rune]int) map[rune]int {
	var keys []rune
	symbolmap := make(map[rune]int)

	for key := range mapping {
		if !unicode.IsLetter(key) && !unicode.IsDigit(key) && !unicode.IsSpace(key) {
			keys = append(keys, key)
		}
	}

	for _, key := range keys {
		symbolmap[key] = mapping[key]
	}

	return symbolmap
}

func GetVocabularySize(wordmap map[string]int) int {
	return len(wordmap)
}

func GetCountStopWords(tweets []Tweet) map[string]int {
	stopwordmap := make(map[string]int)
	re := regexp.MustCompile(stopWords())

	for _, currTweet := range tweets {
		for _, currToken := range currTweet.Tokens {
			if re.MatchString(currToken) {
				stopwordmap[currToken]++
			}
		}
	}

	return stopwordmap
}

func GetWordFrequency(tweets []Tweet) map[string]int {
	wordmap := make(map[string]int)

	for _, currTweet := range tweets {
		for _, currToken := range currTweet.Tokens {
			wordmap[currToken]++
		}
	}

	return wordmap
}

func GetCharacterFrequency(tweets []Tweet) map[rune]int {
	runemap := make(map[rune]int)

	for _, currTweet := range tweets {
		for _, char := range currTweet.Text {
			// Increment the count of the character
			runemap[char]++
		}
	}

	return runemap
}

func GetTweetFrequency(tweets []Tweet) map[int]map[string]int {
	monthnames := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	tweetfrequencymap := make(map[int]map[string]int)

	for _, currTweet := range tweets {
		year := currTweet.Date_created.Year
		month := monthnames[currTweet.Date_created.Month-1]
		if tweetfrequencymap[year] == nil {
			tweetfrequencymap[year] = make(map[string]int)
		}
		tweetfrequencymap[year][month]++
	}

	return tweetfrequencymap
}

func ParseRecord(record []string) Tweet {
	tokens := tokenize(record[3])
	var parsedDate Date
	var err error

	if parsedDate, err = parseDate(record[2], "2006-01-02"); err != nil {
		parsedDate, err = parseDate(record[2], "02/01/2006")
		if err != nil {
			// Default if all parse failed
			parsedDate = Date{
				Year: 1999,
				Month: 1,
				Day: 1,
			}
		}
	}

	newTweet := Tweet{
		Date_created: parsedDate,
		Text:         record[3],
		Tokens:       tokens,
		Word_count:   len(tokens),
	}

	return newTweet
}
