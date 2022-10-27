//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) WordCount 2022
//
// Author: Tyler(UnclassedPenguin)
//    URL: https://unclassed.ca
// GitHub: https://github.com/UnclassedPenguin/wordcount.git
// Description: This is a wordcount program inspired by the Tour Of Go on go.dev
//              It is complete overkill, but was just messing around and
//              practicing with the language. Fun regardless.
//
//-------------------------------------------------------------------------------
//-------------------------------------------------------------------------------

// WordCount.
// Usage:
// wordcount -f file  -w | -n [-r] [-l]
// or:
// go run main.go -f file -w | -n [-r] [-l]
// -f is the file you'd like to cound the words in.
// -w sorts by (w)ords alphabetically
// -n sorts by (n)umber of times a word is found
// -r is reversed order. So high to low or low to high.
// -l is lowercase. It converts all words to lowercase before counting

package main

import (
  "fmt"
  "sort"
  "strings"
  "io/ioutil"
  "regexp"
  "flag"
  "os"
)

// This function actually counts the words. Creates a map,
// checks if the word is there. If it isn't, it adds it.
// If it is, it increments the value it has been seen.
func WordCount(s string, lowercase bool) map[string]int {
  str := strings.Fields(s)

  words := make(map[string]int)

  for _, v := range str {
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

// This function uses regex to clear non alphabetic characters
func clearString(str string) string {
  var nonAlphaRegex = regexp.MustCompile(`[^a-zA-Z ]+`)
  return nonAlphaRegex.ReplaceAllString(str, "")
}

// This creates a slice and sorts it, and then prints out the slice
// in order, and uses that to get the value to print alongside it.
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

// This creates a slice and sorts by number of times
// the word has been seen.
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

  // Flags for command line options
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

  flag.Usage = func() {
    w := flag.CommandLine.Output()
    fmt.Fprintf(w, "Description of %s:\n\nThis is a way overcomplicated way to count the occurences of words in a text file.\nInspired by one of the excercises on the Tour of Go on go.dev\n\nUsage:\n\nwordcount -f file  -w | -n [-r] [-l]\n\n", os.Args[0])
    flag.PrintDefaults()
  }

  flag.Parse()

  var str []byte
  var err error

  if file != "" {
    // Read the file
    str, err = ioutil.ReadFile(file)
    if err != nil {
      fmt.Println("Error reading file:\n", err)
      os.Exit(1)
    }
  } else {
    fmt.Println("You didn't specify a file. Please use -f with a file to read from. Or -h for usage.")
    os.Exit(1)
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
