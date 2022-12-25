package format

import "testing"

func Test_formatString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test format string",
			args: args{
				s: "ezra.ai/cluster",
			},
			want: "ezra_ai_cluster",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatString(tt.args.s); got != tt.want {
				t.Errorf("formatString() = %v, want %v", got, tt.want)
			}
		})
	}
}
