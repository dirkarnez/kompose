//go:build js && wasm
// +build js,wasm,!cgo

package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"path/filepath"
	"strings"
	"syscall/js"

	"github.com/kubernetes/kompose/pkg/app"
	"github.com/kubernetes/kompose/pkg/kobject"
)

func GetFileName(path string) string {
	noExtension := strings.TrimSuffix(path, filepath.Ext(path))
	return noExtension[strings.LastIndex(noExtension, "\\")+1:]
}

func Convert(composeFile string) (string) {
	return "111"
}

func convertComposeFileToKubernetes_JavaScript(this js.Value, args []js.Value) interface{} {
	// buffer := make([]byte, args[0].Length())

	// js.CopyBytesToGo(buffer, args[0])

	// 计算md5的值
	//bytes, _ := dirkCaf.Convert(buffer)

	// dst := js.Global().Get("Uint8Array").New(len(bytes))
	// js.CopyBytesToJS(dst, bytes)

	return Convert(args[0].String())
}

func main() {
	js.Global().Set("convertComposeFileToKubernetes", js.FuncOf(convertComposeFileToKubernetes_JavaScript))
	select {} // block the main thread forever
}
