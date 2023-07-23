package generator

import (
	"bytes"
	"github.com/alecthomas/participle/v2"
	"github.com/matryer/is"
	"os"
	"testing"
)

var testData = []byte(`

[uuid(26d34152-879f-4065-bea2-3daa2cfadfb8), version(1.0)]
library WebView2 {

[v1_enum]
typedef enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME {
    /// Auto color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO,

    /// Light color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT,

    /// Dark color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK
} COREWEBVIEW2_PREFERRED_COLOR_SCHEME;


[v1_enum]
typedef enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME1 {
    /// Auto color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO1 = 1,

    /// Light color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT1 = 2,

    /// Dark color scheme.
    COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK1 = 3,
} COREWEBVIEW2_PREFERRED_COLOR_SCHEME;


[v1_enum]
typedef enum COREWEBVIEW2_PREFERRED_COLOR_SCHEME2 {
   /// Auto color scheme.
   COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO2 = 1 << 1,

   /// Light color scheme.
   COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT2 = 1 << 2,

   /// Dark color scheme.
   COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK2 = 1 << 3
} COREWEBVIEW2_PREFERRED_COLOR_SCHEME;

}`)

func TestEnum(t *testing.T) {

	i := is.New(t)

	var buf bytes.Buffer
	buf.Write(testData)

	idl, err := Parser.Parse("", &buf, participle.Trace(os.Stdout))
	i.NoErr(err)

	err = idl.Process()
	i.NoErr(err)

	err = idl.Generate("./test/enum")
	i.NoErr(err)
}
