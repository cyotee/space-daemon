package fixtures

import (
	"fmt"
	"os"

	"github.com/FleekHQ/space-daemon/config"
	"github.com/FleekHQ/space-daemon/core/env"
	. "github.com/onsi/gomega"
	"github.com/phayes/freeport"
)

// GetTestConfig returns a ConfigMap instance instantiated using the env variables
func GetTestConfig() (*config.Flags, config.Config, env.SpaceEnv) {
	ipfsPort, err := freeport.GetFreePort()
	Expect(err).NotTo(HaveOccurred())

	flags := &config.Flags{
		Ipfsaddr:             fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", ipfsPort),
		Ipfsnode:             true,
		Ipfsnodeaddr:         fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", ipfsPort),
		DevMode:              false,
		ServicesAPIURL:       os.ExpandEnv("$SERVICES_API_URL"),
		VaultAPIURL:          os.ExpandEnv("$VAULT_API_URL"),
		VaultSaltSecret:      os.ExpandEnv("$VAULT_SALT_SECRET"),
		ServicesHubAuthURL:   os.ExpandEnv("$SERVICES_HUB_AUTH_URL"),
		TextileHubTarget:     os.ExpandEnv("$TXL_HUB_TARGET"),
		TextileHubMa:         os.ExpandEnv("$TXL_HUB_MA"),
		TextileThreadsTarget: os.ExpandEnv("$TXL_THREADS_TARGET"),
		TextileHubGatewayUrl: os.ExpandEnv("$TXL_HUB_GATEWAY_URL"),
		TextileUserKey:       os.ExpandEnv("$TXL_USER_KEY"),
		TextileUserSecret:    os.ExpandEnv("$TXL_USER_SECRET"),
		SpaceStorageSiteUrl:  os.ExpandEnv("SPACE_STORAGE_SITE_URL"),
	}

	// env
	spaceEnv := env.New()

	// load configs
	return flags, config.NewMap(flags), spaceEnv
}
