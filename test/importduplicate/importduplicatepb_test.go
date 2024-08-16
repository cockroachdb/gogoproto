// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: importduplicate.proto

package importduplicate

import (
	fmt "fmt"
	_ "github.com/cockroachdb/gogoproto/gogoproto"
	github_com_cockroachdb_gogoproto_jsonpb "github.com/cockroachdb/gogoproto/jsonpb"
	github_com_cockroachdb_gogoproto_proto "github.com/cockroachdb/gogoproto/proto"
	proto "github.com/cockroachdb/gogoproto/proto"
	_ "github.com/cockroachdb/gogoproto/test/importduplicate/proto"
	_ "github.com/cockroachdb/gogoproto/test/importduplicate/sortkeys"
	go_parser "go/parser"
	math "math"
	math_rand "math/rand"
	testing "testing"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func TestMapAndSortKeysProto(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMapAndSortKeys(popr, false)
	dAtA, err := github_com_cockroachdb_gogoproto_proto.Marshal(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &MapAndSortKeys{}
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

func TestMapAndSortKeysJSON(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMapAndSortKeys(popr, true)
	marshaler := github_com_cockroachdb_gogoproto_jsonpb.Marshaler{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	msg := &MapAndSortKeys{}
	err = github_com_cockroachdb_gogoproto_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Json Equal %#v", seed, msg, p)
	}
}
func TestMapAndSortKeysProtoText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMapAndSortKeys(popr, true)
	dAtA := github_com_cockroachdb_gogoproto_proto.MarshalTextString(p)
	msg := &MapAndSortKeys{}
	if err := github_com_cockroachdb_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestMapAndSortKeysProtoCompactText(t *testing.T) {
	seed := time.Now().UnixNano()
	popr := math_rand.New(math_rand.NewSource(seed))
	p := NewPopulatedMapAndSortKeys(popr, true)
	dAtA := github_com_cockroachdb_gogoproto_proto.CompactTextString(p)
	msg := &MapAndSortKeys{}
	if err := github_com_cockroachdb_gogoproto_proto.UnmarshalText(dAtA, msg); err != nil {
		t.Fatalf("seed = %d, err = %v", seed, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("seed = %d, %#v !Proto %#v", seed, msg, p)
	}
}

func TestMapAndSortKeysGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedMapAndSortKeys(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		t.Fatal(err)
	}
}

//These tests are generated by github.com/cockroachdb/gogoproto/plugin/testgen
