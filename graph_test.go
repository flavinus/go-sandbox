package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinks(t *testing.T) {

	g := newGraph()

	g.addNode(1)
	g.addNode(2)
	g.addNode(3)
	g.addNode(4)

	g.addLink(1, 2)
	g.addLink(1, 4)
	g.addLink(2, 3)

	assert.True(t, g.hasLink(1, 2))
	assert.True(t, g.hasLink(2, 1))
	assert.False(t, g.hasLink(1, 3))
	assert.False(t, g.hasLink(3, 1))

	g.delLink(1, 2)
	assert.False(t, g.hasLink(1, 2))
	assert.False(t, g.hasLink(2, 1))
}

func TestShorterPath(t *testing.T) {

	g := newGraph()

	g.addNode(1)
	g.addNode(2)
	g.addNode(3)
	g.addNode(4)
	g.addNode(5)
	g.addNode(6)
	g.addNode(7)
	g.addNode(8)

	g.addLink(1, 2)
	g.addLink(1, 3)
	g.addLink(1, 4)
	g.addLink(2, 5)
	g.addLink(3, 4)
	g.addLink(3, 6)
	g.addLink(5, 6)
	g.addLink(6, 7)

	distances := g.getDistances(1)
	assert.Equal(t, 0, distances[1])
	assert.Equal(t, 1, distances[2])
	assert.Equal(t, 1, distances[3])
	assert.Equal(t, 1, distances[4])
	assert.Equal(t, 2, distances[5])
	assert.Equal(t, 2, distances[6])
	assert.Equal(t, 3, distances[7])

	assert.Equal(t, 3, len(g.getShorterPath(1, 5)))
	assert.Equal(t, 4, len(g.getShorterPath(1, 7)))
	assert.Equal(t, 0, len(g.getShorterPath(1, 8)))

}
