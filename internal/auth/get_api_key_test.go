package auth_test // use auth_test if you want black‑box testing; change to `auth` for white‑box

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantKey   string
		wantErr   error
		wantErrIs error // for errors.Is comparison (optional)
	}{
		{
			name:    "missing Authorization header",
			headers: http.Header{},
			wantKey: "",
			wantErr: auth.ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth.GetAPIKey(tt.headers)

			// compare the returned key
			if got != tt.wantKey {
				t.Fatalf("expected key %q, got %q", tt.wantKey, got)
			}

			// compare the error – either exact match or errors.Is when ErrNoAuthHeaderIncluded is used
			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				// for the defined sentinel error, use errors.Is
				if errors.Is(err, auth.ErrNoAuthHeaderIncluded) {
					if !errors.Is(err, tt.wantErr) {
						t.Fatalf("expected error %v, got %v", tt.wantErr, err)
					}
				} else if err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("did not expect error, got %v", err)
			}
		})
	}
}
