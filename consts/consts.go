package consts

import "path/filepath"

const FilePackage = "boss.json"

const FolderDependencies = "modules"
const FolderBossHome = ".boss"

const BinFolder string = ".bin"
const BplFolder string = ".bpl"
const DcpFolder string = ".dcp"
const DcuFolder string = ".dcu"

const BossConfigFile = "boss.cfg.json"

const MinimalDependencyVersion string = ">0.0.0"

var EnvBossBin = filepath.Join(".", FolderDependencies, BinFolder)

const XmlTagNameProperty string = "PropertyGroup"
const XmlTagNamePropertyAttribute string = "Condition"
const XmlTagNamePropertyAttributeValue string = "'$(Base)'!=''"

const XmlTagNameLibraryPath string = "DCC_UnitSearchPath"

const Version string = "v2.1.2"

const BossInternalDir = "{internal}"

const BplIdentifierName = "BplIdentifier.exe"
