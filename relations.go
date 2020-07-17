package main

import "regexp"

var (
	blockedOnRegex = regexp.MustCompile("Blocked on #([0-9]+)")
	blockedByRegex = regexp.MustCompile("Blocked by #([0-9]+)")
)

const (
	blockedIssueLabel  = "status/blocked/issue"
	blockingIssueLabel = "status/blocking/issue"
)

func addRelations(g *graph) error {
	return nil
}
