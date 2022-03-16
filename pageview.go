/*
   Copyright © 2019, 2022 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/
package pageview

//lint:file-ignore ST1017 - I prefer Yoda conditions
//lint:file-ignore ST1005 - I prefer Capitalisation

import (
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	// Name of the binary tool.
	wkHTMLconverterBinary = "wkhtmltoimage"
)

var (
	// R/O RegEx to extract a filename's extension.
	wkExtRE = regexp.MustCompile(`(\.\w+)([\?\#].*)?$`)

	// Path/filename of the `wkhtmltoimage` executable.
	wkHTMLToImageBinary = func() string {
		// Check whether we can find the binary:
		if cmd, err := exec.LookPath(wkHTMLconverterBinary); nil == err {
			return cmd
		}
		return ""
	}()

	// Max. age of cached page images (in seconds);
	// `0` (zero) disables the age check.
	wkImageAge time.Duration = 0

	// Directory to store the generated images;
	// defaults to the current path/directory.
	wkImageDirectory = func() string {
		result, _ := filepath.Abs("./")
		return result
	}()

	// Type/Format of the generated images.
	// We use `PNG` format because it scales better.
	wkImageType = `png`

	// Height of the image to generate.
	wkImageHeight = 768

	// Quality of the image to generate.
	wkImageQuality = 100

	// Width of the image to generate.
	wkImageWidth = 1024

	// Flag whether to allow JavaScript in retrieved pages.
	wkJavaScript bool

	// R/O RegEx to find all non alpha/digits in URLs.
	wkReplaceNonAlphasRE = regexp.MustCompile(`\W+`)

	// User Agent to use when queuing external sites.
	wkUserAgent string
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

	if 10240 > fi.Size() {
		// Empty and small (< 10KB) files are ignored.
		// File sizes smaller than ~10KB indicate some kind of
		// error during retrieval of the web page or rendering it.
		// Valid preview images take approximately between 10 to 100 KB
		// depending on the respective web page (e.g. number and size
		// of embedded images).
		return false
	}

	if 0 < wkImageAge {
		maxTime := fi.ModTime().Add(wkImageAge * time.Second)
		// files too old are ignored
		return time.Now().Before(maxTime)
	}

	return true
} // exists()

// `fileExt()` returns the filename extension of `aURL`.
//
//	`aURL` The URL to process.
func fileExt(aURL string) string {
	result := wkExtRE.FindStringSubmatch(aURL)
	if 1 < len(result) {
		return result[1]
	}

	return ""
} // fileExt()

// `sanitise()` returns `aURL` with all non alpha/digits removed.
//
//	`aURL` The URL to sanitise.
func sanitise(aURL string) string {
	return wkReplaceNonAlphasRE.ReplaceAllLiteralString(aURL, ``)
} // sanitise()

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

// CreateImage generates an image of `aURL` and stores it in
// `ImageDirectory()`, returning the file name of the saved image
// or an error in case of problems.
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

	result := sanitise(aURL) + `.` + wkImageType
	fName := filepath.Join(wkImageDirectory, result)
	// Check whether we've already got an image so
	// we might avoid additional network traffic:
	if exists(fName) {
		return result, nil
	}

	var (
		// Declare variables here so we can use them in different
		// contexts/closures below.
		err       error
		imageData []byte
		response  *http.Response
	)

	// Exclude certain filetypes from preview generation:
	ext := strings.ToLower(fileExt(aURL))
	switch ext {
	case ".amr", ".arj", ".avi", ".azw3",
		".bak", ".bibtex", ".bz2",
		".cfg", ".com", ".conf", ".csv",
		".db", ".deb", ".doc", ".docx", ".dia",
		".epub", ".exe", ".flv", ".gz",
		".ics", ".iso", ".jar", ".json",
		".md", ".mobi", ".mp3", ".mp4", ".mpeg",
		".odf", ".odg", ".odp", ".ods", ".odt", ".otf", ".oxt",
		".pas", ".pdf", ".ppd", ".ppt", ".pptx",
		".rip", ".rpm", ".spk", ".sxg", ".sxw",
		".ttf", ".vbox", ".vmdk", ".vcs", ".wav",
		".xls", ".xpi", ".xsl", ".zip":
		return "", errors.New("Excluded filename extension: " + ext)

	case ".gif", ".jpeg", ".jpg", ".png", ".svg":
		if response, err = http.Get(aURL); /* #nosec G107 */ nil != err {
			return "", err
		}
		defer response.Body.Close()
		result = sanitise(aURL) + `.` + ext
		fName = filepath.Join(wkImageDirectory, result)

	default:
		if imageData, err = generateImage(aURL); nil != err {
			// Either `wkhtmltoimage` produced an error
			// or it took too long and was canceled.
			return "", err
		}
	}

	file, err := os.OpenFile(fName,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640) // #nosec G302
	if nil != err {
		return "", err
	}
	defer file.Close()

	if 0 < len(imageData) {
		_, err = file.Write(imageData)
	} else if (nil != response) && (0 < response.ContentLength) {
		_, err = io.Copy(file, response.Body)
	} else {
		_ = os.Remove(fName)
		return "", errors.New("No image data received")
	}

	if nil != err {
		// In case of errors during write we delete the file
		// ignoring possible errors and return the write error.
		_ = os.Remove(fName)
		return "", err
	}

	// Everything went well!
	return result, nil
} // CreateImage()

