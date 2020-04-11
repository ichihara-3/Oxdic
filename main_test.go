package main

import (
	"os"
	"testing"
)

func TestTranslate(t *testing.T) {
	t.Run("translate test", func(t *testing.T) {
		app_id := os.Getenv("OXDIC_APP_ID")
		app_key := os.Getenv("OXDIC_APP_KEY")
		tests := []struct {
			term string
			lang string
		}{
			{"test", "en-us"},
			{"swimming", "en-us"},
		}

		for _, test := range tests {
			_, err := search(test.term, test.lang, app_id, app_key)
			if err != nil {
				t.Errorf("cannnot search %s", err)
			}

		}

	})

	t.Run("run test", func(t *testing.T) {
		tests := []struct {
			args     []string
			endpoint string
			want     int
		}{
			{[]string{}, "", -1},
			{[]string{"hello"}, "", 0},
			{[]string{}, "", 0},
			{[]string{}, "invalid_endpoint", -1},
		}

		for i, test := range tests {
			if test.endpoint != "" {
				os.Setenv("OXDIC_ENDPOINT", test.endpoint)
			}

			result := run(test.args)
			if test.want != result {
				t.Errorf(" %d result is wrong. want=%d, actual=%d", i, test.want, result)
			}
		}
	})
}
