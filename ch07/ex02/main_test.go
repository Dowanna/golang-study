package ex02

import (
	"os"
	"testing"
)

func TestCountWriter(t *testing.T) {

	for _, td := range []string{
		"hello there\n",
		"hoge fuga piyo\n",
	} {
		w, count := CountingWriter(os.Stdout)
		_, err := w.Write([]byte(td))

		if err != nil {
			t.Errorf("error! %v", err)
		}

		if int64(len(td)) != *count {
			t.Errorf("error! expected %d, got %d\n", len(td), *count)
		}
	}
}
