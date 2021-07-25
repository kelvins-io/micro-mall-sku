package proto

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

var _proto_ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x99\xbf\x4e\xc2\x50\x14\xc6\xbf\x5b\x18\x9a\x30\xd8\x41\x85\xc5\xa4\xf1\x0d\x78\x83\x4a\xea\xa0\xa3\x8c\x26\x22\x50\x23\xc4\x36\x6d\x8a\xec\xc4\xd9\x07\x70\x72\x37\x2e\x6e\x46\x27\xe3\x62\xe2\xe6\xe8\x13\x38\xb8\xf8\x06\x9a\xde\x7b\x88\x25\x70\x35\x57\x34\x9a\x7a\x7e\x89\xfe\x86\x72\x3e\x7a\x5a\xb8\x7f\x0a\x00\xd1\x18\x06\x75\xa0\x06\xc0\x86\xb2\x53\xc1\x4c\x6c\xfa\x9b\xc2\x22\x3b\x59\x9e\x54\x8c\x16\x22\xb4\x11\x22\x9c\x9d\xc5\x30\x0c\xc3\x30\xcc\x2f\x23\x94\x6c\xcd\xbc\xcf\x30\xcc\x3f\x26\x1b\x1f\x5c\xb2\x47\x1e\x29\x0b\x3a\x6e\x91\xcb\xb9\x1a\x87\xec\x92\x3d\xf2\x48\x59\xd0\xeb\x2c\x72\x99\x6c\x93\x1d\xb2\x4b\xf6\xc8\x23\x65\x1a\xb4\x04\x6d\x3e\x04\xbd\xb3\xa0\x1d\x8a\x70\xc8\x2e\xd9\xfb\x99\x6b\xc3\x30\x45\x20\xfb\xba\x54\x11\xa1\x8f\x2e\xd2\x89\xfd\x7b\x0b\x03\xec\xa1\x8d\x14\x5d\xf4\xd0\x42\x22\x8f\x1f\x22\x0e\xf7\xeb\xcd\x6e\x1c\x25\x2a\xe0\x66\xd5\x3c\x21\x8a\x7d\xbf\x13\xc6\x9d\x6c\xf1\xb1\x71\x77\x72\x5a\xbd\xb8\x5d\xfb\x4a\x4a\xf0\x0d\x29\x49\x2f\xdf\xcd\x4b\xf6\x6f\x49\x9b\xd0\x43\x8c\x44\x7b\x2d\xc4\xe6\xa3\x59\x75\xfe\x3a\xb8\xaf\xcf\x3b\xdb\x67\xb2\x03\xa3\x84\x60\xce\x84\x89\xfe\x85\x1c\x64\x17\xb5\xd5\x07\x18\x6a\xbb\xb7\x9e\x8e\x4d\x6a\xf3\xbd\xdf\x5f\x9f\xaf\xa4\x97\xf2\xcc\x0d\xea\x83\xb9\xea\x27\xfa\x2e\xc9\xd9\x62\x59\x53\x3b\x94\x9f\x9d\x14\x83\x0f\xee\xfc\x95\x69\x7d\xbe\xff\x85\x87\xca\x56\xff\x48\x9e\xbf\x61\x46\x30\x77\xc6\x8c\xfb\x3f\x86\xe7\x7f\x86\x29\x2e\x25\xa5\x5a\xb6\xff\x5f\xd7\x3f\xff\x67\x18\xa6\xc0\x88\xb2\xdf\xf4\x1b\xef\x0f\x04\xa7\xb0\x68\x21\xb0\x3b\x2e\xf8\x64\x21\x20\x72\x3f\x18\xf2\x42\x80\x61\xfe\x1e\x6f\x01\x00\x00\xff\xff\x03\xcb\x5c\xf2\x04\x20\x00\x00")

func proto_ds_store() ([]byte, error) {
	return bindata_read(
		_proto_ds_store,
		"proto/.DS_Store",
	)
}

