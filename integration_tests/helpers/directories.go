package helpers

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/FleekHQ/space-daemon/grpc/pb"
	. "github.com/onsi/gomega"
)

func CreateEmptyFolder(ctx context.Context, client pb.SpaceApiClient, path string) {
	_, err := client.CreateFolder(ctx, &pb.CreateFolderRequest{
		Path:   path,
		Bucket: "personal",
	})
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
}

func CreateLocalStringFile(strContent string) *os.File {
	content := []byte(strContent)
	tmpfile, err := ioutil.TempFile("", "*-localStringFile.txt")
	ExpectWithOffset(1, err).NotTo(HaveOccurred(), "Failed to create local string file")
	if err != nil {
		defer tmpfile.Close()
	}

	_, err = tmpfile.Write(content)
	ExpectWithOffset(1, err).NotTo(HaveOccurred(), "Failed to write string content")

	return tmpfile
}
