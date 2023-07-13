package jsonutil2_test

import (
	"fmt"
	"jsonutil/json/jsonutil2"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	b := []byte("{\"id\":2022,\"info\":{\"name\":\"xiaoming\",\"age\":18},\"habit\":{\"exercise\":{\"football\":\"true\",\"basketball\":\"false\"}},\"class\":[1,\"2\",\"true\",[1,\"2\",\"true\",{\"k1\":\"v1\",\"k2\":\"v2\"}]],\"ptr\":<nil>,\"nil\":[]}")
	res := S{}
	jsonutil2.Unmarshal(b, &res)
	fmt.Println()
}
