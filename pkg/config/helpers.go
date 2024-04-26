package config

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	libp2p_crypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/bacalhau-project/bacalhau/pkg/config/types"
	baccrypto "github.com/bacalhau-project/bacalhau/pkg/lib/crypto"
	"github.com/bacalhau-project/bacalhau/pkg/lib/network"
	"github.com/bacalhau-project/bacalhau/pkg/logger"
)

func Get[T any](c Context, key string) (T, error) {
	raw := c.User().Get(key)
	if raw == nil {
		return zeroValue[T](), fmt.Errorf("value not found for %s", key)
	}

	var val T
	val, ok := raw.(T)
	if !ok {
		err := c.ForKey(key, &val)
		if err != nil {
			return zeroValue[T](), fmt.Errorf("value not of expected type, got: %T: %w", raw, err)
		}
	}

	return val, nil
}

func zeroValue[T any]() T {
	var zero T
	return zero
}

// KeyAsEnvVar returns the environment variable corresponding to a config key
func KeyAsEnvVar(key string) string {
	return strings.ToUpper(
		fmt.Sprintf("%s_%s", environmentVariablePrefix, environmentVariableReplace.Replace(key)),
	)
}

// WritePersistedConfigs will write certain values from the resolved config to the persisted config.
// These include fields for configurations that must not change between version updates, such as the
// execution store and job store paths, in case we change their default values in future updates.
func WritePersistedConfigs(configFile string, resolvedCfg types.BacalhauConfig) error {
	// a viper config instance that is only based on the config file.
	viperWriter := viper.New()
	viperWriter.SetTypeByDefaultValue(true)
	viperWriter.SetConfigFile(configFile)

	// read existing config if it exists.
	if err := viperWriter.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	var fileCfg types.BacalhauConfig
	if err := viperWriter.Unmarshal(&fileCfg, DecoderHook); err != nil {
		return err
	}

	// check if any of the values that we want to write are not set in the config file.
	var doWrite bool
	var logMessage strings.Builder
	set := func(key string, value interface{}) {
		viperWriter.Set(key, value)
		logMessage.WriteString(fmt.Sprintf("\n%s:\t%v", key, value))
		doWrite = true
	}
	emptyStoreConfig := types.JobStoreConfig{}
	if fileCfg.Node.Compute.ExecutionStore == emptyStoreConfig {
		set(types.NodeComputeExecutionStore, resolvedCfg.Node.Compute.ExecutionStore)
	}
	if fileCfg.Node.Requester.JobStore == emptyStoreConfig {
		set(types.NodeRequesterJobStore, resolvedCfg.Node.Requester.JobStore)
	}
	if fileCfg.Node.Name == "" && resolvedCfg.Node.Name != "" {
		set(types.NodeName, resolvedCfg.Node.Name)
	}
	if doWrite {
		log.Info().Msgf("Writing to config file %s:%s", configFile, logMessage.String())
		return viperWriter.WriteConfig()
	}
	return nil
}

func GetClientID(c Context) (string, error) {
	return loadClientID(c)
}

func loadInstallationUserIDKey(c Context) (string, error) {
	key := c.User().GetString(types.UserInstallationID)
	if key == "" {
		return "", fmt.Errorf("config error: user-installation-id-key not set")
	}
	return key, nil
}

func GetInstallationUserID(c Context) (string, error) {
	return loadInstallationUserIDKey(c)
}

// loadClientID loads a hash identifying a user based on their ID key.
func loadClientID(c Context) (string, error) {
	key, err := loadUserIDKey(c)
	if err != nil {
		return "", fmt.Errorf("failed to load user ID key: %w", err)
	}

	return convertToClientID(&key.PublicKey), nil
}

const (
	sigHash = crypto.SHA256 // hash function to use for sign/verify
)

// convertToClientID converts a public key to a client ID:
func convertToClientID(key *rsa.PublicKey) string {
	hash := sigHash.New()
	hash.Write(key.N.Bytes())
	hashBytes := hash.Sum(nil)

	return fmt.Sprintf("%x", hashBytes)
}

func ClientAPIPort(c Context) uint16 {
	return uint16(c.User().GetInt(types.NodeClientAPIPort))
}

func ClientAPIHost(c Context) string {
	return c.User().GetString(types.NodeClientAPIHost)
}

