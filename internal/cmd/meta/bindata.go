// Code generated by go-bindata. DO NOT EDIT.
// sources:
// meta.tmpl (3.881kB)

package main

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

var _metaTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x5d\x6b\xe3\x46\x14\x7d\xd7\xaf\xb8\x35\xa6\x48\x41\x1d\xf5\x39\x8b\x1f\x96\x64\xa1\x4b\x69\x76\x69\x43\x29\x98\x10\x26\xf2\x95\x33\x58\x9a\x51\x67\x46\x4a\x85\xaa\xff\x5e\xe6\x43\x96\x64\xcb\x8e\xc3\x2e\x85\x85\x7d\x31\x68\x34\xf7\x9e\x73\xc6\xe7\x1e\x4d\x92\xc0\x8d\xd8\x20\x6c\x91\xa3\xa4\x1a\x37\xf0\xd4\xc0\x56\xec\x9f\xa1\x66\x14\x18\xd7\x28\x39\xcd\x93\xb4\xd8\x24\x05\x6a\xfa\x0e\x6e\x3f\xc1\xdd\xa7\x7b\xf8\x70\xfb\xf1\x9e\x04\x25\x4d\x77\x74\x8b\xd0\xb6\x40\xee\x68\x81\xd0\x75\x41\xc0\x8a\x52\x48\x0d\x61\xb0\xd8\x32\xfd\x5c\x3d\x91\x54\x14\xc9\x5f\x15\xe5\x2f\x22\x51\x5a\x48\xba\xc5\xc5\x99\x77\x89\x6e\x4a\x54\x8b\x20\x0a\x82\xb6\xfd\x09\x58\x06\xe4\x0f\x94\x35\x4b\x6d\x77\x00\x80\x24\x01\xbf\x22\xef\x9b\x12\x81\x29\xd0\xcf\x08\xca\xaf\x81\x69\x00\x99\x90\x13\x5a\xa6\x30\x15\x5c\xe9\x69\xed\xca\xee\x56\x64\xbc\x18\x2e\x46\x85\x8b\xc8\xd2\x40\xbe\xb1\xe2\x0c\xb6\xe3\x39\xc5\xf6\x6b\xf3\xd8\x1e\x77\x5c\xb7\xc7\x1d\x2d\x1e\xe2\x06\x6d\x0b\xcb\x5b\xaa\x29\x5c\xaf\x80\x58\xf8\x9a\x4a\xe0\x42\xbf\xcf\x73\xf1\x82\x1b\x5f\xfc\x3e\xd5\x4c\x70\x58\x41\x41\xcb\xb5\xd2\x92\xf1\xed\x83\xd2\xb2\x4a\x75\xdb\xb5\x96\xbd\xa4\x7c\x8b\xb0\xdc\xc5\xb0\xac\x6d\xb3\x1b\x5a\xd2\x27\x96\x33\xdd\xf4\x67\xb3\x68\xdb\xe5\xae\xeb\x16\xd7\xb0\x2f\xed\xe2\xb1\xf4\x2e\x70\xf0\x74\x82\xfd\x99\x32\xa9\xa6\xd0\x97\xb3\xf0\x3d\x7a\x0a\xfe\x10\x89\xd3\x63\xc4\xef\xe0\x5f\x48\x69\x81\xf9\x0d\x55\xd8\x75\xd7\xd0\xda\x8d\xe3\x66\xd8\xc4\xb0\x7c\x34\xed\x96\x75\xdf\x68\xaf\x07\x9b\x23\x45\x7d\xbd\x57\x65\x1e\x8f\x74\x8e\x65\x3a\x5f\x7c\x99\xcc\xa9\x7d\x17\x4e\x99\x65\xd6\xee\xf9\x7e\xa9\xa6\x4b\x74\x99\x99\x37\x7f\x7c\x8e\xc0\x8a\x32\xc7\x02\xb9\x56\x7b\x5f\x12\xff\x2e\x8c\x48\x90\x55\x3c\x85\x30\x85\xab\x9b\x9c\x21\xd7\x51\x5f\x17\x52\xe7\x35\x27\x3c\x86\x92\x32\x09\x84\x10\xf7\x1c\xc1\x93\x10\xb9\x17\xc5\x32\x78\x8c\x41\xec\x8c\x8c\x13\x96\x5d\xbb\x6e\x0f\xef\xcc\xb6\xe1\x28\x24\xea\x4a\x72\xc8\x68\xae\xd0\x09\xe9\xc7\xfe\x63\x06\x5c\x38\x50\xc6\xcb\x4a\xc7\xf0\x82\x20\x78\xde\x00\x47\xdc\x80\x16\x90\x3e\x63\xba\x03\xd7\x97\xf4\x3c\x72\xe4\xa1\x29\x8a\x60\xb5\x82\x9f\x8f\x91\xb4\xac\x7a\xa0\x23\xea\x33\x76\x1f\x68\xff\x70\x09\x6f\x93\x06\x8f\x31\x58\x2f\xb8\x7f\xd8\x0a\x18\xea\x2e\x84\x5b\xd7\x47\x88\xb3\xa8\x03\xb2\xfb\x1d\x8b\xec\x82\x57\x27\xd1\x26\x98\x61\xe8\xd7\x66\xe6\xd0\x9b\xef\xc2\x69\xfc\x85\x2a\xdb\x03\x9b\x83\x2e\xc6\x2c\xbe\xc3\xdc\x6b\xf7\x82\xf1\x0d\xfe\xe3\x62\x90\x98\x90\xfc\x8d\x96\x6e\xb3\xef\x3e\x75\xb8\xf5\x6d\x49\xa5\xc2\xd1\x09\xce\x28\x08\x45\xa9\x95\x31\xee\x95\x8b\xe1\xcf\xd6\x1d\xe1\xd5\x79\xdd\x31\xa0\x94\x42\x46\x5e\xb8\x44\x55\xe5\xda\x88\xfd\xf1\x7c\x5d\xeb\x7d\x55\xd3\xbc\x42\x65\x0a\x0a\xba\xc3\x70\x94\x20\xf6\x43\x9b\xd1\x14\xdb\x2e\x3a\x61\x1a\xcb\xf8\x72\xd3\xbc\x1a\xa6\xb3\x66\x4a\x05\xd7\x8c\x57\x87\x46\xfa\x3a\x80\xeb\x9a\xfc\x8a\xcd\x5b\x71\xdd\xa1\xf9\x5a\x58\x41\x4d\xfe\x34\x2b\xe3\x79\xf5\xf7\x84\xc1\x72\x26\xc0\x6b\x18\x1d\xea\x7e\x55\xec\xc6\xb6\x1b\xbe\xeb\xb3\x46\x96\xf8\x77\xc5\x24\x6e\xa6\x7e\xae\xed\x29\xac\x7a\x62\xce\x40\xb3\x16\x7e\x98\xb0\xdb\x77\xf3\x7d\x58\x36\x9f\x1f\x9c\xe5\xb1\xbf\x1d\xdc\xe1\xcb\x07\x29\xcd\xf9\xfe\xee\x8b\xc3\x33\x70\xd1\x68\xe8\x0f\x3e\x05\x2c\x3b\x8c\x58\x63\x5d\x72\x72\x34\x57\x43\x26\x8e\xb6\x9f\xda\x5b\x93\xf0\xfc\xa4\x9e\x61\xe6\x45\x3b\x84\xd8\x88\x0f\xba\xc9\x8d\xeb\xb5\x4f\xea\x90\x57\x6e\xed\x5b\xcd\xab\xe1\xa2\xf1\xd6\xbc\x3a\xa9\xfb\x95\xbc\x3a\x59\xf7\x7f\xe5\xd5\xe8\x6e\xb5\x1e\x6e\x45\x5f\x23\x98\x4e\x74\xfe\x9e\x40\xdf\x13\xe8\xad\x09\xf4\x5f\x00\x00\x00\xff\xff\xf8\xf3\x57\x5c\x29\x0f\x00\x00")

func metaTmplBytes() ([]byte, error) {
	return bindataRead(
		_metaTmpl,
		"meta.tmpl",
	)
}

func metaTmpl() (*asset, error) {
	bytes, err := metaTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta.tmpl", size: 3881, mode: os.FileMode(0644), modTime: time.Unix(1573453598, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd2, 0x4f, 0x4c, 0x64, 0x15, 0xe8, 0x5, 0x83, 0xed, 0xa, 0xe, 0xe5, 0x3d, 0x8e, 0x2a, 0xd3, 0xcd, 0xf7, 0x22, 0xa1, 0x4b, 0xf3, 0x84, 0x82, 0x1a, 0x61, 0xdd, 0xf3, 0x77, 0x7d, 0x68, 0x50}}
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
	"meta.tmpl": metaTmpl,
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
	"meta.tmpl": &bintree{metaTmpl, map[string]*bintree{}},
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
