# localizegen

```shell
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

Generate Localization tools for Mobile Application an Android, iOS

## Download

### Binary

- [macOS](https://github.com/prongbang/localizegen/blob/master/binary/macos/localizegen?raw=true)
- [Linux](https://github.com/prongbang/localizegen/blob/master/binary/linux/localizegen?raw=true)
- [Mindows](https://github.com/prongbang/localizegen/blob/master/binary/windows/localizegen.exe?raw=true)

## Flag

- `-platform`   : `android` or `ios`        *require
- `-target`     : Target path
- `-locale`     : `en`, `th`, etc
- `-document`   : Google sheet document id  *require
- `-sheet`      : Google sheet id           *require
- `-filename`   : Custom file name

## Android

- Generate by language

```shell script
$ localizegen -platform android -target ./app/src/main/res -locale en -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0
```

- Generate all language supported

```shell script
$ localizegen -platform android -target ./app/src/main/res -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0
```

## iOS

- Generate by language

```shell script
$ localizegen -platform ios -target ./ProjectName/Locallization -locale en -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0
```

- Generate all language supported

```shell script
$ localizegen -platform ios -target ./ProjectName/Locallization -document 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU -sheet 0
```

## Setup

- 1. Create a Google Sheet

![Localize](/docs/sheet-localize.png)

- 2. Extract from the link the `DocumentId` and `SheetId` values

```shell script
https://docs.google.com/spreadsheets/d/<DocumentId>/edit#gid=<SheetId>
```

Example

```shell script
https://docs.google.com/spreadsheets/d/1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU/edit#gid=0

DocumentId  = 1r91ECV-As0XtuqGKXU7dXnoY4og9XPBoCqwRcdio6EU
SheetId     = 0
```
