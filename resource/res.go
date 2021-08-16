package resource

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _resource_default_zz_lua = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\xdf\x8b\xdb\x30\x0c\x7e\xcf\x5f\x21\xfa\xe2\x0d\xdc\x91\x96\x95\x0d\x46\x1f\x76\xec\x18\x07\xf7\xe3\xa5\x83\xed\xe9\xb0\x1b\x25\x35\x75\x94\xe2\x38\xdb\x85\xd1\xff\x7d\xc8\x72\x7a\x5d\xbb\xc7\x05\x22\x7f\x92\x3f\x7d\xb6\x64\x45\x63\x3d\xf6\xb0\x86\xdf\x05\x00\x40\xe8\x7e\x25\x07\x16\x1a\x16\xa5\x86\x25\xff\x2b\x0d\xab\x52\xc3\x87\x15\xc7\x4a\x38\xea\x44\xdd\xee\x4c\xe8\x31\x0a\x5d\x0d\xb1\xfe\xd8\xda\xf7\xb0\xed\xbc\x37\x11\xd7\xd9\x7f\x6e\x90\x30\x18\xff\xbc\x75\x4a\x5f\xb1\x60\x0d\x13\xcf\x3a\x62\x82\x37\xd1\xd1\x82\x91\x75\x64\xc2\xa8\xa6\xd3\x0e\x26\x44\x17\x5d\x47\xd3\x79\x54\x61\x9d\x76\x8f\x9f\x8a\xa2\x76\xe8\xab\xd7\x2a\xe6\x73\xd8\xec\x10\xea\x6e\x08\x70\x73\xff\x74\x03\x71\x3c\x60\x0f\x26\x20\x6c\xee\x1e\x7f\x70\x48\x83\xd8\x87\xdb\x2f\x77\xdf\x1e\x04\x1b\xaa\xe0\xfe\xe9\xf1\x2b\x7b\xef\xae\x84\x36\xb7\xdf\x37\x17\x42\x1c\xd2\x20\x56\x84\x04\x4f\x42\xec\x89\x90\xe4\xa5\x9b\x3b\x8a\x4a\x6a\x7a\xfd\x14\x37\xf3\xcd\xb2\x7c\xcb\x95\xff\x34\xe1\xdc\x95\x46\x9c\x6d\x9e\x07\x4e\x3a\xf3\xf9\x85\x22\xd2\xd0\x72\x42\x8f\xf1\x6f\x5a\xec\xaa\x0e\xf6\x38\x82\xf5\x9d\x4d\x57\x8d\xf8\x12\xe1\x10\xb0\x76\x2f\x57\x7a\x2a\x3a\x1a\x99\xa9\xb4\xca\x4b\x8b\x95\x1b\xda\xec\xf8\x8e\x1a\x81\xff\x4e\x65\x6d\xa5\x55\x5e\x24\x35\x3b\x9c\x2a\x30\xa5\xe6\x87\xee\x5d\x43\xd2\x28\x46\x58\xa5\xb1\xa1\x8c\x27\xd2\x1e\xc7\xdc\xcd\x3d\x8e\xc2\x90\x71\x28\x8e\x45\x51\x99\x68\xce\x47\x01\x80\x7b\x21\xf4\xd9\x38\x63\xfa\xcc\xca\xf2\xd9\x6e\x05\xd0\xe0\xfd\xec\x24\x4f\x43\x6b\x31\xe4\x13\x78\x8b\x39\x5c\x4d\x7a\x3b\x50\x25\x9b\x79\x1a\x53\x31\x29\xb0\x2c\x4f\x02\x7d\x0c\x8e\x9a\x0b\x01\xa4\xc6\xbb\x7e\xf7\x1f\x60\x1a\xfb\x3f\x01\x00\x00\xff\xff\xd0\x21\x4c\x6d\xba\x03\x00\x00")

func resource_default_zz_lua() ([]byte, error) {
	return bindata_read(
		_resource_default_zz_lua,
		"resource/default.zz.lua",
	)
}

var _resource_english_txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x4b\xaa\xc2\x30\x14\x06\xe0\xf9\xbf\x97\x0b\x27\x1d\xf4\x31\xbc\x22\xa9\xf6\x21\xe2\x01\xa5\x99\x99\x4a\xaa\x6d\x5a\xc4\x83\x4a\xb3\x03\x77\xe0\xcc\x2d\xba\x04\xbf\x6e\x69\xc4\x04\x19\x4d\x82\xfc\x14\x38\x18\xae\x43\x82\x66\x1c\x78\x4a\x77\x54\x46\x98\xeb\x52\x36\xe9\x8d\x86\x08\xac\x62\x7a\x66\x8f\xab\x5e\x41\x54\x4c\x87\x6c\xbf\x75\x67\xf4\x5e\x5c\xe1\x2f\xb9\x13\x14\x15\xeb\xbe\x5a\x77\x9a\x31\xdd\xbd\xc7\x8c\x06\x16\x0b\x1c\x6d\x8b\x7f\xdb\x42\x81\xf0\xa7\xa0\x08\x11\xe1\xfb\x79\xbf\x7e\x01\x00\x00\xff\xff\x07\x7d\x13\xd4\x7e\x00\x00\x00")

