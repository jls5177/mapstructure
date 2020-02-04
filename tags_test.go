package mapstructure

import "testing"

type TagBasic struct {
	Vstring string      `json:"string"`
	Vint    int         `json:"int"`
	Vuint   uint        `json:"uint"`
	Vbool   bool        `json:"bool"`
	Vfloat  float64     `json:"float"`
	Vextra  string      `json:"extra"`
	vsilent bool        `json:"silent"`
	Vdata   interface{} `json:"data"`
}

type TagBasicInline struct {
	Test TagBasic `json:",inline"`
}

type TagEmbeddedSquash struct {
	TagBasic `mapstructure:",squash"`
}

func decodeWithFallback(input, output interface{}, meta *Metadata) error {
	config := &DecoderConfig{
		Metadata: meta,
		Result:   output,
		FallbackToJSONTags:true,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func TestDecode_TagBasicInline(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{
		"string": "foo",
		"Vuint": 32,
	}

	var result TagBasicInline
	err := decodeWithFallback(input, &result, nil)
	if err != nil {
		t.Fatalf("got an err: %s", err.Error())
	}

	if result.Test.Vstring != "foo" {
		t.Errorf("vstring value should be 'foo': %#v", result.Test.Vstring)
	} else if result.Test.Vuint == 32 {
		t.Errorf("vuint value should not be set")
	}
}

func TestDecode_TagBasicSquash(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{
		"string": "foo",
		"Vuint": 32,
	}

	var result TagEmbeddedSquash
	err := decodeWithFallback(input, &result, nil)
	if err != nil {
		t.Fatalf("got an err: %s", err.Error())
	}

	if result.Vstring != "foo" {
		t.Errorf("vstring value should be 'foo': %#v", result.Vstring)
	} else if result.Vuint == 32 {
		t.Errorf("vuint value should not be set")
	}
}
