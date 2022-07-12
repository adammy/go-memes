package server

// Config defines configuration for the Server.
type Config struct {
	// Port defines the port the server runs on.
	Port uint

	// LocalAssets defines if the Server should serve the local assets.
	LocalAssets bool
}
