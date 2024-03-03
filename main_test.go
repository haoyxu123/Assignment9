package main

import (
	"testing"
)

// BenchmarkEntirePipeline measures the performance of the entire image processing pipeline.
func BenchmarkEntirePipeline(b *testing.B) {
	imagePaths := []string{
		"C:/Assignment9/go_21_goroutines_pipeline/images/audi1.jpg",
		"C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg",
		"C:/Assignment9/go_21_goroutines_pipeline/images/benz2.jpg",
		"C:/Assignment9/go_21_goroutines_pipeline/images/bmw.jpg",
	}
	newWidth := uint(500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1, newWidth)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if !success {
				b.Error("Failed to process an image")
			}
		}
	}
}
