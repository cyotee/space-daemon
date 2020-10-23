package integration_tests

import (
	"context"
	"path/filepath"

	"github.com/FleekHQ/space-daemon/grpc/pb"
	"github.com/FleekHQ/space-daemon/integration_tests/fixtures"
	. "github.com/FleekHQ/space-daemon/integration_tests/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App Uploads", func() {
	var (
		app *fixtures.RunAppCtx
	)

	BeforeEach(func() {
		app = fixtures.RunApp()
	})

	AfterEach(func() {
		app.Shutdown()
	})

	Context("when app is initialized", func() {
		BeforeEach(func() {
			InitializeApp(app)
		})

		It("should create empty folder successfully", func() {
			ctx := context.Background()
			CreateEmptyFolder(ctx, app.Client(), "/Empty Folder")

			res, err := app.Client().ListDirectory(ctx, &pb.ListDirectoryRequest{
				Path:   "/",
				Bucket: "personal",
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(res.Entries).To(HaveLen(1))
			Expect(res.Entries[0].Name).To(Equal("Empty Folder"))
		})

		XIt("should upload and download files successfully", func() {
			ctx := context.Background()
			file := CreateLocalStringFile("random file content")

			fileName := filepath.Base(file.Name())

			_, err := app.Client().AddItems(ctx, &pb.AddItemsRequest{
				SourcePaths: []string{file.Name()},
				TargetPath:  "/",
				Bucket:      "personal",
			})
			Expect(err).NotTo(HaveOccurred())
			ExpectFileExists(ctx, app.Client(), "", fileName)

			// try uploading to a folder
			CreateEmptyFolder(ctx, app.Client(), "/Top Folder")
			_, err = app.Client().AddItems(ctx, &pb.AddItemsRequest{
				SourcePaths: []string{file.Name()},
				TargetPath:  "/Top Folder",
				Bucket:      "personal",
			})
			Expect(err).NotTo(HaveOccurred())
			ExpectFileExists(ctx, app.Client(), "Top Folder", fileName)
		})
	})
})
