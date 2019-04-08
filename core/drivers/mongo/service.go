package dbtools

import (
	"bufio"
	"bytes"
	"fmt"
)

func (m *MongoDB) backup(database, collection string) {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err := m.DumpCollectionTo(database, collection, writer)
	if err != nil {
		fmt.Errorf("Error while dumping collection=%v/%v: %v", database, collection, err)
	}
}

func (m *MongoDB) restore() {

}
