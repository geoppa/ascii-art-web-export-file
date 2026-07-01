package validation

import "testing"

// TestValidateText checks if the validation logic correctly filters printable ASCII inputs
func TestValidateText(t *testing.T) {
	tests := []struct {
		name        string
		inputText   string
		shouldError bool
	}{
		{
			name:        "Valid standard English text",
			inputText:   "Hello World 123!",
			shouldError: false,
		},
		{
			name:        "Invalid non-ASCII characters",
			inputText:   "Καλημέρα (Greek characters)",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateText(tt.inputText)
			if (err != nil) != tt.shouldError {
				t.Errorf("ValidateText() error = %v, shouldError = %v", err, tt.shouldError)
			}
		})
	}
}

// TestValidateBanner ensures only known system styles are approved
func TestValidateBanner(t *testing.T) {
	tests := []struct {
		name        string
		bannerName  string
		shouldError bool
	}{
		{
			name:        "Valid banner style target",
			bannerName:  "standard",
			shouldError: false,
		},
		{
			name:        "Invalid unknown banner asset",
			bannerName:  "custom-font",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBanner(tt.bannerName)
			if (err != nil) != tt.shouldError {
				t.Errorf("ValidateBanner() error = %v, shouldError = %v", err, tt.shouldError)
			}
		})
	}
}
