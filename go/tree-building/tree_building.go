package tree

import (
	"errors"
	"sort"
)

//Record contains the ID and Parent for the tree
type Record struct {
	ID     int
	Parent int
}

//Node holds the tree
type Node struct {
	ID       int
	Children []*Node
}

//Build is the function for constructing the tree
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	var tree = map[int]*Node{}

	for id, r := range records {
		if id != r.ID || r.ID < r.Parent || r.ID > 0 && r.ID == r.Parent {
			return nil, errors.New("not in sequence or has invaid parent")
		}

		subtree := &Node{ID: r.ID}
		tree[r.ID] = subtree

		if r.ID > 0 {
			p := tree[r.Parent]
			p.Children = append(p.Children, tree[r.ID])
		}

	}
	return tree[0], nil
}
