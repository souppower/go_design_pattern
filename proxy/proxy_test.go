package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := PrinterProxy{Name: "A"}
	name := proxy.GetPrinterName()

	if name != "A" {
		t.Errorf("Expected name to equal %s, but got %s.\n", "A", name)
	}

	proxy.SetPrinterName("B")
	name = proxy.GetPrinterName()
	if name != "B" {
		t.Errorf("Expected name to equal %s, but got %s.\n", "B", name)
	}

	result := proxy.Print("C")
	if result != "B:C" {
		t.Errorf("Expected result to equal %s, but got %s.\n", "B:C", result)
	}

}
