package offline_storage

import (
	"os"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/oleksandr/bonjour"
)

const channelName = "offline_storage"

// OfflineStoragePlugin implements flutter.Plugin and handles method.
type OfflineStoragePlugin struct{}

var _ flutter.Plugin = &OfflineStoragePlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *OfflineStoragePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	channel.HandleFunc("getHomeDir", p.handleHomeDirectory)
	channel.HandleFunc("getMakeDir", p.handleMakeDirectory)
	channel.HandleFunc("advertiseServer", p.handleMdnsAdvertisement)
	return nil
}

func (p *OfflineStoragePlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

func (p *OfflineStoragePlugin) handleMakeDirectory(arguments interface{}) (reply interface{}, err error) {
	_path := arguments.(string)
	_err := os.Mkdir(_path, 0777)
	return nil, _err
}

func (p *OfflineStoragePlugin) handleMdnsAdvertisement(arguments interface{}) (reply interface{}, err error) {
	_port := arguments.(int32)
	_, _err := bonjour.Register("Frinks MDNS Service", "_frinks._tcp", ".local", int(_port), []string{"txtv=1", "app=test"}, nil)
	if _err != nil {
		return nil, err
	}
	return nil, err
}

func (p *OfflineStoragePlugin) handleHomeDirectory(arguments interface{}) (reply interface{}, err error) {
	path, err := os.UserHomeDir()
	return path, err
}
