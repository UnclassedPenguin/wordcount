// Wordcount. Usage:
// wordcount -f file -r -l
// or:
// go run main.go -f file -r -l
// -f is the file you'd like to cound the words in.
// -r is reversed order. So high to low or low to high.
// -l is lowercase. It converts all words to lowercase before counting
// -w sorts by words alphabetically
// -n sorts by number of times a word is found

package main

import (
  "fmt"
  "sort"
  "strings"
  "io/ioutil"
  "regexp"
  "flag"
)

func WordCount(s string, lowercase bool) map[string]int {
  str := strings.Fields(s)
  //fmt.Println(str)
  words := make(map[string]int)

  for _, v := range str {
    //fmt.Println("i = ", i)
    //fmt.Println("v = ", v)
    if lowercase {
      v = strings.ToLower(v)
    }

    if strings.Contains(v, "—") {
      splitStrings := strings.Split(v, "—")
      for _, v := range splitStrings {
        v = clearString(v)

        _, ok := words[v]

        if !ok {
          words[v] = 1
        } else {
          words[v] += 1
        }
      }
    } else {
      v = clearString(v)

      _, ok := words[v]

      if !ok {
        words[v] = 1
      } else {
        words[v] += 1
      }
    }
  }

  return words
}

func clearString(str string) string {
  var nonAlphaRegex = regexp.MustCompile(`[^a-zA-Z ]+`)
  return nonAlphaRegex.ReplaceAllString(str, "")
}

func sortWords(m map[string]int, reversed bool) {
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }
  sort.Strings(keys)

  if reversed {
    for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
      keys[i], keys[j] = keys[j], keys[i]
    }
  }

  for _, k := range keys {
    fmt.Println(k, m[k])
  }
}

func sortNumber (m map[string]int, reversed bool) {
  keys := make([]string, 0, len(m))
  for k := range m {
    keys = append(keys, k)
  }

  sort.SliceStable(keys, func(i, j int) bool {
    return m[keys[i]] < m[keys[j]]
  })

  if reversed {
    for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
      keys[i], keys[j] = keys[j], keys[i]
    }
  }

  for _, k := range keys {
    fmt.Println(k, m[k])
  }
}


func main() {
  var file string
  var reversed bool
  var lowercase bool
  var sortByNumber bool
  var sortByWord bool


  flag.StringVar(&file,       "f",    "", "The file to count words from.")
  flag.BoolVar(&reversed,     "r", false, "Wether or not to reverse the order.")
  flag.BoolVar(&lowercase,    "l", false, "Wether or not to lowercase all words.")
  flag.BoolVar(&sortByNumber, "n", false, "Sort by number.")
  flag.BoolVar(&sortByWord,   "w", false, "Sort by words.")

  flag.Parse()

  // Read the file
  str, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Println("Error reading file:\n", err)
  }

  wordList := WordCount(string(str), lowercase)

  if sortByWord && sortByNumber {
    fmt.Println("You can only use one of -n or -w.")
  } else if sortByWord {
    sortWords(wordList, reversed)
  } else if sortByNumber {
    sortNumber(wordList, reversed)
  } else {
    fmt.Println("Please use -n (number) or -w (word) to say wether you want to sort by number or word")
  }
}
