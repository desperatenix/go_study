//nolint:gofumpt
package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	info, err := os.Stat(fromPath)
	if err != nil {
		return err
	}
	if info.IsDir() || !info.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	if offset >= info.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit > info.Size() || limit == 0 {
		limit = info.Size()
	}

	src, err := os.OpenFile(fromPath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}

	dst, err := os.OpenFile(toPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	bytesLeft := info.Size() - offset
	barSize := limit
	if bytesLeft < limit || limit == 0 {
		barSize = bytesLeft
	}

	bar := pb.Full.Start64(barSize)
	defer bar.Finish()
	barReader := bar.NewProxyReader(src)

	_, err = src.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = io.CopyN(dst, barReader, limit)

	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	if err := src.Close(); err != nil {
		return err
	}

	if err := dst.Close(); err != nil {
		return err
	}

	return nil
}
