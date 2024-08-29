package blogrender_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/approvals/go-approval-tests"
	"github.com/exasperlnc/blogrender"
)

var ( 
		aPost = blogrender.Post{
			Title: "hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)
func TestRender(t *testing.T) {	

	postRenderer, err := blogrender.NewPostRender()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	postRenderer, err := blogrender.NewPostRender()

	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i:= 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}