package rpc_test

import "testing"

type EncodingExample struct {
    Testing bool
}

fun TestEncode(t *testing.T){
    expected := "Content_Length: 16\r\n\r\n{\"testing\":true}"
    actual :=   rpc.EncodeMessage(EncodingExample{Testing:true})

    if expected != actual {
        t.Fatalf("Expected: %s, Actual %s", expected, actual)
    }
}
