package render

import (
	"strings"
	"testing"
)

// TestGenerate verifies that the rendering engine correctly transforms input text
// into standard multi-layer structural ASCII art blocks using a mock banner font array.
func TestGenerate(t *testing.T) {
	// Create a minimal mock banner array representing 9 lines per character.
	// For simplicity, we mock character 32 (Space) and character 65 ('A').
	// Formula: start = int(char - 32) * 9 + 1
	// Space (32) index starts at: (32-32)*9 + 1 = 1
	// 'A' (65) index starts at: (65-32)*9 + 1 = 298
	mockBanner := make([]string, 350)

	// Mocking Space lines (rows 1 to 8 inside the banner block)
	for i := 1; i <= 8; i++ {
		mockBanner[i] = "     " // 5 spaces per row for space character
	}

	// Mocking character 'A' lines (rows 298 to 305 inside the banner block)
	mockBanner[298] = "  A  "
	mockBanner[299] = " A A "
	mockBanner[300] = "A   A"
	mockBanner[301] = "AAAAA"
	mockBanner[302] = "A   A"
	mockBanner[303] = "A   A"
	mockBanner[304] = "A   A"
	mockBanner[305] = "A   A"

	// Define table-driven test cases to evaluate diverse rendering pipelines
	tests := []struct {
		name      string // Descriptive label for the specific test case
		inputText string // The simulated input string from the user form
		expected  string // The precise multi-line output structural string expected
	}{
		{
			name:      "Single character render",
			inputText: "A",
			expected:  "  A  \n A A \nA   A\nAAAAA\nA   A\nA   A\nA   A\nA   A",
		},
		{
			name:      "Space handling integration",
			inputText: "A A",
			expected:  "  A         A  \n A A       A A \nA   A     A   A\nAAAAA     AAAAA\nA   A     A   A\nA   A     A   A\nA   A     A   A\nA   A     A   A",
		},

		{
			name:      "Literal newline character expansion",
			inputText: "A\\nA",
			expected:  "  A  \n A A \nA   A\nAAAAA\nA   A\nA   A\nA   A\nA   A\n  A  \n A A \nA   A\nAAAAA\nA   A\nA   A\nA   A\nA   A",
		},
		{
			name:      "Zone01 Special Case: Multiple pure newlines",
			inputText: "\\n\\n",
			expected:  "\n\n",
		},
	}

	// Iterate through all defined table evaluation records
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the actual rendering function with the text input and mock asset
			actual := Generate(tt.inputText, mockBanner)

			// Standardize structural data strings by removing extra system windows carriage returns
			actualClean := strings.ReplaceAll(actual, "\r", "")
			expectedClean := strings.ReplaceAll(tt.expected, "\r", "")

			// Evaluate if the output perfectly matches criteria shapes
			if actualClean != expectedClean {
				t.Errorf("\n[FAILED] Test case: %s\nInput used: %q\nEXPECTED:\n%s\nACTUAL:\n%s\n",
					tt.name, tt.inputText, expectedClean, actualClean)
			}
		})
	}
}
