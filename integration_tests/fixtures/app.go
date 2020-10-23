package fixtures

import (
	"os"
	"path"
	"path/filepath"

	"github.com/FleekHQ/space-daemon/core/keychain"

	"github.com/99designs/keyring"

	"github.com/FleekHQ/space-daemon/app"
	"github.com/FleekHQ/space-daemon/config"
	"github.com/FleekHQ/space-daemon/grpc/pb"
	. "github.com/onsi/gomega"
)

type RunAppCtx struct {
	App            *app.App
	cfg            config.Config
	client         pb.SpaceApiClient
	ClientAppToken string
	ClientMnemonic string
}

func RunApp() *RunAppCtx {
	_, cfg, env := GetTestConfig()
	spaceApp := app.New(cfg, env)
	err := spaceApp.Start()

	Expect(err).NotTo(HaveOccurred(), "space app failed to start")
	Expect(spaceApp.IsRunning).To(Equal(true))

	return &RunAppCtx{
		App:    spaceApp,
		cfg:    cfg,
		client: nil,
	}
}

func (a *RunAppCtx) Shutdown() {
	if a.App != nil {
		// shutdown app
		_ = a.App.Shutdown()

		homeDir, _ := os.UserHomeDir()
		spaceStorPath := filepath.Join(homeDir, ".fleek-space")

		ClearMasterAppToken(spaceStorPath)

		// delete app dir
		_ = os.RemoveAll(spaceStorPath)
	}
}

func ClearMasterAppToken(spaceStorPath string) {
	// clear master token from keystore
	ucd, _ := os.UserConfigDir()
	ring, err := keyring.Open(keyring.Config{
		ServiceName:                    "space",
		KeychainTrustApplication:       true,
		KeychainAccessibleWhenUnlocked: true,
		KWalletAppID:                   "space",
		KWalletFolder:                  "space",
		WinCredPrefix:                  "space",
		LibSecretCollectionName:        "space",
		PassPrefix:                     "space",
		PassDir:                        spaceStorPath + "/kcpw",
		FileDir:                        path.Join(ucd, "space", "keyring"),
	})
	if err == nil {
		_ = ring.Remove(keychain.AppTokenStoreKey + "_" + keychain.MasterAppTokenStoreKey)
	}
}

func (a *RunAppCtx) IsInitialized() bool {
	return a.ClientAppToken != "" && a.ClientMnemonic != ""
}
