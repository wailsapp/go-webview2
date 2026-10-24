package edge

import (
	"errors"
	"testing"
)

type fakePermissionRequestedEventArgs struct {
	uri      string
	uriErr   error
	kind     CoreWebView2PermissionKind
	kindErr  error
	state    CoreWebView2PermissionState
	putErr   error
	putCalls int
}

func (a *fakePermissionRequestedEventArgs) GetURI() (string, error) {
	return a.uri, a.uriErr
}

func (a *fakePermissionRequestedEventArgs) GetPermissionKind() (CoreWebView2PermissionKind, error) {
	return a.kind, a.kindErr
}

func (a *fakePermissionRequestedEventArgs) PutState(state CoreWebView2PermissionState) error {
	a.state = state
	a.putCalls++
	return a.putErr
}

func newPermissionTestChromium() *Chromium {
	return &Chromium{
		permissions:         make(map[CoreWebView2PermissionKind]CoreWebView2PermissionState),
		globalErrorCallback: func(error) {},
	}
}

func TestPermissionHandlerConsultsCallbackBeforeGlobalState(t *testing.T) {
	args := &fakePermissionRequestedEventArgs{
		uri:  "http://wails.localhost/",
		kind: CoreWebView2PermissionKindMicrophone,
	}
	chromium := newPermissionTestChromium()
	chromium.SetGlobalPermission(CoreWebView2PermissionStateAllow)
	chromium.SetPermission(CoreWebView2PermissionKindMicrophone, CoreWebView2PermissionStateAllow)

	called := false
	chromium.PermissionRequestedCallback = func(uri string, kind CoreWebView2PermissionKind, uriErr error) CoreWebView2PermissionState {
		called = true
		if uri != args.uri {
			t.Errorf("callback URI = %q, want %q", uri, args.uri)
		}
		if kind != args.kind {
			t.Errorf("callback permission kind = %v, want %v", kind, args.kind)
		}
		if uriErr != nil {
			t.Errorf("callback URI error = %v, want nil", uriErr)
		}
		return CoreWebView2PermissionStateDeny
	}

	chromium.handlePermissionRequested(args)

	if !called {
		t.Fatal("permission callback was not called")
	}
	if args.state != CoreWebView2PermissionStateDeny {
		t.Errorf("permission state = %v, want %v", args.state, CoreWebView2PermissionStateDeny)
	}
	if args.putCalls != 1 {
		t.Errorf("PutState calls = %d, want 1", args.putCalls)
	}

	t.Run("nil callback preserves global per-kind and default states", func(t *testing.T) {
		tests := []struct {
			name      string
			configure func(*Chromium)
			want      CoreWebView2PermissionState
		}{
			{
				name: "global state takes precedence",
				configure: func(chromium *Chromium) {
					chromium.SetGlobalPermission(CoreWebView2PermissionStateDeny)
					chromium.SetPermission(CoreWebView2PermissionKindMicrophone, CoreWebView2PermissionStateAllow)
				},
				want: CoreWebView2PermissionStateDeny,
			},
			{
				name: "per-kind state applies without global state",
				configure: func(chromium *Chromium) {
					chromium.SetPermission(CoreWebView2PermissionKindMicrophone, CoreWebView2PermissionStateAllow)
				},
				want: CoreWebView2PermissionStateAllow,
			},
			{
				name:      "default state applies without configured state",
				configure: func(*Chromium) {},
				want:      CoreWebView2PermissionStateDefault,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				args := &fakePermissionRequestedEventArgs{kind: CoreWebView2PermissionKindMicrophone}
				chromium := newPermissionTestChromium()
				test.configure(chromium)

				chromium.handlePermissionRequested(args)

				if args.state != test.want {
					t.Errorf("permission state = %v, want %v", args.state, test.want)
				}
			})
		}
	})
}

func TestPermissionHandlerAppliesCallbackState(t *testing.T) {
	tests := []struct {
		name string
		want CoreWebView2PermissionState
	}{
		{name: "default", want: CoreWebView2PermissionStateDefault},
		{name: "allow", want: CoreWebView2PermissionStateAllow},
		{name: "deny", want: CoreWebView2PermissionStateDeny},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			args := &fakePermissionRequestedEventArgs{kind: CoreWebView2PermissionKindCamera}
			chromium := newPermissionTestChromium()
			chromium.PermissionRequestedCallback = func(string, CoreWebView2PermissionKind, error) CoreWebView2PermissionState {
				return test.want
			}

			chromium.handlePermissionRequested(args)

			if args.state != test.want {
				t.Errorf("permission state = %v, want %v", args.state, test.want)
			}
			if args.putCalls != 1 {
				t.Errorf("PutState calls = %d, want 1", args.putCalls)
			}
		})
	}
}

func TestPermissionHandlerDeniesOnURIError(t *testing.T) {
	uriErr := errors.New("URI unavailable")
	args := &fakePermissionRequestedEventArgs{
		uriErr: uriErr,
		kind:   CoreWebView2PermissionKindMicrophone,
	}
	chromium := newPermissionTestChromium()
	chromium.SetGlobalPermission(CoreWebView2PermissionStateAllow)
	chromium.SetPermission(CoreWebView2PermissionKindMicrophone, CoreWebView2PermissionStateAllow)
	called := false
	chromium.PermissionRequestedCallback = func(uri string, kind CoreWebView2PermissionKind, gotURIErr error) CoreWebView2PermissionState {
		called = true
		if uri != "" {
			t.Errorf("callback URI = %q, want empty URI on lookup error", uri)
		}
		if kind != args.kind {
			t.Errorf("callback permission kind = %v, want %v", kind, args.kind)
		}
		if !errors.Is(gotURIErr, uriErr) {
			t.Errorf("callback URI error = %v, want %v", gotURIErr, uriErr)
		}
		if gotURIErr != nil {
			return CoreWebView2PermissionStateDeny
		}
		return CoreWebView2PermissionStateAllow
	}

	chromium.handlePermissionRequested(args)

	if !called {
		t.Fatal("permission callback was not called")
	}
	if args.state != CoreWebView2PermissionStateDeny {
		t.Errorf("permission state = %v, want %v", args.state, CoreWebView2PermissionStateDeny)
	}
	if args.putCalls != 1 {
		t.Errorf("PutState calls = %d, want 1", args.putCalls)
	}
}
