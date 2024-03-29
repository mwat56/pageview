# PageView

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/mwat56/pageview?status.svg)](https://godoc.org/github.com/mwat56/pageview)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/pageview)](https://goreportcard.com/report/github.com/mwat56/pageview)
[![Issues](https://img.shields.io/github/issues/mwat56/pageview.svg)](https://github.com/mwat56/pageview/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/pageview.svg)](https://github.com/mwat56/pageview/)
[![Tag](https://img.shields.io/github/tag/mwat56/pageview.svg)](https://github.com/mwat56/pageview/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/pageview/blob/main/app/pageview.go)
[![License](https://img.shields.io/github/mwat56/pageview.svg)](https://github.com/mwat56/pageview/blob/main/LICENSE)

- [PageView](#pageview)
	- [Purpose](#purpose)
	- [Installation](#installation)
	- [Usage](#usage)
	- [Libraries](#libraries)
	- [Credits](#credits)
	- [Update](#update)
	- [Licence](#licence)

----

## Purpose

Sometimes yoo don't want just standard web links in your web pages but a preview image showing the page you`re linking to. That is where this small package comes in. It generates – by way of calling the external [wkhtmltoimage](https://wkhtmltopdf.org/index.html) commandline utility – an image of the web page a given URL addresses. Those image files are stored locally and may be used as often as you want.

## Installation

You can use `Go` to install this package for you:

	go get github.com/mwat56/pageview

After that you can `import` it the usual Go way to use the library.

## Usage

There are only two functions you have to worry about:

	// SetImageDir sets the directory to use for storing the
	// generated images, returning an error if `aDirectory` can't be used.
	//
	//	`aDirectory` The directory to store the generated images.
	func SetImageDir(aDirectory string) error { … }

This function must be called before any other one to make sure the generated images end up where you want them to be. The default is the current directory from where the program was run.

To actually create an image you'd call:

	// CreateImage generates an image of `aURL` and stores it in
	// `ImageDir()`, returning the file name of the saved image
	// or an error in case of problems.
	//
	//	`aURL` The address of the web page to process.
	func CreateImage(aURL string) (string, error) { … }

The returned string is the name of the generated image file. If you combine it with the directory returned by `ImageDir()` you get the complete path/filename to locally access the image.

Generating a preview image usually takes between one and five seconds, depending on the actual web-page in question, however, it can take considerably longer. To avoid hanging the program the `wkhtmltoimage` utility is called with an one minute timeout.

And, finally, not all web-pages can be rendered properly and turned into an image. In such case `wkhtmltoimage` just crashes and `CreateImage()` doesn't return a filename but an error.

There are a few more functions which you will barely need; for details refer to the [source code documentation](https://godoc.org/github.com/mwat56/pageview).

## Libraries

The great commandline utility

* [wkhtmltoimage](https://wkhtmltopdf.org/downloads.html)

is  **_required_**  for this package to work.

Under Linux this utility is usually part of your distribution. If not, you can [download wkhtmltoimage](https://wkhtmltopdf.org/downloads.html) from the web and install it. Sometimes the package from the download page above is more recent than the version in your Linux distribution. If in doubt, I'd suggest to test both versions to determine which one works best for you.

## Credits

Part of the code (i.e. the file `wkhtmltoimage.go`) started as a modified version of [go-wkhtmltoimage](https://github.com/ninetwentyfour/go-wkhtmltoimage), `Copyright (c) 2015 ninetwentyfour`

## Update

Over the last few years in most cases [wkhtmltoimage](https://wkhtmltopdf.org/downloads.html) worked just fine. However, once in a while `wkhtmltoimage` produced a `segmentation fault (core dumped)` – reproducible. For a while I thought I could live with it, but over time it happened more often (i.e. with additional URLs). Fiddling around with various commandline options provided no improvement.

In the end I started to look around, searching for alternative approaches – short of writing my own URL retrieval and rendering system. That's when I found [ChromeDP](https://github.com/chromedp/chromedp) and hence the [screenshot](https://github.com/mwat56/screenshot) package came into existence which is supposed to be a replacement for this package.

## Licence

        Copyright © 2019, 2022 M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.

----
