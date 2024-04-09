package parser

import (
	"errors"
	"strings"
	"testing"
)

func TestParseSimple1(t *testing.T) {
	test := `<html>
<body>
	<a href="https://golang.org">Golang</a>
</body>
</html>`

	got, err := Parse(strings.NewReader(test))
	if err != nil {
		t.Fatal(err)
	}
	expected := []Link{
		{"https://golang.org"},
	}

	if len(got) != len(expected) {
		t.Fatalf("got %d links, expected %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("got %v, expected %v", got[i], expected[i])
		}
	}
}

func TestParseSimple2(t *testing.T) {
	test := `<html>

<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>

<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>

</html>`

	got, err := Parse(strings.NewReader(test))
	if err != nil {
		t.Fatal(err)
	}
	expected := []Link{
		{"https://www.twitter.com/joncalhoun"},
		{"https://github.com/gophercises"},
	}

	if len(got) != len(expected) {
		t.Fatalf("got %d links, expected %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("got %v, expected %v", got[i], expected[i])
		}
	}
}

func TestParseNotSimple(t *testing.T) {
	test := `<html>
<body>
	<a href="https://golang.org">Golang</a>
	<a href="http://foo.com">foo</a>
	<a href="mailto:foo@bar.com">foo@bar.com</a>
	<a href="#header">header</a>
	<a href="javascript:alert('hello')">alert</a>
	<a href="/path/to/file/with spaces.html">file with spaces</a>
	<a href="">empty href</a>
	<a name="named-anchor">
		<a href="#named-anchor">named anchor</a>
	</a>
	<a href="http://example.com/<b>bold</b>">HTML</a>
	<a href="https://example.com/<script>alert(1)</script>">HTML with script</a>
	<a href="https://example.com/<foo>">HTML with non-standard tag</a>
	<a href="https://example.com/?q=<&>" class="quoted-chars">quoted chars</a>
</body>
</html>`

	got, err := Parse(strings.NewReader(test))
	if err != nil {
		t.Fatal(err)
	}
	expected := []Link{
		{"https://golang.org"},
		{"http://foo.com"},
		{"mailto:foo@bar.com"},
		{"#header"},
		{"javascript:alert('hello')"},
		{"/path/to/file/with spaces.html"},
		{""},
		{"#named-anchor"},
		{"http://example.com/<b>bold</b>"},
		{"https://example.com/<script>alert(1)</script>"},
		{"https://example.com/<foo>"},
		{"https://example.com/?q=<&>"},
	}

	if len(got) != len(expected) {
		t.Fatalf("got %d links, expected %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("got %v, expected %v", got[i], expected[i])
		}
	}
}

func TestParseNested(t *testing.T) {
	test := `<html>
<body>
	<a href="#">
		Main
		<a href="/dope">
			Internal
		</a>
	</a>
</body>
</html>`

	got, err := Parse(strings.NewReader(test))
	if err != nil {
		t.Fatal(err)
	}
	expected := []Link{
		{"#"},
		{"/dope"},
	}

	if len(got) != len(expected) {
		t.Fatalf("got %d links, expected %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("got %v, expected %v", got[i], expected[i])
		}
	}
}

func TestParseComments(t *testing.T) {
	test := `<html>

<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>

</html>`

	got, err := Parse(strings.NewReader(test))
	if err != nil {
		t.Fatal(err)
	}
	expected := []Link{
		{"/dog-cat"},
	}

	if len(got) != len(expected) {
		t.Fatalf("got %d links, expected %d", len(got), len(expected))
	}

	for i := range got {
		if got[i] != expected[i] {
			t.Fatalf("got %v, expected %v", got[i], expected[i])
		}
	}
}

type ErrReader struct{ Error error }

func (e *ErrReader) Read([]byte) (int, error) {
	return 0, e.Error
}

func TestParseNonUTF8(t *testing.T) {
	test := &ErrReader{errors.New("derp")}

	_, err := Parse(test)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
