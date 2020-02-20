import 'dart:async';

import 'package:flutter/services.dart';

class OfflineStorage {
  static const MethodChannel _channel = const MethodChannel('offline_storage');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }

  static Future<String> get homeDir async {
    final String _homeDir = await _channel.invokeMethod('getHomeDir');
    return _homeDir;
  }

  static Future<String> advertiseServer(int port) async {
    return await _channel.invokeMethod('advertiseServer', port);
  }

  static Future<String> makeDir(path) async {
    final String isDirectoryCreated =
        await _channel.invokeMethod('getMakeDir', path);
    return isDirectoryCreated;
  }
}