func ClientTLSConfig(c Context) types.ClientTLSConfig {
	cfg := types.ClientTLSConfig{
		UseTLS:   c.User().GetBool(types.NodeClientAPIClientTLSUseTLS),
		Insecure: c.User().GetBool(types.NodeClientAPIClientTLSInsecure),
		CACert:   c.User().GetString(types.NodeClientAPIClientTLSCACert),
	}

	if !cfg.UseTLS {
		// If we haven't explicitly turned on TLS, but implied it through
		// the other options, then set it to true
		if cfg.Insecure || cfg.CACert != "" {
			cfg.UseTLS = true
		}
	}

	return cfg
}

func ServerAPIPort(c Context) uint16 {
	return uint16(c.User().GetInt(types.NodeServerAPIPort))
}

func configError(e error) {
	msg := fmt.Sprintf("config error: %s", e)
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func ServerAPIHost(c Context) string {
	host := c.User().GetString(types.NodeServerAPIHost)

	if net.ParseIP(host) == nil {
		// We should check that the value gives us an address type
		// we can use to get our IP address. If it doesn't, we should
		// panic.
		atype, ok := network.AddressTypeFromString(host)
		if !ok {
			configError(fmt.Errorf("invalid address type in Server API Host config: %s", host))
		}

		addr, err := network.GetNetworkAddress(atype, network.AllAddresses)
		if err != nil {
			configError(errors.Wrap(err, fmt.Sprintf("failed to get network address for Server API Host: %s", host)))
		}

		if len(addr) == 0 {
			configError(fmt.Errorf("no %s addresses found for Server API Host", host))
		}

		// Use the first address
		host = addr[0]
	}

	return host
}

func ServerAutoCertDomain(c Context) string {
	return c.User().GetString(types.NodeServerAPITLSAutoCert)
}

func GetRequesterCertificateSettings(c Context) (string, string) {
	cert := c.User().GetString(types.NodeServerAPITLSServerCertificate)
	key := c.User().GetString(types.NodeServerAPITLSServerKey)
	return cert, key
}
func GetRequesterSelfSign(c Context) bool {
	return c.User().GetBool(types.NodeServerAPITLSSelfSigned)
}

func DevstackGetShouldPrintInfo() bool {
	return os.Getenv("DEVSTACK_PRINT_INFO") != ""
}

func DevstackSetShouldPrintInfo() {
	os.Setenv("DEVSTACK_PRINT_INFO", "1")
}

func DevstackEnvFile() string {
	return os.Getenv("DEVSTACK_ENV_FILE")
}

func ShouldKeepStack() bool {
	return os.Getenv("KEEP_STACK") != ""
}

const (
	DockerUsernameEnvVar = "DOCKER_USERNAME"
	DockerPasswordEnvVar = "DOCKER_PASSWORD"
)

type DockerCredentials struct {
	Username string
	Password string
}

func (d *DockerCredentials) IsValid() bool {
	return d.Username != "" && d.Password != ""
}

func GetDockerCredentials() DockerCredentials {
	return DockerCredentials{
		Username: os.Getenv(DockerUsernameEnvVar),
		Password: os.Getenv(DockerPasswordEnvVar),
	}
}

func GetLibp2pConfig(c Context) (types.Libp2pConfig, error) {
	var libp2pCfg types.Libp2pConfig
	if err := c.ForKey(types.NodeLibp2p, &libp2pCfg); err != nil {
		return types.Libp2pConfig{}, err
	}
	return libp2pCfg, nil
}

func GetBootstrapPeers(c Context) ([]multiaddr.Multiaddr, error) {
	bootstrappers := c.User().GetStringSlice(types.NodeBootstrapAddresses)
	peers := make([]multiaddr.Multiaddr, 0, len(bootstrappers))
	for _, peer := range bootstrappers {
		parsed, err := multiaddr.NewMultiaddr(peer)
		if err != nil {
			return nil, err
		}
		peers = append(peers, parsed)
	}
	return peers, nil
}

func GetLogMode(c Context) logger.LogMode {
	mode := c.User().Get(types.NodeLoggingMode)
	switch v := mode.(type) {
	case logger.LogMode:
		return v
	case string:
		out, err := logger.ParseLogMode(v)
		if err != nil {
			log.Warn().Err(err).Msgf("invalid logging mode specified: %s", v)
		}
		return out
	default:
		log.Error().Msgf("unknown logging mode: %v", mode)
		return logger.LogModeDefault
	}
}

func GetAutoCertCachePath(c Context) string {
	return c.User().GetString(types.NodeServerAPITLSAutoCertCachePath)
}

func GetLibp2pTracerPath(c Context) string {
	return c.User().GetString(types.MetricsLibp2pTracerPath)
}

func GetEventTracerPath(c Context) string {
	return c.User().GetString(types.MetricsEventTracerPath)
}

func GetExecutorPluginsPath(c Context) string {
	return c.User().GetString(types.NodeExecutorPluginPath)
}

// TODO idk where this goes yet these are mostly random

func GetDownloadURLRequestRetries(c Context) int {
	return c.User().GetInt(types.NodeDownloadURLRequestRetries)
}

func GetDownloadURLRequestTimeout(c Context) time.Duration {
	return c.User().GetDuration(types.NodeDownloadURLRequestTimeout)
}

func SetVolumeSizeRequestTimeout(c Context, value time.Duration) {
	c.User().Set(types.NodeVolumeSizeRequestTimeout, value)
}

func GetVolumeSizeRequestTimeout(c Context) time.Duration {
	return c.User().GetDuration(types.NodeVolumeSizeRequestTimeout)
}

func GetUpdateCheckFrequency(c Context) time.Duration {
	return c.User().GetDuration(types.UpdateCheckFrequency)
}

func GetStoragePath(c Context) string {
	path := c.User().GetString(types.NodeComputeStoragePath)
	if path == "" {
		return os.TempDir()
	}
	return path
}

func GetDockerManifestCacheSettings(c Context) (*types.DockerCacheConfig, error) {
	if cfg, err := Get[types.DockerCacheConfig](c, types.NodeComputeManifestCache); err != nil {
		return nil, err
	} else {
		return &cfg, nil
	}
}

// PreferredAddress will allow for the specifying of
// the preferred address to listen on for cases where it
// is not clear, or where the address does not appear when
// using 0.0.0.0
func PreferredAddress() string {
	return os.Getenv("BACALHAU_PREFERRED_ADDRESS")
}

func GetStringMapString(c Context, key string) map[string]string {
	return c.User().GetStringMapString(key)
}

func GetClientPublicKey(c Context) (*rsa.PublicKey, error) {
	privKey, err := loadUserIDKey(c)
	if err != nil {
		return nil, err
	}
	return &privKey.PublicKey, nil
}

func GetClientPrivateKey(c Context) (*rsa.PrivateKey, error) {
	privKey, err := loadUserIDKey(c)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

// loadUserIDKey loads the user ID key from whatever source is configured.
func loadUserIDKey(c Context) (*rsa.PrivateKey, error) {
	keyFile := c.User().GetString(types.UserKeyPath)
	if keyFile == "" {
		return nil, fmt.Errorf("config error: user-id-key not set")
	}

	return baccrypto.LoadPKCS1KeyFile(keyFile)
}

func GetLibp2pPrivKey(c Context) (libp2p_crypto.PrivKey, error) {
	return loadLibp2pPrivKey(c)
}

func loadLibp2pPrivKey(c Context) (libp2p_crypto.PrivKey, error) {
	keyFile := c.User().GetString(types.UserLibp2pKeyPath)
	if keyFile == "" {
		return nil, fmt.Errorf("config error: libp2p private key not set")
	}

	keyBytes, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}
	// base64 decode keyBytes
	b64, err := base64.StdEncoding.DecodeString(string(keyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}
	// parse the private key
	key, err := libp2p_crypto.UnmarshalPrivateKey(b64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	return key, nil
}

// GetClientPublicKeyString returns a base64-encoding of the user's public ID key:
// NOTE: must be called after InitConfig() or system will panic.
func GetClientPublicKeyString(c Context) (string, error) {
	userIDKey, err := loadUserIDKey(c)
	if err != nil {
		return "", err
	}

	return encodePublicKey(&userIDKey.PublicKey), nil
}

// encodePublicKey encodes a public key as a string:
func encodePublicKey(key *rsa.PublicKey) string {
	return base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(key))
}
