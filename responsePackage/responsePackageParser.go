package responsePackage

import (
	"bufio"
	"bytes"
	"sync"
)

type responseDBEventPackageParser struct{}

var parser *responseDBEventPackageParser
var once sync.Once

func ResponseDBEventPackageParser() *responseDBEventPackageParser {
	once.Do(func() {
		parser = new(responseDBEventPackageParser)
	})
	return parser
}

func (p *responseDBEventPackageParser) Parse(buffer *bytes.Buffer) ([]*ResponseDBEventPackage, int, error) {
	var err error
	packages := []*ResponseDBEventPackage{}
	buf := bytes.NewBuffer(buffer.Bytes())
	scanner := bufio.NewScanner(buf)
	scanner.Split(ScannerSplit)
	consumeBytesCount := 0
	for scanner.Scan() {
		dbEventPack := new(ResponseDBEventPackage)
		err = dbEventPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			continue
		}

		consumeBytesCount += dbEventPack.TotalLength()

		packages = append(packages, dbEventPack)
	}

	return packages, consumeBytesCount, err
}
