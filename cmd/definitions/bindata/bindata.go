// Code generated by go-bindata. DO NOT EDIT.
// sources:
// definitions/features.toml (1.578kB)
// definitions/fields.toml (714B)
// definitions/info_object_meta.toml (1.112kB)
// definitions/info_storage_meta.toml (1.139kB)
// definitions/operations.toml (9.591kB)
// definitions/pairs.toml (1.616kB)

// +build tools

package bindata

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _definitionsFeaturesToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8b\x13\x41\x10\x85\xef\xf3\x2b\x1e\x7b\x51\x21\x09\x7a\x11\x15\x3c\x8b\x37\x41\x6f\x8b\x84\xca\x74\xc5\x29\xd3\xd3\x35\x54\xd7\x4c\xd8\x7f\x2f\x3d\x93\xd9\x0d\x21\x13\xc4\xdd\x83\xb7\x24\xdd\xf5\xea\xbd\xf7\x35\xb9\x8f\xaa\x99\xb7\x1d\x89\xfd\xac\x02\xe7\xda\xa4\x73\xd1\x84\xcf\xb8\xbb\xbb\xab\x9e\x4e\xb1\x67\xf2\xde\x18\x92\x11\x38\xcb\xaf\xc4\x01\x7b\x35\xf4\x99\x2d\xe3\xd8\x28\x82\xa6\x57\x8e\x23\x25\x47\x76\x93\xda\x31\x0e\xd6\x0d\xd7\x87\xbc\xa9\xaa\xaf\x7b\x78\x23\xf9\x5c\x89\x13\xed\x22\x87\x15\xbc\x61\x64\xb6\x41\x6a\xc6\x51\x62\x44\x52\x87\xb1\xf7\x96\x40\x09\x6c\xa6\x36\xae\x2b\xbf\xe7\xbe\xeb\xd4\x26\xf9\x22\xfc\xe3\x5c\xf5\x48\x19\x92\xdc\x34\xf4\x35\x07\x48\xc2\x97\xef\xdf\xd6\xef\xde\x7e\xdc\x54\x25\x51\x75\x3f\x88\x79\x4f\x71\x1b\xae\x47\x3e\x3b\x5e\xcc\x4c\x8f\x5e\xbd\x21\x47\x50\xce\x25\x7a\x43\x03\x23\x91\xcb\xc0\x28\xe3\xb3\xcf\x5d\x3f\xd5\x92\xe1\x8a\xce\x74\x90\xc0\xc8\xd2\xf6\x91\x9c\x03\xb4\x63\xa3\xe2\xa0\x64\x59\xe3\x4a\x4d\x41\xf2\xd8\x13\x5e\x97\x9e\x02\xef\xa9\x8f\x8e\x1d\x37\x34\x88\xda\x9b\x2b\xf5\x8d\x67\x8c\x28\x07\x86\x5c\x38\xa4\xf4\x70\x6e\x6f\x73\x7d\xe7\x32\x9a\x39\xd6\x53\x82\xa2\x36\xbb\x29\x8d\xd7\xc6\xe4\x5c\x2a\x5c\x9d\x3e\xaf\x10\x25\xfb\x0a\x81\x23\x97\x6f\x94\x02\xb2\x42\xd3\x3f\xe3\xd3\xdd\x6f\xae\x7d\xdb\xb2\x53\x20\xa7\x5b\x28\x2f\xae\x3e\x07\xeb\x24\x85\x47\xa9\xff\x18\xf1\x82\xd5\x17\xc0\x7d\xa9\x7c\x8e\xbe\x8c\x4f\xa8\xb3\x93\xbf\x04\xe8\x28\xe9\x70\x8b\x6e\x39\x7f\x0e\xd2\x39\x5f\xb9\x5e\xb4\x96\x08\x2d\x37\x64\x7d\x42\xad\x6d\x47\x2e\xbb\xc8\x68\x35\xf0\xa7\xd3\xbb\x1f\x15\x31\x08\xcd\xdb\x5a\xf6\x46\x43\x5e\x8d\x2f\x86\x62\xd4\x23\x8c\x29\x4c\xf7\xf6\xa6\x2d\x34\x86\x75\xf6\x87\x78\x9a\x9d\xca\x5e\xa0\x56\xfe\x0d\x6f\xfb\x92\x34\x6f\xa6\x0c\xf5\x86\x6d\xbe\xf3\x77\x40\x3e\xbc\x3f\xf1\xf8\x13\x00\x00\xff\xff\x02\x8a\xc0\x31\x2a\x06\x00\x00")

func definitionsFeaturesTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsFeaturesToml,
		"definitions/features.toml",
	)
}

func definitionsFeaturesToml() (*asset, error) {
	bytes, err := definitionsFeaturesTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/features.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe8, 0xde, 0x64, 0x57, 0x6f, 0xe4, 0xa1, 0x28, 0x9a, 0xdf, 0xf5, 0xa, 0x78, 0xe3, 0xe7, 0xe5, 0x4c, 0x35, 0xce, 0xe9, 0xb8, 0x4b, 0x32, 0xb7, 0xb2, 0xdc, 0x74, 0xc5, 0x98, 0xd5, 0x9f, 0x28}}
	return a, nil
}

