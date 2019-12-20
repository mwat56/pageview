# PageView

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/mwat56/pageview?status.svg)](https://godoc.org/github.com/mwat56/pageview)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/pageview)](https://goreportcard.com/report/github.com/mwat56/pageview)
[![Issues](https://img.shields.io/github/issues/mwat56/pageview.svg)](https://github.com/mwat56/pageview/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/pageview.svg)](https://github.com/mwat56/pageview/)
[![Tag](https://img.shields.io/github/tag/mwat56/pageview.svg)](https://github.com/mwat56/pageview/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/pageview/blob/master/app/pageview.go)
[![License](https://img.shields.io/github/mwat56/pageview.svg)](https://github.com/mwat56/pageview/blob/master/LICENSE)

- [PageView](#pageview)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
	- [Libraries](#libraries)
	- [Credits](#credits)
	- [Licence](#licence)

----

## Purpose

Sometimes yoo don't want just standard web links in your web pages but a preview image showing the page you`re linking to.
That is where this small package comes in.
It generates – by way of calling the external [wkhtmltoimage](https://wkhtmltopdf.org/index.html) commandline utility – an image of the web page a given URL addresses.
Those image files are stored locally and may be used as often as you want.

## Installation

You can use `Go` to install this package for you:

	go get github.com/mwat56/pageview

After that you can `import` it the usual Go way to use the library.

## Usage

There are only two functions you have to worry about:

	// SetCacheDirectory sets the directory to use for storing the
	// generated images returning an error if `aDirectory` can't be used.
	//
	//	`aDirectory` The directory to store the generated images.
	func SetCacheDirectory(aDirectory string) error { … }

This function must be called before any other one to make sure the generated images end up where you want them to be.

To actually create an image you'd call:

	// CreateImage generates an image of `aURL` and stores it in
	// `CacheDirectory` returning the file name of the saved image.
	//
	//	`aURL` The address of the web page to process.
	func CreateImage(aURL string) (string, error) { … }

The returned string is the name of the generated image file.
If you combine it with the directory returned by `CacheDirectory()` you get the complete path/filename to locally access the image.

Generating a preview image usually takes between one and five seconds, depending on the actual web-page in question, however, it can take considerably longer.
To avoid hanging the program the `wkhtmltoimage` utulity is called with an one minute timeout.

And, finally, not all web-pages can be rendered properly and turned into an image.
In such case `wkhtmltoimage` just crashes and `CreateImage()` doesn't return a filename but an error.

There are a few more functions which you will barely need; for details refer to the [source code documentation](https://godoc.org/github.com/mwat56/pageview).

## Libraries

The great commandline utility

* [wkhtmltoimage](https://wkhtmltopdf.org/downloads.html)

is  **_required_**  for this package to work.

Under Linux this utility is usually part of your distribution.
If not, you can [download wkhtmltoimage](https://wkhtmltopdf.org/downloads.html) from the web and install it.
Sometimes the package from the download page above is more recent than the version in your Linux distribution.
If in doubt, I'd suggest to test both versions to determine which one works best for you.

## Credits

Part of the code (i.e. the file `wkhtmltoimage.go`) started as a modified version of [go-wkhtmltoimage](https://github.com/ninetwentyfour/go-wkhtmltoimage), `Copyright (c) 2015 ninetwentyfour`

## Licence

        Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
