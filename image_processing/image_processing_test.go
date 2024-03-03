package imageprocessing

import (
	"image"
	"image/color"
	"testing"
)

// TestReadImage tests the ReadImage function
func TestReadImage(t *testing.T) {
	_, err := ReadImage("C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg")
	if err != nil {
		t.Errorf("ReadImage failed for valid image: %v", err)
	}
}

// TestWriteImage tests the WriteImage function
func TestWriteImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	err := WriteImage("C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg", img)
	if err != nil {
		t.Errorf("WriteImage failed for valid path: %v", err)
	}
}

// TestResize tests the Resize function by verifying that the output image has the expected dimensions
func TestResize(t *testing.T) {
	srcImage := image.NewRGBA(image.Rect(0, 0, 100, 50))

	newWidth := uint(50)

	resizedImage := Resize(srcImage, newWidth)

	expectedHeight := uint(25)
	if resizedImage.Bounds().Dx() != int(newWidth) || resizedImage.Bounds().Dy() != int(expectedHeight) {
		t.Errorf("Resized image has incorrect dimensions: got %dx%d, want %dx%d",
			resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy(), newWidth, expectedHeight)
	}
}

func TestGrayscale(t *testing.T) {
	originalImg := image.NewRGBA(image.Rect(0, 0, 2, 2))
	originalImg.Set(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	originalImg.Set(1, 0, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	originalImg.Set(0, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	originalImg.Set(1, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})

	grayImg := Grayscale(originalImg)

	bounds := grayImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			got := color.GrayModel.Convert(grayImg.At(x, y))
			if _, _, _, a := got.RGBA(); a != 65535 {
				t.Errorf("pixel at (%d, %d) is not grayscale, got: %v", x, y, got)
			}
		}
	}
}

func BenchmarkReadImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReadImage("C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkWriteImage(b *testing.B) {
	img, _ := ReadImage("C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := WriteImage("C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg", img)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkResize tests the performance of the Resize function
func BenchmarkResize(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 1024, 768))
	newWidth := uint(100)
	for i := 0; i < b.N; i++ {
		_ = Resize(img, newWidth)
	}
}

func BenchmarkGrayscale(b *testing.B) {
	img, _ := ReadImage("C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Grayscale(img)
	}
}
