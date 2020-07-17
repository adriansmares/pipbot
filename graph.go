package main

import (
	"github.com/google/go-github/v32/github"
)

type node struct {
	neighbors map[int]*node
	issue     *github.Issue
}

type graph struct {
	nodes map[int]*node
}

func createGraph(issues []*github.Issue) (*graph, error) {
	g := &graph{
		nodes: make(map[int]*node),
	}
	for _, issue := range issues {
		g.nodes[*issue.Number] = &node{
			neighbors: make(map[int]*node),
			issue:     issue,
		}
	}
	addRelations(g)
	return g, nil
}
