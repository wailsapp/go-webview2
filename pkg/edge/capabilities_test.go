package edge

import "testing"

func Test_hasCapability(t *testing.T) {
	type args struct {
		webview2RuntimeVersion string
		capability             Capability
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should support GetAdditionalObjects if version is 113.0.1774.30",
			args: args{
				webview2RuntimeVersion: "113.0.1774.30",
				capability:             GetAdditionalObjects,
			},
			want: true,
		},
		{
			name: "should support GetAdditionalObjects if version is 115.0.1901.177",
			args: args{
				webview2RuntimeVersion: "115.0.1901.177",
				capability:             GetAdditionalObjects,
			},
			want: true,
		},
		{
			name: "should not support GetAdditionalObjects if version is 113.0.1724.0",
			args: args{
				webview2RuntimeVersion: "113.0.1724.0",
				capability:             GetAdditionalObjects,
			},
			want: false,
		},
		{
			name: "should support SwipeNavigation if version is equal to minimum",
			args: args{
				webview2RuntimeVersion: "94.0.992.31",
				capability:             SwipeNavigation,
			},
			want: true,
		},
		{
			name: "should support SwipeNavigation if version is above minimum",
			args: args{
				webview2RuntimeVersion: "115.0.1901.177",
				capability:             SwipeNavigation,
			},
			want: true,
		},
		{
			name: "should not support SwipeNavigation if version is below minimum",
			args: args{
				webview2RuntimeVersion: "93.0.992.31",
				capability:             SwipeNavigation,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCapability(tt.args.webview2RuntimeVersion, tt.args.capability); got != tt.want {
				t.Errorf("hasCapability() = %v, want %v", got, tt.want)
			}
		})
	}
}
