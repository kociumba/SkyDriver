package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// Config holds all configuration values
type Config struct {
	Prediction PredictionConfig `toml:"prediction"`
}

// PredictionConfig holds prediction-specific configuration
type PredictionConfig struct {
	Weights WeightConfig `toml:"weights"`
}

// WeightConfig holds the weights for different prediction factors
type WeightConfig struct {
	PriceSpread        float64 `toml:"price_spread"`
	VolumeImbalance    float64 `toml:"volume_imbalance"`
	OrderImbalance     float64 `toml:"order_imbalance"`
	MovingWeekTrend    float64 `toml:"moving_week_trend"`
	OrderBookPressure  float64 `toml:"order_book_pressure"`
	VolumeFactor       float64 `toml:"volume_factor"`
	ProfitMarginFactor float64 `toml:"profit_margin_factor"`
}

var defaultConfig = Config{
	Prediction: PredictionConfig{
		Weights: WeightConfig{
			PriceSpread:        0.1428571429,
			VolumeImbalance:    0.1428571429,
			OrderImbalance:     0.1428571429,
			MovingWeekTrend:    0.1428571429,
			OrderBookPressure:  0.1428571429,
			VolumeFactor:       0.1428571429,
			ProfitMarginFactor: 0.1428571429,
		},
	},
}

// LoadConfig loads the configuration from the default location
// If no config file exists, it creates one with default values
func LoadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	skyDriverDir := filepath.Join(homeDir, ".skydriver")
	if err := os.MkdirAll(skyDriverDir, 0755); err != nil {
		return nil, err
	}

	configPath := filepath.Join(skyDriverDir, "config.toml")

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config file with comments
		defaultConfigContent := `# SkyDriver Configuration File
# Located in ~/.skydriver/config.toml
# Schema: https://raw.githubusercontent.com/kociumba/SkyDriver/main/config/schema.json

# Prediction weights configuration
# All weights should sum to approximately 1.0
[prediction.weights]
# Weight for price spread between buy and sell prices
price_spread = 0.1428571429

# Weight for volume imbalance between buy and sell volumes
volume_imbalance = 0.1428571429

# Weight for order imbalance between buy and sell orders
order_imbalance = 0.1428571429

# Weight for moving week trend
moving_week_trend = 0.1428571429

# Weight for order book pressure
order_book_pressure = 0.1428571429

# Weight for volume factor
volume_factor = 0.1428571429

# Weight for profit margin factor
profit_margin_factor = 0.1428571429
`
		if err := os.WriteFile(configPath, []byte(defaultConfigContent), 0644); err != nil {
			return nil, err
		}
		return &defaultConfig, nil
	}

	// Load existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to the specified path
func SaveConfig(config *Config, path string) error {
	data, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
