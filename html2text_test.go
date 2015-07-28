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
			"*   (http://www.microshwhat.com/bar/soapy/new)[New repository]",
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
			"hi \n\n hello  (https://google.com)[google] \n\n test\n\nList: \n\n    *  (foo)[Foo] \n    *  (http://www.microshwhat.com/bar/soapy)[Barsoap] \n    * Baz",
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
			"hi hello  (https://google.com)[google] test\n\nList: \n\n    *  (foo)[Foo] \n    *  (http://www.microshwhat.com/bar/soapy/bar/baz)[Bar] \n    * Baz",
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
