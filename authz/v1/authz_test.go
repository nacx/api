// Copyright (c) Tetrate, Inc 2020 All Rights Reserved.

package v1

import "testing"

func TestValidateClusterLocalServiceAccount(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		valid bool
	}{
		{"valid", "foo/bar", true},
		{"empty", "", true},
		{"slash", "/", false},
		{"ns", "foo/", false},
		{"sa", "/foo", false},
		{"word", "foo", false},
		{"more", "foo/bar/more", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Subject{
				TlsIdentity: &Subject_ClusterLocalServiceAccount{ClusterLocalServiceAccount: tt.in},
			}
			err := i.Validate()
			if tt.valid != (err == nil) {
				t.Fatalf("valid=%t, err=%v", tt.valid, err)
			}
		})
	}
}
