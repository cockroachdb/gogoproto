// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oneofembed.proto

package proto

import (
	fmt "fmt"
	_ "github.com/cockroachdb/gogoproto/gogoproto"
	github_com_cockroachdb_gogoproto_jsonpb "github.com/cockroachdb/gogoproto/jsonpb"
	github_com_cockroachdb_gogoproto_proto "github.com/cockroachdb/gogoproto/proto"
	proto "github.com/cockroachdb/gogoproto/proto"
	math "math"
	math_rand "math/rand"
	testing "testing"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestFooProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, false)
	dAtA, err := github_com_cockroachdb_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo{}
	if err := github_com_cockroachdb_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cockroachdb_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestBarProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar(popr, false)
	dAtA, err := github_com_cockroachdb_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar{}
	if err := github_com_cockroachdb_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	littlefuzz := make([]byte, len(dAtA))
	copy(littlefuzz, dAtA)
	for i := range dAtA {
		dAtA[i] = byte(popr.Intn(256))
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_cockroachdb_gogoproto_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestFooJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, true)
	marshaler := github_com_cockroachdb_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Foo{}
	err = github_com_cockroachdb_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestBarJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar(popr, true)
	marshaler := github_com_cockroachdb_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &Bar{}
	err = github_com_cockroachdb_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestFooProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, true)
	dAtA := github_com_cockroachdb_gogoproto_proto.MarshalTextString(p)
	msg := &Foo{}
	if err := github_com_cockroachdb_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFooProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedFoo(popr, true)
	dAtA := github_com_cockroachdb_gogoproto_proto.CompactTextString(p)
	msg := &Foo{}
	if err := github_com_cockroachdb_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBarProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar(popr, true)
	dAtA := github_com_cockroachdb_gogoproto_proto.MarshalTextString(p)
	msg := &Bar{}
	if err := github_com_cockroachdb_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestBarProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedBar(popr, true)
	dAtA := github_com_cockroachdb_gogoproto_proto.CompactTextString(p)
	msg := &Bar{}
	if err := github_com_cockroachdb_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestFooCompare(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedFoo(popr, false)
	dAtA, err := github_com_cockroachdb_gogoproto_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Foo{}
	if err := github_com_cockroachdb_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if c := p.Compare(msg); c != 0 {
		t.Fatalf("%#v !Compare %#v, since %d", msg, p, c)
	}
	p2 := NewPopulatedFoo(popr, false)
	c := p.Compare(p2)
	c2 := p2.Compare(p)
	if c != (-1 * c2) {
		t.Errorf("p.Compare(p2) = %d", c)
		t.Errorf("p2.Compare(p) = %d", c2)
		t.Errorf("p = %#v", p)
		t.Errorf("p2 = %#v", p2)
	}
}
func TestBarCompare(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedBar(popr, false)
	dAtA, err := github_com_cockroachdb_gogoproto_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &Bar{}
	if err := github_com_cockroachdb_gogoproto_proto.Unmarshal(dAtA, msg); err != nil {
		panic(err)
	}
	if c := p.Compare(msg); c != 0 {
		t.Fatalf("%#v !Compare %#v, since %d", msg, p, c)
	}
	p2 := NewPopulatedBar(popr, false)
	c := p.Compare(p2)
	c2 := p2.Compare(p)
	if c != (-1 * c2) {
		t.Errorf("p.Compare(p2) = %d", c)
		t.Errorf("p2.Compare(p) = %d", c2)
		t.Errorf("p = %#v", p)
		t.Errorf("p2 = %#v", p2)
	}
}

//These tests are generated by github.com/cockroachdb/gogoproto/plugin/testgen
