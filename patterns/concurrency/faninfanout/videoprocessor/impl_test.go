package videoprocessor

import (
	"testing"
)

// TestVideoProcessor runs all unit tests for the video processor package.
func TestVideoProcessor(t *testing.T) {
	// Shared setup (if any) goes here.

	// Subtest for ProcessSegment func
	t.Run("Process single segment", func(t *testing.T) {
		segment := VideoSegment{ID: 1}
		encoded := ProcessSegment(segment)
		if encoded.ID != segment.ID {
			t.Errorf("ProcessSegment: Expected ID %d, got %d", segment.ID, encoded.ID)
		}
	})

	// Subtest for FanOutFanIn func
	t.Run("Process multiple segments", func(t *testing.T) {
		segments := []VideoSegment{{ID: 1}, {ID: 2}, {ID: 3}}
		numWorkers := 2
		encodedSegments := FanOutFanIn(segments, numWorkers)
		if len(encodedSegments) != len(segments) {
			t.Errorf("FanOutFanIn: Expected %d encoded segments, got %d", len(segments), len(encodedSegments))
		}
		for _, encoded := range encodedSegments {
			if encoded.EncodedData == nil {
				t.Errorf("FanOutFanIn: EncodedData is nil for segment ID %d", encoded.ID)
			}
		}
	})
}
