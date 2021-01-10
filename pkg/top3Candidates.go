package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	//"sort"
	"strconv"
	"strings"
)

func Top3Candidates(path string) []int {
	voterIdMap := make(map[int]bool)
	voteSumMap := make(map[int]int)
	var candidateIdList []int
	var voteSumList []int

	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vote := strings.Split(scanner.Text(), " ")

		voterId, voterIdErr := strconv.Atoi(vote[0])
		candidateId, candidateIdErr := strconv.Atoi(vote[1])

		if voterIdErr != nil || candidateIdErr != nil {
			fmt.Println("VoterId or CandidateId is not a num")
		} else {
			if voterIdMap[voterId] {
				log.Fatal("Fraud")
			} else {
				voterIdMap[voterId] = true
			}
			candidateVoteSum, candidateExists := voteSumMap[candidateId]
			if candidateExists {
				voteSumMap[candidateId] = candidateVoteSum + 1
			} else {
				voteSumMap[candidateId] = 1
				candidateIdList = append(candidateIdList, candidateId)
			}
		}
	}
	for _, candidate := range candidateIdList {
		voteSumList = append(voteSumList, voteSumMap[candidate])
	}
	sort.Slice(voteSumList, func(i, j int) bool {
		swap := voteSumList[i] > voteSumList[j]
		if swap {
			tmp := candidateIdList[i]
			candidateIdList[i] = candidateIdList[j]
			candidateIdList[j] = tmp
		}
		return swap
	})

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return candidateIdList[0:3]
}
