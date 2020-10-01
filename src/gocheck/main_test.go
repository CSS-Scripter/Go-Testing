package main

import (
    "testing"
    . "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCalculateGocheck(c *C) {
    c.Assert(Calculate(2), Equals, 5)
}

func (s *MySuite) TestTableCalculateGocheck(c *C) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        c.Assert(Calculate(test.input), Equals, test.expected)
    }
}