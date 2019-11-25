/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package pagethumb

//lint:file-ignore ST1017 - I prefer Yoda conditions

/*
	CREDITS:
	This file is a modified/optimised version of
		github.com/ninetwentyfour/go-wkhtmltoimage
		Copyright (c) 2015 ninetwentyfour
*/

import (
	"bytes"
	"errors"
	"image/png"
	"os/exec"
	"strconv"
)

// tImageOptions represent the options to generate the image.
type tImageOptions struct {
	// BinaryPath is the path to your wkhtmltoimage binary. REQUIRED
	//
	// Must be an absolute path e.g `/usr/local/bin/wkhtmltoimage`.
	BinaryPath string

	// Height in pixels of the imaginary screen used to render.
	//
	// Default is calculated from the page content;
	// defaults to `0` (zero) which renders the entire page top to bottom.
	Height int

	// Input is the URL to turn into an image. REQUIRED
	Input string

	// Quality determines the final image quality.
	//
	// Values supported between `1` and `100`. Default is 94.
	Quality int

	// Width in pixels of the imaginary screen used to render.
	//
	// Note that this is used only as a guide line.
	// Default is 1024.
	Width int
}

// `buildParams()` takes the image options set by the user
// and turns them into command flags for `wkhtmltoimage`.
// It returns a list of command flags.
//
//	`aOptions` The commandline options for `wkhtmltoimage`.
func buildParams(aOptions *tImageOptions) (rList []string, rErr error) {
	if 0 == len(aOptions.Input) {
		return rList, errors.New("Input not set")
	}
	if 0 == len(aOptions.BinaryPath) {
		return rList, errors.New("BinaryPath not set")
	}

	rList = []string{
		"-q", // silence extra `wkhtmltoimage` output
		"--disable-plugins",
		"--format",
		wkImageFileType, // `PNG` format because it scales better.
	}
	if 0 < aOptions.Height {
		rList = append(rList, "--height", strconv.Itoa(aOptions.Height))
	}
	if 0 < aOptions.Width {
		rList = append(rList, "--width", strconv.Itoa(aOptions.Width))
	}
	if (0 < aOptions.Quality) && (101 > aOptions.Quality) {
		rList = append(rList, "--quality", strconv.Itoa(aOptions.Quality))
	}
	rList = append(rList, aOptions.Input, "-")

	return
} // buildParams()

// `cleanupOutput()` returns `aImage` with unneeded data removed.
//
//	`aImage` The raw image data to cleanup.
func cleanupOutput(aImage []byte) []byte {
	var buf bytes.Buffer

	decoded, err := png.Decode(bytes.NewReader(aImage))
	for nil != err {
		if aImage = aImage[1:]; 0 == len(aImage) {
			break
		}
		decoded, err = png.Decode(bytes.NewReader(aImage))
	}
	png.Encode(&buf, decoded)

	return buf.Bytes()
} // cleanupOutput()

// `generateImage()` creates an image from an input.
// It returns the image data and any error encountered.
//
//	`aOptions` The commandline options for `wkhtmltoimage`.
func generateImage(aOptions *tImageOptions) (rImage []byte, rErr error) {
	var (
		flags    []string
		rawImage []byte
	)
	if flags, rErr = buildParams(aOptions); rErr != nil {
		return
	}

	//TODO add context with timeout

	cmd := exec.Command(aOptions.BinaryPath, flags...)
	rawImage, rErr = cmd.CombinedOutput()
	rImage = cleanupOutput(rawImage)

	return
} // generateImage()
