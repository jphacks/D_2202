import 'dart:developer';
import 'package:app/handle_api/gcs_api.dart';
import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import '../model.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

// TODO
Future<void> upload_post(UploadInfo info) async {
  String url = "https://${dotenv.get('API_IP')}/api/post/upload";
  Map<String, String> headers = {'content-type': 'application/json'};
  String body = json.encode({
    'name': info.name,
    'latitude': info.latitude,
    'longitude': info.longitude,
    'hash': info.hash,
    'tags': info.tags,
  });
  http.Response resp =
      await http.post(Uri.parse(url), headers: headers, body: body);
  if (resp.statusCode != 200) {
    return;
  } else {
    throw Exception('投稿の送信に失敗しました');
  }
}

// 画像のアップロードをする関数
Future<String> uploadImage(image, hash) async {
  await upload_to_gcs(image, hash);
  return hash;
}

// VisionAIにurlを投げて名前を返す
// TODO
Future<dynamic> sendVisionAI(String hash) async {
  Future.delayed(Duration(seconds: 3));
  var data;
  http.Response response;
  await http
      .get(Uri.parse(
          'http://${dotenv.get('API_IP')}/api/plant_identify?hash=${hash}'))
      .then((value) => {data = jsonDecode(value.body)});
  log(data.toString());
  return data;
}
