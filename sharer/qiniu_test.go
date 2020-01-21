package sharer_test

import (
	"os"
	"path/filepath"
	"sharing/sharer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("QiNiu", func() {
	var qiNiuSharer sharer.QiNiuSharer

	BeforeEach(func() {
		qiNiuSharer = sharer.QiNiuSharer{
			AccessKey:     os.Getenv("QINIU_ACCESS_KEY"),
			SecretKey:     os.Getenv("QINIU_SECERET_KEY"),
			Bucket:        "static",
			Zone:          "zonehuadong",
			UseHTTPS:      false,
			UseCdnDomains: false,
			Domain:        "https://static.zhengxiaowai.cc",
		}
	})

	It("get qiniu sharer alias name", func() {
		Ω(qiNiuSharer.GetName()).Should(Equal("qiniu"))
	})

	It("initialize qiniu config file", func() {
		qiNiuSharer = sharer.QiNiuSharer{}
		err := qiNiuSharer.InitConfig("./testdata/qiniu.json")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(qiNiuSharer).To(MatchAllFields(Fields{
			"AccessKey": Equal("test_access_key"),
			"SecretKey": Equal("test_secret_key"),
			"Bucket": Equal("static"),
			"Zone": Equal("zonehuadong"),
			"UseHTTPS": BeTrue(),
			"UseCdnDomains": BeTrue(),
			"Domain": Equal("https://static.zhengxiaowai.cc"),
		}))
	})

	It("upload small file to qiniu object storage", func() {
		key := filepath.Join("sharing", filepath.Base("./testdata/demo.jpg"))
		downloadURL, err := qiNiuSharer.UploadFile(key, "./testdata/demo.jpg")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(downloadURL).Should(Equal("https://static.zhengxiaowai.cc/sharing/demo.jpg"))
	})

	It("upload large file to qiniu object storage", func() {
		// TODO: support part upload
	})
})