// ImageDir returns the directory used to store the generated images.
func ImageDir() string {
	return wkImageDirectory
} // ImageDirectory()

// SetImageDir sets the directory to use for storing the
// generated images, returning an error if `aDirectory` can't be used.
//
//	`aDirectory` The directory to store the generated images.
func SetImageDir(aDirectory string) error {
	dir, err := filepath.Abs(aDirectory)
	if nil == err {
		wkImageDirectory = dir
	}

	return err
} // SetImageDir()

// ImageFileType returns the type of the image files to generate.
func ImageFileType() string {
	return wkImageType
} // ImageFileType()

// SetImageType changes the type of the generated images.
// The default type is `png`, the other options are `gif`, `jpg` and `svg`.
// Passing an invalid value in `aType` will result in `png` being
// selected.
//
// NOTE: Depending on how your `wkhtmltoimage` binary was compiled not
// all formats might be supported.
//
//	`aType` is the new desired type of the images to generate.
func SetImageType(aType string) {
	switch aType {
	case `gif`, `jpg`, `png`, `svg`:
		wkImageType = aType
	case `jpeg`:
		wkImageType = `jpg`
	default:
		wkImageType = `png`
	}
} // SetImageType()

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

// SetImageQuality changes the quality of the image to be generated.
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

// JavaScript returns whether to allow JavaScript during page retrieval;
// defaults to `false` for safety and speed reasons.
func JavaScript() bool {
	return wkJavaScript
} // JavaScript()

// SetJavaScript determines whether to allow JavaScript during page
// retrieval or not.
//
//	`doAllow` If `false` (i.e. the default) no JavaScript will be available
// during page retrieval, otherwise (i.e. `true`) it will be activated.
func SetJavaScript(doAllow bool) {
	wkJavaScript = doAllow
} // SetJavaScript()

// MaxAge returns the maximum age of cached page images.
func MaxAge() time.Duration {
	return wkImageAge
} // MaxAge()

// SetMaxAge sets the maximum age of cached page images
// before they might get updated.
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
		sanitise(aURL)+`.`+wkImageType)
} // PathFile()

// UserAgent returns the current `User Agent` setting.
func UserAgent() string {
	return wkUserAgent
} // UserAgent()

// SetUserAgent changes the current `User Agent` setting to `aAgent`.
//
// Note: This only affects the HTTP header send to the remote host.
// Unfortunately, there is no way to change the `navigator.userAgent`
// setting by a commandline argument.
// So sites requesting that value will still see the hardcoded
// `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/534.34 (KHTML, like Gecko) wkhtmltoimage Safari/534.34`.
//
//	`aAgent` The new `User Agent` setting.
func SetUserAgent(aAgent string) {
	wkUserAgent = aAgent
} // SetUserAgent()

/* _EoF_ */