var _definitionsFieldsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xb1\x4e\xc3\x30\x14\xdc\xf3\x15\x55\xc7\x0e\x9e\x10\x1b\x0b\x62\x61\x40\x54\x65\x60\x88\x3a\xbc\x24\xd7\xf4\x41\x13\xbb\xcf\x2f\xa2\xf0\xf5\x28\x09\x95\x5f\xa8\xd7\xbb\xf3\xdd\xe9\x9e\xcb\x8a\xf7\x85\x7e\x07\xac\x1e\x56\xeb\xcd\xe3\xc9\xd7\x9f\xcf\x0a\x21\xf5\xb2\x2e\x8a\xb2\xe2\x26\xd1\x51\x85\xfb\xf6\x0f\x8e\x09\x2f\xf7\x89\x69\xa2\xe6\x1e\x40\x24\xc1\x10\x99\xdd\x71\x09\x2c\x48\x84\x72\x07\xf7\x34\x08\x29\xfb\x7e\x14\x70\xdf\xe0\x92\x78\xee\x75\x44\x3b\x28\x99\xd2\x6f\xea\x85\x5a\xbc\x40\x69\x64\xfb\x85\xfe\xfe\x6e\xc2\xa8\x43\xae\x96\x37\x36\xaf\xd5\x07\xea\xc9\xdf\x1f\x0e\x11\x9a\xb1\xf1\x7c\xa3\xb7\x63\xf9\x90\xcb\x08\xc4\x62\xc6\x72\xce\x6d\x89\x65\x66\xc4\xa4\x6c\xb6\x24\x7a\x85\x17\xeb\x1a\x46\x8f\xd9\x08\xfe\x67\x63\x5b\x99\xe1\xd9\xbb\x1d\xa8\xc1\x8c\xe3\x6c\x5e\x1d\x55\x83\xdb\xe1\x3c\x20\x4e\x51\x91\x7f\x90\x59\x20\x4a\x9d\x2b\x10\x95\x6f\x0f\x22\xb6\x45\x54\x6f\x2f\x7d\x95\x8c\xd4\x20\xa7\x9c\xe7\xd7\xa2\xf7\xbb\xb0\xce\x72\x25\x69\x91\xf9\x64\xbf\x01\x00\x00\xff\xff\x5c\xa2\x94\x99\xca\x02\x00\x00")

func definitionsFieldsTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsFieldsToml,
		"definitions/fields.toml",
	)
}

func definitionsFieldsToml() (*asset, error) {
	bytes, err := definitionsFieldsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/fields.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfe, 0xa5, 0x29, 0x70, 0xa5, 0x18, 0x85, 0x7b, 0xfe, 0x73, 0xbf, 0x8c, 0xef, 0x61, 0xf6, 0xa1, 0xa2, 0x5, 0x4c, 0xf, 0xe9, 0x27, 0xee, 0x26, 0x1b, 0x41, 0xe8, 0x33, 0x87, 0xaa, 0x98, 0x7f}}
	return a, nil
}

var _definitionsInfo_object_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xbd\x6e\xdb\x4c\x10\xec\xef\x29\x16\x6a\xdc\x7c\x54\xf5\x25\x9d\x8b\x00\x2a\x12\xc0\x82\x02\xd8\x42\x0a\xc1\xc5\x89\xb7\xa4\x36\xe2\xfd\x64\x77\x69\x9b\x08\xf2\xee\xc1\x1d\xc5\x48\x96\x98\x22\x1d\x6f\x7f\x67\x67\x86\x3b\x9b\x12\x06\x57\xc5\xa6\x11\xd4\x67\xa3\x43\x42\xb8\x87\x05\x05\xfd\xf8\xff\xc2\x38\x94\x9a\x29\x29\xc5\x90\xa3\x9f\x4a\xf1\xa6\xd4\x02\x09\xe8\x01\x61\xec\x84\xd8\x94\xd7\x38\x0e\xe2\xfe\x3b\xd6\xba\x5c\x18\xb3\xab\x63\x50\x0c\x5a\x75\x18\x5a\x3d\xdc\x6c\x38\x17\x78\xf7\xe1\x9c\x15\x65\x0a\xed\x65\x3a\x67\xe6\xf2\xa8\xb6\x9d\x8b\x93\xbb\x8d\xe2\x5b\x8a\xac\x70\x0f\xca\x3d\x5e\x1f\xb7\x58\x98\x2f\xab\xe9\xaa\x3e\xd0\x8f\x1e\xe1\x88\x03\x50\x00\xd1\xc8\xb6\xc5\xa5\xc9\x15\x8f\x9f\x37\xdb\x87\x15\xec\x11\x6c\x00\xbb\x97\xd8\xf5\x8a\x90\xac\x1e\xa0\x8e\x3e\x59\xa5\x7d\x87\xf0\x4a\x7a\x28\x93\xd4\x72\x9b\xf9\x49\xc8\x56\x29\xb4\x20\x83\x28\xfa\xca\x61\x43\x01\x1d\x34\xd4\x8d\xdd\xf2\x1f\x44\x06\x3b\xad\x26\x87\x41\xa9\x21\x64\x90\x84\x75\xfe\x72\xb0\x1f\x40\x90\x5f\xa8\xc6\xa5\xc9\x80\xcd\xae\xb3\xa2\x95\x8f\xae\xe4\xcf\x17\x2b\x79\x5c\x3e\x91\xc7\x52\x43\xe1\x58\x8d\x38\x6e\x39\xb9\x62\xe1\x81\xc2\xf1\x69\x84\x7c\xa2\x42\x06\x9f\x07\x4c\x87\x34\x91\xa1\xbc\x2f\x34\xf6\xd1\x5d\x68\xb3\x29\x89\x75\x74\x78\xcd\xb8\xd9\xf9\xbe\x53\x4a\x96\xb5\x9a\xd3\xe7\x0a\xcb\x7a\x2a\x3e\xeb\x92\x5f\x40\x2e\xdb\xad\x7c\x5e\x80\xc8\x1c\xfe\xbb\xe4\x5f\xb3\x6e\x24\x80\xa4\x07\xe4\xd1\xc3\xef\x34\x8d\x63\x90\xb1\xb3\x4a\x2f\xa7\xa0\xc6\x57\xcb\x4e\x26\x63\xdc\x09\x7c\x8b\x7c\x5c\x11\x83\xc3\xfc\x03\x08\xc4\x00\xbd\x20\xdf\x09\x50\x48\xbd\x2e\xcd\xb8\xe9\x6c\x9e\x6d\xa0\x37\x10\x1d\xba\x3f\x52\x9e\x8c\xe1\x51\xad\xb3\x6a\xdf\xfd\x2a\xc8\x8d\xad\xf1\xe7\xaf\x1b\x8e\x1e\x4b\xd3\xfa\xd4\x53\x00\xa1\x9c\x3c\x06\x93\xc7\xa6\x91\x85\xa7\x0c\x6b\x66\x89\xb7\x69\x37\x92\xf6\xfc\x17\x39\xb6\x82\x7c\xbd\x28\x0f\x9b\x5b\xf3\x3b\x00\x00\xff\xff\x9a\xe9\x67\x93\x58\x04\x00\x00")

