package post

import (
	"HTML/template"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"time"

	"github.com/russross/blackfriday"
)

//A Post represents a blog post, and contains the fields of the post.
type Post struct {
	Title   string
	Image   bool
	Content template.HTML
	Date    time.Time
}

//A Manager contains a map of all Posts loaded into it, against their names, and an array of pointers to Posts which sorts to give the chronological ordering of posts.
type Manager struct {
	postMap     map[string]Post
	SortedPosts []*Post
}

//NewManager creates a new Manager.
func NewManager() *Manager {
	return &Manager{}
}

//Len returns the number of Posts in the Manager
func (pm *Manager) Len() int {
	return len(pm.SortedPosts)
}

//Swaps the elements i and j in the SortedPost slice
func (pm *Manager) Swap(i, j int) {
	pm.SortedPosts[i], pm.SortedPosts[j] = pm.SortedPosts[j], pm.SortedPosts[i]
}

//Less returns the earlier of two indices
func (pm *Manager) Less(i, j int) bool {
	return pm.SortedPosts[i].Date.Before(pm.SortedPosts[j].Date)
}

//ImportFolder takes a folder path and loads all present Post files in that folder.
//A Post file is a markdown file, with its filename providing the post name.
func (pm *Manager) ImportFolder(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, f := range files {
		name := f.Name()
		if name[len(name)-3:] == ".md" {
			p, err := NewPostFromFile(path + f.Name())
			if err == nil {
				pm.Add(*p, name[:len(name)-3])

			}
		}
	}

	return nil
}

//Add takes a post and adds it into the Manager map, if it doens't exist
//Needs to sort the array
func (pm *Manager) Add(p Post, name string) error {
	if _, exists := pm.postMap[name]; !exists {
		return fmt.Errorf("ERROR")
	}
	pm.postMap[name] = p
	pm.SortedPosts = append(pm.SortedPosts, &p)
	sort.Sort(pm)
	return nil
}

//GetPostFromName returns the post of the name,  name
func (pm *Manager) GetPostFromName(name string) (*Post, error) {
	p, ok := pm.postMap[name]
	if !ok {
		return nil, fmt.Errorf("no post")
	}
	return &p, nil
}

//NewPostFromFile takes a filename and returns and *Post and an error
func NewPostFromFile(filePath string) (*Post, error) {
	p := Post{}
	_, name := path.Split(filePath)
	f, err := ioutil.ReadFile(filePath + ".md")
	if err != nil {
		log.Println("file not found")

		return nil, err
	}
	stat, err := os.Stat(filePath + ".md")
	if err != nil {
		log.Println(err)
	}
	p.Date = stat.ModTime()
	p.Image = true
	_, err = os.Open(filePath + ".md")
	if err != nil {
		p.Image = false
	}
	p.Title = name[1:]
	p.Content = template.HTML(blackfriday.MarkdownCommon(f))
	return &p, nil
}
