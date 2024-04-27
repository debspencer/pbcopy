package pbcopier

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	a := assert.New(t)

	pbreader, err := NewPbReader("test-files/test.pb.go")
	a.NoError(err)

	packageName, err := pbreader.FindPackage()
	a.NoError(err)

	copyGo := path.Join(t.TempDir(), "copy.go")
	pbwriter, err := NewPbWriter(copyGo, packageName)
	a.NoError(err)

	pbCopier := NewPbCopier(pbreader, pbwriter)

	err = pbCopier.MkCopy("Base") // name of test-files/test.proto main message
	a.NoError(err)

	pbwriter.Flush()

	newFile, err := os.ReadFile(copyGo)
	a.NoError(err)

	cmpFile, err := os.ReadFile("test-files/copy.go")
	a.NoError(err)

	a.Equal(newFile, cmpFile)
}

func TestCopyFail(t *testing.T) {
	t.Run("reader fail", func(t *testing.T) {
		a := assert.New(t)

		_, err := NewPbReader("file/does/not/exist")
		a.Error(err)
	})

	t.Run("package fail", func(t *testing.T) {
		a := assert.New(t)

		r, err := NewPbReader("test-files/nopackage")
		a.NoError(err)

		_, err = r.FindPackage()
		a.Error(err)
	})

	t.Run("writer fail", func(t *testing.T) {
		a := assert.New(t)

		_, err := NewPbWriter("/file/does/not/exist", "pkg")
		a.Error(err)
	})

	t.Run("copy fail", func(t *testing.T) {
		a := assert.New(t)

		r, err := NewPbReader("test-files/corrupt")
		a.NoError(err)

		pkg, err := r.FindPackage()
		a.NoError(err)
		a.Equal(pkg, "test")

		w, err := NewPbWriter("/tmp/deleteme", pkg)
		a.NoError(err)
		defer os.Remove("/tmp/deleteme")

		copier := NewPbCopier(r, w)

		err = copier.MkCopy("NotFound")
		a.Error(err)
	})

	t.Run("copy fail", func(t *testing.T) {
		a := assert.New(t)

		r, err := NewPbReader("test-files/corrupt")
		a.NoError(err)

		pkg, err := r.FindPackage()
		a.NoError(err)
		a.Equal(pkg, "test")

		w, err := NewPbWriter("/tmp/deleteme", pkg)
		a.NoError(err)
		defer os.Remove("/tmp/deleteme")

		copier := NewPbCopier(r, w)

		t.Run("Struct Not Found", func(t *testing.T) {
			a := assert.New(t)

			err := copier.MkCopy("NotFound")
			a.Error(err)
			a.Equal(r.FindType("NotFound"), "")
		})

		t.Run("Type Not Found", func(t *testing.T) {
			a := assert.New(t)

			a.Equal(r.FindType("NotFound"), "")
		})

		t.Run("Parse Error", func(t *testing.T) {
			a := assert.New(t)

			err := copier.MkCopy("Bad")
			a.Error(err)
		})

		t.Run("Bad Type", func(t *testing.T) {
			a := assert.New(t)

			err := copier.MkCopy("BadType")
			a.Error(err)
		})

		t.Run("Flush Error", func(t *testing.T) {
			a := assert.New(t)

			err := w.Flush() // this will close the file
			a.NoError(err)

			err = w.Flush() // this will error because the file is closed
			a.Error(err)
		})

		t.Run("pointer panic", func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("Did not panic")
				}
			}()
			f := &PbField{}
			f.copyMethod("foo")
		})
	})
}
