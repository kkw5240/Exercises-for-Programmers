package main

import "fmt"

/**
 * 작성자 : 김규완
 * 실행방법 : java kkw5240.class
 * 주의사항 : 기본풀이까지만 했습니다.
 */
func main() {
	var madLib MadLib

	noun, err := promptNoun()
	if err != nil {
		return
	}
	madLib.SetNoun(noun)

	verb, err := promptVerb(err)
	if err != nil {
		return
	}
	madLib.SetVerb(verb)

	adjective, err := promptAdjective(err)
	if err != nil {
		return
	}
	madLib.SetAdjective(adjective)

	adverb, err := promptAdverb()
	if err != nil {
		return
	}
	madLib.SetAdverb(adverb)

	var message = madLib.buildMassage()

	printResultMessage(message)
}

func promptNoun() (string, error) {
	fmt.Print("Enter a noun: ")

	var noun string
	_, err := fmt.Scanln(&noun)
	return noun, err
}

func promptVerb(err error) (string, error) {
	fmt.Print("Enter a verb: ")

	var verb string
	_, err = fmt.Scanln(&verb)
	return verb, err
}

func promptAdjective(err error) (string, error) {
	fmt.Print("Enter an adjective: ")

	var adjective string
	_, err = fmt.Scanln(&adjective)
	return adjective, err
}

func promptAdverb() (string, error) {
	fmt.Print("Enter an adverb: ")

	var adverb string
	_, err := fmt.Scanln(&adverb)
	return adverb, err
}

func printResultMessage(message string) {
	fmt.Println(message)
}

type MadLib struct {
	Verb      string
	Adjective string
	Noun      string
	Adverb    string
}

func (madLib *MadLib) GetVerb() string {
	return madLib.Verb
}
func (madLib *MadLib) SetVerb(verb string) {
	madLib.Verb = verb
}

func (madLib *MadLib) GetAdjective() string {
	return madLib.Adjective
}
func (madLib *MadLib) SetAdjective(adjective string) {
	madLib.Adjective = adjective
}

func (madLib *MadLib) GetNoun() string {
	return madLib.Noun
}
func (madLib *MadLib) SetNoun(noun string) {
	madLib.Noun = noun
}

func (madLib *MadLib) GetAdverb() string {
	return madLib.Adverb
}
func (madLib *MadLib) SetAdverb(adverb string) {
	madLib.Adverb = adverb
}

func (madLib *MadLib) buildMassage() string {
	return fmt.Sprintf(
		"Do you %s your %s %s %s? That's hilarious!",
		madLib.GetVerb(),
		madLib.GetAdjective(),
		madLib.GetNoun(),
		madLib.GetAdverb(),
	)
}
