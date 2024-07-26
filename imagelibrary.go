package main

import (
	"fmt"
)

// Initialize a slice to store image filenames
var gallery = make([]string, 0, 10)

// Upload image function
func uploadImage(filename string) {
	if len(gallery) < cap(gallery) {
		gallery = append(gallery, filename)
		fmt.Printf("add %s to gallery, new gallery len(gallery) is %d\n", filename, len(gallery))
	} else {
		fmt.Println("gallery is full")
	}
}

// Delete image function
func deleteImage(filename string) {
	var deletedfile string
	fmt.Printf("\nb4 deletion of %s from gallery %v\n", filename, gallery)
	for i := 0; i < len(gallery); i++ {
		//fmt.Printf("delete loop at index %d where item is %s checking for %s in %v \n", i, gallery[i], filename, gallery)
		if gallery[i] == filename {
			deletedfile = gallery[i]
			fmt.Println("Found !!! match file to be deleted filename, ", filename, "match, ", gallery[i])
			//fmt.Printf("b4 shift: deleted %s from gallery at index %d, new gallery %v\n", filename, i, gallery)
			//fmt.Println("at shift:gallery[0:i],", gallery[0:i], " , gallery[i+1:len(gallery)+1], ", gallery[i+1:len(gallery)+1])
			//gallery = append(gallery[0:i], gallery[i+1:len(gallery)+1]...)
			for ix, value := range gallery[i+1:] {
				fmt.Printf("gallery check B4 manipulation: value %s at index %d from range %v\n", value, ix, gallery[i+1:])
				fmt.Println("from gallery,", gallery)
				fmt.Println("old gallery[i+ix],", gallery[i+ix])
				gallery[i+ix] = value
				fmt.Printf("gallery check AFTER manipulation: value %s at index %d from range %v\n", value, ix, gallery[i+1:])
				fmt.Println("from gallery,", gallery)
				fmt.Println("new gallery[i+ix],", gallery[i+ix])
			}
			fmt.Println("gallery before shorten to len(gallery-1), ", gallery)
			gallery = gallery[:len(gallery)-1]
			fmt.Println("gallery after shorten to len(gallery-1), ", gallery)
			//fmt.Printf("deleted %s from gallery at index %d, new gallery %v\n", filename, i, gallery)
		}
	}

	if deletedfile == "" {
		fmt.Println(filename, "not found")
	} else {
		fmt.Println("after deletion, gallery contains: ", gallery)
	}
}

// List images function
func listImages() {
	if len(gallery) > 0 {
		fmt.Println("images in the gallery now: ")
		var item string
		for _, item = range gallery {
			fmt.Println(item)
		}
	} else {
		fmt.Println("gallery is empty")
	}
}

func main() {
	fmt.Println("Testing Image Gallery Management System")
	uploadImage("image1.jpg")
	uploadImage("image2.jpg")
	uploadImage("image3.jpg")
	listImages()
	deleteImage("image2.jpg")
	listImages()
	uploadImage("image4.jpg")
	listImages()
	deleteImage("image1.jpg")
	deleteImage("image2.jpg")
	deleteImage("image3.jpg")
	uploadImage("image1.jpg")
	uploadImage("image2.jpg")
	uploadImage("image3.jpg")
	uploadImage("image4.jpg")
	uploadImage("image5.jpg")
	uploadImage("image6.jpg")
	uploadImage("image7.jpg")
	deleteImage("image5.jpg")
	listImages()
}

/*for ix, value := range gallery[i+1:] {
    gallery[i+ix] = value
}
gallery = gallery[:len(gallery)-1]*/

//old implementation
//gallery = append(gallery[0:i], gallery[i+1:len(gallery)+1]...)
