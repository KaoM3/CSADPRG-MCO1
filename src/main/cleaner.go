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
)

// Helper function to clean and split the text into words
func tokenize(text string) []string {
	// Convert text to lowercase and remove non-alphabetical characters
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
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
	stop_words := `\b(i|me|my|myself|we|our|ours|ourselves|you|your|yours|yourself|yourselves|he|him|his|himself|she|her|hers|herself|it|its|itself|they|them|their|theirs|themselves|what|which|who|whom|this|that|these|those|am|is|are|was|were|be|been|being|have|has|had|having|do|does|did|doing|a|an|the|and|but|if|or|because|as|until|while|of|at|by|for|with|about|against|between|into|through|during|before|after|above|below|to|from|up|down|in|out|on|off|over|under|again|further|then|once|here|there|when|where|why|how|all|any|both|each|few|more|most|other|some|such|no|nor|not|only|own|same|so|than|too|very|s|t|can|will|just|don|should|now)\b`
	return stop_words
}

func OrderStringsByValue(mapping map[string]int) {
	var keys []string

	for key := range mapping {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return mapping[keys[i]] < mapping[keys[j]]
	})

	fmt.Println("Words sorted by frequency:")
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, mapping[key])
	}
}

func OrderRunesByValue(mapping map[rune]int) {
	var keys []rune

	for key := range mapping {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return mapping[keys[i]] < mapping[keys[j]]
	})

	fmt.Println("Words sorted by frequency:")
	for _, key := range keys {
		fmt.Printf("%c: %d\n", key, mapping[key])
	}
}

func GetVocabularySize(wordmap map[string]int) int {
	return len(wordmap)
}

func GetCountStopWords(tweets []Tweet) map[string]int {
	stopwordmap := make(map[string]int)
	re := regexp.MustCompile(stopWords())

	for _, currTweet := range tweets {
		for _, currToken := range currTweet.Tokens {
			if(re.MatchString(currToken)) {
				stopwordmap[currToken]++
			}
		}
	}
	OrderStringsByValue(stopwordmap)
	return stopwordmap
}

func GetWordFrequency(tweets []Tweet) map[string]int {
	wordmap := make(map[string]int)

	for _, currTweet := range tweets {
		for _, currToken := range currTweet.Tokens {
			wordmap[currToken]++
		}
	}
	OrderStringsByValue(wordmap)
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
	OrderRunesByValue(runemap)
	return runemap
}

func GetTweetFrequency(tweets []Tweet) map[int]map[string]int {
	monthnames := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	tweetfrequencymap := make(map[int]map[string]int)

	for _, currTweet := range tweets {
		year := currTweet.Date_created.Year
		month := monthnames[currTweet.Date_created.Month - 1]
		if tweetfrequencymap[year] == nil {
			tweetfrequencymap[year] = make(map[string]int)
		}
		tweetfrequencymap[year][month]++
	}

	return tweetfrequencymap
}

func ParseRecord(record []string) Tweet {
	tokens := tokenize(record[3])

	newTweet := Tweet{
		Date_created: parseDate(record[2]),
		Text:         record[3],
		Tokens:       tokens,
		Word_count:   len(tokens),
	}

	return newTweet
}
