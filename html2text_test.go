package html2text

import (
	"testing"
)

func TestText(t *testing.T) {
	testCases := []struct {
		input string
		expr  string
	}{
		{
			`<li>
  <a href="/new" data-ga-click="Header, create new repository, icon:repo"><span class="octicon octicon-repo"></span> New repository</a>
</li>`,
			"*  [New repository](http://www.microshwhat.com/bar/soapy/new)",
		},
		{
			`hi

			<br>

	hello <a href="https://google.com">google</a>
	<br><br>
	test<p>List:</p>

	<ul>
		<li><a href="foo">Foo</a></li>
		<li><a href="http://www.microshwhat.com/bar/soapy">Barsoap</a></li>
        <li>Baz</li>
	</ul>
`,
			"hi \n\n hello [google](https://google.com) \n\n test\n\nList: \n\n    * [Foo](foo) \n    * [Barsoap](http://www.microshwhat.com/bar/soapy) \n    * Baz",
		},
		// Malformed input html.
		{
			`hi

			hello <a href="https://google.com">google</a>

			test<p>List:</p>

			<ul>
				<li><a href="foo">Foo</a>
				<li><a href="/
		                bar/baz">Bar</a>
		        <li>Baz</li>
			</ul>
		`,
			"hi hello [google](https://google.com) test\n\nList: \n\n    * [Foo](foo) \n    * [Bar](http://www.microshwhat.com/bar/soapy/bar/baz) \n    * Baz",
		},
		{
			`<table><tr><th>First Column</th><th>Second Column</th><th>Third Column</th></tr>
			<tr><td>row1column1</td><td>row1column2</td><td>row1column3</td></tr>
			<tr><td>row2column1</td><td>row2column2</td><td>row2column3</td></tr>
			</table>`,
			"First Column: row1column1\nSecond Column: row1column2\nThird Column: row1column3\n\nFirst Column: row2column1\nSecond Column: row2column2\nThird Column: row2column3",
		},
	}

	for _, testCase := range testCases {
		text, err := FromString(testCase.input, "http://www.microshwhat.com/bar/soapy", OmitClasses|OmitIds|OmitRoles)
		if err != nil {
			t.Error(err)
		}

		if testCase.expr != text {
			t.Errorf("Input did not match expression\nInput:\n>>>>\n@@%s@@\n<<<<\n\nOutput:\n>>>>\n##%s##\n<<<<\n\nExpression: ^^%s^^\n", testCase.input, text, testCase.expr)
		} else {
			t.Logf("input:\n\n%s\n\n\n\noutput:\n\n%s\n", testCase.input, text)
		}
	}
}
