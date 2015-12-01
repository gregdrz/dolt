package test

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/nomdl/codegen/test/gen"
	"github.com/attic-labs/noms/types"
)

func TestDef(t *testing.T) {
	assert := assert.New(t)
	cs := chunks.NewMemoryStore()

	def := gen.StructDef{"hi", true}
	st := def.New(cs)

	def2 := st.Def()
	st2 := def.New(cs)

	assert.Equal(def, def2)
	assert.True(st.Equals(st2))

	st3 := gen.NewStruct(cs)
	st3 = st3.SetS("hi").SetB(true)
	assert.Equal("hi", st3.S())
	assert.Equal(true, st3.B())
}

func TestValue(t *testing.T) {
	assert := assert.New(t)
	cs := chunks.NewMemoryStore()

	def := gen.StructDef{"hi", true}
	var st types.Value
	st = def.New(cs)
	st2 := st.(gen.Struct)
	assert.True(st.Equals(st2))
}

func TestType(t *testing.T) {
	assert := assert.New(t)
	cs := chunks.NewMemoryStore()

	def := gen.StructDef{"hi", true}
	st := def.New(cs)
	typ := st.Type()
	assert.EqualValues(0, typ.Ordinal())
	assert.Equal(types.UnresolvedKind, typ.Kind())
}

func TestStructChunks(t *testing.T) {
	assert := assert.New(t)
	cs := chunks.NewMemoryStore()

	st := gen.StructDef{"hi", true}.New(cs)
	chunks := st.Chunks()

	// One chunk for the Type
	assert.Len(chunks, 1)
}
