package radiko

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestAreaID(t *testing.T) {
	areaID, err := AreaID()
	if err != nil {
		t.Errorf("Failed to get area id: %s", err)
	}
	fmt.Println(areaID)
}

func TestProcessSpanNode(t *testing.T) {
	expectedAreaID := "JP13"
	s := `document.write('<span class="` + expectedAreaID + `">TOKYO JAPAN</span>');`

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Errorf("Parse HTML: %s", err)
	}

	areaID := processSpanNode(doc)
	if areaID != expectedAreaID {
		t.Errorf("Failed to process span node.\nAreaID: %s",
			areaID)
	}
}