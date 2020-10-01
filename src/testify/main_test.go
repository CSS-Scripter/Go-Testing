package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalculateTestify(t *testing.T) {
    assert.Equal(t, Calculate(2), 5, "2 + 2 should equal to 4")
}

func TestTableCalculateTestify(t *testing.T) {
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
        assert.Equal(t, Calculate(test.input), test.expected, "Values don't align")
    }
}