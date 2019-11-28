/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package pageview

//lint:file-ignore ST1017 - I prefer Yoda conditions

/*
	CREDITS:
	This file started as a modified/optimised version of
		github.com/ninetwentyfour/go-wkhtmltoimage
		Copyright (c) 2015 ninetwentyfour
*/

import (
	"bytes"
	"context"
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// `buildParams()` takes `aURL` set by the user and prepares the
// required commandline options for `wkhtmltoimage`, returning
// the list of those options.
//
//	`aURL` The remote URL to be handled by `wkhtmltoimage`.
func buildParams(aURL string) (rList []string, rErr error) {
	if 0 == len(wkHTMLToImageBinary) {
		return rList, errors.New("BinaryPath not set")
	}
	if 0 == len(aURL) { //lint:ignore ST1005 - I want this
		return rList, errors.New("Input not set")
	}

	rList = []string{
		"-q",
		"--disable-plugins",
		"--load-error-handling",
		"ignore",
		"--format",
		wkImageFileType,
	}
	if ucd, err := os.UserCacheDir(); (nil == err) && (0 < len(ucd)) {
		rList = append(rList, "--cache-dir", ucd)
	}
	if 0 < wkImageHeight {
		rList = append(rList, "--height", strconv.Itoa(wkImageHeight))
	}
	if 0 < wkImageQuality {
		rList = append(rList, "--quality", strconv.Itoa(wkImageQuality))
	}
	if 0 < wkImageWidth {
		rList = append(rList, "--width", strconv.Itoa(wkImageWidth))
	}
	rList = append(rList, aURL, "-") // i.e. send data to StdOut

	return
} // buildParams()

// `cleanupOutput()` removes unneeded leading data from `aRawData`
// and returns the properly encoded image.
//
//	`aRawData` The raw image data to cleanup.
func cleanupOutput(aRawData []byte) []byte {
	if 0 == len(aRawData) {
		return aRawData
	}
	var buffer bytes.Buffer

	switch wkImageFileType {
	case `gif`:
		decoded, err := gif.Decode(bytes.NewReader(aRawData))
		for nil != err {
			if aRawData = aRawData[1:]; 0 == len(aRawData) {
				return aRawData
			}
			decoded, err = gif.Decode(bytes.NewReader(aRawData))
		}
		opts := gif.Options{NumColors: wkImageQuality}
		_ = gif.Encode(&buffer, decoded, &opts)
		return buffer.Bytes()

	case `jpg`:
		decoded, err := jpeg.Decode(bytes.NewReader(aRawData))
		for nil != err {
			if aRawData = aRawData[1:]; 0 == len(aRawData) {
				return aRawData
			}
			decoded, err = jpeg.Decode(bytes.NewReader(aRawData))
		}
		opts := jpeg.Options{Quality: wkImageQuality}
		_ = jpeg.Encode(&buffer, decoded, &opts)
		return buffer.Bytes()

	case `png`:
		decoded, err := png.Decode(bytes.NewReader(aRawData))
		for nil != err {
			if aRawData = aRawData[1:]; 0 == len(aRawData) {
				return aRawData
			}
			decoded, err = png.Decode(bytes.NewReader(aRawData))
		}
		_ = png.Encode(&buffer, decoded)
		return buffer.Bytes()

	case `svg`:
		// nothing to do here
	}

	return aRawData
} // cleanupOutput()

// `generateImage()` creates an image from an input.
// It returns the image data and any error encountered.
//
//	`aURL` The remote URL to be handled by `wkhtmltoimage`.
func generateImage(aURL string) (rImage []byte, rErr error) {
	var (
		options []string
		rawData []byte
	)
	if options, rErr = buildParams(aURL); rErr != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// For some reason (e.g. network errors) `wkhtmltoimage` sometimes
	// hangs – possibly indefinitely. Therefor we use a timeout to let
	// this function continue. The timeout value should be long enough
	// to allow running both `exec.CommandContext()` and `cmd.Output()`.
	defer cancel()
	cmd := exec.CommandContext(ctx, wkHTMLToImageBinary, options...) //#nosec G204
	if rawData, rErr = cmd.Output(); nil != rawData {
		if rImage = cleanupOutput(rawData); 0 < len(rImage) {
			rErr = nil
		}
	}

	return
} // generateImage()

/* _EoF_ */
