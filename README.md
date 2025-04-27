# localizegen 🌍

> Generate localization files for mobile applications (Android, iOS, Flutter) from Google Sheets.

[![Homebrew](https://img.shields.io/badge/Homebrew-available-green.svg)](https://brew.sh)
[![Platform](https://img.shields.io/badge/platform-Android%20%7C%20iOS%20%7C%20Flutter-blue.svg)](https://github.com/prongbang/localizegen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/prongbang/localizegen.svg)](https://golang.org)

```
   __              ___                     
  / /__  _______ _/ (_)__ ___ ___ ____ ___ 
 / / _ \/ __/ _ `/ / /_ // -_) _ `/ -_) _ \
/_/\___/\__/\_,_/_/_//__/\__/\_, /\__/_//_/
                            /___/          
--> START
# Platform: ios
# Language: en
# Generate: localization
> Create ./App/Locallization/Translate.swift -> Success
> Create ./App/Locallization/en.lproj/Localizable.strings -> Success
<-- DONE
```

## ✨ Features

- 🚀 **Multi-Platform Support** - Generate localization files for Android, iOS, and Flutter
- 📊 **Google Sheets Integration** - Manage translations in a familiar spreadsheet interface
- 🌐 **Multiple Languages** - Support for unlimited languages and locales
- 🔄 **Automatic Generation** - Generate all language files with a single command
- 🛠️ **Customizable Output** - Configure target paths and filenames
- ⚡ **Fast and Efficient** - Quick generation with minimal configuration

## 📦 Installation

### Via Homebrew (Recommended)

```shell
brew update
brew tap prongbang/homebrew-formulae
brew install localizegen
```

### Via Go

```shell
go install github.com/prongbang/localizegen
```

### Binary Downloads

Download the latest binary for your platform:

- 🍎 [macOS](https://github.com/prongbang/localizegen/blob/master/localizegen?raw=true)
- 🐧 [Linux](https://github.com/prongbang/localizegen/blob/master/binary/linux/localizegen?raw=true)
- 🪟 [Windows](https://github.com/prongbang/localizegen/blob/master/binary/windows/localizegen.exe?raw=true)

## 🚀 Quick Start

### 1. Setup Google Sheet

Create a Google Sheet with your translations:

![Localize Sheet Example](/docs/sheet-localize.png)

Structure your sheet with:
- Column A: Keys (identifiers)
- Column B+: Translations for each language

### 2. Extract Sheet IDs

From your Google Sheet URL, extract the `DocumentId` and `SheetId`:

```
https://docs.google.com/spreadsheets/d/<DocumentId>/edit#gid=<SheetId>
```

Example:
```
https://docs.google.com/spreadsheets/d/1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU/edit#gid=0
```
- `DocumentId`: `1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU`
- `SheetId`: `0`

### 3. Generate Localization Files

Choose your platform and run the appropriate command:

#### 🤖 Android

Generate all languages:
```shell
localizegen -platform android \
            -target ./app/src/main/res \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

Generate specific language:
```shell
localizegen -platform android \
            -target ./app/src/main/res \
            -locale en \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

#### 🍎 iOS

Generate all languages:
```shell
localizegen -platform ios \
            -target ./ProjectName/Localization \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

Generate specific language:
```shell
localizegen -platform ios \
            -target ./ProjectName/Localization \
            -locale en \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

#### 🦋 Flutter

Generate all languages:
```shell
localizegen -platform flutter \
            -target ./lib/l10n \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

## ⚙️ Command Options

| Flag | Description | Required |
|------|-------------|----------|
| `-platform` | Target platform (`android`, `ios`, `flutter`) | ✅ |
| `-document` | Google Sheet document ID | ✅ |
| `-sheet` | Google Sheet tab ID | ✅ |
| `-target` | Output directory path | ❌ |
| `-locale` | Specific language code (e.g., `en`, `th`) | ❌ |
| `-filename` | Custom output filename | ❌ |

## 📁 Output Structure

### Android
```
android/
 ├── values/
 │   └── strings.xml
 ├── values-de/
 │   └── strings.xml
 ├── values-es/
 │   └── strings.xml
 └── ...
```

### iOS
```
ios/
 ├── Localizables.swift
 ├── de.lproj/
 │   └── Localizable.strings
 ├── en.lproj/
 │   └── Localizable.strings
 └── ...
```

### Flutter
```
flutter/
 ├── keys_localizations.dart
 └── sources_localizations.dart
```

## 💡 Examples

### Complete Android Project Setup
```shell
# Navigate to your Android project
cd MyAndroidApp

# Generate all localizations
localizegen -platform android \
            -target ./app/src/main/res \
            -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU \
            -sheet 0
```

### iOS with Custom Path
```shell
# Generate iOS localizations with custom path
localizegen -platform ios \
            -target ./MyApp/Resources/Localizations \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

### Flutter Single Language
```shell
# Generate only English for Flutter
localizegen -platform flutter \
            -target ./lib/l10n \
            -locale en \
            -document YOUR_DOCUMENT_ID \
            -sheet YOUR_SHEET_ID
```

## 🔧 Troubleshooting

### Common Issues

1. **Permission Denied**: Make sure your Google Sheet is publicly accessible
2. **Invalid Sheet ID**: Double-check the extracted DocumentId and SheetId
3. **Missing Dependencies**: Ensure Go is properly installed for source installation

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 💖 Support

If you find this tool helpful, please consider buying me a coffee:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

---
