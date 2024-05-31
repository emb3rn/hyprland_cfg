package main

import (
	"fmt"
	"image"
	"image/color"
	//"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

/* import "rsc.io/quote" */

func decodeImage() (image.Image){
	file, err := os.Open("./IMG_3947.png")
	if err != nil {log.Fatal(err)}
	
	img, err := png.Decode(file)
	if err != nil{log.Fatal(err)}
	
	defer file.Close()
	return img
}

func caclulateColors(img image.Image) ([]color.Color){
	size := (img.Bounds().Max.X*img.Bounds().Max.Y)
	colorArray := make([]color.Color, size) 

	idx := 0
	for j := 0; j < img.Bounds().Max.Y; j++{
		for i := 0; i < img.Bounds().Max.X; i++{
			colorArray[idx] = img.At(i,j)
			idx++
		}
	}

	return colorArray
}

func calculateLight(colors []color.Color) ([]int, int){
	lightArray := make([]int, len(colors))
	lightMax, lightMin := -1, math.MaxInt32 

	for i, color := range colors{
		r, g, b, _ := color.RGBA()
		sum := int(r+g+b)
		lightArray[i] = sum

		if sum > lightMax{lightMax = sum}
		if sum < lightMin{lightMin = sum}
	}

	return lightArray, lightMax 
}

func createASCII(lightArray []int, maxLight int, characters []byte, compression int) (string){
	asciiArray := make([]byte, len(lightArray))	
	
	i, j := 0, 0
	for i < len(lightArray)-compression{ //Calculate light percent of max (0-1), multiply by length of chars, cast to int, get correct char
		combinedSum := 0
		
		for range compression{
			combinedSum += lightArray[i]
			i++	
		}
		
		//For some reason compression no worky, fix soon
		lightPercentage := (float32(combinedSum/compression) / float32(maxLight))
		char := characters[int(lightPercentage * float32(len(characters)-1))]
		asciiArray[j] = char
		j++
	}
	
	asciiStr := string(asciiArray)
	return asciiStr		
}

func main(){
	ASCII_CHARACTERS := []byte("$EFLlv!;,.") //bright->dark
	IMAGE_COMPRESSION := 2

	decodedImage := decodeImage()
	imageColors := caclulateColors(decodedImage)
	imageLights, lightMax := calculateLight(imageColors)
	asciiString := createASCII(imageLights, lightMax, ASCII_CHARACTERS, IMAGE_COMPRESSION)
	
	idx := 0
	for j := 0; j < decodedImage.Bounds().Max.Y/IMAGE_COMPRESSION; j++{
		for i := 0; i < decodedImage.Bounds().Max.X/IMAGE_COMPRESSION; i++{
			fmt.Printf("%v", asciiString[idx])	
			idx++
		}
		fmt.Printf("\n")
	}


}