var _proto_micro_mall_sku_proto_license = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\x6d\x8f\xe4\x36\x72\xfe\x2e\x40\xff\xa1\xd0\x40\x70\xd3\x80\x76\xb0\x6b\x9f\x2f\xc9\x39\x08\x30\xf1\xae\xed\xc9\xad\x77\x9c\x99\xd9\x2c\xf6\x23\x5b\x2a\x75\xd3\x4b\x91\x3a\x92\xea\x5e\xfd\xa4\xfc\x8d\xfc\xb2\xa0\xaa\x48\xbd\x74\xf7\xf8\x82\x20\xf7\x69\xa6\xf5\x42\xd6\xeb\x53\x4f\x15\xf5\xd0\xa3\x85\x27\xd7\xc6\x93\xf2\x08\xef\x75\x8d\x36\x20\xdc\x6c\x1e\x9e\xde\x6f\xb6\x70\x84\x6f\x6f\x5f\x97\x45\x59\x3c\x1f\x74\x80\x17\x9e\x8d\x07\x84\x4d\xfa\xb5\xd9\x82\xea\x7b\xa3\x31\x40\x74\xa0\xec\x08\xce\xeb\xbd\xb6\xca\xc0\xc9\xf9\x2f\xe0\xda\xb2\x50\x43\x3c\x38\x1f\x0e\xba\x4f\xef\x3e\xe4\x47\x3e\x39\xff\x65\xb3\x85\xd3\xc1\x05\x04\x77\xb2\xe8\x57\xab\x3b\xbf\xd9\xc2\x41\x05\xe8\x8d\xaa\xb1\x81\x78\xc0\xb2\x68\x9d\x31\xee\xa4\xed\x1e\x0c\x3f\x44\xff\x59\x17\x75\x8d\xa0\x9a\xdf\x54\x8d\x36\x92\x28\xb4\x4c\xed\xfa\xd1\xeb\xfd\x21\xe6\x07\x5a\xe7\xf9\x46\x16\xa0\x2c\x48\x82\x3f\x93\xc2\x49\x9f\x06\x06\xdb\x60\x7a\xea\xaa\xfa\x47\xf4\x41\x3b\x9b\x0d\xf5\x66\x0b\x3f\x79\x65\x23\xb8\x16\x7e\x98\xf6\x4b\x0f\xdf\x42\x56\x04\xf6\xf4\x50\x80\xcf\x6e\x00\x45\xa6\x31\xcd\x49\x37\x58\x81\x77\xa3\x32\x71\x7c\xd5\x7a\xc4\xaa\x2c\xac\xb3\xaf\xf0\x6b\x6d\x86\xa0\x8f\x58\x41\x18\x76\xa2\xa5\xda\x19\x4c\x0a\x63\x35\xe9\xd1\x0c\x5e\x45\x12\xc6\xb5\x6b\x85\x2b\x32\x41\xe3\xca\x82\xae\x4e\x16\x63\x45\x01\xd4\x96\xee\x7a\xec\xbd\x6b\x86\x1a\x57\x16\x61\x97\x80\xb6\xb4\x94\xc6\x50\x01\xea\x78\x40\x0f\xca\x38\x8b\xe0\x3c\x90\x37\x94\x67\x75\x15\x2d\x56\x3b\x63\xb0\x8e\xfa\x88\xec\xf0\xef\x65\x8b\x1d\x6f\x11\xbd\xb2\xc1\xa8\x88\x15\xa8\x46\xf5\xb1\x02\x65\x22\xfa\x4a\x6e\xb4\xce\x77\x15\x74\xae\xd1\xed\x58\xf1\xd2\xde\x2b\xbb\xc7\x33\x0f\x01\x8b\x54\xd1\x55\x8f\xbb\x11\x6a\x8f\x2a\x92\xd3\x1b\xf4\xfa\xa8\xa6\x9d\x03\xdc\x6c\xde\xce\x97\xe8\xa5\xb0\xd9\xc2\x4e\xb1\x53\x7b\x67\x25\x7c\x60\xad\x6a\x92\xb7\x66\x79\x1b\x1d\xa2\xd7\xbb\x21\xb2\xa6\xb5\xeb\xba\xc1\xea\x5a\x45\x4c\xd6\xc8\x66\x5e\x1b\x4b\xd9\x06\xe6\x8d\xb3\xc0\x21\x07\x61\xcf\x2e\xac\xe0\xa4\xe3\x41\x2e\x78\x77\xd4\x81\x6e\xaa\xb8\x58\x78\xbd\xa8\xf3\xb4\xd0\xb9\x3e\xf2\x0e\xc5\xd0\xcb\xa2\x86\x83\x32\x06\x76\x53\xb4\xa4\x88\xa6\xe5\xe2\x8b\x49\x9d\xcc\xd0\xb0\x19\x7a\xf4\xe4\x9b\x2b\xaa\x8a\x2e\x66\xfc\x9e\x94\x96\x57\x30\x5b\xae\x37\x6a\xfc\x9d\x57\x6e\xe9\xf1\x6f\x16\xc9\xf2\xab\x8a\x94\xaa\x7f\xe7\x4c\x91\x6c\xee\x65\xaf\xda\x28\xdd\x05\x06\x9a\x46\xcc\x66\xa3\xa7\xf8\x6d\xca\x62\x27\xc2\x4f\x52\xb0\xa5\xc9\x42\xd8\xed\x5c\xa3\xb1\xa1\xa4\xb8\xe2\xfc\x00\xed\xe0\xad\x0e\x07\x6c\x40\xd6\xc8\x58\xe2\xfc\xcb\x89\x2a\x02\x05\x4e\xd3\x4e\x7d\x21\x41\x49\xda\x80\xc6\x54\xe0\xda\x16\x3d\xbf\x1a\x94\x21\x55\x0f\xea\x88\xd0\x29\x32\x02\x05\x9b\xee\x7a\xe7\xe3\xdf\x8c\x44\x09\x19\xb6\xfb\xb7\x0b\xbb\x3f\xb9\xc1\xd7\x08\x3f\xb8\x06\x67\xe3\x3f\x1f\x10\x22\xfa\x0e\x36\x8b\xdb\x1b\xe8\x50\xd9\x90\x82\x16\x5b\xf4\x9e\x2c\xc5\xb1\x71\x35\x13\x48\xe4\x4e\x7d\xa1\xe4\xe4\xb4\xa6\x80\xd4\xce\x72\x2a\xe8\xc8\xe2\x51\x6c\xaa\xa3\xd2\x86\xdc\x54\x16\x8d\xab\x87\x0e\x6d\x14\xe3\x34\x18\x6a\xaf\x77\xf4\xfe\xc1\x9d\xd8\x36\x8c\x0e\x97\x5b\x2d\xc2\x45\xed\x3d\x72\xf5\x29\x0b\xce\xac\x06\x41\x41\xa7\xea\x83\xb6\xf8\xca\xa3\x6a\x38\x20\x08\x19\xb3\xd0\x4b\x0b\x5c\xcf\x68\xe3\xec\xbe\x2c\x38\x63\x51\xd5\x87\xd5\xdb\xeb\x47\x39\x4c\x26\x59\xe6\xac\x0c\xb3\x84\x65\xe1\x31\xa0\x3f\xa2\x58\x52\x0a\x44\x74\x10\x54\xd4\x81\x95\xd3\x01\xdc\xce\xe8\xbd\x58\x61\x37\x72\xbd\x23\x2b\x5c\x2a\x52\x16\x2f\x69\xa2\x2d\x28\x0b\xda\x92\x7b\x64\x21\x8f\xbd\x0b\x3a\x3a\x3f\x82\x47\x15\x9c\x55\x3b\x33\x42\xad\x4c\x3d\x10\x28\x37\x62\x33\xf4\x9d\x8e\xa0\x2d\x7e\xed\xa9\x9a\x1e\x91\xfd\x54\x3b\x7b\x44\xab\x29\x6b\x54\x5d\x63\x08\x24\x16\xa5\x64\x2b\x35\x80\x2c\x44\x7f\x67\x25\x29\x99\xb4\x1d\x84\x08\x2c\xe0\xe9\xd2\x77\x14\x91\x7f\xdc\xc2\x3b\x49\x5e\x8a\x8f\x1f\xbd\xeb\xa6\x02\xcb\x91\x7a\x0b\x1f\x52\xe5\xa1\xf7\xad\xea\x04\x24\xe7\xcc\xb2\x6e\x71\xab\x2c\xa8\x1c\xd9\x51\x32\x9a\xf6\x75\x7e\x02\xe0\xd5\xe6\xf2\x22\xf3\x14\x36\xa1\xf6\x54\x8c\x1a\xec\x14\xc1\x2b\xe9\x41\xae\x22\xae\xc0\x57\x2a\xe8\xd4\x48\x60\x3a\x04\xb6\x17\xa0\x6d\x9c\x0f\x0c\xba\xbd\x77\x9d\x8b\x8c\xe7\xcd\x50\xc7\x20\x05\x09\x1b\x68\x49\x1b\xf2\x6a\x59\xac\x83\x85\x02\xca\x0d\x11\xf0\x6b\xef\xc9\xa4\xbd\xd7\xb4\x0e\x79\x20\x84\x05\x3a\x64\x2d\x6f\xc9\x46\xd8\x47\x50\xa1\x2c\xd2\x4b\x66\x84\x10\xd9\x7b\x54\x0f\xb5\x25\x85\xe2\x81\x82\x85\xf1\x49\x87\xc9\x8e\x09\x47\x49\xd7\x04\x87\xec\xf0\xbc\xf8\x1f\xc2\x42\xf3\x6a\xa6\x0e\xa1\x5a\xa0\x13\x3d\x00\x01\x6b\x8f\x31\x40\x36\x1c\xb9\xa5\x2c\xb4\x8d\xc8\xa5\x7f\x50\x86\x6c\xd0\xa3\x8f\xe3\x2d\x7c\x70\x19\x6d\xf3\xa6\x3a\x88\x28\x62\xc0\xeb\x68\xc7\xa8\x22\x70\x07\x0b\xb4\x73\x3e\x83\x9d\xa0\x30\x41\x05\xc7\x01\xc9\x71\x06\xea\x29\x58\x54\xaa\xf3\x53\xf5\x4b\xf7\x1b\x6c\xb5\x15\x14\x7f\x22\xbe\xe2\x2c\x7c\xc3\xd2\x5e\x17\x93\xdc\x30\x9b\x87\xe3\x6b\xca\x72\x3c\xa2\x05\xdd\x42\x18\xea\x83\x84\x09\x17\x0a\x6d\x6b\x33\x34\x2f\x14\x0a\xda\x2a\xbb\x89\x58\xd9\xc2\x4f\x53\xc1\x26\x8b\xfa\xde\x63\x12\xa1\xf7\xee\xa0\x77\x7a\x01\x2f\x1c\x58\x33\xeb\x4d\x75\x9d\x70\x3b\x40\xa3\xc9\x8e\x64\x90\x29\xfc\xa6\x1d\xc8\x5a\xbf\x07\x5b\x6c\xbb\x93\x0e\x58\x16\x27\x37\x98\x46\x5c\xa0\x66\xa8\x4a\x36\xe2\xcc\xfd\x8e\x32\x37\xa2\xa7\xb5\xde\x62\x6f\xdc\x48\x5e\x59\x96\x90\x2b\xb7\x97\xa5\x84\x7c\x5f\x16\x13\x42\x68\x67\xab\x35\x87\x59\x64\xc2\x39\x29\xba\x64\x44\x0c\x7b\x63\x59\x9c\xd4\x28\x0e\x61\xd5\xfe\x77\xef\x2e\x73\x7b\x37\xd2\x3a\xce\x62\x59\x2c\x42\xe9\xb3\x1b\x2a\x38\x1d\x30\x5d\xa1\x36\x45\xb8\x26\x39\x7c\x06\xb9\xe6\x8c\x84\xa5\x10\x72\x64\xd1\x1e\x7d\x20\x8c\xe3\xda\x48\xc5\x29\x57\x3f\x82\x4f\x65\xa5\x7d\x4a\x5a\x53\x04\x58\x0a\x21\xca\x86\x81\x60\xe6\x48\xdc\x1b\x2c\xc6\x93\x20\xe7\x1d\xbf\x93\x01\xa4\x76\xb6\xd1\xfc\x66\x26\x1a\x29\xe9\x5d\x3b\xc5\x35\xc1\x04\x47\x4a\xc5\xf8\xdd\x0d\x21\x42\x24\x0e\x5d\x16\x14\x17\x57\x9c\x95\xa1\xfe\x25\x1f\xa8\x73\x4b\xb2\x26\x6b\x97\x26\xe2\x15\x52\xaa\xbd\xb9\xa9\xb7\x1c\x3d\x7f\xda\xc2\x5d\x9c\x9f\x7a\x64\xc4\xb9\x9d\x25\xf3\x18\x15\x81\x5a\xca\xa1\xb3\x52\x4d\x02\x2f\xa9\xf6\x19\x29\xe6\xd6\x80\x48\x92\x31\xcb\x56\x48\x80\x82\x83\x6c\x4a\xe9\xd4\x11\x86\x9c\x2d\x58\x16\x7f\x8b\x16\x54\xa4\xe6\x09\x89\xc0\x08\xaa\xe6\x25\x26\x63\x73\xc9\x26\x92\x43\x01\x29\x64\xa6\x67\x23\x45\xfc\x1a\x41\x37\x68\xa3\x6e\xb5\x34\xb2\x84\xdd\x29\x02\x36\x4b\x8b\x7c\xe0\x45\x6f\x37\xb3\x49\x6a\x35\x50\x1c\x9d\x9b\xa3\x4d\x70\xfc\x72\x8f\x20\xe6\xa0\x50\xac\x95\xf7\x23\x28\x2e\x59\xda\xa2\x8d\x65\x71\xb9\xe7\x8b\x24\x21\x91\x0a\xf0\x58\xeb\x5e\x33\x04\x4f\x7b\x10\x52\x94\x85\xb0\x3d\x51\xec\x4a\xb1\xff\xc7\x2d\x7c\xe2\xe6\x2e\x72\xd9\xfd\xd5\xbb\x23\x5a\x65\x6b\xa1\x1a\x6f\x75\x60\x8c\x46\x4f\x37\xf3\x83\x0b\x86\x77\x92\x4b\xb2\xab\x18\x62\xee\xec\x39\xff\x9b\xab\xb5\x5e\x6e\x4c\x6c\x5b\xe0\x6c\x46\x79\xf1\x01\x33\xff\x99\x4c\x7a\x4c\xed\xc1\x79\x3f\x90\x6e\x4e\x6d\x86\x58\x86\x4c\x30\x0d\x0c\x32\x14\x73\xf0\x2c\xe0\x77\xee\xfc\x2e\x2a\xfd\x82\xaf\xdc\x84\xad\x5c\x65\xd4\x98\xcb\x31\x6b\xb1\xd2\x60\x41\x0b\xe0\x82\x15\xa4\xc4\xd1\x5d\x87\x8d\x56\x11\xcd\x48\xbc\x18\x6b\x6c\x28\x3a\x03\x12\xc2\xd4\x58\x5d\x6b\xfa\x03\x24\x02\x3d\x0f\x41\x16\x5a\x38\x2b\xb1\xfa\x54\x16\xf7\x4f\x1b\xf8\xb7\xbb\xa7\xfb\x27\x96\xed\xd3\xfd\xf3\xcf\x0f\x1f\x9f\xe1\xd3\xdd\xe3\xe3\xdd\x87\xe7\xcf\xd3\xcc\x20\xa3\x94\x54\x71\xa3\xb1\xa9\x52\x91\xd4\x76\x5f\x65\x36\x44\xa5\xba\xd3\xd2\x01\x88\x58\xc9\xdf\xa9\x2f\xa6\x3e\x4f\xdb\xd6\x6b\xbb\xc7\x8e\xf3\xb8\x43\x5f\x1f\x94\x8d\x6a\xa7\x8d\xa6\x90\xf2\xd0\xea\x68\x69\xab\x96\x18\x9c\xe2\x09\x85\xa6\x00\xf6\xd0\x0f\xbe\x77\xdc\xe2\xfc\xfc\x0e\xde\x7d\x78\xbe\x7f\x7c\x07\x8f\xf7\x4f\x7f\x81\xbb\x27\x78\x7e\xe0\xab\xff\xf1\xf1\xee\xfd\xfd\xf3\x67\x78\xf8\x91\x7f\x3e\x3c\xde\xff\x74\xff\xe1\xee\x3d\x7c\x7a\x78\xfc\x0b\x29\xcb\x1a\xc2\xe7\x87\x8f\x54\xe4\x74\x80\xb7\xf7\x4f\x3f\xbc\xbf\xbb\xff\xe5\xdd\x23\xbd\x92\xd5\x26\x5f\x86\xa8\x23\x91\x7f\xc6\xe8\x40\xc6\xd6\x44\x8e\xd2\xbc\x44\x28\xe1\xd4\x73\x2d\xb8\xc7\xd5\xe8\x5d\x10\x92\xdd\xb8\xf6\x05\xb2\xff\xcb\x62\xe1\xa6\x66\xca\x22\x4e\xb9\x7f\xda\xc2\xfb\xc9\xac\xc2\x9c\x93\xb9\x6e\xe1\x23\xbf\x65\x1d\xd4\xda\xd7\x43\x17\x22\xa5\xa2\xc4\xd9\x90\x6f\x19\xdc\x2b\x43\x42\x39\x3f\x56\x65\x91\x4b\x20\x45\x17\xd1\xb1\x9b\xc9\x8f\x60\x71\x6f\xf4\x9e\x82\x6a\x5b\x49\x3c\xab\x5a\xc0\x76\xa2\x15\x55\xe2\x38\xab\xe6\x58\x46\x14\x5c\x08\x65\x72\xe8\xec\x0c\x6a\xda\x36\xda\x23\xad\x13\x7a\xac\xb5\x32\x1c\x39\x8c\xa0\xf4\x7f\xea\x37\x02\xfe\x75\x48\x36\x6e\x54\xa7\xf6\x38\x71\xc3\xfa\xa0\x48\x0c\x2a\x9e\x5e\x0b\x2e\x07\x22\x34\x18\x06\x13\x2f\x12\x94\x96\x4b\xcc\xe4\x7a\xd1\xbb\x8c\x5a\x58\x06\x6d\xde\x9c\xa4\x37\x2e\x08\x5d\xdc\x3b\xd7\x9c\x34\x31\x5c\x1e\x84\x86\xe8\xfa\x5e\xed\x91\x6c\xd4\xf5\x03\x89\xd6\x2a\x6d\x06\x8f\xc2\x0c\x4c\x3b\xd8\x7a\x22\x43\xa4\x43\x6a\x97\x33\x1f\x21\x7a\x81\xbe\x5e\x29\x2b\xdb\x51\xb3\xc9\x81\x69\x56\x1e\x37\xd9\xe3\xc9\xfa\xd6\x51\xc9\xef\x7b\x33\xe6\x78\xc3\xaf\x8c\x2a\x89\x81\xf0\xdc\x44\x9d\x26\xde\x19\x84\x4f\xcd\x8b\x72\x60\xfd\xf3\x16\xee\x6a\x8a\xbe\x09\xbf\x9f\x09\xd5\xac\x3c\x01\xf7\x6d\x05\x2a\xb2\xfc\x51\x77\x28\x9c\x63\x06\x29\xc5\x49\x91\xf9\xd1\x22\x1d\xaa\x34\x70\xe1\xfb\xec\x7e\xe2\x51\x01\x46\x37\x78\xa8\x0d\x2a\x2f\xa3\x0f\xef\xf1\xe8\x44\x56\x35\x8b\x71\x9e\x5c\xd3\xa8\xc1\xb5\x40\x8a\x08\x2e\xa7\xbe\x56\x18\x53\x20\x59\x7f\x7f\x94\xc6\x28\xfc\xe2\xd8\xef\x1a\x11\x5a\x50\x2c\xea\x73\x38\xe2\x52\x4d\xc5\xb2\xc0\xb6\xa5\xdc\x99\x07\xcc\xeb\xfc\x8b\x0e\xdc\x8e\xd8\x4f\xf2\x8d\x80\x67\xb2\x88\x6b\x57\xe5\x37\x4d\x76\xaf\xd6\x9b\x14\x0c\xd9\x14\xb3\xc6\x62\xcc\x54\x55\xa2\x13\x41\xa2\xfa\x92\xc8\x85\xaa\xa3\x3e\x6a\x86\x5d\xa3\x43\x5c\xf7\x4a\x6f\x16\xd0\xbd\x58\x87\xc9\x85\xd0\x8c\x4b\x5e\xb6\x1e\xbd\x5e\x90\x29\xdb\x40\xe3\xb8\x28\xb9\x29\xab\x0e\xce\x3a\x42\x7a\x7a\x21\x24\xfd\x88\xa7\x2e\x1d\x97\x0a\x15\xc7\xa8\x60\xe3\xcc\x06\x28\x7e\x39\x4e\x2c\xb3\x5a\x7a\x9e\x60\x8c\x24\x1c\x6f\xcb\xe2\xc3\x0b\xad\xb2\x0e\x33\xf5\x26\x30\x6a\x5b\xac\xe3\x62\x59\xc1\x5b\xde\x9c\x07\x9f\x73\x4e\x84\x25\x14\x6e\x5a\xa5\x99\xb7\x6f\x28\x3a\xe4\x57\x83\xca\x68\xbb\xdf\x6c\xcf\xdc\xc2\x39\x29\x8c\x81\x92\x67\x55\xb1\x59\x05\x8e\x24\x35\x32\x18\x3b\xbb\xe7\x7a\x4a\xf9\x9f\xfa\xb9\x14\x93\x6b\x4e\x53\x16\x89\x93\x9c\x17\x0c\x76\x02\xbb\x2d\x83\x4e\x74\x62\xec\x4c\x43\x26\xe3\x4e\x3e\x2f\x8b\x89\xb5\xbf\x79\xbd\x5d\xa6\x39\x03\x5d\x9a\xdf\xde\xd5\x92\xf8\x97\xca\xc1\xac\x9b\x1a\xa2\xeb\x54\xd4\xb5\x32\xc4\x47\xfe\xaf\xea\xc1\x5a\xbb\x45\xba\x4f\x39\xda\xd0\x76\x4c\x7e\x09\x2e\x05\xa2\x38\xb2\x9d\x5d\x46\xb0\x82\xda\xbb\x10\x5e\x71\xc9\x4c\xe5\x64\xa0\x98\xe1\x0b\x15\xa8\xbd\xd2\x36\xc4\x35\xfd\x9b\x07\x29\x48\xe0\x82\x7b\x09\xd4\x6b\xed\x66\x59\x64\xca\x42\x65\x47\xd8\x5b\xb2\x51\x5c\xd8\x51\x8e\x02\xe8\xbf\x09\xa2\x41\x10\x5a\x0a\x61\x92\xbc\x2c\xa6\xed\x12\x11\x5c\x12\x22\xc9\x80\x6e\x97\x16\x7d\x01\xae\x98\x84\xa6\xd1\x4d\xc8\xd3\x7f\xe7\xe1\xa0\x7c\x43\xff\x8b\x9f\xdf\x6c\xe1\xdf\x07\xaf\x43\xa3\x93\xc9\xfe\x13\xed\x20\x70\xfa\x13\xf5\xa2\x96\x64\x78\xaf\x4e\xb7\x70\x47\x15\x4a\xb0\xc1\x79\x08\x83\xa6\xde\xcd\xc8\xc9\xcc\x39\xb8\xe7\x46\x7b\xe7\xdd\x40\xe9\xe4\xac\x19\x33\x51\xad\xdd\xe0\xd3\x4c\x07\x7e\x5b\x6c\x4d\x5d\x37\xf7\x48\x6b\xd2\xe0\x31\xe8\x46\x4a\x9f\xa6\x67\x74\x7d\x98\x9d\x44\x61\xcc\x93\x38\xc2\xfc\xde\xeb\x4e\xf9\x11\x76\x43\xd0\xc4\x0d\xab\x99\xe0\x08\xde\x19\x75\x4a\xb6\x52\x71\xbd\x35\x1f\x2b\x70\x9c\xd0\x42\xb5\xb3\xad\xd1\x75\x7c\xe5\xda\x57\xa9\x3e\x8a\xd3\xb8\xec\xa2\xd4\xd4\xb3\xb9\xc5\x47\xcb\xd8\xf4\x21\x79\xe4\x07\x9e\xa6\xca\x03\x96\x7e\x31\x49\x0a\xb9\x6b\x2f\x8b\xfb\x15\x5e\x3d\x29\xc3\x35\xed\x27\xe7\x1a\x86\xbb\xb9\x80\x8a\x68\xd8\x88\x03\xae\x10\x16\x01\x5f\x70\x43\x24\x43\xf1\xad\x50\xbb\x1e\xaf\x90\x1e\x50\x2d\xd1\x90\x5c\x20\x73\x50\xe6\x79\x54\x59\x84\x61\xf7\x1b\x01\x61\xa2\x0b\x1e\xff\x3a\x68\x8f\x32\x84\xe3\xae\x04\xad\x32\x99\xa7\xcf\x68\xe9\xa6\x03\x17\x32\x57\x76\xa1\xea\x7b\xef\x7a\x4f\x08\xb7\xb2\x76\x4a\x8b\x3c\x29\x90\xdd\xc3\xe0\x8f\xdc\x38\xa7\x22\x97\x65\xbb\x60\xd1\x1c\xb5\xdf\xf0\x50\xc1\x79\x8b\x63\xf8\x03\xfc\x88\x44\x88\xee\xad\x90\x28\x59\x95\x27\xb7\xad\xf3\xf5\xbc\xe4\x65\x9f\x26\xd3\x5f\xe4\xf3\x8b\x4c\xb1\xe6\x98\xa6\x70\x8c\xae\xca\xa7\x21\x47\xa5\x8d\x24\xa4\x9f\x18\xd6\x0e\x81\xbc\x1c\x4d\x9a\xad\x7b\xac\x79\x80\x23\x51\x14\x92\xd5\x64\xd0\x8e\xe1\x5a\x33\xb4\xa2\x95\x4b\xee\xa0\x66\xfd\x5a\xc4\xcc\x65\x68\x49\x6d\xeb\xc1\x7b\xa9\xd6\xb5\xb3\x36\xd9\x91\xf3\x9d\x09\xdc\x05\xfa\xc9\xe0\x47\xf5\x3d\x2a\x66\x48\x8b\xa7\xfe\x1f\x9c\xf1\xed\x16\x7e\xd1\xa1\x46\x63\x94\x45\x37\x08\xcf\xe2\x79\xed\x04\x76\xe7\x86\xd7\x01\x0e\x68\x18\xde\x29\xec\x06\x9b\x7c\x45\xaa\x57\x22\xde\x39\x52\xee\x28\x1c\x5b\xe7\x3b\x6c\x04\x4e\xd6\x94\xd6\x62\x8d\x21\x28\x3f\x72\x45\x64\x2e\xa6\x23\x2c\x96\x15\x59\xff\xb8\x85\xb7\xd8\x6a\xab\xb3\x3e\x9b\xcf\x6e\xd8\x90\x25\x9f\x57\x84\x4a\x2e\xc7\x03\xe3\x17\xb9\x69\x29\xfe\xba\x3d\x1a\xfa\x1e\xbd\x30\xf3\x13\xb1\x76\x45\xdc\x56\x86\x9f\x7c\x3c\xd3\xe8\xa3\x6e\x06\x32\x3c\xf1\x47\x69\xb3\x38\x68\xc6\xb2\x48\x15\x90\xc2\x2a\x15\xbe\x34\xb8\x13\x6f\x77\xbd\x19\xe9\x1e\xbb\x36\xb1\xdb\x65\x34\x57\x6b\x1a\x58\x16\x3f\x92\x18\xf3\x0e\x7c\x9e\x9f\x35\x64\x18\x91\x39\x96\x6c\x9f\xcf\xa3\xf9\x48\x94\x82\x33\x2c\xce\x47\x61\x47\x5d\x20\x81\x6e\x92\x89\xeb\xab\xb3\xf9\x11\x91\x69\x74\xc3\x2d\xd0\xa6\xa9\xe9\x9e\x53\xac\x99\xac\x5c\x95\xc5\x26\xbd\x94\x87\xc2\x37\x7a\x2b\x59\x45\x16\xab\x40\x3a\x3f\x41\xf8\xdc\x05\xf2\x0c\x8b\xd0\x8e\x6b\x3c\x5f\x94\xca\x53\x16\x9d\xb2\x2a\x15\xc2\x1c\xcb\xa2\xd0\x3c\xb8\xdd\x8d\x53\x63\x7a\xd6\x97\x3a\x0f\x37\x5a\x6f\xe5\xdb\x93\x70\xd0\x3d\xb7\x6e\xad\x6e\xe3\x08\x3d\x7a\xfe\x98\xe4\xe6\xbb\xd7\xff\xb0\xe5\x26\xcd\xf9\x09\x6e\x09\x5f\xa3\xb2\x32\x54\x39\x28\x4f\xa6\x95\xc5\xf4\x16\x76\x68\xb1\xd5\xd4\xaa\x95\xc5\xb4\xf2\x99\x74\x12\x80\xdf\x6d\x65\x04\x4a\x0a\x7e\xa4\x48\xcb\xb4\x28\xeb\x7a\xf1\x5d\x06\xf9\xfd\xa4\xc6\xc0\x5c\x61\x31\xb8\xf7\x48\x5d\x4c\x3d\xcd\xa2\x13\xa1\xbb\x32\x3c\x70\x6c\x10\xa3\x4e\x12\x57\x53\xf5\xe4\xe1\x20\x39\x8d\x1b\x45\x1e\xfb\x45\xf4\x2d\xfa\x34\xc1\x92\x96\xdd\x63\xe8\x9d\x0d\x9a\x3a\x30\x3e\xb3\x21\x95\x06\x7a\x4d\xa6\xc6\xa2\xd7\x9f\xb6\xf0\xcb\xe2\x04\x98\x74\x7f\x7e\xb9\x45\xd1\x61\xf1\xe9\xcc\x7f\xff\x17\x7c\xf3\xfa\xf5\x77\xc4\x32\x3c\x11\xb8\xb2\x78\x74\x01\xed\x2d\xfc\x3a\x4f\xcf\xd6\x67\x35\x54\x79\xaa\x45\x1b\x57\x9d\x7f\x12\xb1\x54\x5f\x8e\x75\x29\x87\x97\x47\xd4\xcb\xe3\x99\xb5\xb5\xe4\x7c\x54\x3e\x47\x58\x1e\x48\x9f\xd1\x4f\xf9\x08\xea\x85\x39\x24\x55\x79\x77\x79\x32\x0f\x3f\xbb\x13\x1e\xd3\x60\x9e\x02\x79\x5c\x9e\x77\xf3\xcc\xf8\x1c\x29\x05\x08\xd6\xea\x9e\x6b\x4b\x34\x7f\x31\x92\xcd\xdf\x2b\xc9\x87\x55\xbf\xe4\xcb\x8b\xef\xb7\xa8\x57\x66\xd2\xa9\x39\x0e\x65\xdc\xb0\xfa\x8c\x2b\x1d\xb0\xce\xdf\x71\x9d\x31\x83\xf9\x9b\xac\xb9\x93\xf8\x33\xe7\xf5\x4c\xf3\xe7\xae\x9e\x8f\xc0\x08\x78\x4e\x2a\xa1\x0e\x77\x26\xe7\xa2\x91\x97\xe5\x73\xb1\x6b\x5f\xac\x48\x8b\xf5\xf0\xf4\x7e\x23\x2d\xd9\xb8\xd8\x48\x92\xc7\x85\x7c\x70\x9c\xa8\x26\xfd\x20\x8b\x5e\xdd\xed\x7b\x01\x02\xf1\x84\x9c\x44\xf0\x57\x67\xf2\xa6\x0c\xc8\x79\x16\xd5\x2e\xbe\x05\x69\xb5\x0f\x91\x28\x80\xda\x7b\xd5\x1f\x40\xed\xdc\x71\x31\xef\x95\xd7\xca\x62\x73\xf6\x71\xd9\xbf\x68\x1b\xd0\x27\xb5\xf3\x0c\x90\xa5\x23\x92\xf1\xaf\xac\x9a\x80\x7c\xde\x39\x4b\xed\x4e\x7c\x7a\xa8\x22\x19\x87\x54\x25\x72\x4a\x04\x77\x6f\x46\x08\xba\xd3\x46\xf9\xec\x94\xf4\xea\x59\x40\xf3\xa7\x3b\x09\xa7\xa6\xb0\xe3\x95\xb8\x15\x9a\xdd\xb1\x0e\x00\x99\x8c\x93\x23\x82\x9c\x41\x4c\x3d\xc4\x60\x0d\x86\x90\xa2\xee\xc2\x87\x07\x15\x60\x87\xc8\xc7\x5c\xde\x1d\x05\x8e\x92\x43\x79\x9d\x7b\x2a\x0b\x92\x15\x37\x0f\x4f\xf7\xdb\xdc\xfb\x52\x47\x46\x05\x4f\x0c\x41\x19\x98\x0d\xe5\xf1\xa8\x51\xda\xfc\x1a\x7d\x9c\x41\xa6\xf7\x8e\xaa\xfe\x6d\x59\xfc\x4f\x00\x00\x00\xff\xff\x62\xf6\x45\x48\xe5\x28\x00\x00")

