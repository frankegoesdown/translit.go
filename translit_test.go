// Copyright (c) 2011 Dmitry Chestnykh
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package translit

import "testing"

func TestToLatin(t *testing.T) {
	tests := map[string]string{
		"частица": "chastica",
		"Частица": "Chastica",
		"проЩай":  "proSchaj",
		"ЩИ":      "SCHI",
		"Щи":      "Schi",
		"щи":      "schi",
		"Лебедь, Рак и Щука":     "Lebed, Rak i Schuka",
		"Это ДОЛЖНО работать":    "Eto DOLZHNO rabotat",
		"English must work тоже": "English must work tozhe",
	}
	for k, v := range tests {
		tr := ToLatin(k, RussianASCII)
		if tr != v {
			t.Errorf("\"%s\" expected %q got %q", k, v, tr)
		}
	}
}
