// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/members.tpl
// template/pkg.tpl
// template/type.tpl
package main

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _templateMembersTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\x5d\x8b\xd4\x30\x14\x7d\x9f\x5f\x71\x08\x08\x6d\xa1\xf9\x0f\xe2\x07\x88\xae\x2b\xe2\x9b\x08\x4d\x37\xb7\x6e\xb4\x4d\x6a\x92\x7d\x18\x32\xf9\xef\x92\x64\x3a\x4e\xa7\x5d\xb7\x0f\xa1\xb9\x1f\xe7\x9e\x73\x6e\x42\x80\xa4\x41\x69\x02\x9b\x68\xea\xc9\x3a\x86\x18\x0f\x87\x10\x60\x85\xfe\x49\xe0\x77\x25\x9c\xa2\x21\x40\x0d\xd0\xc6\xa3\x7a\x54\x52\x92\x2e\x39\xf0\x3a\xc6\x03\xce\xdf\x09\x5d\x08\x18\x14\x8d\xf2\xb3\x98\x08\x1c\x6d\x8c\x1d\x42\xc0\x6c\x95\xf6\x03\xd8\x2b\xc7\xc0\x4e\x68\x9a\x86\xa5\xdc\xa5\x33\x7d\x65\xc4\xa8\xf4\xef\xf7\xc6\x7e\x3b\xce\x04\x9e\xcf\x9b\xb2\x52\xda\xae\x21\xbf\x67\xb8\x14\xf6\xc7\x99\xde\x2a\x37\x8f\xe2\x58\x38\x9c\x31\x7e\x54\x21\x6c\xc1\x4b\xcf\x0a\xaa\xde\x65\x46\xa3\x7b\x96\xc9\xfe\xc8\x2d\x48\x0b\xd2\x72\x37\xbe\x22\x90\xcd\x89\x71\x9b\x38\xad\x99\x15\xbf\xb2\xdb\xef\xa6\x9e\xa4\x24\x09\xbe\xc7\xb1\x5a\xf6\x68\x86\xb4\xa0\x76\xbb\x21\x61\x09\xb4\x60\x28\xed\x0d\xfc\xa3\x72\x59\x18\xaf\x0f\xff\x53\x50\x48\x28\x77\x3f\x7b\x65\xb4\x18\x97\x67\xf1\x92\xc8\xf4\x04\xaa\xa5\xa9\x6e\x9a\x06\xb7\xe2\xfe\xcd\x5a\x05\x2d\x69\x49\xf6\x8d\x99\x26\xd2\xde\x81\x9f\xff\x3e\x29\x4d\xee\x82\x50\x58\x09\x2d\x51\xd1\x1f\x54\x79\x21\x3c\x09\xce\x47\x0d\x76\xdf\xff\xa2\x07\x7f\x47\x5e\xb0\x1a\x37\x63\x57\x3c\xbf\xd2\x40\x16\xd9\x10\xc2\xc7\xa7\x9e\xac\x26\x4f\x0e\xaf\xbf\x7c\x80\x34\x0f\x4f\x69\xb6\x48\x2a\x30\x18\x9b\x8b\xb2\xb9\xd9\xeb\x74\xeb\x26\xf2\x42\x0a\x2f\xba\x92\xe0\xec\x8a\xe4\x45\xe2\xb3\xbb\xbe\xb6\x61\x65\xc9\xf5\xe5\x6f\x00\x00\x00\xff\xff\xc7\x7f\x5f\xfc\xc9\x03\x00\x00")

func templateMembersTplBytes() ([]byte, error) {
	return bindataRead(
		_templateMembersTpl,
		"template/members.tpl",
	)
}

func templateMembersTpl() (*asset, error) {
	bytes, err := templateMembersTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/members.tpl", size: 969, mode: os.FileMode(420), modTime: time.Unix(1580009840, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatePkgTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x41\x6f\xc2\x30\x0c\x85\xef\xfd\x15\x4f\xe5\x02\x91\x9c\x5f\xc1\x79\x97\x4d\x3b\x93\x51\x83\xa2\x95\x10\x35\x19\x12\x4a\xfc\xdf\xa7\x9a\x96\xb6\xbe\xd8\xfe\xda\x97\xe7\x57\x0a\x3a\xbe\xf8\xc0\x68\xa3\x3b\xff\xba\x2b\xa7\x16\x22\x4d\x53\x0a\x06\x17\xae\x0c\x3b\xf3\x11\x03\xc0\xfb\xcb\xfe\xe1\x93\xff\xe9\xf9\xeb\x19\x39\x61\x9f\xee\x43\xe6\xee\xb5\x58\x6d\x87\xc3\x24\x99\x64\x99\x6f\xb1\x77\x99\xd1\xe6\x67\xe4\x16\x16\xab\x37\x39\x74\xf3\x46\x44\xcd\x42\xc6\xd1\x5f\x60\x13\x87\xe4\xb3\x7f\xf0\xd1\x65\x27\xd2\xec\x76\xf8\x9c\x11\xbe\x5d\xff\xc7\x49\xd5\x15\x1f\xee\xc6\xa8\x18\x6f\x40\xc5\x91\xd3\x79\xf0\x31\xfb\x7b\x40\x7d\xfd\x41\x5a\xdb\x36\x55\xdd\x66\xdc\xba\x62\x95\xa7\xe2\x54\x0a\xc1\xaa\x19\x89\x9c\x50\x61\x8c\x51\xa6\xce\x24\x62\x8c\x41\x85\xa2\xf5\x15\x24\x82\xd9\x87\x34\x26\x89\xac\x12\x2f\xd3\x7f\x00\x00\x00\xff\xff\x6d\x8c\xee\x9b\x9f\x01\x00\x00")

