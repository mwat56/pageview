/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package pagethumb

//lint:file-ignore ST1017 - I prefer Yoda conditions

import (
	"encoding/base64"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	h2i "github.com/ninetwentyfour/go-wkhtmltoimage"
)

const (
	// Type/Format of the generated thumbnails
	wkImageFileType string = `png`
)

var (
	// Path/filename of the `wkhtmltoimage` executable.
	wkHTMLToImageBinary = ""

	// Max. age of cached thumbnail images (in seconds);
	// `0` (zero) disables the age check.
	wkPageThumbAge time.Duration = 0

	// Directory to store the generated thumbnails
	wkPageThumbDirectory = ""
)

// `exists()` returns whether there is a usable file cached.
//
// This function uses the global `wkPageThumbAge` value to determine
// whether an already existing file is considered to be too old.
func exists(aFilename string) bool {
	fi, err := os.Stat(aFilename)
	if (nil != err) || fi.IsDir() {
		return false
	}

	if 0 >= fi.Size() {
		// empty files are ignored
		return false
	}

	if 0 < wkPageThumbAge {
		maxTime := fi.ModTime().Add(wkPageThumbAge * time.Second)
		return maxTime.Before(time.Now())
	}

	return true
} // exists()

// `init()` trieds to set the binary's path.
func init() {
	if cmd, err := exec.LookPath("wkhtmltoimage"); nil == err {
		wkHTMLToImageBinary = cmd
	}
	wkPageThumbDirectory, _ = filepath.Abs("./")
} // init()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// Image generates an image of `aURL` and stores it in
// `CacheDirectory` returning the file name of the saved image.
//
//	`aURL` The address of the web page to receive.
func Image(aURL string) (string, error) {
	if 0 == len(wkHTMLToImageBinary) {
		// we can't do anything without the executable
		return "", &exec.Error{
			Name: "wkhtmltoimage",
			Err:  exec.ErrNotFound,
		}
	}

	result := ThumbCode(aURL) + `.` + wkImageFileType
	fName := filepath.Join(wkPageThumbDirectory, result)
	// Look whether we've already got a thumbnail
	if exists(fName) {
		return result, nil
	}

	c := h2i.ImageOptions{
		BinaryPath: wkHTMLToImageBinary,
		Format:     wkImageFileType,
		Height:     742, // empirically determined …
		Input:      aURL,
		Quality:    100,
		Width:      1024,
	}
	imageData, _ := h2i.GenerateImage(&c)

	file, err := os.OpenFile(fName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640) // #nosec G302
	if nil != err {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(imageData); nil != err {
		return "", err
	}

	return result, nil
} // Image()

// MaxAge returns the maximimum age of cached page images.
func MaxAge() time.Duration {
	return wkPageThumbAge
} // MaxAge()

// SetMaxAge sets the maximimum age of cached page images.
//
// Usually you want this property at its default value (`0`, zero)
// which disables an age check because you want an image of the page
// at the time you linked to it
//
//	`aLengthInSeconds` is the age a page thumb can have before
// requesting it again.
func SetMaxAge(aLengthInSeconds time.Duration) {
	if 0 < aLengthInSeconds {
		wkPageThumbAge = aLengthInSeconds
	} else {
		wkPageThumbAge = 0
	}
} // SetMaxAge()

// CacheDirectory returns the directory to store the generated thumbnails.
func CacheDirectory() string {
	return wkPageThumbDirectory
} // CacheDirectory()

// SetCacheDirectory sets the directory to use for storing the
// generated thumbnails returning an error if `aDirectory` can't be used.
//
//	`aDirectory` The directory to store the generated thumbnails.
func SetCacheDirectory(aDirectory string) error {
	dir, err := filepath.Abs(aDirectory)
	if nil == err {
		wkPageThumbDirectory = dir
	}

	return err
} // SetCacheDirectory()

// ThumbCode returns the encoded `aURL`.
//
//	`aURL` The URL to encode.
func ThumbCode(aURL string) string {
	return base64.StdEncoding.EncodeToString([]byte(aURL))
} // ThumbCode()

// PathFile returns the complete local path/file of `aURL`.
func PathFile(aURL string) string {
	fName := ThumbCode(aURL) + `.` + wkImageFileType
	return filepath.Join(wkPageThumbDirectory, fName)
} // PathFile()

/* _EoF_ */
