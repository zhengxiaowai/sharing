package sharer

import (
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

type DemoSharer struct{}

func (d DemoSharer) InitConfig(confPath string) error {
	return nil
}

func (d DemoSharer) UploadFile(key string, filePath string) (string, error) {
	return "", nil
}

func (d DemoSharer) GetName() string {
	return "demo"
}

func Test_mustGetUserHomeDir(t *testing.T) {
	g := NewGomegaWithT(t)
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		g.Expect(mustGetUserHomeDir()).Should(Equal(os.Getenv("HOME")))
	}
}

func Test_getSharerConf(t *testing.T) {
	g := NewGomegaWithT(t)
	demoSharer := DemoSharer{}
	g.Expect(getSharerConf(demoSharer)).
		Should(Equal(filepath.Join(mustGetUserHomeDir(), defaultConfDirectory, "demo.json")))
}

func Test_makePublicURL(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(makePublicURL("https://example.com", "images/demo.png")).
		Should(Equal("https://example.com/images/demo.png"))
}
