package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Group struct {
	MemberCnt   int
	GroupYayCnt int
}

func NewGroup(yays string) *Group {
	memberyays := strings.Split(yays, "\n")
	groupyays := map[string]int{}

	for _, member := range memberyays {
		yays := strings.Split(member, "")
		for _, yay := range yays {
			if _, ok := groupyays[yay]; !ok {
				groupyays[yay] = 1
				continue
			}

			groupyays[yay]++
		}
	}

	return &Group{
		MemberCnt:   len(memberyays),
		GroupYayCnt: len(groupyays),
	}
}

func ParseGroups(reader io.Reader) ([]*Group, error) {
	groups := []*Group{}
	curGroup := []string{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if scanner.Text() == "" {
			groups = append(groups, NewGroup(strings.Join(curGroup, "\n")))
			curGroup = []string{}
			continue
		}

		curGroup = append(curGroup, scanner.Text())
	}

	groups = append(groups, NewGroup(strings.Join(curGroup, "\n")))

	if scanner.Err() != nil {
		return groups, scanner.Err()
	}

	return groups, nil
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	yayCnt := 0
	groups, err := ParseGroups(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(groups))

	for _, g := range groups {
		yayCnt = yayCnt + g.GroupYayCnt
	}

	fmt.Println(yayCnt)
}
