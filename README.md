# adbutil
Android Debug Bridge Utility

adb is a powerful utility, but it's extremely barebones and repetitive in it's vanilla state. It has some builtin support for scripting, but this project ultimately seeks to make it more robust in general.

# status
Device management is in version 0.1 and further work isn't expected on that logic, but it's robustness needs to be tested going forward.

`adb install <myapk>.apk` is the first critical command to implement now that command framework is prepared, then `adb shell am start` and other general package management commands.

When this tool is complete, the user should only have to run `adbutil.exe` and then call simple commands such as:
- `ls apk` (to get apks in pwd)
- `install <apk>`,
- `ls pkg` (to get installed package ids)
- `uninstall <packageid>`
- `ls activity pkg` (to automatically locate an installed package's activities)
- `run <activity>`
- `stop <activity>`
*as well as*, importantly, commands for setting 'working' values so that <apk>, <packageid>, and <activity> may be omitted.
  
# setup
This project uses a vanilla go setup, and assumes that you already have adb in your PATH.

In order to use this project, clone it into your %GOHOME%/src/ directory

