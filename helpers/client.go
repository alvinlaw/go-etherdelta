package helpers

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Client client
var Client *ethclient.Client

var defaultProviderURI = "https://cloudflare-eth.com"

// SetClientProviderURI ETH provider client URI
func SetClientProviderURI(providerURI string) {
	// Create an IPC based RPC connection to a remote node
	client, err := ethclient.Dial(providerURI)
	Client = client
	if err != nil {
		panic(err)
	}

	//log.Infof("Connected EtherDelta client to provider: %s", providerURI)
}

func init() {
	providerURIEnv := os.Getenv("ETH_PROVIDER_URI")

	if providerURIEnv != "" {
		SetClientProviderURI(providerURIEnv)
	} else {
		SetClientProviderURI(defaultProviderURI)
	}
}