func definitionsInfo_object_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsInfo_object_metaToml,
		"definitions/info_object_meta.toml",
	)
}

func definitionsInfo_object_metaToml() (*asset, error) {
	bytes, err := definitionsInfo_object_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/info_object_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbe, 0xfc, 0xc9, 0x91, 0x3, 0x38, 0x82, 0x77, 0xaa, 0x97, 0xb1, 0x0, 0x25, 0xe, 0xb3, 0x46, 0xd3, 0x3, 0x6b, 0x3, 0xd8, 0xf3, 0xa, 0x4c, 0x8a, 0x82, 0x90, 0x2f, 0xa0, 0x14, 0x15, 0x2a}}
	return a, nil
}

var _definitionsInfo_storage_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd2\xbd\x72\xab\x30\x10\x05\xe0\x5e\x4f\xa1\xa1\xe7\x56\x77\xd2\xf9\x11\x5c\xa5\xf4\xa4\x58\xa3\xc5\xd9\x09\xfa\x99\xd5\x12\x43\x32\x79\xf7\x8c\x84\x21\x63\x4c\x1c\xdb\x94\x08\x9d\xef\x2c\x48\x3b\x08\x01\x9d\x29\x5d\x6b\xf7\xc8\xa5\x85\x8e\x6c\x6b\x5f\x94\xf4\x01\xf5\x46\x17\xe4\xa4\x50\x06\x63\xc5\x14\x84\xbc\x4b\x6b\x5b\xe8\xf4\x10\xd3\x43\x2c\x6a\x72\xe3\x8a\x0f\xc8\x90\x76\xfe\x2b\x94\x1a\xf5\x48\x1f\xb8\x68\x3f\xfd\xbf\xa6\xa7\x58\xa2\x03\xf2\x55\x5e\xbc\x40\xf3\x68\x49\x0e\x4f\x55\x8b\x35\x95\x0f\xfd\xbd\x7c\xda\x36\xa8\xb5\x67\x9d\x84\x19\x5a\xa3\x54\xaf\xeb\xd4\x4c\xcc\xd8\xc6\x57\xf9\xe1\x07\x8b\xc2\xe4\x0e\xe9\x9d\xf5\xef\xb8\xae\x31\x09\xb3\x42\xdb\x36\x42\x01\x58\xee\xbd\x43\x59\x4e\xc1\xd3\x2d\x4a\xbf\x7f\xc2\x7e\x2d\x79\x64\xfc\x0c\xe6\x6f\x30\x58\x93\x43\xa3\xf7\xbd\x8e\xe2\x19\x0e\xc8\x8b\x3e\xb9\x9b\xfc\x61\xdb\x0d\xbe\x03\x8b\x97\x27\x82\x5d\xf0\x2c\x7a\xa3\x85\x5b\x54\x6a\x17\xfb\x28\x68\x4b\x8b\x02\x06\x04\xce\xea\x91\x6b\xa8\xf0\xf3\xeb\x62\x88\xe7\x1c\xda\x9e\x32\xb9\x16\xa3\x1e\xa8\x69\x9e\x91\xcc\xc3\x1c\x3d\xbf\x95\x86\xf8\xef\x81\x8e\x4c\xb2\xf2\xc6\x64\xe2\xec\x34\xbf\x03\x00\x00\xff\xff\x41\xde\xf8\x57\x73\x04\x00\x00")

func definitionsInfo_storage_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsInfo_storage_metaToml,
		"definitions/info_storage_meta.toml",
	)
}

func definitionsInfo_storage_metaToml() (*asset, error) {
	bytes, err := definitionsInfo_storage_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/info_storage_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0xdc, 0x18, 0x27, 0xc0, 0x7d, 0xb9, 0x57, 0x8b, 0xe9, 0x90, 0xfd, 0xf8, 0xb3, 0x2, 0x40, 0x16, 0xce, 0x94, 0xf9, 0x18, 0x51, 0x52, 0x7a, 0xff, 0x6b, 0x10, 0x2d, 0x50, 0x8, 0x2e, 0xb0}}
	return a, nil
}

