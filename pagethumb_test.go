/*
   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
                  All rights reserved
              EMail : <support@mwat.de>
*/

package pagethumb

import (
	"testing"
)

const (
	tmpImageDirectory = "/tmp/"
)

func TestPageThumb(t *testing.T) {
	SetCacheDirectory(tmpImageDirectory)
	u1 := "http://dev.mwat.de/"
	n1 := sanitise(u1) + `.png`
	u2 := "http://www.mwat.de/"
	n2 := sanitise(u2) + `.png`
	u3 := "http://www.mwat.de/index"
	n3 := sanitise(u3) + `.png`
	u4 := "http://matthias.mwat.de/index"
	n4 := sanitise(u4) + `.png`

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
		{" 3", args{u3}, n3, false},
		{" 4", args{u4}, n4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateImage(tt.args.aURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("PageThumb() error = %v,\nwantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PageThumb() = %v,\nwant %v", got, tt.want)
			}
		})
	}
} // TestPageThumb()
