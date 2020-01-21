package sharer_test

import (
	"os"
	"path/filepath"
	"sharing/sharer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("S3", func() {
	var s3Sharer sharer.S3Sharer

	BeforeEach(func() {
		s3Sharer = sharer.S3Sharer{
			AccessKey:  os.Getenv("S3_ACCESS_KEY"),
			SecretKey:  os.Getenv("S3_SECRET_KEY"),
			EndPoint:   "sgp1.digitaloceanspaces.com",
			BucketName: "sharing-sg-demo",
			UseSSL:     true,
			UseDomain:  "https://sharing-sg-demo.sgp1.digitaloceanspaces.com",
		}
	})

	It("get s3 sharer alias name", func() {
		Ω(s3Sharer.GetName()).Should(Equal("s3"))
	})

	It("initialize s3 config file", func() {
		s3Sharer = sharer.S3Sharer{}
		err := s3Sharer.InitConfig("./testdata/s3.json")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(s3Sharer).To(MatchAllFields(Fields{
			"AccessKey": Equal("test_access_key"),
			"SecretKey": Equal("test_secret_key"),
			"EndPoint": Equal("sgp1.digitaloceanspaces.com"),
			"BucketName": Equal("sharing-sg-demo"),
			"UseSSL": BeTrue(),
			"UseDomain": Equal("https://sharing-sg-demo.sgp1.digitaloceanspaces.com"),
		}))
	})

	It("upload small file to s3 object storage", func() {
		key := filepath.Join("sharing", filepath.Base("./testdata/demo.jpg"))
		downloadURL, err := s3Sharer.UploadFile(key, "./testdata/demo.jpg")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(downloadURL).Should(Equal("https://sharing-sg-demo.sgp1.digitaloceanspaces.com/sharing/demo.jpg"))
	})

	It("upload large file to s3 object storage", func() {
		// TODO: support part upload
	})
})
