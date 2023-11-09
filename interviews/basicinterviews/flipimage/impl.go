package flipimage

func flipInvertImage(image [][]int) [][]int {
	for _, row := range image {
		for i := 0; i*2 < len(row); i++ {
			// Flip and invert the element in one go.
			// XOR with 1 will invert the bits, 0 becomes 1 and 1 becomes 0.
			row[i], row[len(row)-1-i] = row[len(row)-1-i]^1, row[i]^1
		}
	}
	return image
}

func flipInvertImage2(image [][]int) [][]int {
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
