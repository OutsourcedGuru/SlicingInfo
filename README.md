# SlicingInfo
An executable to display slicing information from the indicated GCODE file.

## Syntax
This assumes that you've installed it per the instructions below so that it will be in your path.

```
$ SlicingInfo PathToGcodeFile
```

## Installation
The installation of this depends upon whether or not you also have a computer based upon the OS X operating system and further, whether or not you already have the Go compiler itself.

### Mac
Here are the instructions for installing this executable if you are on an Apple-based computer.

#### If you have Go installed:

```
$ cd /usr/local/go/bin
$ sudo curl https://github.com/OutsourcedGuru/SlicingInfo/raw/master/bin/SlicingInfo SlicingInfo
$ cd ~/Desktop
$ which SlicingInfo
/usr/local/go/bin/SlicingInfo
$ SlicingInfo *RC_MyFile*.gcode
Slicer:          Cura_SteamEngine 2.3.1
Layers:          239
Quality:         low
Profile:         Low Quality Robo C2
Filament size:   1.75
Hotend temp:     190
Bed temp:        0
Supports:        False
Retraction:      True
Jerk:            True
Speed 1st layer: 10
Print speed:     50
Travel speed:    80
Infill pattern:  cubic

Finished.
```

#### If you have don't have Go installed:

```
$ cd /usr/local/bin
$ sudo curl https://github.com/OutsourcedGuru/SlicingInfo/raw/master/bin/SlicingInfo SlicingInfo
$ cd ~/Desktop
$ which SlicingInfo
/usr/local/SlicingInfo
$ SlicingInfo *RC_MyFile*.gcode
Slicer:          Cura_SteamEngine 2.3.1
Layers:          239
Quality:         low
Profile:         Low Quality Robo C2
Filament size:   1.75
Hotend temp:     190
Bed temp:        0
Supports:        False
Retraction:      True
Jerk:            True
Speed 1st layer: 10
Print speed:     50
Travel speed:    80
Infill pattern:  cubic

Finished.
```

### Windows or UNIX
Here are the instructions for building this executable if you are on a different operating system.

#### Install Go
The first step is to [install the Go language compiler](https://golang.org).

It's then usual to create a Go working folder under your user's profile.

```
# These two are optional, depending upon whether
# or not you did this during the Go installation
$ mkdir -p ~/go/src
$ cd ~/go/src
$ go get github.com/OutsourcedGuru/SlicingInfo/
```

This should download everything required and build it for you. Assuming that you installed Go correctly earlier and it's in your path, you should then be able to run it as in the instructions above.