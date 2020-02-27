// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// api.txt
// api_test.txt
package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x5f\x6f\xda\x30\x10\x7f\x4e\x3e\xc5\x4d\xea\x03\xa9\xaa\xe4\x9d\x25\x48\x0c\x36\x69\x0f\xab\xba\xb2\x3d\x4d\xd3\x64\x92\x03\xa2\x39\x8e\xe7\x5c\x06\x14\xf1\xdd\x27\xdb\x49\x08\xd0\xd0\x52\x5a\xed\x2d\xbe\xdc\xdd\xef\xcf\xf9\xe4\x20\x18\x96\x94\xc3\x1c\x05\x2a\x46\x98\xc0\x32\xa5\x05\x7c\x41\x62\x43\x99\xc2\x82\x48\x16\xfd\x20\x98\xa7\xb4\x28\xa7\x7e\x9c\x67\x01\xae\xd6\x0f\x0f\xeb\x20\x43\x62\x4c\xa6\xae\x64\xf1\x6f\x36\x47\x08\x43\xf0\xef\xaa\xef\xc1\xc0\x75\xf5\xf9\x73\x26\x73\x45\xf5\x51\x31\x31\x47\xb8\x4a\x45\x82\xab\x1b\xb8\x22\x36\xe5\x08\xfd\x08\xfc\x6f\xfa\xab\xd0\x69\x41\x30\x52\xc8\x08\xc1\x84\xdc\x59\x29\x62\xb0\x11\x13\x08\xc3\xaa\xcc\x1f\x31\x79\xcb\x32\x8d\xd4\x4b\xa6\x70\x5d\xfc\xe1\xfe\xf8\x83\x07\x3d\x54\x0a\x50\xa9\x5c\x79\xb0\x71\x9d\x5f\x37\xfa\x00\x11\x24\x53\xff\xe3\x0a\xe3\x5d\xfd\x58\xe5\xd2\xb4\x9c\x10\x23\xcc\x50\x18\x96\x4e\x3a\x33\x05\xef\x22\x10\x29\xd7\x1d\x1c\x85\x54\x2a\xe1\x3a\xdb\x53\xdd\x5a\x14\xf7\xfb\x55\xd5\x5b\xd7\x0d\x02\x0d\xd9\xd6\xd5\x50\x78\x73\x55\x6d\x16\x13\x52\x65\x4c\x2e\xad\xa5\x19\xd9\x0e\x78\x92\x8a\x79\x05\x0e\x85\x49\x82\x8d\xbb\xcb\xb0\x75\x9f\x52\xe4\x89\x99\x94\x69\x66\x75\x5b\x39\xbd\x56\x6e\xab\xd5\x75\x17\x88\x57\x0d\xb6\xeb\xff\xbe\x03\x0a\x8b\x92\x53\x27\x63\xeb\xc8\xce\xa2\x82\x32\xb2\xb1\xbe\xb1\xe9\x4e\xa1\x64\x0a\x0f\x27\x76\xc6\xf0\x13\x9c\xa1\x02\xdd\xd7\x1f\xf1\xbc\xc0\x9e\xe7\x3a\x76\x0a\x26\xf6\xb5\x44\xb5\xbe\xcf\x97\x87\x08\x26\x7e\x34\x84\x7b\x24\x95\xe2\xdf\x17\x3b\x57\xd7\xbf\x89\x77\x55\x7a\xd4\x59\xb0\xd9\xd6\xd2\x93\xe9\x23\xc2\x6b\x72\xdd\x77\xb0\xce\x18\x72\x6e\x1d\x68\x05\x9e\x5e\x86\x5d\x46\xed\xd4\x8f\x9f\xcf\xd6\x96\x2f\x8b\xf6\xbd\x30\xec\x8f\xa9\x0f\x39\x3f\xe3\x6a\xcc\x72\x05\xba\xb1\x7f\x8b\x2b\xea\x79\xd5\x6f\x63\x62\xff\xb4\x8b\x75\xdf\xc8\xd4\xb7\xae\x41\xcc\xc4\x90\x73\x18\x0c\xde\x1f\xe2\x36\xc0\x1a\xd9\x39\xf6\x22\x02\x26\x25\x8a\xe4\xd8\xa6\x1b\xb0\xa4\x3c\x43\xda\x10\x6e\x6e\x72\x7b\x3a\xdf\x65\x72\xc1\x52\xdb\xea\xff\xb3\xd4\x16\xfb\xb2\xa5\x7e\x62\xab\x2d\xc4\xe3\x5b\x3d\x46\x8e\x2f\x37\xce\x56\x3f\xcf\xb8\x33\x7d\xb1\xad\x2f\xf5\xa5\x79\x76\x4c\xfc\xe0\xe1\x31\x08\xa7\x6c\x69\x76\xbd\x39\x76\x49\x2d\x5e\x41\x6b\x7b\x81\x5f\x4b\xed\xde\x9e\x84\x21\xa0\x48\xb4\xd6\x7f\x01\x00\x00\xff\xff\xac\x1f\x93\x53\x40\x09\x00\x00")

