package dialect

import (
  "github.com/jdkato/prose/tokenize"
  "github.com/c9s/inflect"
  "regexp"
  "strings"
  "fmt"
)

type Detector interface {
  Corpus () (Corpus)
  Flag(p Product) (string)
}

type NaiveDetector struct {
  corpus Corpus
}

func (d NaiveDetector) Corpus () (Corpus) {
   return d.corpus
}

func (d NaiveDetector) Flag(p Product) (string) {
  hasAmericanWord, hasBritishWord := false, false
  americanWord, britishWord := "", ""

  for w, _ := range *words(p) {
    if !hasAmericanWord {
      if _, ok := d.Corpus().AmericanWords[w]; ok {
        hasAmericanWord = true
        americanWord += w
        continue
      }
    }
    if !hasBritishWord {
      if _, ok := d.Corpus().BritishWords[w]; ok {
        hasBritishWord = true
        britishWord += w
        continue
      }
    }
    if hasAmericanWord && hasBritishWord {
      break
    }
  }
  var output string = "Unknown"
  if hasAmericanWord && hasBritishWord {
    output = "Mixed British and American English"
    americanWord = americanWord + " "
  } else if hasAmericanWord {
    output = "American English"
  } else if hasBritishWord {
    output = "British English"
  }
  return fmt.Sprintf("%-35s %s%s", output, americanWord, britishWord)
}


func words(p Product) *map[string]bool {
  words := make(map[string]bool)
  santize(p.Name, &words)
  santize(p.Description, &words)

  return &words;
}

func santize(s string, words *map[string]bool) {
  var word = regexp.MustCompile(`^[A-Za-z]+$`)

  for _, w := range tokenize.TextToWords(s) {
    if word.MatchString(w) {
      (*words)[inflect.Singularize(strings.ToLower(w))] = true
    }
  }
}