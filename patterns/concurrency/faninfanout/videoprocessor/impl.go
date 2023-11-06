package videoprocessor

import (
	"sync"
)

type VideoSegment struct {
	ID     int
	Frames []Frame
}

type Frame struct {
	// Frame data
}

type EncodedSegment struct {
	ID          int
	EncodedData []byte
}

func ProcessSegment(segment VideoSegment) EncodedSegment {
	// Simulate encoding logic.
	// For testing, I create a dummy EncodedData.
	dummyData := []byte{1, 2, 3} // Dummy data to represent encoded data.

	return EncodedSegment{
		ID:          segment.ID,
		EncodedData: dummyData,
	}
}

func FanOutFanIn(segments []VideoSegment, numWorkers int) []EncodedSegment {
	segmentChan := make(chan VideoSegment, len(segments))
	encodedChan := make(chan EncodedSegment, len(segments))

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for segment := range segmentChan {
				encodedChan <- ProcessSegment(segment)
			}
		}()
	}

	for _, segment := range segments {
		segmentChan <- segment
	}
	close(segmentChan)

	go func() {
		wg.Wait()
		close(encodedChan)
	}()

	var encodedSegments []EncodedSegment
	for encoded := range encodedChan {
		encodedSegments = append(encodedSegments, encoded)
	}

	return encodedSegments
}
