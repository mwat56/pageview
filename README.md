# pageview

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/mwat56/pageview?status.svg)](https://godoc.org/github.com/mwat56/pageview)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/pageview)](https://goreportcard.com/report/github.com/mwat56/pageview)
[![Issues](https://img.shields.io/github/issues/mwat56/pageview.svg)](https://github.com/mwat56/pageview/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/pageview.svg)](https://github.com/mwat56/pageview/)
[![Tag](https://img.shields.io/github/tag/mwat56/pageview.svg)](https://github.com/mwat56/pageview/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/pageview/blob/master/app/pageview.go)
[![License](https://img.shields.io/github/mwat56/pageview.svg)](https://github.com/mwat56/pageview/blob/master/LICENSE)

## Purpose

Sometimes go don't want just standard web links in your web pages but a preview image showing the page you`re linking to.
That is where this small package comes in.
It generates – by way of calling the [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html) commandline utility – an image of the web page a given URL addresses.
Those image files are stored locally and may be used several times.

	//TODO

## Installation

You can use `Go` to install this package for you:

	go get github.com/mwat56/pageview

## Usage

	//TODO

## Libraries

The great commandline utility

* [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html)

is  **_required_**  for this package to work.

## Licence

        Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                        All rights reserved
                    EMail : <support@mwat.de>

> This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 3 of the License, or (at your option) any later version.
>
> This software is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
>
> You should have received a copy of the GNU General Public License along with this program. If not, see the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) for details.
