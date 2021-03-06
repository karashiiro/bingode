// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TownTable struct {
	_tab flatbuffers.Table
}

func GetRootAsTownTable(buf []byte, offset flatbuffers.UOffsetT) *TownTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TownTable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TownTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TownTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TownTable) Towns(obj *Town, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TownTable) TownsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TownTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TownTableAddTowns(builder *flatbuffers.Builder, Towns flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Towns), 0)
}
func TownTableStartTownsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TownTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Town struct {
	_tab flatbuffers.Table
}

func GetRootAsTown(buf []byte, offset flatbuffers.UOffsetT) *Town {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Town{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Town) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Town) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Town) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Town) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Town) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Town) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Town) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Town) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TownStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func TownAddId(builder *flatbuffers.Builder, Id uint32) {
	builder.PrependUint32Slot(0, Id, 0)
}
func TownAddNameEn(builder *flatbuffers.Builder, NameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(NameEn), 0)
}
func TownAddNameFr(builder *flatbuffers.Builder, NameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(NameFr), 0)
}
func TownAddNameDe(builder *flatbuffers.Builder, NameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NameDe), 0)
}
func TownAddNameJa(builder *flatbuffers.Builder, NameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(NameJa), 0)
}
func TownEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
