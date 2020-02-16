package offline_storage

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "offline_storage"

// OfflineStoragePlugin implements flutter.Plugin and handles method.
type OfflineStoragePlugin struct{}

var _ flutter.Plugin = &OfflineStoragePlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *OfflineStoragePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	return nil
}

func (p *OfflineStoragePlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}
