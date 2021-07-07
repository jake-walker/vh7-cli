$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"

$url        = 'https://github.com/quark-links/quark-cli/releases/download/v__VERSION__/quark-cli___VERSION___Windows_i386.zip'   # download url, HTTPS preferred
$url64      = 'https://github.com/quark-links/quark-cli/releases/download/v__VERSION__/quark-cli___VERSION___Windows_x86_64.zip' # 64bit URL here (HTTPS preferred) or remove - if installer contains both (very rare), use $url

$packageArgs = @{
  packageName   = $env:ChocolateyPackageName
  unzipLocation = $toolsDir
  fileType      = 'exe'
  url           = $url
  url64bit      = $url64

  softwareName  = 'quark-cli*'

  checksum      = '__CHECKSUM_32__'
  checksumType  = 'sha256'
  checksum64    = '__CHECKSUM_64__'
  checksumType64= 'sha256'
}

Install-ChocolateyZipPackage $packageName $url $toolsDir $url64 -checksum $checksum -checksumType $checksumType -checksum64 $checksum64 -checksumType64 $checksumType64
