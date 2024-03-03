package formatting

import (
	"bytes"
	"fmt"
	"html/template"
)

type Path string

const (
	Index      Path = "/"
	Generate   Path = "/generate"
	ResetIndex Path = "/reset"
	Help       Path = "/help"
)

const navbar = `
<nav>
    <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
        <div class="relative flex items-center justify-between h-16">
            <div class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
                <div class="sm:ml-6">
                    <div class="flex space-x-4">
                        %s
                    </div>
                </div>
            </div>
            <div class="ml-auto flex items-center">
                <div class="flex space-x-4">
                    %s
                </div>
            </div>
        </div>
    </div>
</nav>
`

const (
	IndexButtonSelected = `<a href="/" class="bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium">Index</a>`
	IndexButton         = `<a href="/" class="text-gray-600 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium">Index</a>`

	GenerateButtonSelected = `<a href="/generate" class="bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium">Generate</a>`
	GenerateButton         = `<a href="/generate" class="text-gray-600 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium">Generate</a>`

	ResetIndexButtonSelected = `<a href="/reset" class="bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium">Reset</a>`
	ResetIndexButton         = `<a href="/reset" class="text-gray-600 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium">Reset</a>`

	GithubButton = `<a target="_blank" href="https://github.com/ValeryVerkhoturov/science-generation" class="text-gray-600 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium">Github</a>`

	HelpButtonSelected = `<a href="/help" class="bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium">Help</a>`
	HelpButton         = `<a href="/help" class="text-gray-600 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium">Help</a>`
)

func NewNavbarTitle(path Path) template.HTML {
	leftButtons := make([]string, 0)
	rightButtons := make([]string, 0)
	switch path {
	case Index:
		leftButtons = append(leftButtons, IndexButtonSelected, GenerateButton, ResetIndexButton)
		rightButtons = append(rightButtons, GithubButton, HelpButton)
	case Generate:
		leftButtons = append(leftButtons, IndexButton, GenerateButtonSelected, ResetIndexButton)
		rightButtons = append(rightButtons, GithubButton, HelpButton)
	case ResetIndex:
		leftButtons = append(leftButtons, IndexButton, GenerateButton, ResetIndexButtonSelected)
		rightButtons = append(rightButtons, GithubButton, HelpButton)
	case Help:
		leftButtons = append(leftButtons, IndexButton, GenerateButton, ResetIndexButton)
		rightButtons = append(rightButtons, GithubButton, HelpButtonSelected)
	}

	return newNavbarTitle(leftButtons, rightButtons)
}

func newNavbarTitle(leftButtons []string, rightButtons []string) template.HTML {
	return template.HTML(fmt.Sprintf(navbar, concatButtons(leftButtons), concatButtons(rightButtons)))
}

func concatButtons(buttons []string) string {
	var buffer bytes.Buffer
	for _, button := range buttons {
		buffer.WriteString(button)
	}
	return buffer.String()
}
