/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package pageview

import (
	"testing"
)

func Test_fileExt(t *testing.T) {
	type args struct {
		aURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{" 1", args{""}, ""},
		{" 2", args{"image.gif"}, ".gif"},
		{" 3", args{"document.txt"}, ".txt"},
		{" 4", args{"document.txt.doc"}, ".doc"},
		{" 5", args{"http://example.com/page.html?view=print"}, ".html"},
		{" 6", args{"http://example.com/sometopic?show=all&lang=en"}, ""},
		{" 5", args{"http://example.com/page.md?view=print#top"}, ".md"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExt(tt.args.aURL); got != tt.want {
				t.Errorf("extension() = %v, want %v", got, tt.want)
			}
		})
	}
} // Test_fileExt()

func TestCreateImage(t *testing.T) {
	SetImageDirectory("/tmp/")
	SetMaxAge(60)
	SetImageFileType(`png`)
	u1 := "http://dev.mwat.de/"
	n1 := sanitise(u1) + `.` + wkImageFileType
	u2 := "http://www.mwat.de/"
	n2 := sanitise(u2) + `.` + wkImageFileType
	u3 := "http://www.mwat.de/index.pl"
	u4 := "http://bla.mwat.de/index.shtml"
	n4 := sanitise(u4) + `.` + wkImageFileType

	type args struct {
		aURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{" 1", args{u1}, n1, false},
		{" 2", args{u2}, n2, false},
		{" 3", args{u3}, "", true},
		{" 4", args{u4}, n4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateImage(tt.args.aURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateImage() error = %v,\nwantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateImage() = %v,\nwant %v", got, tt.want)
			}
		})
	}
} // TestCreateImage()
