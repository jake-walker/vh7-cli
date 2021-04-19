$ErrorActionPreference = 'Stop';
$toolsDir   = "$(Split-Path -parent $MyInvocation.MyCommand.Definition)"

$url        = 'https://github.com/jake-walker/vh7-cli/releases/download/v2.1.0/vh7-cli-v2.1.0-windows-386.zip' # download url, HTTPS preferred
$url64      = 'https://github.com/jake-walker/vh7-cli/releases/download/v2.1.0/vh7-cli-v2.1.0-windows-386.zip' # 64bit URL here (HTTPS preferred) or remove - if installer contains both (very rare), use $url

$packageArgs = @{
  packageName   = $env:ChocolateyPackageName
  unzipLocation = $toolsDir
  fileType      = 'exe'
  url           = $url
  url64bit      = $url64

  softwareName  = 'vh7-cli*'

  checksum      = 'c088d7a9ba17ffe4ab9b46094f1d2ba4'
  checksumType  = 'md5'
  checksum64    = 'cd9a8c0d0fdc093763c63280c6f52d66'
  checksumType64= 'md5'
}

Install-ChocolateyZipPackage $packageName $url $toolsDir $url64 -checksum $checksum -checksumType $checksumType -checksum64 $checksum64 -checksumType64 $checksumType64
