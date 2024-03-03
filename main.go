package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"log"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			job := Job{InputPath: p, OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				log.Printf("failed to read image %s: %v", p, err)
				continue
			}
			job.Image = img
			out <- job
		}
		close(out)
	}()
	return out
}

//	func resize(input <-chan Job) <-chan Job {
//		out := make(chan Job)
//		go func() {
//			// For each input job, create a new job after resize and add it to
//			// the out channel
//			for job := range input { // Read from the channel
//				job.Image = imageprocessing.Resize(job.Image)
//				out <- job
//			}
//			close(out)
//		}()
//		return out
//	}
func resize(input <-chan Job, newWidth uint) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Resize(job.Image, newWidth)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			err := imageprocessing.WriteImage(job.OutPath, job.Image)
			if err != nil {
				log.Printf("failed to write image %s: %v", job.OutPath, err)
				out <- false
				continue
			}
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

	imagePaths := []string{"C:/Assignment9/go_21_goroutines_pipeline/images/audi1.jpg",
		"C:/Assignment9/go_21_goroutines_pipeline/images/benz1.jpg",
		"C:/Assignment9/go_21_goroutines_pipeline/images/benz2.jpg",
		"C:/Assignment9/go_21_goroutines_pipeline/images/bmw.jpg",
	}
	newWidth := uint(500)

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1, newWidth)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
