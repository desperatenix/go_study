package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	TmpDirName string = "tmp"
)

var ErrNoSuchFileOrDirectory = "stat ./there_is_no_file: no such file or directory"

type CopySuite struct {
	suite.Suite
	tmpDirPath  string
	dstFilePath string
}

func (s *CopySuite) SetupTest() {
	s.tmpDirPath = path.Join(os.TempDir(), TmpDirName)
	if err := os.Mkdir(s.tmpDirPath, 0o755); err != nil {
		s.T().Fail()
	}
	s.dstFilePath = path.Join(s.tmpDirPath, "xxx")
}

func (s *CopySuite) TearDownTest() {
	if err := os.RemoveAll(s.tmpDirPath); err != nil {
		s.T().Fail()
	}
}

func (s *CopySuite) TestCopy() {
	err := Copy("./testdata/input.txt", s.dstFilePath, 0, 0)
	s.Require().NoError(err)

	info, err := os.Stat("./testdata/input.txt")
	if err != nil {
		s.T().Fail()
	}
	resultInfo, err := os.Stat(s.dstFilePath)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(info.Size(), resultInfo.Size())
}

func (s *CopySuite) TestCopyWithLimit() {
	err := Copy("./testdata/input.txt", s.dstFilePath, 0, 10)
	s.Require().NoError(err)

	resultInfo, err := os.Stat(s.dstFilePath)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(int64(10), resultInfo.Size())
}

func (s *CopySuite) TestCopyWithLimitOffset() {
	err := Copy("./testdata/input.txt", s.dstFilePath, 6000, 1000)
	s.Require().NoError(err)

	resultInfo, err := os.Stat(s.dstFilePath)
	if err != nil {
		s.T().Fail()
	}
	s.Require().Equal(int64(617), resultInfo.Size())
}

func (s *CopySuite) TestNoFile() {
	err := Copy("./there_is_no_file", s.dstFilePath, 0, 0)
	s.Require().Error(err)
	s.Require().EqualError(err, ErrNoSuchFileOrDirectory)
}

func (s *CopySuite) TestSpecialFile() {
	err := Copy("/dev/null", s.dstFilePath, 0, 0)
	s.Require().Error(err)
	s.Require().Equal(err, ErrUnsupportedFile)
}

func (s *CopySuite) TestOffsetExceeds() {
	err := Copy("./testdata/input.txt", s.dstFilePath, 10000, 0)
	s.Require().Error(err)
	s.Require().Equal(err, ErrOffsetExceedsFileSize)
}

func TestCopy(t *testing.T) {
	suite.Run(t, new(CopySuite))
}
