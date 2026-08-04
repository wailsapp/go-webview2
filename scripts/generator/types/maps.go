package types

var idlTypeToGoType = map[string]string{
	"IUnknown":               "IUnknown",
	"EventRegistrationToken": "EventRegistrationToken",
	"LPWSTR":                 "string",
	"LPCWSTR":                "string",
	"HRESULT":                "uintptr",
	"UINT64":                 "uint64",
	"UINT32":                 "uint32",
	// Sized, not Go's int/uint. These appear as out-parameters, where the generated method
	// declares a local of the mapped type and hands the callee its address -- and a 64-bit local
	// receiving a 32-bit write keeps its zeroed high half, so the sign never extends. GetExitCode
	// returned 3221225477 for an exit code of -1073741819. Seven out-parameters were affected, of
	// which GetKeyEventLParam (bit 31 set on every key-up) and GetExitCode (negative NTSTATUS) are
	// wrong for ordinary inputs rather than only extreme ones.
	//
	// Lowercase "int" is the spelling the IDL actually uses for six of the seven and was absent
	// from this map entirely, so it fell through to Go's int by passthrough.
	"UINT":                   "uint32",
	"INT":                    "int32",
	"int":                    "int32",
	"INT32":                  "int32",
	"INT64":                  "int64",
	"BOOL":                   "bool",
	"BYTE":                   "uint8",
	"DWORD":                  "uint32",
	"double":                 "float64",
}

func IdlTypeToGoType(input string) string {
	result := idlTypeToGoType[input]
	if result == "" {
		return input
	}
	return result
}

// byValueArgument gives, per IDL type, the expression that passes a BY-VALUE in-parameter of that
// type to ComProc.Call. %s is the variable name. Only types that are neither a mapped scalar nor
// an enum reach here: the handle typedefs and the aggregates.
//
// The rule being encoded is the Windows x64 calling convention's, and it is not "aggregates go by
// reference". An aggregate of exactly 1, 2, 4 or 8 bytes is passed IN A REGISTER as an integer of
// that width; anything else is passed by reference. So POINT (8) and COREWEBVIEW2_COLOR (4) are
// register-passed while RECT (16) is not, and a single rule cannot cover both -- which is how
// every one of these ended up as &address, the answer that is only ever right for the large ones.
//
// Reinterpreting the struct through a same-width integer copies its layout rather than re-deriving
// the field order by hand, and it is a plain read, so unlike &address it also has no
// unsafe.Pointer lifetime question.
//
// The 8-byte entries assume a 64-bit uintptr. On windows/386 they would truncate, because an 8-byte
// by-value aggregate occupies two stack slots there and cannot be one Call argument at all -- the
// same shape of limit as the floats. That assumption is not new: PutPerformanceCount already passes
// a uint64 as one uintptr, so the generated binding has always been 64-bit-only in practice.
var byValueArgument = map[string]string{
	// Register-sized aggregates.
	"EventRegistrationToken": "uintptr(*(*uint64)(unsafe.Pointer(&%s)))", // struct{ value int64 }
	"POINT":                  "uintptr(*(*uint64)(unsafe.Pointer(&%s)))", // struct{ X, Y int32 }
	"COREWEBVIEW2_COLOR":     "uintptr(*(*uint32)(unsafe.Pointer(&%s)))", // struct{ A, R, G, B byte }
}

// uintptrTypedef lists the IDL types the binding declares as `type X uintptr` in com.tmpl. They are
// integers, so both questions this file answers have one obvious answer: pass the value, and use 0
// as the zero value.
//
// Only HWND currently appears as a by-value parameter and only HANDLE, HWND and HCURSOR as
// out-parameters; the rest are listed because for a uintptr typedef there is exactly one right
// answer, so pre-classifying them is not the guessing that byRefAggregate exists to prevent.
//
// VARIANT is deliberately ABSENT even though com.tmpl declares it `type VARIANT uintptr`. A real
// VARIANT is a 16-byte tagged union, com.tmpl says so itself ("NOTE: For sure, this is wrong!"), and
// the IDL only ever uses VARIANT* -- so a by-value VARIANT should hit the generator's error and make
// someone look, not quietly become an integer.
var uintptrTypedef = map[string]bool{
	"HANDLE":    true,
	"HBRUSH":    true,
	"HCURSOR":   true,
	"HICON":     true,
	"HINSTANCE": true,
	"HMENU":     true,
	"HMODULE":   true,
	"HWND":      true,
}

// byRefAggregate are the by-value in-parameter types that really are passed by address, with the
// width that makes them so. They are listed rather than left to a default so that a type in
// NEITHER table is a generator error instead of a silent guess -- the failure mode this entire
// family is made of. Adding a struct to the IDL should cost one line here and a look at its size.
var byRefAggregate = map[string]int{
	"RECT":                             16, // struct{ Left, Top, Right, Bottom int32 }
	"COREWEBVIEW2_PHYSICAL_KEY_STATUS": 24, // 2x UINT32 + 4x BOOL
}
