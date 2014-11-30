package post

//TODO: Write some tests

import (
	"HTML/template"
	"fmt"
	"testing"
	"time"
)

func testNewManager(t *testing.T) {

	//	m := NewManager()

}

func TestSort(t *testing.T) {
	m := NewManager()
	p1 := Post{
		Content: template.HTML("Text."),
		Title:   "Post 1",
		Image:   false,
		Date:    time.Unix(0, 0),
	}
	fmt.Println(p1)

	if err := m.Add(p1, p1.Title); err != nil {
		t.Fatal(err)

	}
	p2 := Post{
		Content: template.HTML("Text."),
		Title:   "Post 2",
		Image:   false,
		Date:    time.Unix(1000, 0),
	}

	if err := m.Add(p2, p2.Title); err != nil {
		t.Fatal(err)
	}
	fmt.Println(*m)
	if false {
		t.Error("NOO")
	}
	fmt.Println(*m.SortedPosts[0])
	fmt.Println(p1)
	fmt.Println(*m.SortedPosts[0] == p1)
}
