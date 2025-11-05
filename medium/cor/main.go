package main

import "log"

type SupportHandler interface {
	handleRequest(*SupportIssue)
}

type SupportIssue struct {
	title      string
	complexity int
}

func NewIssue(title string, complexity int) *SupportIssue {
	return &SupportIssue{title: title, complexity: complexity}
}

type AbstractSupportHandler struct {
	next      SupportHandler
	canHandle func(*SupportIssue) bool
	doHandle  func(*SupportIssue)
}

func (ash *AbstractSupportHandler) handleRequest(issue *SupportIssue) {
	if ash.canHandle(issue) {
		ash.doHandle(issue)
		return
	}

	if ash.next != nil {
		log.Printf("[W] can't handle (complexity = %d)", issue.complexity)
		ash.next.handleRequest(issue)
	} else {
		log.Printf("[E] Issue %q can not be handled by anyone in the chain", issue.title)
	}
}

// concreate handlers
type Tier1Support struct {
	AbstractSupportHandler
}

func CreateTier1Support(next SupportHandler) *Tier1Support {
	s := &Tier1Support{}
	s.next = next
	s.canHandle = func(issue *SupportIssue) bool {
		return issue.complexity <= 1
	}
	s.doHandle = func(issue *SupportIssue) {
		log.Printf("Tier1Support: Resolved issue %q", issue.title)
	}

	return s
}

type Tier2Support struct {
	AbstractSupportHandler
}

func CreateTier2Support(next SupportHandler) *Tier2Support {
	s := &Tier2Support{}
	s.next = next
	s.canHandle = func(issue *SupportIssue) bool {
		return issue.complexity <= 2
	}
	s.doHandle = func(issue *SupportIssue) {
		log.Printf("Tier2Support: Resolved issue %q", issue.title)
	}

	return s
}

type Tier3Support struct {
	AbstractSupportHandler
}

func CreateTier3Support(next SupportHandler) *Tier3Support {
	s := &Tier3Support{}
	s.next = next
	s.canHandle = func(issue *SupportIssue) bool {
		return issue.complexity <= 3
	}
	s.doHandle = func(issue *SupportIssue) {
		log.Printf("Tier3Support: Resolved issue %q", issue.title)
	}

	return s
}

func main() {
	log.Println("==================================================")
	log.Println("[+] Chain of Responcibility")
	t3 := CreateTier3Support(nil)
	t2 := CreateTier2Support(t3)
	t1 := CreateTier1Support(t2)
	supportChain := t1

	simpleIssue := NewIssue("Forgotten password", 1)
	intermediateIssue := NewIssue("System running slow", 2)
	hardIssue := NewIssue("Data loss in database", 3)
	impossibleIssue := NewIssue("Quantom server anomaly", 5)

	supportChain.handleRequest(simpleIssue)
	supportChain.handleRequest(intermediateIssue)
	supportChain.handleRequest(hardIssue)
	supportChain.handleRequest(impossibleIssue)
}
