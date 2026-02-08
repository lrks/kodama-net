# kodama-net
LAN内のECHONET Liteをデバッグするためのツール。

## Usage
```
$ ./bin/kodama-net --help
kodama-net: ECHONET Lite explorer

Usage:
  kodama-net [flags]
  kodama-net [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  discovery   Discover ECHONET Lite devices on the local network
  help        Help about any command
  version     Print version information

Flags:
  -h, --help   help for kodama-net

Use "kodama-net [command] --help" for more information about a command.
```

### Discoveryコマンド
LAN内のECHONET Liteデバイスを探索します。
`--detail` オプションを付けると「状変アナウンスプロパティマップ」「Setプロパティマップ」「Getプロパティマップ」を取得し、さらにGetプロパティマップに記述されたEPCの値を取得します。

```
# 実行例
$ ./bin/kodama-net discovery
Sending discovery packet. waiting for responses (timeout: 20.0s)...
Discovered device: IPAddress=192.0.2.111, EOJ=0ef001
Discovered device: IPAddress=192.0.2.111, EOJ=05ff01
Discovered device: IPAddress=192.0.2.222, EOJ=0ef001
Discovered device: IPAddress=192.0.2.222, EOJ=013001
```

<details>
<summary>実行例（--detailオプション付き）</summary>

