// The jpeg command read a PNG image from the standard input
// and writes it as a JPEG image to the standard output.

package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG dcoder
	"io"
	"os"
)

func main() {
	if err := toJPEC(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format=", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
