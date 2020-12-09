package main

import (
	"strings"
	"testing"
)

func TestNewGroup(t *testing.T) {
	cases := []struct {
		Desc, Input            string
		MemberCnt, GroupYayCnt int
	}{
		{
			Desc:        "1 member, 1 yay",
			Input:       "a",
			MemberCnt:   1,
			GroupYayCnt: 1,
		},
		{
			Desc:        "4 members, 1 yay",
			Input:       "a\na\na\na",
			MemberCnt:   4,
			GroupYayCnt: 1,
		},
		{
			Desc:        "1 member, 4 yays",
			Input:       "abcd",
			MemberCnt:   1,
			GroupYayCnt: 4,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual := NewGroup(tc.Input)

			if tc.MemberCnt != actual.MemberCnt {
				t.Errorf("Expected %v members, Got %v", tc.MemberCnt, actual.MemberCnt)
			}

			if tc.GroupYayCnt != actual.GroupYayCnt {
				t.Errorf("Expected %v group yays, Got %v", tc.GroupYayCnt, actual.GroupYayCnt)
			}
		})
	}
}

func TestParseGroups(t *testing.T) {
	cases := []struct {
		Desc, Input string
		GroupCnt    int
		GroupYayCnt []int
	}{
		{
			Desc:     "5 Groups",
			GroupCnt: 5,
			Input: `abc

a
b
c

ab
ac

a
a
a
a

b
`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			actual, err := ParseGroups(strings.NewReader(tc.Input))
			if err != nil {
				t.Error(err)
			}

			if tc.GroupCnt != len(actual) {
				t.Errorf("Expected %v groups, Got %v", tc.GroupCnt, len(actual))
			}
		})
	}
}
