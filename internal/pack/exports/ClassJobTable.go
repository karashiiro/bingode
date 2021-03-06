// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package exports

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ClassJobTable struct {
	_tab flatbuffers.Table
}

func GetRootAsClassJobTable(buf []byte, offset flatbuffers.UOffsetT) *ClassJobTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ClassJobTable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ClassJobTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ClassJobTable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ClassJobTable) ClassJobs(obj *ClassJob, j int) bool {
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

func (rcv *ClassJobTable) ClassJobsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ClassJobTableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ClassJobTableAddClassJobs(builder *flatbuffers.Builder, ClassJobs flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ClassJobs), 0)
}
func ClassJobTableStartClassJobsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ClassJobTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type ClassJob struct {
	_tab flatbuffers.Table
}

func GetRootAsClassJob(buf []byte, offset flatbuffers.UOffsetT) *ClassJob {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ClassJob{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ClassJob) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ClassJob) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ClassJob) Id() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ClassJob) MutateId(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *ClassJob) NameEn() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ClassJob) NameFr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ClassJob) NameDe() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ClassJob) NameJa() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ClassJobStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ClassJobAddId(builder *flatbuffers.Builder, Id uint32) {
	builder.PrependUint32Slot(0, Id, 0)
}
func ClassJobAddNameEn(builder *flatbuffers.Builder, NameEn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(NameEn), 0)
}
func ClassJobAddNameFr(builder *flatbuffers.Builder, NameFr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(NameFr), 0)
}
func ClassJobAddNameDe(builder *flatbuffers.Builder, NameDe flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NameDe), 0)
}
func ClassJobAddNameJa(builder *flatbuffers.Builder, NameJa flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(NameJa), 0)
}
func ClassJobEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
