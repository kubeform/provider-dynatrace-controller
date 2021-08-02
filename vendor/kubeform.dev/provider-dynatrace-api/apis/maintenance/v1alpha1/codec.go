/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecMetadata{}).Type1()):           WindowSpecMetadataCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecSchedule{}).Type1()):           WindowSpecScheduleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScheduleRecurrence{}).Type1()): WindowSpecScheduleRecurrenceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScope{}).Type1()):              WindowSpecScopeCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecMetadata{}).Type1()):           WindowSpecMetadataCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecSchedule{}).Type1()):           WindowSpecScheduleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScheduleRecurrence{}).Type1()): WindowSpecScheduleRecurrenceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScope{}).Type1()):              WindowSpecScopeCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type WindowSpecMetadataCodec struct {
}

func (WindowSpecMetadataCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WindowSpecMetadata)(ptr) == nil
}

func (WindowSpecMetadataCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WindowSpecMetadata)(ptr)
	var objs []WindowSpecMetadata
	if obj != nil {
		objs = []WindowSpecMetadata{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecMetadata{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WindowSpecMetadataCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WindowSpecMetadata)(ptr) = WindowSpecMetadata{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WindowSpecMetadata

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecMetadata{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WindowSpecMetadata)(ptr) = objs[0]
			} else {
				*(*WindowSpecMetadata)(ptr) = WindowSpecMetadata{}
			}
		} else {
			*(*WindowSpecMetadata)(ptr) = WindowSpecMetadata{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WindowSpecMetadata

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecMetadata{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WindowSpecMetadata)(ptr) = obj
		} else {
			*(*WindowSpecMetadata)(ptr) = WindowSpecMetadata{}
		}
	default:
		iter.ReportError("decode WindowSpecMetadata", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WindowSpecScheduleCodec struct {
}

func (WindowSpecScheduleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WindowSpecSchedule)(ptr) == nil
}

func (WindowSpecScheduleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WindowSpecSchedule)(ptr)
	var objs []WindowSpecSchedule
	if obj != nil {
		objs = []WindowSpecSchedule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecSchedule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WindowSpecScheduleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WindowSpecSchedule)(ptr) = WindowSpecSchedule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WindowSpecSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WindowSpecSchedule)(ptr) = objs[0]
			} else {
				*(*WindowSpecSchedule)(ptr) = WindowSpecSchedule{}
			}
		} else {
			*(*WindowSpecSchedule)(ptr) = WindowSpecSchedule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WindowSpecSchedule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecSchedule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WindowSpecSchedule)(ptr) = obj
		} else {
			*(*WindowSpecSchedule)(ptr) = WindowSpecSchedule{}
		}
	default:
		iter.ReportError("decode WindowSpecSchedule", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WindowSpecScheduleRecurrenceCodec struct {
}

func (WindowSpecScheduleRecurrenceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WindowSpecScheduleRecurrence)(ptr) == nil
}

func (WindowSpecScheduleRecurrenceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WindowSpecScheduleRecurrence)(ptr)
	var objs []WindowSpecScheduleRecurrence
	if obj != nil {
		objs = []WindowSpecScheduleRecurrence{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScheduleRecurrence{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WindowSpecScheduleRecurrenceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WindowSpecScheduleRecurrence)(ptr) = WindowSpecScheduleRecurrence{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WindowSpecScheduleRecurrence

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScheduleRecurrence{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WindowSpecScheduleRecurrence)(ptr) = objs[0]
			} else {
				*(*WindowSpecScheduleRecurrence)(ptr) = WindowSpecScheduleRecurrence{}
			}
		} else {
			*(*WindowSpecScheduleRecurrence)(ptr) = WindowSpecScheduleRecurrence{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WindowSpecScheduleRecurrence

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScheduleRecurrence{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WindowSpecScheduleRecurrence)(ptr) = obj
		} else {
			*(*WindowSpecScheduleRecurrence)(ptr) = WindowSpecScheduleRecurrence{}
		}
	default:
		iter.ReportError("decode WindowSpecScheduleRecurrence", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WindowSpecScopeCodec struct {
}

func (WindowSpecScopeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WindowSpecScope)(ptr) == nil
}

func (WindowSpecScopeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WindowSpecScope)(ptr)
	var objs []WindowSpecScope
	if obj != nil {
		objs = []WindowSpecScope{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScope{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WindowSpecScopeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WindowSpecScope)(ptr) = WindowSpecScope{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WindowSpecScope

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScope{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WindowSpecScope)(ptr) = objs[0]
			} else {
				*(*WindowSpecScope)(ptr) = WindowSpecScope{}
			}
		} else {
			*(*WindowSpecScope)(ptr) = WindowSpecScope{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WindowSpecScope

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WindowSpecScope{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WindowSpecScope)(ptr) = obj
		} else {
			*(*WindowSpecScope)(ptr) = WindowSpecScope{}
		}
	default:
		iter.ReportError("decode WindowSpecScope", "unexpected JSON type")
	}
}
