package main

import (
	"fmt"
	"github.com/truongnhukhang/hellogo/disjointSet"
)

func disjointSetTest() {
	b := "b"
	c := "c"
	d := "d"
	e := "e"
	f := "f"
	g := "g"
	h := "h"
	disjoint := disjointSet.New()
	disjoint.MakeSet(b)
	disjoint.MakeSet(c)
	disjoint.MakeSet(d)
	disjoint.MakeSet(e)
	disjoint.MakeSet(f)
	disjoint.MakeSet(g)
	disjoint.MakeSet(h)

	disjoint.Union(h, b)
	disjoint.Union(c, e)
	set := disjoint.FindSet(b)
	fmt.Println(set.Key)
	disjoint.Union(e, b)
	fmt.Println(disjoint.FindSet(b).Key)
	fmt.Println(disjoint.FindSet(h).Key)
	fmt.Println(disjoint.FindSet(e).Key)
}
