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
			name: "should support getAdditionalObjects if version is 113.0.1774.30",
			args: args{
				webview2RuntimeVersion: "113.0.1774.30",
				capability:             getAdditionalObjects,
			},
			want: true,
		},
		{
			name: "should support getAdditionalObjects if version is 115.0.1901.177",
			args: args{
				webview2RuntimeVersion: "115.0.1901.177",
				capability:             getAdditionalObjects,
			},
			want: true,
		},
		{
			name: "should not support getAdditionalObjects if version is 113.0.1724.0",
			args: args{
				webview2RuntimeVersion: "113.0.1724.0",
				capability:             getAdditionalObjects,
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
