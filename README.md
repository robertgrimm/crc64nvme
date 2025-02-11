# crc64nvme

crc64nvme calculates a CRC64 checksum using the NVME polynomial.

## Build Instructions

```
go build
```

## Usage

crc64nvme reads data from STDIN and outputs the checksum.

```
$ echo 'Hello, World!' | crc64nvme | base64
SoXXbx67KpE=
```