var _definitionsOperationsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4d\x8f\xdb\x38\x12\xbd\xfb\x57\x14\x34\x87\x5c\xba\x8d\x39\x0f\x90\x43\x26\x9d\xc5\x04\xc8\xd7\x4e\x67\x76\x0e\x41\xa3\x4d\x4b\x25\x8b\x1b\x89\x54\x48\xca\x8e\xf7\xd7\x2f\x8a\x45\x49\xd4\x97\xad\xf4\xf4\xee\x65\xf7\x12\xb8\x25\xb2\xaa\x58\x7c\xf5\xea\x91\xca\x17\x51\xd7\xa8\x32\x34\x0f\x9b\x0c\x6d\x6a\x64\xed\xa4\x56\xf0\x12\x12\x69\xc1\x15\x08\x52\x39\x34\xb9\x48\x11\x72\x6d\xe0\x95\x1f\x0d\x06\x4b\xe1\x30\x03\x5d\xa3\x11\x34\xc1\x6e\x93\xcd\xa6\xb3\xb5\xd5\xf5\x36\x35\x28\x1c\x3e\xf2\xa3\x87\x4d\x2d\x8c\xa8\x2c\xbc\x84\x2f\x49\x2d\x5c\x91\x3c\x6c\x0c\xda\xa6\x74\xfc\x48\x27\x13\xf7\x49\xb2\x39\xc9\xb2\x04\xb6\x03\x42\x01\x9b\x02\xbd\xff\x27\xa6\x6e\xbb\xd9\xfc\xf4\x13\xfc\x8a\x85\x38\x4a\x6d\x36\x9b\x5b\x78\xed\x07\x86\x00\xef\x7f\xfb\xf8\xc7\xbb\xbb\xc9\x64\xb1\x2f\x31\x18\x80\x93\x74\x05\xd4\xda\x4a\xef\xf1\x67\x10\x2a\x03\x2b\xff\x85\xf0\xf3\x76\xc1\xda\x87\x8f\x9f\xc1\xa0\x6b\x8c\x22\x8b\x68\x8c\x36\x20\x38\x4b\xc1\x26\x7e\x97\xd6\x6d\x37\x00\xb7\x70\x8f\xe6\x28\x53\xec\x22\x29\x30\xfd\xea\x7d\x64\x58\xa2\xc3\x78\x96\xcc\x79\xa2\xdd\x6e\x68\xd9\xc3\x3c\x9e\x8c\x8c\xd2\x38\x4a\x92\xcf\x50\xc8\x4b\xaa\x95\x43\xe5\xc0\xe9\x69\xb2\x92\x38\xff\x3a\xb9\x81\xc4\xd0\x3f\xb4\xdc\xd1\x4e\xa8\xe4\x61\xc3\x93\x1e\x2b\x9d\x21\x39\x61\x53\x93\xfd\xd5\x55\x25\xdd\xc5\xc0\x78\x88\x5f\x74\x2e\x95\xb4\x45\x14\x58\x6d\x74\x8a\xd6\x8e\x23\xbb\xe0\x7c\x5f\xea\xf4\xeb\x5a\x9c\xfe\x4a\x83\x97\x60\x1a\x2c\x45\x28\xf5\x4f\x9e\x01\xa4\xa0\xf0\x04\xde\xd8\x3c\x4c\x5b\x5c\x71\x78\x7f\x05\x56\xae\x10\x0e\x0a\x61\x41\x09\x27\x8f\x08\xb6\xa9\x6b\x6d\x9c\x5f\xfc\x4e\x1f\xd1\x78\xdc\xec\x20\xd3\x68\xd5\x0b\x07\x1f\xde\xbc\xb9\x23\x68\x30\x0e\xc7\xb6\x2d\x68\x03\x4a\xcf\xf9\x68\x0d\x14\xe2\x88\xd7\x9c\x3d\x05\xeb\xd1\x66\x30\xd4\xc3\x5e\xcc\x01\xca\x0f\x18\x00\x9d\x93\x7d\x11\xdf\x37\x90\xec\x65\x76\x1d\xe6\xde\xd2\x18\x1e\xba\xda\x4b\x75\x31\xa6\x30\x84\x03\xb1\x04\x43\x5f\x7f\x8b\x85\xb7\x97\x99\x5d\xe7\xbc\x94\xd6\x5d\xf2\x4c\xef\x5b\xb7\x7b\x2c\xb5\x3a\x50\x4a\x5c\x21\xed\x82\xf7\x61\x0a\xf6\xf2\x42\x18\xa9\xae\xe5\xda\x5a\x7b\xad\xeb\xf3\xb6\x9f\xc4\x79\xab\xcf\x83\x72\xb2\x26\xa5\xc5\x67\xd6\x5d\x28\x22\x5d\x9f\x29\x75\x1f\x19\x29\xda\x40\xd5\x94\x4e\xd6\x3d\x65\x4b\xe5\xdd\x5b\xc6\xe7\x4c\x0b\x20\x0b\x5a\x95\x67\xb6\xa5\x15\x7a\x14\xfa\x27\xf4\x47\x5b\x93\x31\xc8\xef\x3e\x7e\x78\xf1\xb9\x2b\x8f\x16\xd9\x1c\x0b\x28\xad\x6e\xb1\xaa\xdd\x19\x32\x69\x30\x75\xda\x9c\x29\x2e\xff\x36\x97\x25\x5a\x30\x98\x36\xc6\xca\x23\x96\x67\xb6\xfb\x87\x45\xd3\x99\x93\x55\x5d\x62\x45\x60\xbd\x60\x50\x78\xf2\xae\xcf\xb1\x2d\xd8\x9f\x69\xa9\x95\xc5\x32\x67\xbb\xaf\xd9\x00\x79\x65\xe4\xf7\x06\x42\xd9\x05\x0e\xd9\xbd\x31\x86\x53\xf8\x5e\x67\xf8\x56\x1d\x45\x29\xb3\xdd\xb6\x4d\xcf\x65\xd2\xc9\xac\x1b\x12\xc3\x73\xb3\xce\xc4\xc1\x7f\x9e\x79\x22\x97\x11\xfb\xdc\xc2\x2b\xb0\x4d\x4a\x4d\x28\x6f\x02\xf8\x74\xcd\x6d\x02\x6c\xa1\x9b\x32\x83\x3d\x91\x0d\x6d\xa1\xc3\x1b\x38\x15\x32\x2d\xa0\x42\xa1\xec\xc8\xec\x0b\xdb\x71\x12\xb9\xae\xd0\x89\x4c\x38\x11\x59\xf1\xa8\x15\x15\x52\x8a\xad\x49\x3b\x20\x32\x05\xd2\x4e\xae\xac\xb5\xbb\x76\xd3\xb7\xdd\xc4\xa8\x8d\x65\x72\x6a\x66\xda\x9f\x32\x69\x66\x09\x62\xa1\xe5\x6d\xbe\xe4\xe8\xd2\x62\x6d\x84\x7f\xa3\xc1\x3e\xba\x30\x8d\xe2\xf3\x3f\xa7\xfd\xf5\x06\x92\xc6\x94\xcb\x8c\xe0\xa7\x41\x6e\x74\x05\x02\x0e\xf2\x88\x0a\x1a\x53\x12\xa0\x68\xfa\xb4\xfe\xbd\xef\x1f\x6c\xab\x53\x28\xb0\xd7\x4e\x33\xac\x02\xc3\x53\x80\x60\xf0\x5b\x23\x8d\x54\x07\x5e\x21\xbd\x6a\x4c\xd9\x82\xa2\x22\x74\xaf\x4b\xf9\x7b\x7d\xc4\x6d\x37\x87\xf2\x4d\x3f\x7e\x98\x7f\x69\x52\xdf\xba\xae\x92\x2d\x79\x65\x6a\xad\xf8\xd7\x93\xc9\x96\x1d\xcf\x71\xe3\x45\x4a\x5d\x9e\x36\x65\xcf\xf7\x3c\xf6\xaf\xb0\xa7\x37\xf1\x7f\xf6\x9c\x2d\x19\x46\x40\x57\x31\xc1\xde\x7f\x89\x3e\x59\x25\x08\xe3\x56\xd7\x4b\x3b\x61\xe9\x80\x10\x59\x8c\xd8\xb5\x7b\xfa\x5c\x07\x85\xce\xe0\xd2\x61\xb6\x0f\xf4\x87\x39\x6d\x92\x9a\x5e\x60\x47\xeb\x58\x29\xb2\xfb\x40\xaf\x08\x6d\xa9\x32\xfc\x3e\x91\xda\x37\x90\xd0\xe4\xa9\xdc\xf4\x4f\x67\xf2\x1d\x40\x73\x2d\xd2\x76\x5c\x1c\x21\x34\x75\xa9\x45\x16\x44\x95\xb2\xce\x34\xa9\xeb\x55\xe5\xcc\x02\x68\xd6\x8c\x22\x5f\x88\xcd\x8b\xf2\x2b\x71\x79\x61\xee\xcd\x8e\x75\xf9\x62\x22\x87\x29\xab\x67\xa4\x79\x1b\x4f\x2d\x0e\x6b\x71\xfe\x49\x1c\x70\x06\xe2\xa1\x14\x5b\x32\x30\x42\x65\xba\xe2\x6d\xdf\x76\x1e\x22\xdc\xd3\xdf\xcf\x05\x79\xb2\x75\xe5\x06\xc7\x47\xfd\x44\xbc\x77\xa1\x33\xd2\x39\xf2\x55\x20\xb7\x35\xa6\x32\x97\x29\xe8\x3c\xb7\x78\x15\xe8\x3c\xea\xfa\xa1\x92\x22\xa0\xc0\x0c\x8a\x59\x09\x95\x24\x9b\xd9\xad\xfb\x9d\xc6\x6f\x37\x9b\x3b\xac\x0d\xa6\xb4\x81\xbf\x50\x1b\x84\x7b\xa7\x8d\x38\xe0\x6f\x9f\x3f\x7f\xba\x97\x07\x85\x06\xa4\xb2\x0e\x45\xd6\x66\x20\x38\xa2\x1c\xf8\x9f\xd7\x36\xee\xa2\xfe\xaa\x8d\x3e\xca\x8c\xf6\xee\x24\xce\x2d\x85\xa7\x42\x81\x37\x1d\x6d\xc3\x4c\xa4\x7f\x6f\xd0\x9c\x29\x46\x8a\xf5\x77\x14\xd9\x38\xd2\x20\x2c\xa6\x39\x21\x07\x95\x90\xca\x09\xa9\xa2\xc2\xb6\xbc\xf4\x56\x90\x30\x53\xb7\x46\x7a\xb8\xae\x10\xc0\xc1\x12\xe7\x4e\xa8\x14\x87\xbb\xad\x44\x35\xbe\x15\xa3\x19\xe8\xb5\x70\xec\x90\xbb\xe4\xbc\xc3\xd0\x41\xc5\x4a\x67\x43\xc3\x07\x5c\x60\x96\x03\x3a\x10\xe0\x35\xc9\xd4\xb0\x07\x4e\x2b\xd7\x9e\xb6\x22\x62\xae\x0b\x9c\x26\xca\x72\xea\xd6\x42\xa3\x32\x34\x4c\x70\xbd\xfb\xa1\x33\xc9\xae\xc2\xdc\x75\xf4\x35\xda\x70\xde\xef\x60\x21\x4e\xff\x14\xe0\xb5\x90\x26\x54\x6e\x5f\x8e\xcb\x30\x6f\xf7\xaa\x93\xbe\x5e\x8f\x2f\x0a\xdf\x3b\x1e\xee\x85\x6e\xd6\xfe\x7e\xb2\xf8\x35\xc8\x3a\xb6\x2c\x2f\x8a\x5d\x1e\xf6\x48\x3b\x30\x10\xb7\x5d\x38\xd2\x82\xcc\xb0\xaa\x35\x31\x5a\x70\xdc\x0b\xb4\x76\x89\xe5\x49\x9c\x6d\x4b\xa9\x4a\x96\xcc\xa9\x3c\x3c\x18\x0a\xd4\xab\xf0\x88\xa6\x53\xc5\xdc\x3d\x3f\x68\xf7\x86\xf8\x76\x17\x8f\x1f\x2e\x6b\xfe\x7e\xb1\x93\xa9\xa1\xf4\xa3\x5d\x6c\xe5\xde\x3c\xec\x82\xff\xb4\x31\x86\xb2\xd0\x81\xaf\x9d\x35\x02\x1a\x3d\x4e\x1e\x36\xa5\x4e\x45\x09\x2f\xc1\x99\x06\x47\xee\x18\xe1\x97\x20\xc3\x1d\x9e\x01\x33\x68\x71\x72\x19\x41\x21\x4c\x2e\x92\xbe\x97\xcc\x9f\x52\x47\x5f\x06\x5a\x20\x64\x98\x8b\xa6\x74\xb0\x7b\x27\xad\x3f\x7c\xf8\x53\xc7\x68\x70\x0f\x88\x6e\xd8\x9d\x34\x3b\xff\x31\x43\x37\xce\x27\x9e\xb7\xc0\x6b\xfc\x7f\x48\xe3\x1a\x51\xd2\x90\xd8\xd8\x70\xcb\x76\xf7\x4e\xb8\x1d\xb1\x7b\x49\x25\xc8\x96\x77\x73\x7b\x65\x50\x2c\x5c\xf5\xd3\x1b\xef\x9c\x4e\x58\x2f\x2c\x84\xbd\x99\x39\xf0\x9f\x06\xc9\x1e\xf5\x53\x92\x90\xfa\x31\x15\x65\xb9\x17\xe9\xd7\x99\xf6\x3a\x8c\xa7\x65\xfc\xeb\xda\x64\x89\x11\x06\x48\x59\x23\x60\xa2\x8f\x47\x94\x6f\xa1\xce\x20\x6a\x09\xa9\xaf\xdf\x79\x41\x13\x8b\x19\x8b\x2a\xf3\x93\x5e\x7d\x7a\x1b\x26\x8d\x87\x89\x34\xc5\xda\x41\x7f\x08\x05\x0a\x9e\x64\x4f\xf0\x4d\xa1\xcf\xed\x8e\x75\xe2\x0a\xb4\x87\x6b\x5f\xa9\xdf\xc8\x2c\x08\x8f\x65\x42\x0b\xb5\x20\xa9\x72\x0d\x3a\x8f\xee\xb9\xa7\x20\xa7\x59\x57\x57\x54\x48\x35\x3e\x24\xbf\xf6\x53\xfc\xa9\x35\x93\x79\x8e\xbe\xee\x3b\xd8\x07\x05\x4b\xe0\xee\xdf\x0e\xf2\xb2\xf8\xed\x4d\xe6\xa1\x4c\xa5\x3a\xc4\xb1\x90\x2c\x17\x2e\x2d\x66\x32\xea\x55\xe2\xec\xb5\x55\xfc\xf9\xac\xcf\xef\x65\xec\x2e\xe4\x97\xa5\xa8\x3f\xf1\xf2\xc7\x03\x5f\x43\xd3\x8c\xfe\xe9\xc7\x8d\x48\xa3\xb1\x68\x2c\xd4\xc2\x5a\xcf\xe6\x3b\xa9\xb7\xa4\xb5\x90\xea\x7d\xed\xd5\xc3\xdc\xec\x0b\x77\x10\xf1\x28\xc9\x13\x9f\x7c\x09\x31\xeb\x7a\x72\x1b\xd1\x7f\xc5\xe5\x7b\x9e\x41\x04\x39\x48\x17\xe2\xd8\x8e\xb3\xf4\xbf\xfa\x65\x6d\x74\x3b\xc3\x08\xfb\xb1\xeb\x99\x27\x5d\xcd\x70\xdf\x43\x92\xfa\xc1\xa9\xc1\x6f\x0d\xda\x71\xe3\x7f\x2c\x9c\xab\x1f\xad\x3f\xc0\xac\xd3\x82\xf7\x6d\xe3\xbf\x7a\x9c\x15\x8d\x2b\x50\x39\x99\xfa\xb7\xdb\x05\xaf\x54\xdd\xdf\xe8\x84\xe2\xff\xe6\x37\x57\xba\x9b\x5f\x78\x77\x53\xeb\x6f\x12\xf7\x67\x68\x2c\xd1\x89\x37\x05\x9e\x27\xd0\x51\x45\x52\x1d\xf7\x81\x74\x79\xb0\xf3\x3d\x11\xbf\xd7\xd2\x8c\x29\xd9\xe0\xb7\xb8\xe1\x5d\x8b\x3e\x70\xd5\xf2\x61\x77\xcc\x30\xcf\x16\x7d\xdb\xc3\x2f\xae\xa2\x94\x6a\xf5\xc7\x7b\x1a\x9b\x74\x73\xa2\x9b\x08\x7a\x30\x4b\xc7\x4e\x98\xc3\xe4\x38\x3e\xdf\xd3\xfe\x1c\xb4\x74\xb2\x18\xf5\xb0\x01\xe1\xee\xc8\xfa\xce\x03\x7f\xc7\x0e\x76\xa1\x3b\xed\xc3\xb5\x0a\x55\x2c\x51\xca\xde\xea\xb2\x71\x18\xf4\xde\x2d\xbc\xcd\xfb\x19\x4a\xb7\x4c\x70\x13\x1a\xfd\x3b\xf2\x19\x5a\xeb\x52\x30\xfd\x37\x0e\xb6\xc6\xa1\xb0\x9d\x5f\x3c\x95\xf4\x4f\xa5\x25\xcd\x79\xae\xa2\xe9\x53\x57\xe1\x98\xe1\xcb\x75\x30\x74\xc0\xb0\x5e\xe2\xcc\x86\x31\x71\x49\xeb\x5a\xe3\xb6\x65\xdf\xb9\x7b\xf5\x40\xc9\xa7\x02\x07\xdf\x1a\x3c\x33\x7a\x07\x6d\x51\x77\x4c\xc8\xdc\x16\x47\x18\xed\x08\xb9\x62\xf2\x09\x92\x96\x48\x34\x88\xd9\xdb\x38\xb4\xd7\x23\x2d\x7d\x64\x99\xec\xf1\x05\x39\x0a\xd7\x18\xe4\xa8\xda\x88\x5a\xbd\x37\xa4\xe9\xfe\xc4\x66\x5b\x52\x45\xe5\xff\x8f\x91\x3f\x0c\xb7\x96\xfa\x03\xdb\x11\xc3\x55\xd5\xbf\x03\x00\x00\xff\xff\x39\x2f\x1d\x6a\x77\x25\x00\x00")

func definitionsOperationsTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsOperationsToml,
		"definitions/operations.toml",
	)
}

func definitionsOperationsToml() (*asset, error) {
	bytes, err := definitionsOperationsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/operations.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xaa, 0x9e, 0x5d, 0xcf, 0x93, 0x19, 0x98, 0xa2, 0xa2, 0xcd, 0xd5, 0xba, 0x92, 0xa4, 0x9e, 0xbb, 0xc6, 0xec, 0x7b, 0x33, 0xa3, 0x37, 0x3d, 0xa3, 0x3f, 0x22, 0x84, 0x8f, 0xfe, 0xc4, 0x99, 0x5e}}
	return a, nil
}

var _definitionsPairsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xcb\x6e\xdb\x40\x0c\xbc\xeb\x2b\x08\x9f\x5a\x20\x70\x2f\x6d\x0f\x05\x7a\x0b\x8a\x14\x48\x91\x36\x8f\x53\x10\xd8\x94\x96\xb2\x58\xaf\x96\x2a\x97\xb2\xa3\x7e\x7d\xb1\x2b\xbf\x90\xc4\x79\xf4\x26\x88\xe4\x0c\x67\x76\x76\x6f\x2b\x09\x46\xc1\x66\xad\xfb\x74\x57\xd8\xd0\x11\x7c\x85\x49\x34\xe5\xb0\x98\x14\xc5\xae\x9c\x2a\x77\x85\xa3\x1a\x7b\x6f\x58\xfa\xd4\x66\xda\xd3\x91\x11\x0e\x3d\x1a\x4b\x98\x99\x2c\x29\x3c\x06\x76\x14\x2b\xe5\x2e\xb5\xe4\xdf\x1d\x55\x5c\x0f\x60\x0d\xc1\xe1\x38\xe4\x71\xa8\x45\xc1\x73\xb4\x8c\xae\xe4\x28\x18\xa3\x7f\x35\x6a\x23\x6b\x30\x81\x4e\x65\xc5\x8e\x60\x8f\x90\x81\x23\xe9\x8a\x2b\x82\xf4\x69\xa2\xb8\xa0\x44\x43\xc1\x75\xc2\xc1\xfe\x97\x64\x3b\xff\x1c\xc5\x7d\xc7\x4a\x7b\x02\xe3\x96\xa6\xa7\xbd\x66\xe5\x8f\x78\x26\x93\x62\x4b\xb5\x6e\x28\x64\xab\x7a\xf5\xa0\x64\xbd\x06\x72\x50\x0e\xa0\x84\x55\x03\x6b\xf6\x1e\x46\xf0\x69\x51\x9c\x52\xa7\x54\xa1\x91\xfb\x02\x37\x91\x60\xfe\xab\x27\x1d\xae\x78\x11\xce\xae\xaf\x7f\x5e\x12\xba\x39\x70\x88\x46\xe8\x40\x6a\x98\x5f\x26\x8c\x39\x60\x70\x30\x1f\x41\xe6\xc0\x31\xd3\x39\xaa\x3c\x2a\x39\x40\x5d\xf4\x2d\x05\x9b\x16\x69\xad\xe2\x96\x83\x91\x56\xd4\x99\xe8\x5e\xce\xf7\xfd\xcf\xd4\x93\xce\x6f\xd6\x8a\x3b\x10\x7c\xce\xd1\x7e\x88\xcb\x66\x78\xa9\xb2\xee\x37\x45\x65\x3b\xf4\x8c\xc9\x01\x5b\x7a\x13\xe6\x66\x16\xd2\x60\x02\x90\xf2\x37\x55\x0f\x37\xbf\xc8\x3f\xc7\xdd\x1f\x60\xed\x4b\xd0\x70\xc8\x91\x95\xba\x8e\x74\x90\x24\x0e\xf6\xf9\xe3\xd1\x25\xc6\xee\x2c\xc9\x1a\x8e\xa0\xf4\xa7\xa7\x68\x27\xbb\xcd\xf2\xf9\x46\xa2\x65\xca\x5b\x6e\xd9\x8c\x94\x54\x8b\x52\x4a\x81\x4b\xb4\x6d\xef\x8d\x3b\x54\x9b\xb1\x7b\xea\x6a\xb3\xcc\x2a\xf4\xbe\xc4\x6a\xb9\x2f\xd7\x7d\xa8\xde\xdd\xde\x95\x83\xd1\xfb\xc9\x53\xf7\xfd\xc8\xd6\xeb\x06\x0d\x4c\x9c\x00\xad\x48\x07\x48\x61\x86\xf5\xb8\x0c\x38\x34\x84\x5a\xa5\x85\x28\xbd\x56\xd9\xd7\xc8\x7f\xe9\xb5\x96\xa4\xde\x97\x0c\x91\xe0\x87\x91\xcd\x73\xcb\x46\x0e\x36\x4f\x57\x66\x4f\x8c\x6b\xd1\xe5\xcc\xb1\xbe\x18\x87\x83\xab\x96\x12\x91\xe6\xc0\xb1\x1e\x49\xd9\xc9\x46\xb1\x74\x34\x5e\xdd\x71\x9f\x32\x69\xf7\x68\xbc\xa2\xdd\x39\x39\xd6\x69\xb1\x5d\x03\xae\xce\x2e\x6e\xce\x4f\x53\x23\x06\xc0\x32\x8a\xef\x8d\xa0\x43\x6b\x0e\x9a\xb6\x58\x9b\x93\x48\x50\x1f\x80\x6b\x08\x62\x10\xc9\x9e\x84\xbb\x09\x7c\x0f\xd1\x06\x3f\x9a\x36\x26\x78\xe7\xd6\x46\x41\x9c\x16\xdf\x44\xa1\x8e\x0f\x0b\x90\x15\x04\x27\xeb\x08\x9d\x47\xab\x45\xdb\x93\x6c\x44\x49\x0d\xae\x58\x34\xbd\x07\x8e\x6a\x4e\xaf\x4e\xa4\x0e\x15\x8d\xfc\x30\x3e\x06\xff\x02\x00\x00\xff\xff\x52\xa0\xcf\x26\x50\x06\x00\x00")

func definitionsPairsTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsPairsToml,
		"definitions/pairs.toml",
	)
}

func definitionsPairsToml() (*asset, error) {
	bytes, err := definitionsPairsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/pairs.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd4, 0x36, 0xbf, 0xd6, 0x11, 0x42, 0x49, 0x2c, 0x3, 0x7, 0x3e, 0x5e, 0x79, 0xa3, 0x1c, 0xbc, 0x2a, 0xde, 0x96, 0x2d, 0xfd, 0xc9, 0xb0, 0xe5, 0xdb, 0x60, 0x21, 0xd5, 0x3f, 0xf, 0xc6, 0x1b}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"definitions/features.toml":          definitionsFeaturesToml,
	"definitions/fields.toml":            definitionsFieldsToml,
	"definitions/info_object_meta.toml":  definitionsInfo_object_metaToml,
	"definitions/info_storage_meta.toml": definitionsInfo_storage_metaToml,
	"definitions/operations.toml":        definitionsOperationsToml,
	"definitions/pairs.toml":             definitionsPairsToml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"definitions": {nil, map[string]*bintree{
		"features.toml":          {definitionsFeaturesToml, map[string]*bintree{}},
		"fields.toml":            {definitionsFieldsToml, map[string]*bintree{}},
		"info_object_meta.toml":  {definitionsInfo_object_metaToml, map[string]*bintree{}},
		"info_storage_meta.toml": {definitionsInfo_storage_metaToml, map[string]*bintree{}},
		"operations.toml":        {definitionsOperationsToml, map[string]*bintree{}},
		"pairs.toml":             {definitionsPairsToml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
