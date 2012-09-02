package goquery

import (
	"exp/html"
	"testing"
)

func TestEach(t *testing.T) {
	var cnt int

	sel := Doc().Root.Find(".hero-unit .row-fluid").Each(func(i int, n *Selection) {
		cnt++
		t.Logf("At index %v, node %v", i, n.Nodes[0].Data)
	}).Find("a")

	if cnt != 4 {
		t.Errorf("Expected Each() to call function 4 times, got %v times.", cnt)
	}
	if len(sel.Nodes) != 6 {
		t.Errorf("Expected 6 matching nodes, found %v.", len(sel.Nodes))
	}
}

func TestEachEmptySelection(t *testing.T) {
	var cnt int

	sel := Doc().Root.Find("zzzz")
	sel.Each(func(i int, n *Selection) {
		cnt++
	})
	if cnt > 0 {
		t.Error("Expected Each() to not be called on empty Selection.")
	}
	sel2 := sel.Find("div")
	if sel2.Nodes != nil {
		t.Error("Expected Find() on empty Selection to return an empty Selection.")
	}
}

func TestMap(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content")
	vals := sel.Map(func(i int, s *Selection) string {
		n := s.Get(0)
		if n.Type == html.ElementNode {
			return n.Data
		}
		return ""
	})
	for _, v := range vals {
		if v != "div" {
			t.Error("Expected Map array result to be all 'div's.")
		}
	}
	if len(vals) != 3 {
		t.Errorf("Expected Map array result to have a length of 3, found %v.", len(vals))
	}
}
