package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_getInputFile(t *testing.T) {
	tests := []struct {
		name    string
		want    inputFile
		wantErr bool
		osArgs  []string
	}{
		{"Default arguments", inputFile{"test.yml"}, false, []string{"cmd", "test.yml"}},
		{"Missing file name", inputFile{}, true, []string{"cmd"}},
	}
	for _, tt := range tests {
		actualOsArgs := os.Args
		defer func() {
			os.Args = actualOsArgs
		}()

		os.Args = tt.osArgs
		t.Run(tt.name, func(t *testing.T) {
			got, err := getInputFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("getInputFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInputFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Valid extension .yaml", args{"test.yaml"}, true, false},
		{"Valid extension .yml", args{"test.yml"}, true, false},
		{"Invalid extension .xyz", args{"test.xyz"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isValidExtension(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("isValidExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isValidExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Non existing file", args{"nonexisting.yml"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileExists(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