func templatePkgTplBytes() ([]byte, error) {
	return bindataRead(
		_templatePkgTpl,
		"template/pkg.tpl",
	)
}

func templatePkgTpl() (*asset, error) {
	bytes, err := templatePkgTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/pkg.tpl", size: 415, mode: os.FileMode(420), modTime: time.Unix(1580009840, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateTypeTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xcb\x8e\xd4\x30\x10\xbc\xe7\x2b\x4a\x5e\x90\x66\x24\xe2\x0f\x58\x69\x0f\x2b\x2d\xcb\x01\xc1\x01\x01\x17\x84\x14\xb3\xe9\x0c\xad\x4d\x1c\x63\x9b\x47\xd4\xf1\xbf\x23\x27\x99\x99\x24\xe0\xc3\xa8\xa7\x5c\x55\x5d\xa9\x44\x04\x35\x35\x6c\x09\x2a\x0e\x8e\x14\x52\x2a\x8a\xe2\xe6\x06\x22\xd0\xef\x4d\x47\xd3\x0f\xca\x94\x0a\x00\x19\xe5\x06\xf4\x03\xfa\x2d\xdb\x1a\xea\xbe\x65\x13\xb2\xe8\x50\x89\xe8\x4f\xb6\x26\xdf\x0e\x6c\x4f\x29\x55\x30\xf9\xee\x28\x02\xb2\xf5\x64\x20\x82\xdf\x1c\xbf\xe3\x90\x37\x7d\xa0\x86\x3c\xd9\x27\x0a\xd0\x47\x2c\xf6\xf9\xdc\x3b\x47\xc6\x07\xf4\xf6\x16\x22\x25\x5e\x38\x4f\xbf\x70\x7b\x07\xa5\x2e\x31\xe6\x28\x25\xbc\xb1\x27\x82\xde\xe0\xe7\x3b\x6e\x16\x69\x99\xd2\x2b\xac\x62\x6c\x89\x0b\xe9\x0e\x1a\xff\x31\x71\x9e\x6d\x6c\xa0\x5e\x06\x05\xf5\x65\x0a\x20\x82\x9c\xff\x81\x83\x6b\xcd\x30\xb5\x93\xa5\xff\xb0\xbf\x4e\xec\x83\x08\x5a\xb6\xcf\x8f\xbd\xff\x38\xb8\x99\x7a\xdc\x3c\xc3\xaa\x9d\x3c\xe6\xfe\xe7\x96\xf5\x3b\xea\xbe\x91\x0f\xfb\x5c\x23\x1e\x99\xda\x1a\x23\x26\xc7\x11\x0f\x14\x9e\x3c\xbb\xc8\xbd\xc5\xb8\xa3\x96\xd3\x39\x0f\x17\x60\xf9\xb7\x0a\x92\x37\x72\x78\xfd\xc7\xf5\x3e\x52\x7d\xc9\xba\xb3\xab\x8c\xe3\xcf\xe4\x03\xf7\xb6\xc2\x88\x10\x3d\xdb\x53\xc6\x45\x8c\xe3\x37\xbe\xff\xe9\xa0\xf3\xcb\xdf\xe7\x00\x50\x3d\xb3\xad\xf7\xaa\xeb\x37\xb6\x55\x5d\xdb\x58\x21\x91\x3a\xd7\x9a\x48\x50\xdd\x5c\x8d\xca\xcb\x8a\x5d\x75\xcb\xf8\x37\x00\x00\xff\xff\x38\x1b\xc7\xd9\xda\x02\x00\x00")

func templateTypeTplBytes() ([]byte, error) {
	return bindataRead(
		_templateTypeTpl,
		"template/type.tpl",
	)
}

func templateTypeTpl() (*asset, error) {
	bytes, err := templateTypeTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/type.tpl", size: 730, mode: os.FileMode(420), modTime: time.Unix(1580009840, 0)}
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
	"template/members.tpl": templateMembersTpl,
	"template/pkg.tpl":     templatePkgTpl,
	"template/type.tpl":    templateTypeTpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"members.tpl": &bintree{templateMembersTpl, map[string]*bintree{}},
		"pkg.tpl":     &bintree{templatePkgTpl, map[string]*bintree{}},
		"type.tpl":    &bintree{templateTypeTpl, map[string]*bintree{}},
	}},
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
