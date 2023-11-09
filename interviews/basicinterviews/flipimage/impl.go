package flipimage

func flipInvertImage(image [][]int) [][]int {
	// Flip the image horizontally
	for _, row := range image {
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
	}

	// Invert the image
	for i := range image {
		for j := range image[i] {
			// Invert each element: 0 to 1 and 1 to 0
			image[i][j] ^= 1
		}
	}
	return image
}
