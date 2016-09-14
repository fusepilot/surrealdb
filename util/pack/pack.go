// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pack

import (
	"reflect"
	"time"

	"github.com/ugorji/go/codec"

	"github.com/abcum/surreal/sql"
)

var jh codec.JsonHandle
var ch codec.CborHandle
var bh codec.BincHandle
var mh codec.MsgpackHandle

func init() {

	// JSONHandle

	jh.Canonical = true
	jh.CheckCircularRef = false
	jh.AsSymbols = codec.AsSymbolDefault
	jh.SliceType = reflect.TypeOf([]interface{}(nil))
	jh.MapType = reflect.TypeOf(map[string]interface{}(nil))

	// CBORHandle

	ch.Canonical = true
	ch.CheckCircularRef = false
	ch.AsSymbols = codec.AsSymbolDefault
	ch.SliceType = reflect.TypeOf([]interface{}(nil))
	ch.MapType = reflect.TypeOf(map[string]interface{}(nil))

	// BINCHandle

	bh.Canonical = true
	bh.CheckCircularRef = false
	bh.AsSymbols = codec.AsSymbolDefault
	bh.SliceType = reflect.TypeOf([]interface{}(nil))
	bh.MapType = reflect.TypeOf(map[string]interface{}(nil))

	bh.SetBytesExt(reflect.TypeOf(time.Time{}), 1, extTime{})
	bh.SetBytesExt(reflect.TypeOf(sql.All{}), 2, extSqlAll{})
	bh.SetBytesExt(reflect.TypeOf(sql.Asc{}), 3, extSqlAsc{})
	bh.SetBytesExt(reflect.TypeOf(sql.Desc{}), 4, extSqlDesc{})
	bh.SetBytesExt(reflect.TypeOf(sql.Null{}), 5, extSqlNull{})
	bh.SetBytesExt(reflect.TypeOf(sql.Void{}), 6, extSqlVoid{})
	bh.SetBytesExt(reflect.TypeOf(sql.Empty{}), 7, extSqlEmpty{})
	bh.SetBytesExt(reflect.TypeOf(sql.Ident{}), 8, extSqlIdent{})
	bh.SetBytesExt(reflect.TypeOf(sql.Table{}), 9, extSqlTable{})
	bh.SetBytesExt(reflect.TypeOf(sql.Thing{}), 10, extSqlThing{})
	bh.SetBytesExt(reflect.TypeOf(sql.Field{}), 11, extSqlField{})
	bh.SetBytesExt(reflect.TypeOf(sql.Group{}), 12, extSqlGroup{})

	// PACKHandle

	mh.WriteExt = true
	mh.Canonical = true
	mh.RawToString = true
	mh.CheckCircularRef = false
	mh.AsSymbols = codec.AsSymbolDefault
	mh.SliceType = reflect.TypeOf([]interface{}(nil))
	mh.MapType = reflect.TypeOf(map[string]interface{}(nil))

	mh.SetBytesExt(reflect.TypeOf(time.Time{}), 1, extTime{})
	mh.SetBytesExt(reflect.TypeOf(sql.All{}), 2, extSqlAll{})
	mh.SetBytesExt(reflect.TypeOf(sql.Asc{}), 3, extSqlAsc{})
	mh.SetBytesExt(reflect.TypeOf(sql.Desc{}), 4, extSqlDesc{})
	mh.SetBytesExt(reflect.TypeOf(sql.Null{}), 5, extSqlNull{})
	mh.SetBytesExt(reflect.TypeOf(sql.Void{}), 6, extSqlVoid{})
	mh.SetBytesExt(reflect.TypeOf(sql.Empty{}), 7, extSqlEmpty{})
	mh.SetBytesExt(reflect.TypeOf(sql.Ident{}), 8, extSqlIdent{})
	mh.SetBytesExt(reflect.TypeOf(sql.Table{}), 9, extSqlTable{})
	mh.SetBytesExt(reflect.TypeOf(sql.Thing{}), 10, extSqlThing{})
	mh.SetBytesExt(reflect.TypeOf(sql.Field{}), 11, extSqlField{})
	mh.SetBytesExt(reflect.TypeOf(sql.Group{}), 12, extSqlGroup{})

}

// ToJSON encodes a data object to a JSON byte slice.
func ToJSON(src interface{}) (dst []byte) {
	codec.NewEncoderBytes(&dst, &jh).Encode(src)
	return
}

// FromJSON decodes a JSON byte slice into a data object.
func FromJSON(src []byte, dst interface{}) {
	codec.NewDecoderBytes(src, &jh).Decode(dst)
	return
}

// ToCBOR encodes a data object to a CBOR byte slice.
func ToCBOR(src interface{}) (dst []byte) {
	codec.NewEncoderBytes(&dst, &ch).Encode(src)
	return
}

// FromCBOR decodes a CBOR byte slice into a data object.
func FromCBOR(src []byte, dst interface{}) {
	codec.NewDecoderBytes(src, &ch).Decode(dst)
	return
}

// ToBINC encodes a data object to a BINC byte slice.
func ToBINC(src interface{}) (dst []byte) {
	codec.NewEncoderBytes(&dst, &bh).Encode(src)
	return
}

// FromBINC decodes a BINC byte slice into a data object.
func FromBINC(src []byte, dst interface{}) {
	codec.NewDecoderBytes(src, &bh).Decode(dst)
	return
}

// ToPACK encodes a data object to a MsgPack byte slice.
func ToPACK(src interface{}) (dst []byte) {
	codec.NewEncoderBytes(&dst, &mh).Encode(src)
	return
}

// FromPACK decodes a MsgPack byte slice into a data object.
func FromPACK(src []byte, dst interface{}) {
	codec.NewDecoderBytes(src, &mh).Decode(dst)
	return
}