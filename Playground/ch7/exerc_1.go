package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	ranker := r.Ranking()
	for _, v := range ranker {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func (l *League) MatchResult(t1 string, s1 int, t2 string, s2 int) {

	if _, ok := l.Teams[t1]; !ok {
		return
	}
	if _, ok := l.Teams[t2]; !ok {
		return
	}
	if s1 == s2 {
		return
	}

	if s1 > s2 {
		l.Wins[t1]++
	} else {
		l.Wins[t2]++
	}
	fmt.Println(l.Wins)
}

func (l League) Ranking() []string {
	ranking := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		ranking = append(ranking, k)
	}
	sort.Slice(ranking, func(i, j int) bool {
		return l.Wins[ranking[i]] > l.Wins[ranking[j]]
	})
	return ranking
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	RankPrinter(l, os.Stdout)
}
