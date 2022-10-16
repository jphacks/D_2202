// detectWeb gets image properties from the Vision API for an image at the given file path.
package main

import (
	"context"
	"fmt"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func detectWeb(file string) error {
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

	web, err := client.DetectWeb(ctx, image, nil)
	if err != nil {
		return err
	}

	print("client detected web-info successfully\n")

	fmt.Println("Web properties:")
	if len(web.FullMatchingImages) != 0 {
		fmt.Println("\tFull image matches:")
		for _, full := range web.FullMatchingImages {
			fmt.Printf("\t\t%s\n", full.Url)
		}
	}
	if len(web.PagesWithMatchingImages) != 0 {
		fmt.Println("\tPages with this image:")
		for _, page := range web.PagesWithMatchingImages {
			fmt.Printf("\t\t%s\n", page.Url)
		}
	}
	if len(web.WebEntities) != 0 {
		fmt.Println("\tEntities:")
		fmt.Println("\t\tEntity\t\tScore\tDescription")
		for _, entity := range web.WebEntities {
			fmt.Printf("\t\t%-14s\t%-2.4f\t%s\n", entity.EntityId, entity.Score, entity.Description)
		}
	}
	if len(web.BestGuessLabels) != 0 {
		fmt.Println("\tBest guess labels:")
		for _, label := range web.BestGuessLabels {
			fmt.Printf("\t\t%s\n", label.Label)
		}
	}

	return nil
}
