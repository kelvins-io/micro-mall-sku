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

var _proto_micro_mall_sku_proto_sku_business_sku_swagger_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x8f\xdb\xc4\x17\x7f\xcf\xa7\xb0\xfc\xff\x3f\xa2\xa6\x14\xc4\x43\xdf\x4c\xd6\xb4\x51\xdb\xcd\x2a\x4e\x04\x08\x55\xd6\xc4\x3e\xce\x4e\x63\xcf\xb8\x73\xd9\x12\xd0\x4a\xf4\xa1\xdc\xa4\xd2\x87\xaa\x94\xa7\x4a\xe5\x22\x78\x01\xc1\x0b\xa0\x7e\x1e\xb2\x2a\xdf\x02\x8d\xe3\x24\xbe\x6f\xe2\x64\x77\x13\x69\x2d\xad\xb4\x19\xcf\x39\x73\x6e\xbf\x73\xce\x9c\xe4\xd3\x86\xa6\xe9\xfc\x01\x1a\x0e\x81\xe9\xd7\x35\xfd\xda\x95\xab\xfa\x1b\x6a\x0d\x13\x8f\xea\xd7\x35\xf5\x5e\xd3\x74\x81\x85\x0f\xea\x7d\xc8\xa8\xa0\xcd\x00\x3b\x8c\xda\x01\xf2\x7d\x9b\x8f\xa4\x3d\x5d\x54\xff\x0d\x24\xc7\x04\x38\x57\x1f\xae\x44\xcb\x11\x37\x4d\xd3\x8f\x80\x71\x4c\x89\xe2\x11\xff\xab\x11\x2a\x34\x0e\x42\x6f\x68\xda\x71\x74\xa6\x43\x09\x97\x01\x70\xfd\xba\xf6\xd1\x94\x0a\x85\xa1\x8f\x1d\x24\x30\x25\xcd\x7b\x9c\x12\xb5\xf7\x6e\xb4\x37\x64\xd4\x95\xce\x92\x7b\x91\x38\xe4\x0b\x65\x9a\x47\x6f\x2a\xf9\x9a\x98\x1c\x01\x11\x94\x8d\x9b\x2e\xb8\xd2\x11\xf3\x1d\x8a\x44\x26\x3f\x2a\x1b\xc9\x20\x40\x6c\xac\xe4\x3f\xf9\xea\x87\xc9\x17\x4f\x26\xaf\x9e\x4e\x7e\xfd\x2e\x56\x2f\xda\x42\x43\x60\xd1\xf1\x6d\x57\x6d\xb3\x46\xf2\xdd\xd8\x1c\x16\xb0\x23\xec\x80\xbd\x17\x9d\xd3\x9e\x9d\x9b\x24\x66\xc0\x43\x4a\x38\xf0\xd4\xb1\x9a\xa6\x5f\xbb\x7a\x35\xb3\xa4\x69\xba\x0b\xdc\x61\x38\x14\xb1\x45\x0d\x8d\x4b\xc7\x01\xce\x3d\xe9\x6b\x33\x4e\x57\x12\xec\xa7\x2a\x38\x87\x10\xa0\x1c\x33\x4d\xd3\xff\xcf\xc0\x53\x7c\xfe\xd7\x74\xc1\xc3\x04\x2b\xbe\x3c\xe5\xd0\x8c\xe4\xdd\xf8\x0c\x3d\xc5\xe9\x38\xf1\xe9\x38\x79\xb8\xee\x82\x87\xa4\x2f\x4e\x57\x84\x68\x92\xc0\xc7\x21\x38\x02\x5c\x0d\x18\xa3\x6c\xae\xcf\xba\xea\x30\x49\x04\x0e\xc0\x54\x4c\x2b\xe4\x6e\x14\x68\xa0\x87\x88\xa1\x00\x04\xb0\x45\xbc\x4d\x9f\x8c\x3a\x04\x05\x11\x4c\x06\xd4\x1d\x67\xe5\xc5\xa4\xec\x0d\x83\xfb\x12\x33\x50\x51\x23\x98\x84\xb3\x74\xdb\x7d\x09\x5c\x2c\xa3\xfd\xdd\x84\xf6\x02\x0d\xb3\x7a\x17\x44\xf7\x82\xeb\xdd\x46\x92\x5b\x6c\xc7\x02\xd8\xf9\x98\xa7\x40\x37\x84\x72\xd0\xbd\xfe\xe6\xaf\xc9\x93\x6f\x27\xaf\x9e\xff\xfb\xf4\x15\x1f\xc9\xc9\x97\xcf\x5f\xbf\xfc\x65\x35\xf4\xdd\x00\x61\x8d\xe4\x6d\x75\xe8\xee\x00\x6f\x21\xf4\x25\xe6\xe2\xa7\x04\x73\xfc\x90\x86\x36\x76\x8b\x61\x77\x5f\x02\xab\xc2\x9d\x87\x7c\x9e\x05\x9e\x18\x87\x53\xc6\x82\x61\x32\xcc\x12\x7b\x94\x05\x48\x59\x57\xc7\x44\xbc\xf3\xb6\x7e\x61\x28\x0a\xa5\xb0\xd1\x03\x34\x4e\x95\x2f\xca\xcb\xa1\xf4\xcf\xdf\x5f\x9f\xbc\xf8\x73\xf2\xec\xf3\xc9\xd3\x87\xab\x21\xe8\x40\x0a\xe3\x01\x1a\x5b\x23\xb9\x43\x08\x5a\x08\x7d\x89\xa0\xf8\xd9\xf2\xaa\x95\xf4\xd8\x56\x15\x2c\x06\x5c\x50\x06\x4b\x37\x8a\x0f\xbf\x9f\xfc\xf8\xb8\x4e\xa3\xd8\x9d\x1e\xb4\x8b\x9d\x62\x56\xf4\x4b\xd0\xc5\xcf\x96\x83\x2e\xef\xb7\x2d\x80\x5e\xc8\x14\x56\xc4\xb8\xc9\x65\x18\xfa\x10\x00\x59\xfa\x96\xf6\xfa\xe5\x4f\x93\x47\x8f\x54\xab\x18\x15\xba\xc9\x1f\x2f\x4e\x3e\xfb\x79\x35\x14\x5a\xf3\x43\xad\x91\x3c\x88\x45\xd9\x21\x28\x16\xca\x7f\x89\xc7\xf8\xd9\x72\x3c\x96\x38\xef\x22\x40\x39\x1f\xd0\x24\x64\x5d\x8c\x53\xa2\x31\xcf\x40\x7a\x06\x49\xb5\xa0\xb3\xee\x99\x0e\xee\x81\xb3\xb8\x71\xe9\x31\xa4\x71\x06\x34\xd1\x7e\x5b\x32\x3f\x0b\xa5\x4c\x17\x5e\xe8\xe3\x23\xe4\x4b\x38\x85\x30\x15\xd4\x8b\xe6\x7d\x30\x16\x09\xc5\x8f\x0b\xb3\x51\x2a\xfc\xd6\x50\x11\x32\x0c\x96\xd6\xcf\xa1\x6e\xa9\x7a\x98\x08\x18\x02\x2b\xd3\x0f\x13\xf1\xd6\xb5\x62\xae\x01\x70\x8e\x86\xa7\xd9\xad\x90\xd4\x05\x81\xb0\x9f\x4b\x7b\x33\x52\xc4\x18\x4a\x43\x45\xc7\x02\x82\xec\xfe\x52\x1c\x24\x43\xaa\x18\xe0\x85\x7e\x4a\xc2\xa7\x45\x83\x80\x92\x79\xb2\x5b\xc3\x6b\x45\xc6\x5f\xa6\x9c\x8a\x96\x22\x2c\xb6\x3c\x1f\x2e\x6d\xf5\x53\x35\x2d\x99\xf1\xac\xa1\x71\x66\x2e\xb3\x51\xd7\x26\x25\x9f\xcb\x6c\x12\xc1\xc6\xd6\x21\x0d\x4b\xbc\x5d\x54\xb5\xed\x00\x44\x36\xb5\x2e\x71\x66\x67\x46\x7f\x47\x91\xaf\x63\xe6\x4d\x44\x96\x8a\xd1\xd5\x75\xc8\xc4\x76\xa1\x9d\x30\xb7\xe3\xf6\xa2\xcc\x93\x03\x4a\x7d\x40\xa4\x34\x31\xc6\xaf\x57\x30\x51\xc1\xac\x6a\x07\xa2\xd0\x1a\xc9\xb9\x57\xdb\xc4\xa3\x75\x33\x4e\x3a\x9a\xf7\xa2\x0c\xb9\x8e\xfe\x8a\x77\x55\xe2\xaf\xca\xcf\x28\xa0\x92\x94\x5a\xaf\xba\x22\x66\xc6\x59\xab\x2a\x1e\xc1\x78\x1d\xb5\xe3\x31\xde\x46\x64\x4f\x26\x0e\x29\x6c\xc1\x90\x0b\x36\xa1\xf5\x4b\xde\xf9\xa7\xc5\x38\x90\x6a\x06\x65\x3a\xdd\xad\xe1\x16\x1a\xda\xf2\x2c\xbc\x12\xda\x38\xdc\x5c\x2d\x9c\xab\xdb\x9b\xb2\xc8\xa9\x9b\x11\x54\x07\x22\x83\x54\x67\xac\xb7\xba\xa6\xd1\x33\x93\xb7\xbc\xfe\xc1\x5e\x66\x65\xcf\xbc\x6d\xa6\x57\x8c\xfe\x5e\xbb\x97\x5c\x38\xe8\xf7\x6c\xe3\x7d\xe3\xc3\x99\xf4\xf3\x76\x3c\x71\xa1\x9a\x9d\x55\xaa\x4e\x7e\x10\xb6\x5e\x3e\x59\xbd\xd4\x94\x27\xc7\xe2\xa2\x2c\xd2\x76\x5f\xf2\x94\xb4\xdb\xb6\xab\xee\x17\x8c\x8f\xb7\xa9\xe4\x9f\x2a\x7f\xd9\x5c\x67\x07\x2a\xf3\x4e\xf5\x87\xa5\x73\xcf\x6d\x8a\x96\x0b\x6e\x10\x67\x97\xa2\x5a\x89\xd9\xea\xb7\x5a\xa6\x65\x25\x93\xac\xd9\xed\x76\xba\xa9\x54\x6d\x99\x5d\x7b\xbf\xd3\xb3\xcd\x0f\xda\x56\x2f\xf7\x26\xb7\x7a\xc7\xec\xb6\x6e\x1a\xfb\xbd\x62\x9a\xf9\xdb\xdc\x1b\xeb\x66\xe7\xa0\x98\x26\x7a\x93\x5f\xbd\xd5\x2f\xd9\x7e\xab\x5f\xbc\x68\xdc\xe9\xf4\x67\x72\xed\x77\xfa\x37\x6e\x26\x37\xf4\xba\xc6\xbe\x65\xb4\x7a\xed\xce\xbe\xfd\x9e\xd1\xbe\x6d\xee\x55\x96\x99\x99\xe5\x4a\xfd\x92\xcb\xf2\x17\xd4\xb6\xc6\x33\xb8\x95\xe9\x42\x86\x9d\x5a\x84\xb3\xdf\x35\xad\x4c\xc8\xe5\xc0\xae\x4d\xec\x02\x77\xea\xa9\x49\xd5\x1d\x14\xe7\x53\xc2\x72\x22\x87\xa1\x8f\xa1\xde\x00\x0a\x09\x18\x52\x36\xde\xfc\x10\xca\xa1\x7e\xdd\x99\x98\x4f\x59\x65\xa4\xd5\x17\x8a\x87\xe0\x60\x2f\xfe\x31\x59\x5d\x0f\xdb\x3e\x26\xa3\x5a\x8e\x12\x48\x9c\x81\x52\x9b\xbc\x13\x26\xa5\xdd\xe4\x7d\xad\x91\xe1\xbf\xf8\xe1\x61\xfc\x2d\x4e\xf4\x5d\xea\xc9\xb3\xdf\x4f\x1e\xff\x56\x91\xce\xaa\x46\xe7\x5b\x78\x35\x5d\x27\x65\x72\xfc\x49\x3d\xba\x43\x94\xef\xd4\x57\x4b\x42\xb6\xa3\x02\xaa\x3c\x29\x2c\xc9\xc5\xad\x88\xf6\x6a\x0d\xc0\xf7\x6c\x1f\x7b\xe7\x5a\x6a\xd6\x1d\x23\x9c\x55\x7f\x7a\x0e\xb7\xb1\x53\xdb\xbb\xea\xef\x1b\xb7\xa9\x0f\x4e\xab\xd2\x50\x7f\xc7\x8d\xff\x02\x00\x00\xff\xff\x91\x46\x95\x79\x0c\x2d\x00\x00")

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
	"proto/micro_mall_sku_proto/sku_business/sku.swagger.json": proto_micro_mall_sku_proto_sku_business_sku_swagger_json,
	"proto/micro_mall_users_proto/users/users.swagger.json":    proto_micro_mall_users_proto_users_users_swagger_json,
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"proto": &_bintree_t{nil, map[string]*_bintree_t{
		"micro_mall_sku_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"sku_business": &_bintree_t{nil, map[string]*_bintree_t{
				"sku.swagger.json": &_bintree_t{proto_micro_mall_sku_proto_sku_business_sku_swagger_json, map[string]*_bintree_t{}},
			}},
		}},
		"micro_mall_users_proto": &_bintree_t{nil, map[string]*_bintree_t{
			"users": &_bintree_t{nil, map[string]*_bintree_t{
				"users.swagger.json": &_bintree_t{proto_micro_mall_users_proto_users_users_swagger_json, map[string]*_bintree_t{}},
			}},
		}},
	}},
}}
