package helpers

import (
	"context"
	"errors"

	"github.com/FleekHQ/space-daemon/grpc/pb"
	. "github.com/onsi/gomega"
)

func ExpectFileExists(ctx context.Context, client pb.SpaceApiClient, remotePath string, remoteFileName string) {
	res, err := client.ListDirectory(ctx, &pb.ListDirectoryRequest{
		Path:   remotePath,
		Bucket: "personal",
	})
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	ExpectWithOffset(1, res.Entries).NotTo(BeEmpty(), "file at remote path not found")

	for _, item := range res.Entries {
		if item.Name == remoteFileName {
			return
		}
	}
	// Not Found
	ExpectWithOffset(1, errors.New("file at remote path not found")).NotTo(HaveOccurred())
}
