package ex17

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	td := []struct {
		node     string
		name     []string
		expected int
	}{
		{`
  <html>
    <div>
      <div>'inner nested string'</div>
      <img src='innerImage.jpg'/>
      <img src='innerImage.jpg'/>
      'hogehoge and piyo'
			<h1>'h1だよ'</h1>
			<h1>'h1だよ'</h1>
    </div>
    <img src='here.gif' />
		<h1>'h1だよ'</h1>
    <img src='here.gif' />
  </html>`, []string{"img"}, 4},
		{`
  <html>
    <div>
      <div>'inner nested string'</div>
      <img src='innerImage.jpg'/>
      <img src='innerImage.jpg'/>
      'hogehoge and piyo'
			<h1>'h1だよ'</h1>
			<h1>'h1だよ'</h1>
    </div>
    <img src='here.gif' />
		<h1>'h1だよ'</h1>
    <img src='here.gif' />
  </html>`, []string{"h1"}, 3},
		{`
  <html>
    <div>
      <div>'inner nested string'</div>
      <img src='innerImage.jpg'/>
      <img src='innerImage.jpg'/>
      'hogehoge and piyo'
			<h1>'h1だよ'</h1>
			<h1>'h1だよ'</h1>
    </div>
    <img src='here.gif' />
		<h1>'h1だよ'</h1>
		<h2>'h2だよ'</h2>
    <img src='here.gif' />
  </html>`, []string{"h1", "h2", "h3"}, 4},
	}

	for _, test := range td {
		doc, _ := html.Parse(strings.NewReader(test.node))
		res := ElementsByTagName(doc, test.name[0:]...)

		if len(res) != test.expected {
			t.Errorf("expected %d elements, got %d\n", test.expected, len(res))
		}
	}
}
