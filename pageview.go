/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package pageview

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

const (
	// Name of the binary tool.
	wkHTMLconverterBinary = "wkhtmltoimage"
)

var (
	// Path/filename of the `wkhtmltoimage` executable.
	wkHTMLToImageBinary = func() string {
		if cmd, err := exec.LookPath(wkHTMLconverterBinary); nil == err {
			return cmd
		}
		return ""
	}()

	// Max. age of cached page images (in seconds);
	// `0` (zero) disables the age check.
	wkImageAge time.Duration = 0

	// Directory to store the generated images.
	wkImageDirectory = func() string {
		result, _ := filepath.Abs("./")
		return result
	}()

	// Type/Format of the generated images.
	// We use `PNG` format because it scales better.
	wkImageFileType = `png`

	// Height of the image to generate.
	wkImageHeight = 768

	// Quality of the image to generate.
	wkImageQuality = 100

	// Width of the image to generate.
	wkImageWidth = 1024

	// R/O RegEx to find all non alpha/digits in URLs.
	wkReplaceNonAlphas = regexp.MustCompile(`\W+`)
)

// `exists()` returns whether there is an usable file cached.
//
// This function uses the `MaxAge()` value to determine whether
// an already existing file is considered to be too old.
func exists(aFilename string) bool {
	fi, err := os.Stat(aFilename)
	if (nil != err) || fi.IsDir() {
		return false
	}

	if 0 >= fi.Size() {
		// empty files are ignored
		return false
	}

	if 0 < wkImageAge {
		maxTime := fi.ModTime().Add(wkImageAge * time.Second)
		// files too old are ignored
		return time.Now().Before(maxTime)
	}

	return true
} // exists()

// `sanitise()` returns `aURL` with all non alpha/digits removed.
//
//	`aURL` The URL to sanitise.
func sanitise(aURL string) string {
	return wkReplaceNonAlphas.ReplaceAllLiteralString(aURL, ``)
} // sanitise()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// CreateImage generates an image of `aURL` and stores it in
// `CacheDirectory()`, returning the file name of the saved image.
//
//	`aURL` The address of the web page to process.
func CreateImage(aURL string) (string, error) {
	if 0 == len(wkHTMLToImageBinary) {
		// We can't do anything without the executable.
		return "", &exec.Error{
			Name: wkHTMLconverterBinary,
			Err:  exec.ErrNotFound,
		}
	}

	result := sanitise(aURL) + `.` + wkImageFileType
	fName := filepath.Join(wkImageDirectory, result)
	// Check whether we've already got an image so
	// we might avoid additional network traffic:
	if exists(fName) {
		return result, nil
	}

	imageData, err := generateImage(aURL)
	if nil != err {
		// Either `wkhtmltoimage` produced an error
		// or it took too long and was canceled.
		return "", err
	}

	file, err := os.OpenFile(fName,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640) // #nosec G302
	if nil != err {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(imageData); nil != err {
		// In case of errors during write we delete the file
		// ignoring possible errors and return the write error.
		_ = file.Close()
		file = nil
		_ = os.Remove(fName)
		return "", err
	}

	return result, nil
} // CreateImage()

// ImageDirectory returns the directory used to store the generated images.
func ImageDirectory() string {
	return wkImageDirectory
} // ImageDirectory()

// SetImageDirectory sets the directory to use for storing the
// generated images, returning an error if `aDirectory` can't be used.
//
//	`aDirectory` The directory to store the generated images.
func SetImageDirectory(aDirectory string) error {
	dir, err := filepath.Abs(aDirectory)
	if nil == err {
		wkImageDirectory = dir
	}

	return err
} // SetImageeDirectory()

// ImageFileType returns the type of the image fles to generate.
func ImageFileType() string {
	return wkImageFileType
} // ImageFileType()

// SetImageFileType changes the type of the generated images.
// The default type is `png`, the other options are `gif`, `jpg` and `svg`.
// Passing an invalid value in `aType` will result in `png` being
// selected.
//
// NOTE: Depending on how your `wkhtmltoimage` binary was compiled not
// all formats might be supported.
//
//	`aType` is the new desired type of the images to generate.
func SetImageFileType(aType string) {
	switch aType {
	case `gif`, `jpg`, `png`, `svg`:
		wkImageFileType = aType
	default:
		wkImageFileType = `png`
	}
} // SetImageFileType()

// ImageHeight is the height in pixels of the imaginary screen used to render.
// The default value is `768`.
//
// The value `0` (zero) renders the entire page top to bottom,
// calculating the actual height from the page content.
func ImageHeight() int {
	return wkImageHeight
} // ImageHeight()

// SetImageHeight sets the height of the images to generate.
// The default value is `768`.
//
//	`aHeight` The new height of the images to generate.
func SetImageHeight(aHeight int) {
	if 0 < aHeight {
		wkImageHeight = aHeight
	} else {
		wkImageHeight = 0
	}
} // SetImageHeight()

// ImageQuality returns the desired image quality.
func ImageQuality() int {
	return wkImageQuality
} // ImageQuality

// SetImageQuality changes the quality of the genereated image.
// Values supported between `1` and `100`; default is `100`.
//
//	`aQuality` the new desired image quality.
func SetImageQuality(aQuality int) {
	if (0 < aQuality) && (101 > aQuality) {
		wkImageQuality = aQuality
	} else {
		wkImageQuality = 0 // i.e. ignore it
	}
} // SetImageQuality()

// ImageWidth is the width in pixels of the imaginary screen used to render.
// The default value is `1024`.
// Note that this is used only as a guide line.
func ImageWidth() int {
	return wkImageWidth
} // ImageWidth()

// SetImageWidth sets the width of the images to generate.
// The default value is `1024`.
//
//	`aWidth` The new width of the images to generate.
func SetImageWidth(aWidth int) {
	if 0 < aWidth {
		wkImageWidth = aWidth
	} else {
		wkImageWidth = 0
	}
} // SetImageWidth()

// MaxAge returns the maximimum age of cached page images.
func MaxAge() time.Duration {
	return wkImageAge
} // MaxAge()

// SetMaxAge sets the maximimum age of cached page images.
//
// Usually you want this property at its default value (`0`, zero)
// which disables an age check because you want an image of the page
// at the time you linked to it.
//
//	`aLengthInSeconds` is the age a page image can have before
// requesting it again.
//lint:ignore ST1011 - it's a proper argument's name
func SetMaxAge(aLengthInSeconds time.Duration) {
	if 0 < aLengthInSeconds {
		wkImageAge = aLengthInSeconds
	} else {
		wkImageAge = 0
	}
} // SetMaxAge()

// PathFile returns the complete local path/file of `aURL`.
//
// NOTE: This function does not check whether the file for `aURL`
// actually exists in the local filesystem.
//
//	`aURL` The address of the web page to process.
func PathFile(aURL string) string {
	return filepath.Join(wkImageDirectory,
		sanitise(aURL)+`.`+wkImageFileType)
} // PathFile()

/* _EoF_ */
