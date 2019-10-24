package analyzer

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/eadium/contents-analyzer/brackets"
)

type Ingredient struct {
	Name        string        `json: "name"`
	Ingredients *[]Ingredient `json: "ingredients"`
}

type argError struct {
	arg  string
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}

// Analyze analyses product ingredients
func Analyze(s string) (*[]Ingredient, error) {
	bracketsBalanced, _ := brackets.Bracket(s)
	if !bracketsBalanced {
		log.Println("ALERT! Brackets are not balanced!")
		return nil, &argError{"SYNTAX_ERROR", "Brackets are not balanced"}
	}
	letters := strings.Split(s, "")
	ings := make([]Ingredient, 0)
	err := parse(letters, &ings)
	if err != nil {
		return nil, err
	}
	// json, _ := json.Marshal(ings)
	// ioutil.WriteFile("./prod_contents.json", json, 0644)
	return &ings, nil
}

// Parse parses
func parse(letters []string, ings *[]Ingredient) error {
	reg, err := regexp.Compile("[^a-zA-Z0-9А-Яа-я[:space:]]+")
	bracketsReg, err1 := regexp.Compile("[)]+")
	separatorReg, err2 := regexp.Compile("[,.:]+")
	recursiveReg, err3 := regexp.Compile("[(]+")
	if err != nil || err1 != nil || err2 != nil || err3 != nil {
		return err
	}
	curWord := ""
	for i := 0; i < len(letters); i++ {
		if separatorReg.MatchString(letters[i]) || i == len(letters)-1 {
			if i == len(letters)-1 || len(curWord) <= 2 {
				curWord += letters[i]
				if reg.MatchString(curWord) == true {
					curWord = ""
					continue
				}
			}
			*ings = append(*ings, Ingredient{strings.TrimSpace(bracketsReg.ReplaceAllString(curWord, "")), nil})
			curWord = ""
			continue
		}

		if recursiveReg.MatchString(letters[i]) {
			closePos := findClosingParen(letters, i+1)
			substring := letters[i+1 : closePos-1]
			subIngs := make([]Ingredient, 0)
			err := parse(substring, &subIngs)
			if err != nil {
				return err
			}
			*ings = append(*ings, Ingredient{strings.TrimSpace(bracketsReg.ReplaceAllString(curWord, "")), &subIngs})
			if i < len(letters)-1 {
				i = closePos - 1
			} else {
				i = closePos + 1
			}
			curWord = ""
		}
		curWord += letters[i]
	}
	return nil
}

func findClosingParen(text []string, openPos int) int {
	closePos := openPos
	counter := 1
	for counter > 0 {
		c := text[closePos]
		closePos++
		if c == "(" {
			counter++
		} else if c == ")" {
			counter--
		}
	}
	return closePos
}
