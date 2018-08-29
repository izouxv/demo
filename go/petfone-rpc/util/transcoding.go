package util

import (
	"encoding/base64"
	log "github.com/cihub/seelog"
	"github.com/pkg/errors"
)

type Source struct {
	Company byte //公司
	Scene   byte //场景
	Client  byte //客户端
	Reserve byte //保留位
}

func (s Source) MarShalBinary() ([]byte, error) {
	var out []byte
	out = append(out, s.Company)
	out = append(out, s.Scene)
	out = append(out, s.Client)
	out = append(out, s.Reserve)
	return out, nil
}

func (s *Source) UnMarShalBinary(b []byte) error {
	if len(b) != 4 {
		return errors.New("is not 4 bytes")
	}
	s.Company = b[0]
	s.Scene = b[1]
	s.Client = b[2]
	s.Reserve = b[3]
	return nil
}

func (s *Source) Base64() (string, error) {
	str := ""
	b, err := s.MarShalBinary()
	if err != nil {
		return "", err
	}
	str = base64.StdEncoding.EncodeToString(b)
	return str, nil
}

func (s *Source) UnBase64(in string) error {
	b, err := base64.StdEncoding.DecodeString(in)

	if err != nil {
		log.Error("Source transCoding is failed ", err)
		return err
	}
	if len(b) != 4 {
		return errors.New("")
	}
	s.Company = b[0]
	s.Scene = b[1]
	s.Client = b[2]
	s.Reserve = b[3]
	return nil
}

func CheckSource(source string) bool {
	so := Source{}
	so.UnBase64(source)
	if so.Company == byte(2) {
		return false
	}
	return true
}