func resource_english_txt() ([]byte, error) {
	return bindata_read(
		_resource_english_txt,
		"resource/english.txt",
	)
}

var _resource_expression_zz_lua = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x52\xc1\x6a\xdc\x30\x10\xbd\xef\x57\x0c\xbe\x4c\x0a\xda\x45\x4a\x7a\x88\x29\x39\x84\x10\xc8\xad\x90\x04\x7a\x96\xe3\xd9\x8d\x88\x2c\x2f\x92\xdc\xd6\x04\xff\x7b\x19\x49\xb6\xd7\xf4\x90\xdb\x93\xde\x1b\xcf\x7b\xcf\x8a\xba\xb1\x14\xe0\x0e\x3e\x77\x00\x00\xbe\xff\x93\x0e\x4a\x80\x92\x02\xae\xa5\x80\x1b\x29\xa0\x96\x93\x48\xfc\x7e\x0f\x2f\x4f\x3f\x7f\xc1\xc3\xd3\xfd\xf3\xfd\xc3\xeb\xe3\x33\xbc\x3c\xbe\xfe\x48\xd4\xdb\xbb\xf6\x81\x62\x1a\xc7\x21\x1e\x6f\x51\x20\xae\x63\x67\xed\xa3\x89\xa6\x77\xe0\x86\xae\x21\x9f\xee\x97\xcb\x32\xe5\x5a\x3a\xf2\xcc\xb4\xdb\x1d\x0d\xd9\x76\x35\x16\xc7\x73\xb6\x89\xc6\x45\x14\x80\xd1\xb8\xb1\xc0\xd0\x69\x6b\x0b\x6e\xcc\xa9\xa0\x96\xde\x4c\xa7\xed\xd5\xf7\x94\xe3\x1b\x66\x27\x80\x47\xdb\xeb\x2c\xe8\x87\xc6\x12\x23\xb6\x7e\x95\x34\x80\xbf\xb5\xbf\x3c\x92\x1b\xba\xb4\x83\xf2\x8c\x8e\x14\x4d\x47\xf3\xd7\x00\x9b\xbe\xb7\x28\x80\x0d\x75\x14\xa2\xee\xce\xac\x1b\x49\xfb\x59\xcf\x81\x92\x3a\x98\x93\x4b\x11\x18\x50\xcb\xfc\xe0\x0a\x2e\x45\x7d\xd0\x18\x58\xf1\x41\x23\x4e\x5c\x43\xab\xa3\x5e\x4a\xc8\xcd\xe5\x1a\xdc\x60\xed\xa6\x87\x62\x08\x00\xd5\xf5\xa1\xae\x15\x93\xea\x20\x65\xcd\x60\x5f\x1f\xd4\xed\x0d\x0a\x94\x28\x70\xaf\x50\xa0\xc2\x34\x50\xf6\x72\xcc\xbb\x4f\xac\xc6\x0a\x05\x56\x4d\x95\x04\x02\xab\xbf\xe9\xcc\xbb\xaa\xd9\x22\x70\xe2\xf2\x46\xa4\x00\xe6\xa6\x65\x39\xe7\xde\xd8\x4b\x45\xac\xfc\xdc\xdf\x46\xb3\x94\xba\xea\x96\x32\xbf\x12\xf2\xd5\x7f\x9a\x0b\x3e\x44\x6f\xdc\x69\x5b\x99\xa5\x18\xc9\xe7\xbf\x7b\xb2\x26\xbc\x97\xb4\x32\x3f\xbd\x7f\x01\x00\x00\xff\xff\x89\x82\x00\xfe\x13\x03\x00\x00")

func resource_expression_zz_lua() ([]byte, error) {
	return bindata_read(
		_resource_expression_zz_lua,
		"resource/expression.zz.lua",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"resource/default.zz.lua": resource_default_zz_lua,
	"resource/english.txt": resource_english_txt,
	"resource/expression.zz.lua": resource_expression_zz_lua,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resource": &_bintree_t{nil, map[string]*_bintree_t{
		"default.zz.lua": &_bintree_t{resource_default_zz_lua, map[string]*_bintree_t{
		}},
		"english.txt": &_bintree_t{resource_english_txt, map[string]*_bintree_t{
		}},
		"expression.zz.lua": &_bintree_t{resource_expression_zz_lua, map[string]*_bintree_t{
		}},
	}},
}}
