package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bacalhau-project/bacalhau/pkg/config"
	"github.com/bacalhau-project/bacalhau/pkg/config/types"
	"github.com/bacalhau-project/bacalhau/pkg/repo"
	"github.com/bacalhau-project/bacalhau/pkg/storage/util"
)

// V2Migration updates the repo so that nodeID is no longer part of the execution and job store paths.
// It does the following:
// - Generates and persists the nodeID in the config if it is missing, which is the case for v2 repos
// - Adds the execution and job store paths to the config if they are missing, which is the case for v3 repos
// - Renames the execution and job store directories to the new name if they exist
var V2Migration = repo.NewMigration(
	repo.RepoVersion2,
	repo.RepoVersion3,
	func(r repo.FsRepo) error {
		currentCtx, currentCfg, err := readConfig(r)
		if err != nil {
			return err
		}
		repoPath, err := r.Path()
		if err != nil {
			return err
		}
		// we load the config to resolve the libp2p node id. Loading the config this way will also
		// use default values, args and env vars to fill in the config, so we can be sure we are
		// reading the correct libp2p key in case the user is overriding the default value.
		newCtx := config.New()
		if _, err := os.Stat(filepath.Join(repoPath, config.FileName)); err == nil {
			if err := newCtx.Load(filepath.Join(repoPath, config.FileName)); err != nil {
				return err
			}
		}
		r.EnsureRepoPathsConfigured(newCtx)
		resolvedCfg, err := newCtx.Current()
		if err != nil {
			return err
		}
		libp2pNodeID, err := getLibp2pNodeID(newCtx)
		if err != nil {
			return err
		}

		doWrite := false
		var logMessage strings.Builder
		set := func(key string, value interface{}) {
			currentCtx.Set(key, value)
			logMessage.WriteString(fmt.Sprintf("\n%s:\t%v", key, value))
			doWrite = true
		}

		if currentCfg.Node.Compute.ExecutionStore.Path == "" {
			// persist the execution store in the repo
			executionStore := resolvedCfg.Node.Compute.ExecutionStore

			// handle an edge case where config.yaml has store config entry but no path,
			// which will override the resolved config path and make it empty as well
			if executionStore.Path == "" {
				executionStore.Path = filepath.Join(repoPath, "compute_store", "executions.db")
			}

			// if execution store already exist with nodeID, then rename it to the new name
			legacyStoreName := filepath.Join(repoPath, libp2pNodeID+"-compute")
			newStorePath := filepath.Dir(executionStore.Path)
			if _, err := os.Stat(legacyStoreName); err == nil {
				// expecting compute_store for newStorePath
				if err := os.Rename(legacyStoreName, newStorePath); err != nil {
					return err
				}
			} else if err = os.MkdirAll(newStorePath, util.OS_USER_RWX); err != nil {
				return err
			}
			set(types.NodeComputeExecutionStore, executionStore)
		}

		if currentCfg.Node.Requester.JobStore.Path == "" {
			// persist the job store in the repo
			jobStore := resolvedCfg.Node.Requester.JobStore

			// handle an edge case where config.yaml has store config entry but no path,
			// which will override the resolved config path and make it empty as well
			if jobStore.Path == "" {
				jobStore.Path = filepath.Join(repoPath, "orchestrator_store", "jobs.db")
			}

			// if job store already exist with nodeID, then rename it to the new name
			legacyStoreName := filepath.Join(repoPath, libp2pNodeID+"-requester")
			newStorePath := filepath.Dir(jobStore.Path)
			if _, err := os.Stat(legacyStoreName); err == nil {
				if err := os.Rename(legacyStoreName, newStorePath); err != nil {
					return err
				}
			} else if err = os.MkdirAll(newStorePath, util.OS_USER_RWX); err != nil {
				return err
			}
			set(types.NodeRequesterJobStore, jobStore)
		}

		if currentCfg.Node.Name == "" {
			set(types.NodeName, libp2pNodeID)
		}

		if doWrite {
			currentCtx.User().SetConfigFile(filepath.Join(repoPath, config.FileName))
			return currentCtx.User().WriteConfig()
		}
		return nil
	})
