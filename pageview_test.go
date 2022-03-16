/*
   Copyright © 2019, 2022 M.Watermann, 10247 Berlin, Germany
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
		{" 6", args{"https://github.com/mwat56/Nele/blob/master/README.md#nele-blog"}, ".md"},
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
	SetImageDir("/tmp/")
	SetMaxAge(60)
	SetImageType(`png`)
	//
	u1 := "http://dev.mwat.de/dw/"
	n1 := sanitise(u1) + `.` + wkImageType
	//
	u2 := "http://www.mwat.de/"
	// n2 := sanitise(u2) + `.` + wkImageType
	n2 := ""
	//
	u3 := "http://www.mwat.de/index.pl"
	n3 := sanitise(u3) + `.` + wkImageType
	//
	u4 := "http://bla.mwat.de/index.shtml"
	// n4 := sanitise(u4) + `.` + wkImageFileType
	n4 := ""
	//
	u5 := "http://www.mwat.de/Self/"
	// n5 := sanitise(u5) + `.` + wkImageType
	n5 := ""

	tests := []struct {
		name    string
		aURL    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{" 1", u1, n1, false},
		{" 2", u2, n2, true},
		{" 3", u3, n3, false},
		{" 4", u4, n4, true},
		{" 5", u5, n5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateImage(tt.aURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateImage() error = »%v«,\nwantErr »%v«", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateImage() = %v,\nwant %v", got, tt.want)
			}
		})
	}
} // TestCreateImage()
