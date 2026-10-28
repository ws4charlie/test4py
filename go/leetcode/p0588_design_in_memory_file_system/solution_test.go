package p0588

import (
	"slices"
	"testing"
)

func TestFileSystem(t *testing.T) {
	fs := Constructor()

	if got := fs.Ls("/"); len(got) != 0 {
		t.Fatalf(`Ls("/") on empty fs = %v, want []`, got)
	}

	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	if got, want := fs.Ls("/"), []string{"a"}; !slices.Equal(got, want) {
		t.Errorf(`Ls("/") = %v, want %v`, got, want)
	}
	if got, want := fs.ReadContentFromFile("/a/b/c/d"), "hello"; got != want {
		t.Errorf("ReadContentFromFile = %q, want %q", got, want)
	}

	fs.AddContentToFile("/a/b/c/d", " world")
	if got, want := fs.ReadContentFromFile("/a/b/c/d"), "hello world"; got != want {
		t.Errorf("after append, ReadContentFromFile = %q, want %q", got, want)
	}

	fs.Mkdir("/a/b/c/bb")
	fs.AddContentToFile("/a/b/c/aa", "x")
	if got, want := fs.Ls("/a/b/c"), []string{"aa", "bb", "d"}; !slices.Equal(got, want) {
		t.Errorf(`Ls("/a/b/c") = %v, want %v (sorted)`, got, want)
	}
	if got, want := fs.Ls("/a/b/c/d"), []string{"d"}; !slices.Equal(got, want) {
		t.Errorf(`Ls on a file = %v, want %v`, got, want)
	}
}
