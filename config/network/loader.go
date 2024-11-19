package network

import "sync"

var (
	instance *NetworkConfig
	once     sync.Once
	mu       sync.RWMutex // Mutex to access thread-safe
)

// Automatically initialize the singleton instance
func init() {
	once.Do(func() {
		instance = &NetworkConfig{} // Initialize the instance
	})
}

// Set value for thread-safe
func Set(config NetworkConfig) {
	mu.Lock()
	defer mu.Unlock()
	*instance = config
}

// Get value
func Get() NetworkConfig {
	mu.RLock()
	defer mu.RUnlock()
	return *instance
}
