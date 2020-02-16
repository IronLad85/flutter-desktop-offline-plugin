#import "OfflineStoragePlugin.h"
#if __has_include(<offline_storage/offline_storage-Swift.h>)
#import <offline_storage/offline_storage-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "offline_storage-Swift.h"
#endif

@implementation OfflineStoragePlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftOfflineStoragePlugin registerWithRegistrar:registrar];
}
@end
