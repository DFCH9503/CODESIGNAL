package main

import(
	"fmt"
	"slices"
)

func electionWinners(votes []int, k int) int{
	maxVotesBefore := slices.Max(votes)
	possibleWinners := 0

	if k == 0{
		for _, vote := range votes{
			if vote == maxVotesBefore{
				possibleWinners++
			}
		}
		if possibleWinners == 1{
			return 1
		}else{
			return 0
		}
	}

	for _, vote := range votes{
		if maxVotesBefore < vote + k{
			possibleWinners++
		}
	}
	return possibleWinners
}

func main(){
	votes := []int{2, 3, 5, 2}
	k := 3

	fmt.Printf("With the votes of each candidate represented by %v and a remaining of %d vote/s, there could be %d winner/s", votes, k, electionWinners(votes, k) )
}