```
$ ./bin/kodama-net discovery --detail -t 5
Sending discovery packet. waiting for responses (timeout: 5.0s)...
Discovered device: IPAddress=192.0.2.111, EOJ=0ef001
Discovered device: IPAddress=192.0.2.111, EOJ=05ff01
Discovered device: IPAddress=192.0.2.222, EOJ=0ef001
Discovered device: IPAddress=192.0.2.222, EOJ=013001
Getting detailed info for device 192.0.2.111#0ef001...
  EOJ info:
    X1=0e X2=f0 X3=01
    Name=ノードプロファイル
    NameEN=Node profile
    ShortName=nodeProfile
  Property maps:
    AnnouncePropertyMap (0x9d):
      EPC=0x80
        Name=動作状態
        NameEN=Operating status
        ShortName=operatingStatus
        Description=ノードの動作状態を示す
        DescriptionEN=Indicates the node operating status.
      EPC=0xd5
        Name=インスタンスリスト通知
        NameEN=Instance list notification
        ShortName=instanceListNotification
        Description=自ノード内インスタンスに構成変化があった時のインスタンスリスト
        DescriptionEN=Instance list when self-node instance configuration is changed.
    SetPropertyMap (0x9e):
    GetPropertyMap (0x9f):
      EPC=0x80
        Name=動作状態
        NameEN=Operating status
        ShortName=operatingStatus
        Description=ノードの動作状態を示す
        DescriptionEN=Indicates the node operating status.
      EPC=0x82
        Name=Version情報
        NameEN=Version information
        ShortName=version
        Description=通信ミドルウェアが適用しているECHONET LiteのVersion、および通信ミドルウェアがサポートする電文タイプを示す
        DescriptionEN=Indicates ECHONET Lite version used by communication middleware and message types supported by communication middleware.
      EPC=0x83
        Name=識別番号
        NameEN=Identification number
        ShortName=id
        Description=オブジェクトを、ドメイン内で一意に識別するための番号
        DescriptionEN=Number to identify the node implementing the device object in the domain.
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常の発生状況を示す
        DescriptionEN=This property indicates whether a fault has occurred or not.
      EPC=0x8a
        Name=メーカコード
        NameEN=Manufacturer code
        ShortName=manufacturer(MC)
        Description=3バイトで指定
        DescriptionEN=3-byte manufacturer code
      EPC=0x8d
        Name=製造番号
        NameEN=Production number
        ShortName=serialNumber
        Description=ASCIIコードで指定
        DescriptionEN=This property indicates the production number using ASCII code.
      EPC=0x9d
        Name=状変アナウンスプロパティマップ
        NameEN=Status change announcement property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x9e
        Name=Setプロパティマップ
        NameEN=Set property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x9f
        Name=Getプロパティマップ
        NameEN=Get property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0xd3
        Name=自ノードインスタンス数
        NameEN=Number of self-node instances
        ShortName=selfNodeInstances
        Description=自ノードで保持するインスタンスの総数
        DescriptionEN=Total number of instances held by self-node
      EPC=0xd4
        Name=自ノードクラス数
        NameEN=Number of self-node classes
        ShortName=selfNodeClasses
        Description=自ノードで保持するクラス総数
        DescriptionEN=Total number of classes held by self-node
      EPC=0xd6
        Name=自ノードインスタンスリストS
        NameEN=Self-node instance list S
        ShortName=selfNodeInstanceListS
        Description=自ノード内インスタンスリスト
        DescriptionEN=Self-node instance list
      EPC=0xd7
        Name=自ノードクラスリストS
        NameEN=Self-node class list S
        ShortName=selfNodeClassListS
        Description=自ノード内クラスリスト
        DescriptionEN=
  Property values:
    EPC=0x9f PDC=14  Value=0x0d808283888a8d9d9e9fd3d4d6d7
    EPC=0x80 PDC=1   Value=0x30
    EPC=0x82 PDC=4   Value=0x010d0100
    EPC=0x83 PDC=17  Value=0xfe0000e2437562652d4a31000000000000
    EPC=0x88 PDC=1   Value=0x42
    EPC=0x8a PDC=3   Value=0x0000e2
    EPC=0x8d PDC=12  Value=0x434342384138304141343932
    EPC=0x9d PDC=3   Value=0x0280d5
    EPC=0x9e PDC=1   Value=0x00
    EPC=0xd3 PDC=3   Value=0x000001
    EPC=0xd4 PDC=2   Value=0x0002
    EPC=0xd6 PDC=4   Value=0x0105ff01
    EPC=0xd7 PDC=3   Value=0x0105ff
Getting detailed info for device 192.0.2.111#05ff01...
  EOJ info:
    X1=05 X2=ff X3=01
    Name=コントローラ
    NameEN=Controller
    ShortName=controller
  Property maps:
    AnnouncePropertyMap (0x9d):
      EPC=0x80
        Name=動作状態
        NameEN=Operation status
        ShortName=operationStatus
        Description=ON/OFFの状態を示す
        DescriptionEN=This property indicates the ON/OFF status.
      EPC=0x81
        Name=設置場所
        NameEN=Installation location
        ShortName=installationLocation
        Description=設置場所を示す
        DescriptionEN=This property indicates the installation location
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常(センサトラブル等)の発生状況を示す
        DescriptionEN=This property indicates whether a fault (e.g. a sensor trouble) has occurred or not.
    SetPropertyMap (0x9e):
      EPC=0x81
        Name=設置場所
        NameEN=Installation location
        ShortName=installationLocation
        Description=設置場所を示す
        DescriptionEN=This property indicates the installation location
    GetPropertyMap (0x9f):
      EPC=0x80
        Name=動作状態
        NameEN=Operation status
        ShortName=operationStatus
        Description=ON/OFFの状態を示す
        DescriptionEN=This property indicates the ON/OFF status.
      EPC=0x81
        Name=設置場所
        NameEN=Installation location
        ShortName=installationLocation
        Description=設置場所を示す
        DescriptionEN=This property indicates the installation location
      EPC=0x82
        Name=規格Version情報
        NameEN=Standard version information
        ShortName=protocol
        Description=対応するAPPENDIXのリリース番号を示す
        DescriptionEN=This property indicates the release number of the corresponding Appendix.
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常(センサトラブル等)の発生状況を示す
        DescriptionEN=This property indicates whether a fault (e.g. a sensor trouble) has occurred or not.
      EPC=0x8a
        Name=メーカコード
        NameEN=manufacturer code
        ShortName=manufacturer
        Description=3バイトで指定
        DescriptionEN=3-byte manufacturer code
      EPC=0x9d
        Name=状変アナウンスプロパティマップ
        NameEN=Status change announcement property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x9e
        Name=Setプロパティマップ
        NameEN=Set property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x9f
        Name=Getプロパティマップ
        NameEN=Get property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
  Property values:
    EPC=0x9f PDC=9   Value=0x08808182888a9d9e9f
    EPC=0x80 PDC=1   Value=0x30
    EPC=0x81 PDC=1   Value=0x00
    EPC=0x82 PDC=4   Value=0x00004c00
    EPC=0x88 PDC=1   Value=0x42
    EPC=0x8a PDC=3   Value=0x0000e2
    EPC=0x9d PDC=4   Value=0x03808188
    EPC=0x9e PDC=2   Value=0x0181
Getting detailed info for device 192.0.2.222#0ef001...
  EOJ info:
    X1=0e X2=f0 X3=01
    Name=ノードプロファイル
    NameEN=Node profile
    ShortName=nodeProfile
  Property maps:
    AnnouncePropertyMap (0x9d):
      EPC=0x80
        Name=動作状態
        NameEN=Operating status
        ShortName=operatingStatus
        Description=ノードの動作状態を示す
        DescriptionEN=Indicates the node operating status.
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常の発生状況を示す
        DescriptionEN=This property indicates whether a fault has occurred or not.
      EPC=0xd5
        Name=インスタンスリスト通知
        NameEN=Instance list notification
        ShortName=instanceListNotification
        Description=自ノード内インスタンスに構成変化があった時のインスタンスリスト
        DescriptionEN=Instance list when self-node instance configuration is changed.
    SetPropertyMap (0x9e):
      EPC=0xbf
        Name=個体識別情報
        NameEN=Unique identifier data
        ShortName=uid
        Description=ドメイン内で、各ノードを一意に識別するための2バイトデータ
        DescriptionEN=2 byte data to identify each node in a domain
      EPC=0xff
        Name=ユーザ定義領域
        NameEN=User defined area
        ShortName=userDefinedArea
        Description=ユーザ定義領域
        DescriptionEN=User defined area
    GetPropertyMap (0x9f):
      EPC=0x80
        Name=動作状態
        NameEN=Operating status
        ShortName=operatingStatus
        Description=ノードの動作状態を示す
        DescriptionEN=Indicates the node operating status.
      EPC=0x82
        Name=Version情報
        NameEN=Version information
        ShortName=version
        Description=通信ミドルウェアが適用しているECHONET LiteのVersion、および通信ミドルウェアがサポートする電文タイプを示す
        DescriptionEN=Indicates ECHONET Lite version used by communication middleware and message types supported by communication middleware.
      EPC=0x83
        Name=識別番号
        NameEN=Identification number
        ShortName=id
        Description=オブジェクトを、ドメイン内で一意に識別するための番号
        DescriptionEN=Number to identify the node implementing the device object in the domain.
      EPC=0xd3
        Name=自ノードインスタンス数
        NameEN=Number of self-node instances
        ShortName=selfNodeInstances
        Description=自ノードで保持するインスタンスの総数
        DescriptionEN=Total number of instances held by self-node
      EPC=0xd4
        Name=自ノードクラス数
        NameEN=Number of self-node classes
        ShortName=selfNodeClasses
        Description=自ノードで保持するクラス総数
        DescriptionEN=Total number of classes held by self-node
      EPC=0xd6
        Name=自ノードインスタンスリストS
        NameEN=Self-node instance list S
        ShortName=selfNodeInstanceListS
        Description=自ノード内インスタンスリスト
        DescriptionEN=Self-node instance list
      EPC=0xd7
        Name=自ノードクラスリストS
        NameEN=Self-node class list S
        ShortName=selfNodeClassListS
        Description=自ノード内クラスリスト
        DescriptionEN=
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常の発生状況を示す
        DescriptionEN=This property indicates whether a fault has occurred or not.
      EPC=0x89
        Name=異常内容
        NameEN=Fault description
        ShortName=faultDescription
        Description=異常内容
        DescriptionEN=Fault content
      EPC=0x8a
        Name=メーカコード
        NameEN=Manufacturer code
        ShortName=manufacturer(MC)
        Description=3バイトで指定
        DescriptionEN=3-byte manufacturer code
      EPC=0x8c
        Name=商品コード
        NameEN=Product code
        ShortName=productCode
        Description=ASCIIコードで指定
        DescriptionEN=Identifies the product using ASCII code.
      EPC=0x8d
        Name=製造番号
        NameEN=Production number
        ShortName=serialNumber
        Description=ASCIIコードで指定
        DescriptionEN=This property indicates the production number using ASCII code.
      EPC=0x9d
        Name=状変アナウンスプロパティマップ
        NameEN=Status change announcement property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x8e
        Name=製造年月日
        NameEN=Production date
        ShortName=productionDate
        Description=4バイトで指定
        DescriptionEN=4-byte production date code
      EPC=0x9e
        Name=Setプロパティマップ
        NameEN=Set property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x9f
        Name=Getプロパティマップ
        NameEN=Get property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0xbf
        Name=個体識別情報
        NameEN=Unique identifier data
        ShortName=uid
        Description=ドメイン内で、各ノードを一意に識別するための2バイトデータ
        DescriptionEN=2 byte data to identify each node in a domain
  Property values:
    EPC=0x9f PDC=17  Value=0x110100012120002020010101000103030a
    EPC=0x80 PDC=1   Value=0x30
    EPC=0x82 PDC=4   Value=0x010d0300
    EPC=0x83 PDC=17  Value=0xfe0000060105388d3dfffe035c7f0ef001
    EPC=0xd3 PDC=3   Value=0x000001
    EPC=0xd4 PDC=2   Value=0x0002
    EPC=0xd6 PDC=4   Value=0x01013001
    EPC=0xd7 PDC=3   Value=0x010130
    EPC=0x88 PDC=1   Value=0x42
    EPC=0x89 PDC=2   Value=0x0000
    EPC=0x8a PDC=3   Value=0x000006
    EPC=0x8c PDC=12  Value=0x4d41432d3930304946000000
    EPC=0x8d PDC=12  Value=0x323533323231333138330000
    EPC=0x9d PDC=4   Value=0x038088d5
    EPC=0x8e PDC=4   Value=0x07e90502
    EPC=0x9e PDC=3   Value=0x02bfff
    EPC=0xbf PDC=2   Value=0x0001
Getting detailed info for device 192.0.2.222#013001...
  EOJ info:
    X1=01 X2=30 X3=01
    Name=家庭用エアコン
    NameEN=Home air conditioner
    ShortName=homeAirConditioner
  Property maps:
    AnnouncePropertyMap (0x9d):
      EPC=0x80
        Name=動作状態
        NameEN=Operation status
        ShortName=operationStatus
        Description=ON/OFFの状態を示す
        DescriptionEN=This property indicates the ON/OFF status.
      EPC=0x81
        Name=設置場所
        NameEN=Installation location
        ShortName=installationLocation
        Description=設置場所を示す
        DescriptionEN=This property indicates the installation location
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常(センサトラブル等)の発生状況を示す
        DescriptionEN=This property indicates whether a fault (e.g. a sensor trouble) has occurred or not.
      EPC=0x8f
        Name=節電動作設定
        NameEN=Power-saving operation setting
        ShortName=powerSavingOperation
        Description=機器の節電動作を設定し、状態を取得する
        DescriptionEN=This property indicates whether the device is operating in power-saving mode.
      EPC=0xa0
        Name=風量設定
        NameEN=Air flow rate setting
        ShortName=airFlowLevel
        Description=風量レベルおよび風量自動状態を設定し、設定状態を取得する。風量レベルは8段階で指定
        DescriptionEN=Used to specify the air flow rate or use the function to automatically control the air flow rate, and to acquire the current setting. The air flow rate shall be selected from among the 8 predefined levels.
      EPC=0xb0
        Name=運転モード設定
        NameEN=Operation mode setting
        ShortName=operationMode
        Description=自動/冷房/暖房/除湿/送風/その他の運転モードを設定し、設定状態を取得する
        DescriptionEN=Used to specify the operation mode ('automatic','cooling','heating','dehumidification','air circulator' or 'other'), and to acquire the current setting.
    SetPropertyMap (0x9e):
      EPC=0x80
        Name=動作状態
        NameEN=Operation status
        ShortName=operationStatus
        Description=ON/OFFの状態を示す
        DescriptionEN=This property indicates the ON/OFF status.
      EPC=0x81
        Name=設置場所
        NameEN=Installation location
        ShortName=installationLocation
        Description=設置場所を示す
        DescriptionEN=This property indicates the installation location
      EPC=0x8f
        Name=節電動作設定
        NameEN=Power-saving operation setting
        ShortName=powerSavingOperation
        Description=機器の節電動作を設定し、状態を取得する
        DescriptionEN=This property indicates whether the device is operating in power-saving mode.
      EPC=0x93
        Name=遠隔操作設定
        NameEN=Remote control setting
        ShortName=remoteControl
        Description=公衆回線を介した操作か否かを示す。(0x41、0x42)通信回線の状態が正常か否かを示す。(0x61、0x62)
        DescriptionEN=This property indicates whether remote control is through a public network or not. (0x41, 0x42)This property indicates whether the status of the communication line is normal or not. (0x61, 0x62)
      EPC=0xa0
        Name=風量設定
        NameEN=Air flow rate setting
        ShortName=airFlowLevel
        Description=風量レベルおよび風量自動状態を設定し、設定状態を取得する。風量レベルは8段階で指定
        DescriptionEN=Used to specify the air flow rate or use the function to automatically control the air flow rate, and to acquire the current setting. The air flow rate shall be selected from among the 8 predefined levels.
      EPC=0xa1
        Name=風向自動設定
        NameEN=Automatic control of air flow direction setting
        ShortName=automaticControlAirFlowDirection
        Description=風向き上下左右のAUTO/非AUTOを設定し、設定状態を取得する
        DescriptionEN=Used to specify whether or not to use the automatic air flow direction control function, to specify the plane(s) (vertical and/or horizontal) in which the automatic air flow direction control function is to be used, and to acquire the current setting.
      EPC=0xa3
        Name=風向スイング設定
        NameEN=Automatic swing of air flow setting
        ShortName=automaticSwingAirFlow
        Description=風向スイングOFF/上下/左右/上下左右を設定し、設定状態を取得する
        DescriptionEN=Used to specify whether or not to use the automatic air flow swing function, to specify the plane(s) (vertical and/or horizontal) in which the automatic air flow swing function is to be used, and to acquire the current setting.
      EPC=0xa4
        Name=風向上下設定
        NameEN=Air flow direction (vertical) setting
        ShortName=airFlowDirectionVertical
        Description=上下方向の風向きを5通りのパターンで設定し、設定状態を取得する
        DescriptionEN=Used to specify the air flow direction in the vertical plane by selecting a pattern from among the 5 predefined patterns, and to acquire the current setting.
      EPC=0xa5
        Name=風向左右設定
        NameEN=Air flow direction (horizontal) setting
        ShortName=airFlowDirectionHorizontal
        Description=左右方向の風向きを31通りのパターンで設定し、設定状態を取得する
        DescriptionEN=Used to specify the air flow direction(s) in the horizontal plane by selecting a pattern from among the 31 predefined patterns, and to acquire the current setting.
      EPC=0xb0
        Name=運転モード設定
        NameEN=Operation mode setting
        ShortName=operationMode
        Description=自動/冷房/暖房/除湿/送風/その他の運転モードを設定し、設定状態を取得する
        DescriptionEN=Used to specify the operation mode ('automatic','cooling','heating','dehumidification','air circulator' or 'other'), and to acquire the current setting.
      EPC=0xb3
        Name=温度設定値
        NameEN=Set temperature value
        ShortName=targetTemperature
        Description=温度設定値を設定し、設定状態を取得する
        DescriptionEN=Used to set the temperature and to acquire the current setting.
      EPC=0xd0
        Name=ブザー
        NameEN=Buzzer
        ShortName=beepBuzzer
        Description=ブザー音を発生する
        DescriptionEN=Used to generate a buzzer sound.
    GetPropertyMap (0x9f):
      EPC=0x80
        Name=動作状態
        NameEN=Operation status
        ShortName=operationStatus
        Description=ON/OFFの状態を示す
        DescriptionEN=This property indicates the ON/OFF status.
      EPC=0xa0
        Name=風量設定
        NameEN=Air flow rate setting
        ShortName=airFlowLevel
        Description=風量レベルおよび風量自動状態を設定し、設定状態を取得する。風量レベルは8段階で指定
        DescriptionEN=Used to specify the air flow rate or use the function to automatically control the air flow rate, and to acquire the current setting. The air flow rate shall be selected from among the 8 predefined levels.
      EPC=0xb0
        Name=運転モード設定
        NameEN=Operation mode setting
        ShortName=operationMode
        Description=自動/冷房/暖房/除湿/送風/その他の運転モードを設定し、設定状態を取得する
        DescriptionEN=Used to specify the operation mode ('automatic','cooling','heating','dehumidification','air circulator' or 'other'), and to acquire the current setting.
      EPC=0x81
        Name=設置場所
        NameEN=Installation location
        ShortName=installationLocation
        Description=設置場所を示す
        DescriptionEN=This property indicates the installation location
      EPC=0xa1
        Name=風向自動設定
        NameEN=Automatic control of air flow direction setting
        ShortName=automaticControlAirFlowDirection
        Description=風向き上下左右のAUTO/非AUTOを設定し、設定状態を取得する
        DescriptionEN=Used to specify whether or not to use the automatic air flow direction control function, to specify the plane(s) (vertical and/or horizontal) in which the automatic air flow direction control function is to be used, and to acquire the current setting.
      EPC=0x82
        Name=規格Version情報
        NameEN=Standard version information
        ShortName=protocol
        Description=対応するAPPENDIXのリリース番号を示す
        DescriptionEN=This property indicates the release number of the corresponding Appendix.
      EPC=0x83
        Name=識別番号
        NameEN=Identification number
        ShortName=id
        Description=オブジェクトを固有に識別する番号
        DescriptionEN=A number that allows each object to be uniquely identified.
      EPC=0xa3
        Name=風向スイング設定
        NameEN=Automatic swing of air flow setting
        ShortName=automaticSwingAirFlow
        Description=風向スイングOFF/上下/左右/上下左右を設定し、設定状態を取得する
        DescriptionEN=Used to specify whether or not to use the automatic air flow swing function, to specify the plane(s) (vertical and/or horizontal) in which the automatic air flow swing function is to be used, and to acquire the current setting.
      EPC=0xb3
        Name=温度設定値
        NameEN=Set temperature value
        ShortName=targetTemperature
        Description=温度設定値を設定し、設定状態を取得する
        DescriptionEN=Used to set the temperature and to acquire the current setting.
      EPC=0xa4
        Name=風向上下設定
        NameEN=Air flow direction (vertical) setting
        ShortName=airFlowDirectionVertical
        Description=上下方向の風向きを5通りのパターンで設定し、設定状態を取得する
        DescriptionEN=Used to specify the air flow direction in the vertical plane by selecting a pattern from among the 5 predefined patterns, and to acquire the current setting.
      EPC=0x85
        Name=積算消費電力量計測値
        NameEN=Measured cumulative electric energy consumption
        ShortName=consumedCumulativeElectricEnergy
        Description=機器の積算消費電力量を0.001kWhで示す
        DescriptionEN=This property indicates the cumulative electric energy consumption of the device in increments of 0.001kWh.
      EPC=0xa5
        Name=風向左右設定
        NameEN=Air flow direction (horizontal) setting
        ShortName=airFlowDirectionHorizontal
        Description=左右方向の風向きを31通りのパターンで設定し、設定状態を取得する
        DescriptionEN=Used to specify the air flow direction(s) in the horizontal plane by selecting a pattern from among the 31 predefined patterns, and to acquire the current setting.
      EPC=0x86
        Name=メーカ異常コード
        NameEN=Manufacturer's fault code
        ShortName=manufacturerFaultCode
        Description=各メーカ独自の異常コードを示す
        DescriptionEN=This property indicates the manufacturer-defined fault code.
      EPC=0x88
        Name=異常発生状態
        NameEN=Fault status
        ShortName=faultStatus
        Description=何らかの異常(センサトラブル等)の発生状況を示す
        DescriptionEN=This property indicates whether a fault (e.g. a sensor trouble) has occurred or not.
      EPC=0x89
        Name=異常内容
        NameEN=Fault description
        ShortName=faultDescription
        Description=異常内容
        DescriptionEN=Describes the fault.
      EPC=0x8a
        Name=メーカコード
        NameEN=manufacturer code
        ShortName=manufacturer
        Description=3バイトで指定
        DescriptionEN=3-byte manufacturer code
      EPC=0x9a
        Name=積算運転時間
        NameEN=Cumulative operating time
        ShortName=hourMeter
        Description=現在までの運転時間の積算値を単位1バイト、時間4バイトで示す
        DescriptionEN=This property indicates the cumulative number of days, hours, minutes or seconds for which the device has operated, using 1 byte for the unit and 4 bytes for the time.
      EPC=0x8b
        Name=事業場コード
        NameEN=Business facility code
        ShortName=businessFacilityCode
        Description=3バイトの事業場コードで指定
        DescriptionEN=3-byte business facility code
      EPC=0xbb
        Name=室内温度計測値
        NameEN=Measured value of room temperature
        ShortName=roomTemperature
        Description=室内温度計測値
        DescriptionEN=Measured value of room temperature
      EPC=0x9d
        Name=状変アナウンスプロパティマップ
        NameEN=Status change announcement property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0x9e
        Name=Setプロパティマップ
        NameEN=Set property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
      EPC=0xbe
        Name=外気温度計測値
        NameEN=Measured outdoor air temperature
        ShortName=outdoorTemperature
        Description=外気温度計測値
        DescriptionEN=This property indicates the measured outdoor air temperature.
      EPC=0x8f
        Name=節電動作設定
        NameEN=Power-saving operation setting
        ShortName=powerSavingOperation
        Description=機器の節電動作を設定し、状態を取得する
        DescriptionEN=This property indicates whether the device is operating in power-saving mode.
      EPC=0x9f
        Name=Getプロパティマップ
        NameEN=Get property map
        ShortName=DEL
        Description=プロパティ数が15以下の場合はEPCを列挙、16以上の場合はビットマップで記述する。
        DescriptionEN=Enumuration of EPC in case of the count is less than 16, or bitmap in case of the count is more than 15
  Property values:
    EPC=0x9f PDC=17  Value=0x180d05010d040501000101030900020a03
    EPC=0x80 PDC=1   Value=0x30
    EPC=0xa0 PDC=1   Value=0x35
    EPC=0xb0 PDC=1   Value=0x43
    EPC=0x81 PDC=1   Value=0x00
    EPC=0xa1 PDC=1   Value=0x42
    EPC=0x82 PDC=4   Value=0x00004b00
    EPC=0x83 PDC=17  Value=0xfe0000060105388d3dfffe035c7f013001
    EPC=0xa3 PDC=1   Value=0x31
    EPC=0xb3 PDC=1   Value=0x13
    EPC=0xa4 PDC=1   Value=0x44
    EPC=0x85 PDC=4   Value=0x0054c784
    EPC=0xa5 PDC=1   Value=0x54
    EPC=0x86 PDC=10  Value=0x06000006000000028000
    EPC=0x88 PDC=1   Value=0x42
    EPC=0x89 PDC=2   Value=0x0000
    EPC=0x8a PDC=3   Value=0x000006
    EPC=0x9a PDC=5   Value=0x420013f7b8
    EPC=0x8b PDC=3   Value=0x000029
    EPC=0xbb PDC=1   Value=0x14
    EPC=0x9d PDC=7   Value=0x068081888fa0b0
    EPC=0x9e PDC=13  Value=0x0c80818f93a0a1a3a4a5b0b3d0
    EPC=0xbe PDC=1   Value=0x02
    EPC=0x8f PDC=1   Value=0x42
```
</details>

## Makefileの説明
* `make generate`
  * [MRA](https://echonet.jp/spec_mra_rr3/)から`internal/echonetlite/mra_generated.go`を生成する
  * 事前にMRAのzipファイルをダウンロードし、中身を`tools/mra/data`に展開しておく必要がある
    * `tools/mra/data/metaData.json` などが存在することになる（Git管理外）
* `make build`
  * `bin/kodama-net`を生成する
* `make build-all`
  * `bin/`配下に各種プラットフォーム用のバイナリを生成する