func proto_micro_mall_sku_proto_license() ([]byte, error) {
	return bindata_read(
		_proto_micro_mall_sku_proto_license,
		"proto/micro_mall_sku_proto/LICENSE",
	)
}

var _proto_micro_mall_sku_proto_sku_business_sku_swagger_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\xdd\x8e\xdb\xc6\x15\xbe\xdf\xa7\x20\xd8\x5e\xb4\x40\x60\xb9\x6e\xd1\x0b\xdf\xb1\x12\x1d\x0b\xd9\x5d\x09\x24\xe5\xd4\x28\x0c\x82\x22\x47\xda\x89\x48\x0e\x3d\x3f\xeb\xaa\xc5\x02\x31\x8a\xb4\x71\x10\x37\x01\x5c\xc7\xfd\x31\x82\x3a\x8d\x51\x5f\xd4\x6e\x0a\x14\xb1\xeb\x16\xe8\xcb\x58\xda\xf5\x5b\x04\xa4\x48\x89\xa4\x38\x12\x45\x6a\xb5\x14\x20\x02\x06\x56\xc3\x99\x33\xe7\xef\x9b\x73\x66\xe6\x98\xbf\xde\x13\x04\x91\xdc\x31\xfa\x7d\x80\xc5\xab\x82\x78\xe5\xd2\x65\xf1\x1d\xbf\x0d\xba\x3d\x24\x5e\x15\xfc\xf7\x82\x20\x52\x48\x6d\xe0\xbf\xf7\x30\xa2\xa8\xe6\x40\x13\x23\xdd\x31\x6c\x5b\x27\x03\xa6\x4f\x1a\xfd\xbf\xba\x8c\x40\x17\x10\xe2\xff\xb8\x14\x34\x07\xd4\x04\x41\x3c\x06\x98\x40\xe4\xfa\x34\xc2\x3f\x05\x17\x51\x81\x00\x2a\xee\x09\xc2\x49\x30\xa7\x89\x5c\xc2\x1c\x40\xc4\xab\xc2\x2f\x26\xa3\x0c\xcf\xb3\xa1\x69\x50\x88\xdc\xda\x07\x04\xb9\x7e\xdf\x5b\x41\x5f\x0f\x23\x8b\x99\x39\xfb\x1a\xf4\x88\xcc\x84\xa9\x1d\xff\xc8\xe7\xaf\x06\xdd\x63\xe0\x52\x84\x87\x35\x13\xb9\x3d\x88\x9d\x69\x17\x7f\x0c\x22\x34\xf6\xdb\xd7\x12\x73\x1c\x03\x0f\x7d\x09\x4e\x9f\xbc\x38\x7b\xf1\xf5\xe8\xf5\x83\xd1\xf3\x3f\x86\x02\x06\x5d\x90\x07\x70\xc0\x40\xd3\xf2\xbb\xa9\x03\xf6\xb3\x50\x21\x2a\xc0\xc7\xd0\x04\x7a\x7d\x32\x93\x3a\x60\xcd\x68\xf6\x38\x01\x0c\x88\x87\x5c\x02\x48\x62\x6a\x41\x10\xaf\x5c\xbe\x9c\x6a\x12\x04\xd1\x02\xc4\xc4\xd0\xa3\xa1\x5e\x25\x81\x30\xd3\x04\x84\xf4\x98\x2d\x44\x94\x2e\xc5\xc8\x4f\xc4\x30\x8f\x80\x63\xcc\x11\x13\x04\xf1\xfb\x18\xf4\x7c\x3a\xdf\xab\x59\xa0\x07\x5d\xe8\xd3\x25\x09\xb3\x66\x70\xaf\x84\xf3\x88\x09\x6a\x27\xb1\x5f\x27\x71\x06\x44\x0b\xf4\x0c\x66\xd3\xe5\xc2\xb8\x02\x73\xc1\x2f\x3d\x60\x52\x60\x09\x00\x63\x84\xd7\x27\x13\x66\x2e\x85\x0e\x90\x7d\xaa\x0b\x18\xdf\xcb\x10\x41\xf4\x0c\x6c\x38\x80\x02\x3c\x73\xbd\xc9\x93\x92\xc7\x35\x9c\x00\x31\x5d\x64\x0d\xd3\xfc\x42\x97\xf7\x06\x83\xdb\x0c\x62\xe0\xbb\x0f\xc5\x0c\x9c\xb7\xed\x6e\x33\x40\x68\x1e\x0d\xdc\x8a\x69\x80\x1a\xfd\xb4\xec\x19\xae\x3e\xa3\x7a\x6b\x2f\x4e\x2d\xd4\x65\x06\x0a\x2d\x60\x31\x93\x26\x40\xc8\xf8\x18\x1c\xdf\xfb\xdb\xe8\x77\x9f\x15\xc1\x60\x23\x98\x67\x1b\xf1\x97\xe2\x7c\x87\xbd\xe8\xa9\x38\xf6\xe6\xec\x56\x29\xdc\xd9\x30\x11\xea\xc4\x3e\xe0\xa3\xee\xec\xf7\x2f\x47\x9f\x7d\x31\x7a\xfd\xe8\xed\x83\xd7\x64\xc0\x46\x1f\x3f\x3a\x7b\xf2\x6c\x35\xf8\xbd\x0b\xa8\x3a\x60\xfb\xfe\xa4\xdb\x83\xbc\x19\xd3\x3b\xd0\x45\x0f\x07\x74\xe4\x08\x79\x3a\xb4\xb2\x71\x77\x9b\x01\xbc\x08\x78\x3d\xc3\x26\x69\xe4\xd1\xa1\x37\x21\x4c\x31\x74\xfb\xe9\xc1\x3d\x84\x1d\xc3\x57\xaf\x08\x5d\xfa\xd3\x9f\x88\x3c\x43\xf0\xb8\x1d\x30\xdd\x44\x16\xd0\xed\xa4\x43\xae\x89\x67\x03\x63\x63\x7e\x09\xa2\xc0\x49\xfb\x79\x96\xa4\x49\x43\xa5\x88\x98\xc8\xb6\x81\xe9\x9b\xf7\xda\x54\x03\x0e\xb3\x29\x14\x2f\x6c\x21\xf1\x18\xd5\x8d\x3b\xc6\x30\x77\x1e\xfd\xe6\xd5\x27\xe3\x2f\xbf\x1d\x3d\xfc\xed\xe8\xc1\xdd\xd5\x16\x91\x36\xa3\xd2\x1d\x63\xa8\x0e\xd8\x16\x2d\x22\x33\xa6\x77\x8b\x48\xf4\x54\x3c\x72\xc7\x4d\x56\xa9\xa0\x8d\x01\xa1\x08\x83\xdc\xd9\xf2\xdd\xaf\x46\x5f\xdf\x2f\x92\x2d\x2b\x93\x89\xb6\x31\x5d\x4e\xb3\xbe\x43\x5d\xf4\x54\x1c\x75\xf3\x86\xab\x14\xf6\x08\x30\xb0\x79\x94\x3b\xca\x4d\xe2\xdb\xf8\xf3\xc7\xa7\xff\xfe\x6a\x35\xec\xa9\xc1\x44\x5b\x7a\x58\x34\xcf\xfc\x0e\x7f\xd1\x53\x71\xfc\x65\x99\xae\x82\x08\xac\x91\xa1\x6b\xae\x08\xc3\x49\x08\x9c\x80\x71\xf4\xf9\xa7\xe3\xe7\x4f\xc7\x0f\xbf\x19\xdf\x7f\xf1\x83\xb3\x7f\xbe\x1c\x3d\x7e\xf6\xe6\xd5\x9f\x47\x9f\x3c\x79\xf3\xea\xfe\xe9\x7f\xee\x8e\xfe\xff\xa7\xf1\xa3\x6f\xcf\xbe\xf9\xcd\xe9\x1f\x9e\xfd\xb0\x10\x70\x87\xae\xb9\xdd\xe0\x4d\x09\xb0\x03\x70\xf4\x6c\x07\x80\xe7\xcc\x57\x01\x10\x7b\x18\x9a\xa0\xd6\x83\x36\xc5\x06\x05\xb9\xd1\x3b\xfe\xf4\xde\xe9\xf3\x2f\x4e\x1f\x3e\x3d\x7d\xfe\x97\xb7\x1f\xde\x9b\x80\xf9\xcd\x7f\x5f\x8e\xff\xfa\xbf\xd3\x7b\x1f\x8f\x1f\xff\x63\x35\x80\x5e\x0b\xe7\x57\x07\xac\xed\x33\x74\x23\xbc\x9e\xda\x1e\x84\x72\x24\xd8\x41\x34\x7a\x2a\x0e\x51\xae\xfd\x2a\x81\x51\x1f\x41\x74\x58\x23\xcc\xf3\x6c\xe0\x00\x37\xf7\xb5\xcc\xd9\x93\xa7\xa3\x8f\x3e\x22\x03\x16\x86\xdb\x7f\x7d\x39\xfe\xf0\xef\x2b\x06\xcf\xe9\xa4\x81\x6e\x26\xac\x6c\x11\x32\x33\xf9\xdf\xe1\x32\x7a\x2a\x8e\x4b\x8e\xf5\x2e\x02\x95\xd3\xc2\x88\x18\xaf\xb3\x32\x86\xa0\xbc\xa2\xcb\x7a\x92\x9b\x38\x70\x8d\xce\x90\x51\xf7\x03\x60\xce\x4e\xb4\xc5\x10\xd3\x30\x85\x9a\xa0\xbf\xce\xb0\x9d\xc6\x12\xef\x2c\x3a\x6e\xe3\x63\xc3\x66\x60\xc9\xc0\x84\x57\xcf\x0e\xeb\xbb\x43\x1a\x13\xfc\x24\x73\x39\x4a\xb8\x5f\x09\x11\x41\x8a\x40\x6e\xf9\x4c\x64\x71\xc5\x83\x2e\x05\x7d\x80\x79\xf2\x41\x97\xfe\xf8\x4a\x36\x55\x07\x10\x62\xf4\x97\xe9\x2d\x73\xa8\x05\xa8\x01\xed\xb9\x75\x8f\x7f\xdb\xc0\xb9\x6b\xe0\xe0\x20\xee\x52\xd9\x00\xcf\xb4\x53\xb2\xcc\xc0\x71\x62\x59\x48\x09\xab\x65\x29\x3f\xcf\xd1\x11\xad\xfb\x03\xb3\x35\x4f\xfa\xb9\xb5\x9e\x43\x52\x6e\x41\x45\x09\xa9\x11\xa3\x3a\xc5\x86\x05\x74\x17\x95\x37\x33\xff\x42\x29\x7b\xd1\x46\x9e\xee\x00\x9a\x5e\x39\x73\xa8\xbd\x15\x45\xf5\x03\x7f\x78\x59\x2d\xae\xc3\x79\x7c\x37\x5c\x5d\x8e\x94\xfb\xae\x20\x08\xe7\x8e\xbf\x84\x0c\xa9\x7b\xf9\x82\x2e\xb0\x5c\xe6\x29\xcf\xb2\x4b\xf1\x50\x3d\x42\x5e\x1e\x47\x09\xed\xbd\x69\x7f\xe1\x95\xc0\x54\xc9\x57\x62\x7a\x82\x44\x0f\xd3\x4d\x9e\x25\xbb\x08\xd9\xc0\x70\x57\xd1\xc1\x92\xad\x43\x09\x55\x10\x40\xf5\x4d\xb9\x5d\x8a\x7b\x15\xd0\xe5\x6e\xe7\x21\x1b\x9a\x43\x3d\x64\x67\x45\x7b\x45\x13\xb6\x03\x22\x91\x12\x35\x9f\x56\xe6\x64\x36\x74\x20\xd5\x99\xe7\x01\x6e\xf2\xb0\x3c\x09\x88\x57\x24\x94\xb0\x6c\xb5\xbd\x1c\x03\x32\xbf\x57\x39\x17\xa7\xe1\x2b\xc8\x67\xa0\x60\xce\xb2\x98\x68\x19\x38\x85\xa5\x30\x85\x32\xe4\xb4\xf3\xc4\xf4\x1d\x15\xad\x6c\x28\x41\x58\xaa\xc0\x8c\xea\xa8\x8a\x7a\xea\x06\x17\xb7\x69\x8c\x6a\xba\x3d\x54\x54\xb3\xc9\xd8\xdc\x08\xd2\xff\x52\x1e\xb9\xc4\x73\x16\x6d\x3e\x0c\x07\x31\x97\xab\xbd\xd5\x9c\x79\x45\xc1\x83\xa4\xa4\x82\x40\xcc\x93\xad\x2f\xdf\xcf\x6d\x3e\xc9\x0b\x1d\xa9\xa0\x53\x26\x93\xb7\x32\x9b\x1d\x4f\x67\xe7\x61\x15\x4f\x87\xde\xfa\x36\x7a\x53\x71\xb5\x64\xda\xc1\x63\x54\x04\x2e\x73\x12\xc7\x3e\x62\x5d\x91\x25\x4d\x8e\x9f\x61\x76\xda\x8d\x54\x4b\x43\xde\x97\x93\x2d\x52\xa7\xd1\xd4\xe2\x0d\xed\x8e\xa6\x4b\xef\x4b\x37\x23\xee\xa7\x67\x4d\xb1\xe3\xc2\x68\x2e\xae\x38\xf3\x25\x4d\xe5\xd6\x93\x42\x89\x18\x67\x71\xcc\xde\x62\x14\x4b\xf7\x92\x66\x5b\x32\xc5\x86\x77\x31\x19\x95\x80\x55\x0a\x98\x4b\xf9\xe7\x15\xe8\xec\x76\xbb\xeb\xf5\x13\x6e\x05\x5b\x95\xbc\xe5\xbc\xb7\xbb\xd1\x91\x5e\xa1\x95\x57\xed\xd4\xeb\xb2\xaa\xc6\x57\x51\x59\x51\x5a\x4a\xbc\xa1\x79\x78\x43\xda\x6f\x36\xf4\xb6\xa4\x48\x07\xb2\x26\x27\x5e\x76\x54\x59\xd1\x0f\x5b\x9a\x2e\xff\xbc\xa9\x6a\x73\x6f\xe6\x5a\x0f\x64\xa5\x7e\x5d\x3a\xd4\xb2\xc7\x4c\xdf\xce\xbd\x51\xaf\xb7\xda\xd9\x63\x82\x37\xf3\xad\xef\x75\x38\xdd\xdf\xeb\x64\x37\x4a\x07\xad\x4e\xc4\xd7\x61\xab\xf3\xee\xf5\x78\x07\x4d\x91\x0e\x55\xa9\xae\x35\x5b\x87\xfa\x35\xa9\xb9\x2f\x37\xd2\xc3\xdb\x4a\xb3\x2e\xeb\x37\x64\x45\xf5\xfb\x70\xa7\x4e\x76\x6b\xb7\xf6\x9b\xf5\x9b\xba\x76\xb3\x2d\x07\x43\xd4\x4e\xbb\xdd\x52\xf2\x0d\x6a\x48\x9a\xa4\x5f\x6b\x29\x07\x92\xa6\xcb\x8a\x92\x1e\xd3\x90\x1b\x9d\xba\xa6\x37\x0f\x6f\xc8\x87\x5a\x4b\xb9\xa9\x2b\x72\xbd\xa5\x34\xb2\xb9\xe2\x75\xe6\x8a\xa1\xc8\xaa\xd6\x52\xe4\x9c\xe4\xb9\xbd\x67\xf4\x17\x85\xec\xc8\x49\xb9\x10\x98\x2f\xc8\x0a\x56\xaf\x32\xcb\x40\xe2\x3f\xf8\x86\x6d\xeb\x0a\xdd\xc4\x4c\x96\x26\x27\x18\x74\x99\xd3\xe5\x1f\x99\x58\x88\x75\xed\x95\x62\x11\xbf\x58\xad\x84\x76\x06\x60\x78\x07\xe1\x65\xa9\x69\x49\x2e\xab\xbd\x96\x6f\x6c\xab\xcc\xf1\xed\xec\xa8\x9c\x57\xd1\x9c\xd2\xa7\x0a\xee\x21\x3d\xa3\x0f\x74\x02\x7f\x55\xf0\x22\x75\x19\xe5\x49\x44\xdc\xc4\x96\x7d\x49\xcd\x60\x45\xdd\x3c\x6b\x1d\xac\xf2\x89\xd0\x1c\x9d\x0b\x3a\x0d\x0a\xeb\x36\x56\x1e\x17\xd4\xff\x15\x19\x18\x7d\x83\x62\xe5\x81\x84\x75\xf5\xc2\x83\x2d\x40\xcc\x62\x62\x22\x8b\x99\x61\x95\x4e\x01\x96\x3d\xcf\x86\xfc\x7b\x87\x85\x45\x0b\x06\x05\x7d\x94\x48\x0e\x12\x63\x8b\x17\x2e\x98\xc8\x2e\x5a\x47\x61\x23\xbc\xd0\xd3\x8a\x33\x45\x3c\x60\xc2\x5e\xf8\xe1\x8f\xa2\x16\xd6\x6d\xe8\x0e\x0a\x19\x8a\x26\x8b\x58\xd7\x24\xd4\xec\x6b\x29\xeb\x8d\x08\xeb\x3c\xc2\x7d\xe7\x9c\x42\xe3\x5e\x8a\xfe\xec\xe3\x33\x89\x0a\xfe\xa0\x6a\x7f\x41\xc6\xcc\xbf\xec\x2b\xb4\x8f\x8c\xf6\x26\xaa\x1c\xec\x91\xe2\x1b\x80\xe8\x55\xa7\xdd\x4e\x6e\x1d\xa3\x17\xfb\xad\xf7\x65\x65\xe1\x06\x20\x4d\x7d\xa9\x58\xe1\x15\x55\xe9\x5d\x40\x99\x08\x50\x78\x25\x5f\xab\x77\xe7\x09\x98\xe9\x7b\xe6\x0a\x66\x81\xc0\xb7\xe4\x45\xdd\xbd\x97\x4b\xb8\x17\x15\x4c\x56\x50\xd3\x65\x5c\x3e\x47\x9a\xce\x59\x20\x8d\xf9\x23\xec\xd5\xd2\x08\xdd\xf4\x97\x6e\x7e\x58\xcf\x49\xc5\x5a\x10\xaf\x16\x4b\x00\xec\x9e\x6e\xc3\xde\x46\x93\xc5\xb2\xf7\x6b\xe7\x75\x70\xbb\x81\x6b\x8a\xa2\xb8\xab\xe0\x6e\x2b\x29\xca\x9e\xff\xef\x64\xef\xbb\x00\x00\x00\xff\xff\xf3\xf7\xc8\x71\x7a\x4e\x00\x00")

