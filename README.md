# File Consolidator
Application to de-duplicate files
### Status
This is a convenience utility created for my personal use. It is not currently well tested, so there are no guarantees if you choose to use this application.

**_!!! USE AT YOUR OWN RISK !!!_**
## Build
```
make
```
## Usage
```
bin/consolidator <list of dirs space separated> -<options>
```
### Options
```
-v Verbose output
```
Example
```
bin/consolidator data/first data/second
bin/consolidator data/first data/second -v
```

### Output
```
Starting app

Duplication Report:
--------------------
-> Total Files Scanned: 9
-> Unique Files Found: 3
-> Duplicates Found: 6
```
## Migration
The application will ask if you would like to migrate unique files.
```
Copy unique files to new directory? (yes/no)
yes
```
You will then be asked for the destination directory. This directory must not already exist.
```
Destination directory (must not exist):
data/unique
```
The application will then migrate the files
```
Copying unique files to data/unique
Creating directory: data/unique
migrating file: data/first/sub/sub2/foo3.txt
migrating file: data/first/bar.txt
migrating file: data/second/zed.txt
App finished...
```

### Result
The unique files are now in their respective location relative to the root destiation folder (data3 in this case)
```
$ tree data/unique
data/unique
├── bar.txt
├── sub
│   └── sub2
│       └── foo3.txt
└── zed.txt

2 directories, 3 files
```
