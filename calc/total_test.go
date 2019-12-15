package calc

import "testing"

// TotalSum は加算のための手続き
func TestTotalSum(t *testing.T) {
	var result int

	result = TotalSum(1, 2)
	if result != 3 {
		t.Errorf("totalSum failed. expect:%d, actual:%d", 3, result)
	}

	t.Logf("result is %d", result)
}
