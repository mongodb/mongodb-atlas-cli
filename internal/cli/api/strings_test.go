package api

import "testing"

func TestSplitShortAndLongDescription(t *testing.T) {
	tests := []struct {
		name        string
		description string
		wantShort   string
		wantLong    string
	}{
		{
			name:        "empty string",
			description: "",
			wantShort:   "",
			wantLong:    "",
		},
		{
			name:        "single sentence",
			description: "This is a single sentence.",
			wantShort:   "This is a single sentence.",
			wantLong:    "",
		},
		{
			name:        "two sentences",
			description: "First sentence. Second sentence.",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence.",
		},
		{
			name:        "multiple sentences with spaces",
			description: "First sentence.   Second sentence.    Third sentence.",
			wantShort:   "First sentence.",
			wantLong:    "Second sentence. Third sentence.",
		},
		{
			name:        "sentence without period",
			description: "This is a sentence without period",
			wantShort:   "This is a sentence without period.",
			wantLong:    "",
		},
		{
			name:        "multiple sentences with extra periods",
			description: "This is version 1.2.3. Second sentence. Third sentence.",
			wantShort:   "This is version 1.2.3.",
			wantLong:    "Second sentence. Third sentence.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShort, gotLong := splitShortAndLongDescription(tt.description)
			if gotShort != tt.wantShort {
				t.Errorf("splitShortAndLongDescription() gotShort = %v, want %v", gotShort, tt.wantShort)
			}
			if gotLong != tt.wantLong {
				t.Errorf("splitShortAndLongDescription() gotLong = %v, want %v", gotLong, tt.wantLong)
			}
		})
	}
}
