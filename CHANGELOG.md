# CHANGELOG

## [Unreleased]

- Fixed `CreateWebResourceResponse` never passing the content stream to WebView2 (the stream variable was shadowed), so every response had an empty body.
- Added support for drag and drop of files and folders. [@ayatkyo](https://github.com/ayatkyo) in [#1](https://github.com/wailsapp/go-webview2/pull/1).