import 'dart:async';

import 'package:flutter/services.dart';

class OfflineStorage {
  static const MethodChannel _channel =
      const MethodChannel('offline_storage');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }
}
