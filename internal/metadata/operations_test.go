package metadata

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Metadata{
		Entries: []Entry{
			{Name: "test1", Paths: []string{"a", "long", "path"}, Method: Copy},
			{Name: "test2", Paths: []string{"a", "really", "long", "path"}, Method: Symlink},
		},
	}

	want := Metadata{
		Entries: []Entry{
			{Name: "test1", Paths: []string{"a", "long", "path"}, Method: Copy},
			{Name: "test2", Paths: []string{"a", "really", "long", "path"}, Method: Symlink},
			{Name: "test3", Paths: []string{"a", "small", "path"}, Method: Copy},
		},
	}

	e := Entry{Name: "test3", Paths: []string{"a", "small", "path"}, Method: Copy}
	if err := got.add(e); err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	got := Metadata{
		Entries: []Entry{
			{Name: "test1", Paths: []string{"a", "long", "path"}, Method: Copy},
			{Name: "test2", Paths: []string{"a", "really", "long", "path"}, Method: Symlink},
		},
	}

	want := Metadata{
		Entries: []Entry{
			{Name: "test2", Paths: []string{"a", "really", "long", "path"}, Method: Symlink},
		},
	}

	if err := got.remove("test1"); err != nil {
		t.Fatalf("got %s, want nil", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