func apiTxtBytes() ([]byte, error) {
	return bindataRead(
		_apiTxt,
		"api.txt",
	)
}

func apiTxt() (*asset, error) {
	bytes, err := apiTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.txt", size: 2368, mode: os.FileMode(420), modTime: time.Unix(1582830097, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _api_testTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x5b\x73\xe3\x36\xd2\x7d\x16\x7f\x45\x87\x15\xa7\xc8\x8c\x86\xca\xe4\xfb\xf6\x45\xb1\x54\xe5\xd8\x9a\xca\xec\xc6\x97\xb5\x34\xc9\xd6\x7a\x5d\x53\x10\xd9\x92\xb1\xa6\x00\x1a\x00\x2d\x3b\x2e\xff\xf7\xad\x06\x48\x8a\xb2\xa8\x8b\xe5\xd9\xda\x4a\x1e\x62\x11\x97\xc6\xe9\xc6\xe9\x83\x06\x92\x4e\xe7\x28\x37\x12\xa6\x28\x50\x31\x83\x09\xcc\xb9\xb9\x81\x53\x34\xec\x28\xe3\x70\x63\x4c\xa6\xbb\x9d\xce\x94\x9b\x9b\x7c\x1c\xc5\x72\xd6\xc1\x87\xc7\x3f\xfe\x78\xec\xcc\xd0\x30\x96\x71\x2f\x63\xf1\x2d\x9b\x22\x1c\x1e\x42\x74\x51\xfc\xee\xf7\x3d\x8f\xcf\x32\xa9\x0c\x04\x5e\xcb\x4f\x98\x61\x63\xa6\xb1\xa3\xef\x52\xdf\x6b\xf9\x28\x62\x99\x70\x31\xed\xfc\x5b\x4b\x41\x0d\x93\x99\xa1\x3f\xa9\x9c\xd2\x1f\xa9\xe9\xdf\x0a\x27\x29\xc6\xb6\x5d\x1b\xc5\xc5\xd4\xb6\x1a\xd4\x86\x0b\x3b\xcc\xf0\x19\xfa\x5e\xe8\x79\xf7\x4c\x01\xb5\x9f\x8c\xe1\x7b\x7d\x97\x46\x27\x3f\xdb\xa6\x58\x8a\x09\x9f\x26\x63\x98\xb1\xec\xca\x99\xb8\xe6\xc2\xa0\x9a\xb0\x18\x9f\x9e\xbd\x58\x0a\x6d\x8a\x89\x67\x6c\x86\xd0\x03\x6b\x9e\x1c\xf9\xc8\x53\xbc\x50\x38\xe1\x0f\xd0\xef\xfb\x9e\xd7\xe9\x40\xcf\xfd\x03\x37\x98\x66\xa8\x34\xb5\x31\xad\xf3\x19\x6a\x60\xc5\x5a\xa9\x8c\x59\x9a\x8c\x23\x72\x0b\x26\x3c\x45\x60\xba\xeb\x75\x3a\x4f\x34\x1f\x00\xfc\x5f\xa4\x36\x7e\x17\x7c\x3b\xf0\x86\x3e\xda\x65\xd7\x85\x54\xb6\xeb\x2f\xff\xff\x7f\x3f\x2e\x5a\x3f\x6b\x54\xd4\x9a\x8c\x05\x9b\x61\x6d\x34\xd3\xba\xa9\x9d\xfc\x68\x6a\x1f\x0e\x7f\x3d\x95\x89\xeb\xe2\x9a\x8d\x53\xf4\xbd\x4e\xe7\xd9\x9b\xe4\x22\x86\x54\xb2\xe4\xd8\x3a\x10\x84\xf0\xe4\xb5\x26\x33\x13\x5d\x28\x2e\x4c\x2a\x02\x1f\x6a\xdd\x7e\xe8\xb5\xc8\xad\x36\xa0\x52\xd0\xed\x81\xd4\xd1\x79\x86\x22\xf0\x57\xfd\xa7\xb1\x7c\x62\x07\x7e\xd3\x03\xc1\x53\xb2\xdc\x4a\xe5\x34\xba\x60\x82\xc7\x64\xfa\x98\x09\x21\x0d\xc8\x0c\xc5\x72\x00\x6d\xec\x7c\xbb\x4a\x34\x50\x4a\xaa\x20\x0c\xbd\xd6\xb3\xd7\x4a\x30\x96\x09\xda\xa5\x69\x8d\xe8\x0c\xe7\x27\xae\x29\xa0\x29\xa1\xd7\xa2\xf5\x7a\x50\x8c\x8b\x5c\x67\xf0\x5d\x49\x85\x9d\x31\x4d\xd1\x80\x05\x53\x00\xcb\x15\x33\x5c\x0a\x82\xa6\xe4\x6c\x1d\xbe\x67\xcf\x05\x34\x56\xc8\x0c\x9e\x8c\x83\xa4\x62\x64\x1b\x12\x47\x32\x47\xc4\x36\xc8\xb9\x40\x55\x7c\x85\x10\x10\x26\x24\x53\x76\x07\xb4\x26\x17\x69\x1f\x86\x19\x6d\xc4\x24\xf0\x8f\x2f\x07\x47\xa3\x01\x9c\x1c\x8d\x8e\x7e\x3e\x1a\x0e\xe0\x40\xc3\xf9\xef\x67\x83\x4b\x38\xd0\x7e\x69\xbc\xb0\x1a\xae\x6c\xa1\x0f\xef\x40\xeb\xd0\x6b\x7d\x71\x5b\xd7\x83\x64\x1c\x0d\x1e\x30\x0e\x6c\xab\x42\x93\x2b\x51\xc1\xd7\x68\x46\x7f\x2c\xa1\xdf\x05\xe0\x70\x30\x82\xd1\xa7\xd3\x01\xfc\xf3\xfc\x6c\x00\x9f\x47\xc7\xfe\x5b\x71\x24\x4a\x66\x9b\x82\xb8\x0b\xaa\x93\xcb\xf3\x8b\x7a\xd0\xaa\x58\xbd\x15\x9c\x92\xf3\xc1\x03\xd7\x46\x2f\xe1\xbb\xcb\x51\x3d\x56\x7b\xcc\xd4\x54\x43\x14\x45\x35\xd5\x21\xcc\x76\x1a\x8c\xa5\x4c\xdd\x42\x0b\x07\xdc\xf4\x95\xc0\xfe\x3a\x38\x1e\xc1\xe0\x1f\x9f\x86\xa3\x21\x04\x07\x3a\xf4\x8b\x85\x9a\x7d\x28\xbb\x2a\x1f\xfe\x4e\x0d\x97\x72\x1e\xd8\x1e\x07\x2b\x8a\xa2\x30\x1a\xc6\x4c\x04\xdf\x39\x3c\xab\x1e\x1a\x92\x89\x06\x1f\x6d\xfb\x62\x0b\xee\x59\xca\x93\x26\x6f\xbc\x96\xed\x2a\x83\x59\x0f\x58\x1b\x4a\xaf\x3e\xc0\xc7\xcb\xf3\x53\xc8\xa6\x5f\xac\x59\x0d\xbf\xff\x32\xb8\x1c\xb8\x35\x84\x13\xe5\x6f\x3f\xf8\xc5\xa2\xab\x18\xb9\xe0\x66\x64\x15\x3c\x58\x61\x43\x5d\xd6\xbc\x56\xa6\xef\xd2\x4f\x62\x22\x57\x18\x42\x32\xdc\x3b\xd0\x40\xa7\x15\xfd\xcd\x35\x2a\xfb\xcd\xb4\x9e\x4b\x95\xd0\x6f\xff\x9d\xd7\x6a\xf9\x5a\xa7\x33\x99\x60\xcf\xb2\xa8\x94\x94\x2b\x27\xea\xd7\xf5\x16\xab\xe5\x4b\x2d\x56\xc7\x97\xc7\x90\x82\x2f\xb5\x94\x1a\x7d\x1d\x7a\x2d\x77\x2a\x95\xa1\xa3\xc8\x3b\x9d\xcd\xa4\x36\x53\x85\x04\xa0\x74\x68\x35\x28\x14\x90\x53\xc6\x45\x30\x83\xef\x8b\xf3\x32\x3a\xb5\x21\xe9\x74\xe8\x9b\x12\x3c\xcf\x1c\x41\xba\xbd\xa5\x18\x6e\xd5\xc7\xd8\xe9\xe3\x62\x0e\x34\xe8\x60\xc9\xbd\x4a\x07\x4b\x77\x16\x87\xed\x6a\x70\xa2\xa0\xa0\xd4\xce\x20\x8e\x0b\xf3\x9b\x20\x94\x5a\xe6\x56\xde\xd9\x74\x31\xad\xd9\x72\xa7\xa3\x72\x61\x7d\xd1\x5e\x0b\x1f\xb8\xf9\x8d\xa5\x14\xc8\x59\x74\x99\x8b\x20\xf4\xaa\x38\x1b\x64\x2a\x91\x73\x51\xe5\xa2\x13\xb4\xd5\x68\xec\x8c\xeb\xc4\x5a\x68\x82\xd5\x92\x3a\x1a\x3c\x70\x13\x14\x80\x42\xa2\x83\x79\xcc\x10\x62\x39\xcb\x98\xc2\x11\xfd\x26\x7e\x04\x35\x39\x6a\xc3\x92\x36\x51\x16\x17\x24\x12\xf2\xd8\xcd\x0b\x14\xea\x3c\x35\x6d\xc0\x87\x0c\x63\xb3\x3a\x61\xa9\x64\x98\x04\x7e\x35\xb3\x0b\x07\xf7\x6d\x38\xb8\x87\xf7\x00\x07\xa3\x36\x1c\x8c\xe0\x5f\xc2\x6f\xc3\xb2\xc1\x97\xdf\x15\x9f\x21\x30\x2a\xc7\x70\x71\x1c\xe0\x84\xe5\xa9\xd9\x17\xd6\xf2\xf4\x37\x63\x73\x1d\xd0\xeb\x55\x5d\x25\x4e\xaa\x4c\xf6\x05\x59\x9b\xbb\x2f\x42\x62\x5f\x51\x80\x62\x9a\x68\x88\x99\x80\x31\x02\x13\x8f\x20\x15\x15\x4f\x6c\x62\x50\x41\x32\x06\xe7\x4a\x1b\xb4\x04\x85\x2c\x21\x74\xb2\xb1\x58\x06\x26\x12\x48\xa5\xbc\x05\x92\x0d\xb7\xe0\x29\xcb\x2c\xe7\xd9\x2d\x06\xcd\x05\x76\x58\x29\x8c\xad\xd5\x3e\x8b\x19\x53\xfa\x86\xa5\xc1\xd5\xf5\xf8\xd1\x94\xb1\x59\xa4\x7d\x1b\xbe\xab\x4c\x6f\x4e\x08\x12\x7a\x47\x7a\xe7\xf4\xce\x58\xd6\x41\x71\x66\x96\xa0\x54\x96\x77\x84\xe2\xb5\x26\x52\xc1\x6d\x1b\xee\x09\x8a\x62\x62\x8a\xb0\x40\x47\x73\xf8\x04\xee\xc9\x46\xe5\xe4\xd5\xed\xb5\xed\x58\x62\xc0\xdf\xf0\xd1\xed\xfc\xa5\x1d\xe6\x7e\x0f\xac\x21\xfa\xed\xb7\x69\x8d\xba\x89\x36\xdc\x87\x64\xa4\x20\xe6\x84\xa5\x1a\xbd\x16\x05\xe7\xb9\x62\x2b\x25\xd2\xa2\xbc\xb3\x4e\x9e\x49\x33\xb8\xcb\x59\xfa\x92\xa2\x2e\x22\x15\x3b\x17\x84\x2f\x2e\x61\xd1\x48\xf1\xd9\x30\x63\x31\x16\x2d\x85\x81\x30\x24\xdf\xd6\x0e\x2a\xd8\x19\x86\xde\x0b\x1c\xaf\xc8\x94\x92\x7a\x7f\xd5\x52\x54\x57\x10\xbb\xa1\xa7\xc5\x76\x16\x50\x5e\x45\x9e\xf5\xd6\xaa\x9c\xdf\xc9\x5a\x7d\x13\x97\x5c\x5b\x97\xc8\x4b\xe1\x23\x14\x61\xd5\xb6\x40\x16\x6e\x90\xa0\x2d\x3b\x62\xa7\x87\xa4\x50\x5b\x76\xc5\x0d\xb4\x02\xd6\xe9\x70\x63\x9f\x01\xc0\xdc\x28\x99\x4f\x6f\x00\x59\x7c\xe3\xa4\x04\xe4\x84\x2c\xe5\xb1\xb1\x7a\xc0\xb2\x2c\x7d\x04\x73\x53\x1d\x30\xf6\x70\xa1\x5b\x12\x18\x59\x9f\x46\x37\xff\x04\xa4\x58\x3a\x88\x66\x2c\x73\x2c\x40\x22\xe1\x47\x1a\xb8\x81\x02\x6d\x3b\x97\x12\xa9\x96\xde\x35\x73\xa1\x2b\xfb\x2c\x47\x72\x9b\x7f\xee\xf1\x20\xfa\x8d\xa5\x39\x9e\x4f\x16\x5b\x79\xdf\xd4\x5b\xd1\x86\xce\xcb\xf3\xc9\x90\xc6\xe4\x11\x19\xb6\xa7\x39\x25\x36\xa7\xb6\x1f\x7e\x02\x0e\x87\x90\x47\x67\xf9\xcc\x41\x0e\x7f\x02\xfe\xee\x9d\x5d\x97\x58\xf2\x4d\x50\xe0\xbc\x2a\x2c\x45\x6e\x18\x0f\x23\x3a\xe6\xaf\xc3\xe0\x7e\xd1\xf2\xa9\xf4\x2f\x08\xdb\x90\x37\xb6\x87\x4e\x1f\xca\xd4\x9e\x19\x77\xe4\x4f\x02\xdf\x8e\xee\xc2\x81\xde\x20\x14\x8d\x18\xda\xf0\x4a\x08\x2f\xb5\x44\xf0\x94\x98\xe2\x75\x3a\xc5\x45\x20\xc3\x98\x4f\x78\x0c\x9e\x77\x78\x58\x08\xdf\xb7\x5c\x24\xf8\xd0\x86\x6f\xdd\x90\x6e\x0f\xa2\x91\x2b\xef\xfb\xfd\xe2\xdd\xe5\xf0\xb0\xe8\xb5\xa8\xa0\xdf\xb7\x1f\xe5\x3b\xcc\x4a\xaf\x7d\x83\xb1\x45\x55\xc2\x0c\x83\xf7\x20\xa4\xc1\x2e\x24\x32\xa7\x05\xc6\x8a\xc5\xb7\x68\x34\x70\x57\x97\xb9\x41\x02\x31\x01\x4d\x7c\x87\x31\x9a\x39\xa2\x00\x69\x6e\x50\xcd\xb9\x46\x20\xc2\x5a\x86\x65\x0a\x0d\x26\xc0\x34\x18\x9c\x65\x29\x71\x9f\x59\x1a\x57\xef\x4a\x0b\x30\xc7\x2c\x1b\x92\x76\x3a\x4c\xd0\x83\xab\x1f\xaf\xd7\xf5\x3e\xd5\x5c\xb4\xa5\x32\x41\xfa\x00\xfd\x7e\xbb\xa9\xe3\x47\xb2\xf7\xec\xde\xb2\xf2\x2c\x61\x06\x37\xac\xba\xae\x6b\xcd\x82\x14\xba\x97\x39\xaa\x8b\x27\x32\xdb\xba\x64\x6f\xb1\x4c\x73\xae\x3d\x79\xb5\xe1\xae\xf9\x94\x65\x1f\x5d\xb1\xd1\xef\x3b\x15\xa9\x5e\xcb\x6c\x95\xdc\x5d\x8b\xb8\xbc\x4c\xe3\x3d\x2a\xdd\x88\x23\x58\xa1\x02\x5c\xad\x8d\x79\x58\xd5\x65\x9b\xc6\x3c\xd5\x73\x3a\x45\xb1\xba\x44\x08\xef\xe1\x03\x25\x7b\xdf\x25\xfd\xfb\xf7\x36\x11\xcb\x9a\x8f\xb4\x0f\xc5\x42\xb0\x56\xe6\x5f\xf1\xeb\xb0\x96\x33\x4b\x37\x33\x77\x6b\xb1\xe9\xd0\xe8\xaf\x59\x5c\xda\x46\xab\xcf\x71\xbd\xde\x96\xf9\x7e\xe8\x55\x75\xd7\xd6\xa5\xd6\x5f\x8a\x16\x4a\xb3\x74\xd9\x5a\x6b\xcb\xbe\x3c\xbc\xb8\x93\x00\xa6\x1a\xad\xb5\x17\xaf\x14\x27\x52\x60\x77\x9b\x45\xbf\x3c\xa4\xb9\x36\xba\x3a\xa0\xeb\x8f\x12\xe5\x3d\x6a\xbd\x94\xec\xe6\x5a\xcd\x66\x93\x1b\xd6\xc6\x37\xc5\xab\xcd\xb2\x81\x3a\x9a\x06\x12\xb9\xdd\xc7\x04\xe6\x4a\x8a\x29\x68\xc3\x4c\xae\x21\x96\x09\x76\x61\x2a\x0d\xd5\x04\x73\x26\x8c\x53\xeb\xd2\x51\x77\xe9\xd9\x1a\xbd\xda\xd2\xfe\xd2\xe3\xe3\x82\x63\xeb\x32\x60\x47\x8a\xad\x9b\x6e\x19\x56\x9d\xd5\xc5\xb6\x6c\x90\xc9\xab\x1f\xae\xa3\x6d\x80\x5e\x4d\xc4\xb5\xea\xb8\x2f\x0f\x37\x78\x5b\x7b\x4e\x68\xa8\x56\xb6\xb8\xde\xde\x24\xb2\x5b\x3c\x66\x19\x5f\x70\x28\x17\xae\x88\xc1\xa4\xa8\x08\xa3\x46\xb2\xd6\x69\x70\x89\x46\x71\xbc\xdf\x9f\x08\xdb\x0c\xec\x43\x85\xed\xa0\x5e\x41\x86\x6d\xc6\xf6\xa0\xc3\x0e\x3e\x5b\x5d\xfa\x13\xf3\xe1\x28\x4d\xf7\x3a\x7a\xb6\xcc\xb7\x6c\xf8\xb2\x1b\x11\x3e\xfc\xf9\x35\xe1\x25\xf1\xb7\x46\x77\x0f\x5e\xaf\xb1\xf5\x06\x5a\xaf\xdf\xbb\xf2\x6d\xd3\xd5\x62\x30\xc6\x98\xe5\x54\x28\x67\x1c\xb8\x86\x93\xc1\xf0\xb8\x0d\x57\xdd\x6b\xfa\xd0\x29\x8f\x91\xee\x83\x2c\x4d\x81\x29\xc5\x1e\x01\x53\x9c\xa1\x70\x2f\xa2\xf6\xfa\x66\xef\x59\xeb\xcb\xba\x8d\xec\xe8\x52\xf9\x64\x4b\xb4\x36\x7c\x79\xf9\x9e\x62\x7d\x5c\x93\x7e\x57\xfc\xba\xbc\x3f\xda\x9f\x9b\x73\xad\x61\x27\xde\x94\x6d\x36\x82\xcf\xf6\x85\x63\x91\x74\x9f\x37\x56\xf3\x5b\x73\x6e\xf3\xf4\x26\x01\xde\x7c\x7d\x88\xb6\xe1\x79\x05\x4b\x37\x9b\xda\x83\xa4\x5b\x9d\xdd\xa8\xbc\x9b\x1d\xff\x5f\x0a\xef\x09\xa6\xf8\x06\x0e\x6c\x9e\x5e\xaf\xf8\xb7\x1d\xbe\xdb\x80\xbc\x62\xf3\x37\x9b\xda\x63\xf3\xb7\x7a\x49\x9b\x5f\xfd\xb7\xd9\xaf\x5e\x66\xf4\x5e\xbf\xf1\xae\x8a\xbf\x94\x73\x57\xc6\x9f\xc9\x4b\x39\xf7\x97\x3d\x5d\x18\xd7\x77\x29\x99\xb5\x83\xf4\x8b\x77\xdf\xd7\xd4\x1f\xee\xff\xcf\x11\xb2\xc0\xe0\x5b\xdd\x59\x2c\xf8\x55\x6a\x24\xf7\xe2\xb3\xca\xe0\x7d\x0b\x87\x8d\xb3\xeb\xfc\xdd\xb2\xcc\xab\xe9\xf9\xf5\xce\xcf\x6d\x2e\xbc\xa1\x20\xf8\x6f\x55\x04\xd6\x6e\x8a\xa2\x7c\xdf\x84\x3e\xfc\xf0\xb5\xf9\xbd\x4f\xad\xd1\xc0\x60\x22\xdb\xe1\x21\xa0\x48\xa0\xdf\xff\x4f\x00\x00\x00\xff\xff\x05\xc3\x74\x78\x92\x26\x00\x00")

func api_testTxtBytes() ([]byte, error) {
	return bindataRead(
		_api_testTxt,
		"api_test.txt",
	)
}

func api_testTxt() (*asset, error) {
	bytes, err := api_testTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api_test.txt", size: 9874, mode: os.FileMode(420), modTime: time.Unix(1582830097, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"api.txt":      apiTxt,
	"api_test.txt": api_testTxt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"api.txt":      &bintree{apiTxt, map[string]*bintree{}},
	"api_test.txt": &bintree{api_testTxt, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
