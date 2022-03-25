package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Function Complexity: Medium
// Time: O(n)
func TestValidity(str string) bool {
	words := strings.Split(str, "-")
	num_total := 0
	word_total := 0
	for _, word := range words {
		if word == "" {
			return false
		}
		if _, err := strconv.Atoi(word); err == nil {
			num_total++
		} else {
			word_total++
		}
		if num_total-word_total == 1 || num_total == word_total {
			continue
		} else {
			return false
		}
	}
	return true
}

// Function Complexity: Easy
// Time: O(n)
func AverageNumber(str string) float32 {
	words := strings.Split(str, "-")
	total := 0
	num_total := 0
	for _, word := range words {
		if val, err := strconv.Atoi(word); err == nil {
			num_total++
			total += val
		}
	}
	return float32(total) / float32(num_total)
}

// Function Complexity: Medium
// Time: O(n)
func WholeStory(str string) string {
	words := strings.Split(str, "-")
	text := ""
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			text += word + " "
		}
	}
	return strings.TrimSpace(text)
}

// Function Complexity: Medium
// Time: O(n)
func StoryStats(str string) (string, string, float32, []string) {
	words := strings.Split(str, "-")
	var shortest string
	var longest string
	total_words := 0
	total_len := 0
	var listWord []string
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			if shortest == "" || len(word) < len(shortest) {
				shortest = word
			}
			if longest == "" || len(word) > len(longest) {
				longest = word
			}
			total_words++
			total_len += len(word)
		}
	}
	avg_word := float32(total_len) / float32(total_words)
	avg_word_ciel := int(avg_word) + 1
	avg_word_float := int(avg_word)
	for _, word := range words {
		if _, err := strconv.Atoi(word); err == nil {
			continue
		} else {
			if len(word) == avg_word_float || len(word) == avg_word_ciel {

				listWord = append(listWord, word)
			}
		}
	}
	return shortest, longest, avg_word, listWord
}

func random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabets[rand.Intn(len(alphabets))]
	}
	return string(b)
}

// Function Complexity: Medium
// Time: O(n)
func generate(isCorrect bool) string {
	seed1 := rand.NewSource(time.Now().UnixNano())
	temp := rand.New(seed1)
	temp2 := temp.Intn(10) + 2
	result := ""
	for i := 0; i < temp2; i++ {
		result += fmt.Sprintf("%d-%s-", temp.Intn(1000), random(1+temp.Intn(10)))
	}
	if isCorrect {
		return strings.Trim(result, "-")
	} else {
		return result
	}
}

func main() {
	fmt.Println(generate(false))
	fmt.Println(generate(true))
}
