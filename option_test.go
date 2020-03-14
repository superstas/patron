package patron

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	phttp "github.com/beatlabs/patron/component/http"
)

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func TestRoutes(t *testing.T) {
	type args struct {
		rr []phttp.Route
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"failure due to empty routes", args{rr: []phttp.Route{}}, true},
		{"failure due to nil routes", args{rr: nil}, true},
		{"success", args{rr: []phttp.Route{phttp.NewRoute("/", "GET", nil, true, nil)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New("test", "1.0.0")
			assert.NoError(t, err)
			err = Routes(tt.args.rr)(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMiddlewares(t *testing.T) {
	type args struct {
		mm []phttp.MiddlewareFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"success", args{mm: []phttp.MiddlewareFunc{middleware}}, false},
		{"failure because empty", args{mm: []phttp.MiddlewareFunc{}}, true},
		{"failure because nil", args{mm: nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New("test", "1.0.0")
			assert.NoError(t, err)
			err = Middlewares(tt.args.mm...)(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAliveCheck(t *testing.T) {
	type args struct {
		acf phttp.AliveCheckFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"failure due to nil liveness check", args{acf: nil}, true},
		{"success", args{acf: phttp.DefaultAliveCheck}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New("test", "1.0.0")
			assert.NoError(t, err)
			err = AliveCheck(tt.args.acf)(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestReadyCheck(t *testing.T) {
	type args struct {
		rcf phttp.ReadyCheckFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"failure due to nil liveness check", args{rcf: nil}, true},
		{"success", args{rcf: phttp.DefaultReadyCheck}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New("test", "1.0.0")
			assert.NoError(t, err)
			err = ReadyCheck(tt.args.rcf)(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestComponents(t *testing.T) {
	type args struct {
		c Component
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"failure due to empty components", args{}, true},
		{"failure due to nil components", args{c: nil}, true},
		{"success", args{c: &testComponent{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New("test", "1.0.0")
			assert.NoError(t, err)
			err = Components(tt.args.c)(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSIGHUP(t *testing.T) {
	type args struct {
		handler func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "nil handler", args: args{handler: nil}, wantErr: true},
		{name: "success", args: args{handler: func() {}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := New("test", "1.0.0")
			assert.NoError(t, err)
			err = SIGHUP(tt.args.handler)(s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
