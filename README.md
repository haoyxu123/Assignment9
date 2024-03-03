# Assignment 9: Data Pipelines with Concurrency 

## Clone the GitHub repository for image processing
Clone the repository from github as the mode to achieve this assignment. URL https://github.com/code-heim/go_21_goroutines_pipeline
## Build and run the program in its original form
type go run main. go in the terminal which has 4 outputs shows success
## Add error checking for image file input and output.
In the main.go file, I added the error checking for the loadImage function and saveImage function.
## Replace the four input image files with files of your choosing
Upload four images of cars that are different brands. I did not remove the original images and saved the all outputs in the images folder. All outputs and inputs can be found in that folder. 
## Add unit tests to the code repository.
Create unit tests for all the functions in the image_processing.go. The first one is TestReadImage, the second is func TestWriteImage, the third is func TestResize, and the last one is func TestGrayscale. They are all saved in the file image_processing_test.go. 
## Add benchmark methods for capturing pipeline throughput times
Under the unit tests, there are four different benchmarks for four different functions. There are func BenchmarkReadImage, func BenchmarkWriteImage, func BenchmarkResize,func BenchmarkGrayscale. All unit tests and benchmarks are saved in the file image_processing_test. go.
## Make additional code modifications as you see fit
I modify the function Resize in the image_processing.go. I found the previous Resize function hardcodes the width and height to 500 pixels which will distort my images if the images are not 500*500. The new Resize function introduces flexibility via the newWidth parameter. This approach prevents distortion. 
