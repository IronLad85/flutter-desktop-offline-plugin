# offline_storage

This Go package implements the host-side of the Flutter [offline_storage](https://github.com/IronLad85/flutter-desktop-offline-plugin.git) plugin.

## Usage

Import as:

```go
import offline_storage "github.com/IronLad85/flutter-desktop-offline-plugin.git/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&offline_storage.OfflineStoragePlugin{}),
```
