# PageThumb

[![golang](https://img.shields.io/badge/Language-Go-green.svg)](https://golang.org)
[![GoDoc](https://godoc.org/github.com/mwat56/pagethumb?status.svg)](https://godoc.org/github.com/mwat56/pagethumb)
[![Go Report](https://goreportcard.com/badge/github.com/mwat56/pagethumb)](https://goreportcard.com/report/github.com/mwat56/pagethumb)
[![Issues](https://img.shields.io/github/issues/mwat56/pagethumb.svg)](https://github.com/mwat56/pagethumb/issues?q=is%3Aopen+is%3Aissue)
[![Size](https://img.shields.io/github/repo-size/mwat56/pagethumb.svg)](https://github.com/mwat56/pagethumb/)
[![Tag](https://img.shields.io/github/tag/mwat56/pagethumb.svg)](https://github.com/mwat56/pagethumb/tags)
[![View examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg)](https://github.com/mwat56/pagethumb/blob/master/cmd/pagethumb.go)
[![License](https://img.shields.io/github/mwat56/pagethumb.svg)](https://github.com/mwat56/pagethumb/blob/master/LICENSE)

## Purpose

Sometimes go don't want just standard web links in your web pages but a preview image showing the page you`re linking to.
That is where this small package comes in.
It generates – by way of calling the [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html) commandline utility – an image of the web page a given URL addresses.
Those image files are stored locally and may be used several times.

	//TODO

## Installation

You can use `Go` to install this package for you:

	go get github.com/mwat56/pagethumb

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
