// detectLandmarks gets landmarks from the Vision API for an image at the given file path.
package main

import (
	"context"
	"fmt"
	"os"

	// "os"

	vision "cloud.google.com/go/vision/apiv1"
)

func detectLandmarks(file string) error {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return err
	}
	fmt.Print("client created successfully\n")

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		return err
	}
	print("image file successfully encoded\n")

	annotations, err := client.DetectLandmarks(ctx, image, nil, 10)
	if err != nil {
		return err
	}
	print("client detected landmarks successfully\n")

	if len(annotations) == 0 {
		fmt.Println("No landmarks found.")
	} else {
		fmt.Println("Landmarks:")
		for _, annotation := range annotations {
			fmt.Println(annotation.Description)
		}
	}

	return nil
}
