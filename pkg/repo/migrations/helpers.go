package migrations

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/config/types"
	"github.com/bacalhau-project/bacalhau/pkg/repo"
)

func getConfigContext(r repo.FsRepo) (config.Context, error) {
	repoPath, err := r.Path()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(repoPath, config.FileName)
	c := config.New()
	// read existing config file if it exists
	if err := c.Load(configFile); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return c, nil
}

func configExists(r repo.FsRepo) (bool, error) {
	repoPath, err := r.Path()
	if err != nil {
		return false, err
	}

	configFile := filepath.Join(repoPath, config.FileName)
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func readConfig(r repo.FsRepo) (config.Context, types.BacalhauConfig, error) {
	c, err := getConfigContext(r)
	if err != nil {
		return nil, types.BacalhauConfig{}, err
	}
	cfg, err := c.Current()
	if err != nil {
		return nil, types.BacalhauConfig{}, err
	}
	return c, cfg, nil
}

// haveSameElements returns true if arr1 and arr2 have the same elements, false otherwise.
func haveSameElements(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	elementCount := make(map[string]int)

	for _, item := range arr1 {
		elementCount[item]++
	}

	for _, item := range arr2 {
		if count, exists := elementCount[item]; !exists || count == 0 {
			return false
		}
		elementCount[item]--
	}

	return true
}

func getLibp2pNodeID(c config.Context) (string, error) {
	privKey, err := config.GetLibp2pPrivKey(c)
	if err != nil {
		return "", err
	}
	peerID, err := peer.IDFromPrivateKey(privKey)
	if err != nil {
		return "", err
	}
	return peerID.String(), nil
}
