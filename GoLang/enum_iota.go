package main
import "fmt"

// iota for Enumeration

type BlogCategory int

const (
  Unlisted BlogCategory = iota
  Education
  Entertainment
  Sports
)

// Blog Struct
type Blog struct {
  Title string
  Content string
  Category BlogCategory
}

// Method to print Blog struct properly
func (b *Blog) Format() string {
  return fmt.Sprintf("{title: %s, content: %s, category: %v}", b.Title, b.Content, b.Category)
}


/**
 * Implement a blog struct with enum category
 */

func main() {
  var sportBlog = Blog{"Massive win", "Massive win for the home team in the football match", Sports}

  unlistedBlog := Blog{
  Title: "Unlisted",
  Content: "This blog is unlisted",
  }

  fmt.Println("SportBlog:", sportBlog.Format())
  fmt.Println("Unlisted:", unlistedBlog.Format())
}