func proto_micro_mall_sku_proto_sku_business_sku_swagger_json() ([]byte, error) {
	return bindata_read(
		_proto_micro_mall_sku_proto_sku_business_sku_swagger_json,
		"proto/micro_mall_sku_proto/sku_business/sku.swagger.json",
	)
}

var _proto_micro_mall_users_proto_users_users_swagger_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x4f\x6f\x23\x35\x14\xbf\xe7\x53\x58\x86\x63\xd5\x94\x82\x38\xf4\x16\xa5\x11\x54\xa2\x5b\x94\xa4\x12\x12\xaa\x46\xee\xcc\x4b\xe2\xd5\x8c\x3d\xb5\xdf\x94\x46\xa8\x87\x95\x00\xed\x09\x0e\x5b\x16\x21\xd0\x8a\xbd\x71\x82\x03\x42\x45\xcb\x61\xbf\x0c\xc9\x96\x6f\x81\x3c\x7f\x92\x99\xc9\x24\x9d\xfc\xeb\x26\xa2\x52\xab\xd4\xf6\x7b\x7e\xbf\x9f\x7f\x7e\xcf\x2f\x5f\x56\x08\xa1\xfa\x0b\xd6\xed\x82\xa2\x07\x84\xee\xef\xee\xd1\x1d\x33\xc6\x45\x47\xd2\x03\x62\xe6\x09\xa1\xc8\xd1\x05\x33\xef\x2b\x89\xb2\xea\x71\x5b\x49\xcb\x63\xae\x6b\x05\x1a\x94\xb6\xa2\xe1\xf0\x73\xf4\x77\x37\x1c\x09\x5d\x11\x42\x2f\x41\x69\x2e\x85\x71\x10\x7f\x24\x42\x22\xd1\x80\xb4\x42\xc8\x75\xb8\xa1\x2d\x85\x0e\x3c\xd0\xf4\x80\x7c\x1e\x59\x31\xdf\x77\xb9\xcd\x90\x4b\x51\x7d\xac\xa5\x30\x6b\xcf\xc2\xb5\xbe\x92\x4e\x60\x97\x5c\xcb\xb0\xa7\xc7\x48\xaa\x97\xef\x85\x11\x56\x33\x00\x09\xa1\x5d\xc0\xd4\xbf\x86\x95\xc0\xf3\x98\xea\x9b\xa0\xef\xbe\xbd\x1d\x7c\xf7\xfc\xcd\xcd\xaf\xc3\xa7\xb7\xff\xbc\x7e\x39\x7c\xf2\x7b\x8c\x2c\x5c\x28\x7d\x50\xe1\xce\x47\x8e\x59\xfc\x11\xe0\xa9\x06\x75\x64\xfc\xa7\x56\x29\xd0\xbe\x14\x1a\x74\x66\x17\x42\xe8\xfe\xde\x5e\x6e\x88\x10\xea\x80\xb6\x15\xf7\x31\x66\xad\x46\x74\x60\xdb\xa0\x75\x27\x70\x49\xe2\x69\x37\xe5\x3e\x8a\xd8\xee\x81\xc7\x26\x9c\x11\x42\xdf\x55\xd0\x31\x7e\xde\xa9\x3a\xd0\xe1\x82\x1b\xbf\xf1\x49\xa5\xc2\x6d\xc6\x8e\x69\xc6\xfc\xba\x52\xf4\xf9\x3a\x05\xcd\x67\x8a\x79\x80\xa0\xc6\x07\x12\xfd\xe4\x40\x09\xe6\x85\x22\x0a\xb8\x93\x0f\x9d\x87\x30\x2f\x02\x50\xfd\xfc\x94\x82\x8b\x80\x2b\x30\xdc\x76\x98\xab\x21\x37\x8d\x7d\x3f\x74\xaa\x51\x71\xd1\xcd\x1b\x77\xa4\xf2\x98\x39\x58\xca\x05\x7e\xf8\x01\x2d\xc4\x72\x96\xc2\x82\xac\x9b\x47\x41\x0d\x3f\xba\x05\xea\x92\xdb\x29\x6e\xce\x2a\x69\x3f\x31\x1f\x63\x7d\x79\xa0\xec\x1e\x13\xa8\xab\x4c\x6b\x69\x73\x86\x50\xd5\x3d\xe9\xa7\x35\xe7\x07\xd3\x35\x37\xf8\xfe\x9b\xe1\xd3\xdb\xc1\xd7\x7f\xdc\x3d\xb9\x19\xbc\xfa\xe1\xdf\x67\xaf\x66\x68\xee\x38\xd9\xac\x96\xec\xd5\x32\x5b\x6d\x81\xfc\x8a\x23\x7f\x28\x25\x9e\x4b\x67\x42\x6f\x91\x14\x8b\x66\x52\x4a\x44\x15\xe4\x85\xb8\x52\xf8\x17\x01\x68\x2c\x83\xfe\x3e\xed\x8e\x36\x58\x42\xbf\x81\xc3\xb1\xb4\x6c\x7f\x7b\x39\xfc\xe5\xaf\x48\xbc\x3b\x77\xaf\x6f\x06\x3f\xbd\x88\x86\xde\xfc\xfd\x6c\xf8\xe2\xe7\x32\x1a\x3e\x66\x08\x8a\x33\xb7\x16\xee\xbb\x4d\x1a\xce\x44\xfe\xff\xd3\x70\x0e\xfe\x26\x69\xd8\x8b\x43\x9b\xaf\xe2\x47\x32\xbe\xfb\xf3\xab\xe1\xf3\x1f\x67\x57\xfc\x09\x0e\xb6\x41\xb7\x45\x71\x3f\x94\x6a\x93\x03\xb1\xb6\xec\x2d\x50\x42\x8b\x3b\xa3\x3c\x29\xf5\x8a\xde\x94\x5b\x29\xaf\xb7\xa6\xad\xcd\xcb\x88\x6f\x23\x19\x8e\x3a\xab\x54\x78\xe3\x3e\x28\x8c\xb3\x2e\x3d\x4f\x8a\xd1\xb1\xa4\x72\x63\x72\x93\xe4\xf9\x63\xb0\xc7\x45\xd8\x74\x5e\x3e\x28\xe4\x39\xb1\x51\x5b\x3a\x90\x97\xdf\x0c\x82\x9a\x80\x75\x63\x51\x78\xdc\x9e\xee\xe6\x5d\xe5\x2e\x76\x25\x4f\x5b\x36\xff\xdf\x9f\xdc\x96\x42\x6a\x38\x9b\x03\x6b\x8e\xe4\x42\xc8\xb9\x66\x74\x6e\x7d\x85\xcd\x66\x49\x56\x26\xba\xbd\x2d\x27\x63\xd4\x6c\x97\xc1\x3f\xfb\xbd\xbd\x04\x13\xe9\x7a\x36\x5b\xbc\xe9\x04\x33\xbd\x26\xa5\x09\x31\x6d\xe3\x3a\xfc\x72\x6d\x8d\x7a\xd3\x69\xce\xcf\xa5\x74\x81\x89\x69\xde\x93\xe9\x65\xb8\xdf\x24\x19\x96\x88\xbe\xf0\x95\xbb\xa1\xca\x09\x5b\x37\x2b\x58\x83\x67\x43\x39\x08\x5c\x41\xa2\xbe\xa7\x7b\xda\x08\x59\xa4\x2f\x23\xce\xb8\x2d\x5c\x20\x74\x41\xcd\xe0\xf3\xfd\xfd\xd5\xf0\x99\x32\x15\x70\x85\xd6\xfa\x4e\x9a\x6b\x4b\x5e\x86\xdf\x0e\x3f\x58\x82\xc8\x94\xb5\x25\x04\xb0\x0e\x3a\xd6\x79\x5d\x15\x74\xb9\x46\x50\x16\x73\x9c\xa9\x84\xcf\x92\x42\x0f\x98\x8b\x3d\xcb\x66\xca\xb1\xc4\x44\x11\x2d\xe3\x81\x3b\x20\x90\x63\x7f\xf5\x0a\x5f\xd3\xc5\x41\x76\x35\x37\xde\x39\x24\xb8\x82\x04\xbf\xe2\xc7\x5d\x51\x97\x68\xc5\x31\x95\xdd\xe3\x24\x31\x6c\x1b\xbb\x05\x89\xd9\xd0\x04\xbd\xf2\x0b\x3a\x9d\x94\x2c\x8d\x05\x44\xe4\x36\xa2\x20\x02\x2f\xd3\xd0\xd1\x7a\xb3\x51\x6b\x37\xd2\xed\xfc\xe9\xa7\x87\xb9\x91\xc3\xc6\x27\x8d\xec\x48\xed\xf4\xf0\xa8\x9d\x84\x38\xea\x18\x4d\xa7\xc7\x02\x37\x04\x11\xfb\x9d\x8c\x39\x69\xbd\x16\x8a\xb6\x75\x5a\xaf\x37\x5a\xad\x74\x28\x8d\x66\xf3\xa4\x99\x89\xbf\xd5\x68\x5a\x8f\x4e\xda\x56\xe3\xb3\xa3\x56\x7b\x62\x66\x62\xf4\xb8\xd1\xac\x7f\x5c\x7b\xd4\x2e\xb6\x19\xcd\x46\x33\xb3\x20\x27\xc1\x4d\x62\x1e\xb5\x08\x1b\x56\x4b\x98\x6d\xcb\x40\xe0\xfd\x4a\x2d\xb4\x36\xd0\xac\xf8\x3b\x8f\xb9\x8d\x35\x5c\xad\xe3\x11\x13\x08\x54\x7d\xab\xe8\xdb\x80\x32\x41\xf9\x3d\x29\x16\x32\x04\x8f\x71\x77\x21\x1a\xd6\x53\x94\xb8\xb3\x54\x0d\x16\x97\x1c\xa7\xbf\xb8\x16\x7f\xc9\x45\x7e\x17\x3e\x1e\x5b\x0a\x64\x36\x2e\xfc\x3a\x61\xdd\x75\x3c\x9c\x15\x30\x04\x0b\xf9\x62\xf7\x20\xf0\x9d\xb9\xcd\xb3\x05\xa1\x62\x7e\xaf\x2b\xff\x05\x00\x00\xff\xff\x60\x10\x61\xa2\x4a\x21\x00\x00")

func proto_micro_mall_users_proto_users_users_swagger_json() ([]byte, error) {
	return bindata_read(
		_proto_micro_mall_users_proto_users_users_swagger_json,
		"proto/micro_mall_users_proto/users/users.swagger.json",
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
	"proto/.DS_Store": proto_ds_store,
	"proto/micro_mall_sku_proto/LICENSE": proto_micro_mall_sku_proto_license,
	"proto/micro_mall_sku_proto/sku_business/sku.swagger.json": proto_micro_mall_sku_proto_sku_business_sku_swagger_json,
	"proto/micro_mall_users_proto/users/users.swagger.json": proto_micro_mall_users_proto_users_users_swagger_json,
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
	"proto": &_bintree_t{nil, map[string]*_bintree_t{
		".DS_Store": &_bintree_t{proto_ds_store, map[string]*_bintree_t{
		}},
		"micro_mall_sku_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"LICENSE": &_bintree_t{proto_micro_mall_sku_proto_license, map[string]*_bintree_t{
			}},
			"sku_business": &_bintree_t{nil, map[string]*_bintree_t{
				"sku.swagger.json": &_bintree_t{proto_micro_mall_sku_proto_sku_business_sku_swagger_json, map[string]*_bintree_t{
				}},
			}},
		}},
		"micro_mall_users_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"users": &_bintree_t{nil, map[string]*_bintree_t{
				"users.swagger.json": &_bintree_t{proto_micro_mall_users_proto_users_users_swagger_json, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}
