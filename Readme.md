CIQDB
=====

**CIQDB** is the unofficial [**C**onnect **IQ**](https://developer.garmin.com/connect-iq/programmers-guide/) Debugger.

It takes an already-compiled `.prg` or `.iq` file, make with the Connect IQ SDK, and parses it. It provides any symbols, debug sections, and more pulled from the binary.

CIQDB is written in Go, with no current dependencies. 

Building
---------

To build, all you need is a Go compiler, and to run `go generate && go build`. A makefile is provided to simplify this to just `make`.

Usage
-----

`ciqdb <path/to/file.prg>`. There are no flags or other arguments, currently. Just pass in the path to a compiled CIQ program.

The following is example output, when compiling [crystal-face](https://github.com/warmsound/crystal-face) with CIQ SDK 3.0.0-beta1. Some sections are trimmed for length

```
$ ./ciqdb crystal-face.prg
d000d000 - Head (13 bytes)
    CIQ version: 3.0.0
    App Trial Enabled: false
6060c0de - Entry Points (38 bytes)
  Entry Point 0
    UUID: bde5c058244911e8b4670ed5f89f718b
    Type: WatchFace
    AppName: CrystalApp
    Module: globals
    icon: f8
da7ababe - Data (5186 bytes)
    extends Offset: 0
    static Entry: 0
    parent module: Unknown Symbol ID: 0
    module ID: statics
    app types: 3f
    fields:
        globals_Rez: 564  ModuleDef
        globals_Rez_Drawables: 643  ModuleDef
        globals_Rez_Menus: 682  ModuleDef
        globals_Rez_Fonts: 713  ModuleDef
        globals_Rez_JsonData: 808  ModuleDef
        globals_Rez_Layouts: 839  ModuleDef
        globals_Rez_Strings: 878  ModuleDef
        LeftGoalCurrent: 87  ModuleDef
    extends Offset: 0
    static Entry: 0
    parent module: Unknown Symbol ID: 0
    module ID: globals
    app types: 3f
    fields:
        GoalMeter: 1907 const  ClassDefinition
        DateLine: 1780 const  ClassDefinition
        MoveBar: 2289 const  ClassDefinition
        CrystalApp: 2170 const  ClassDefinition
        DEFAULT_PARAMS: 0  Null
        testGetSegmentScale: 268444038 const  Method
        testGetSegments: 268444150 const  Method
        ThickThinTime: 357 const  ClassDefinition
        CrystalView: 1405 const  ClassDefinition
        GoalMeterMask: 294 const  ClassDefinition
        globals_Rez_Drawables: 0  Null
        <globals_ThickThinTime_<>mSecondsFont>: 268454218 hidden  Method
        <globals_ThickThinTime_<>mSecondsFont>: 268454218 hidden  Method
... //SNIP

c0debabe - Code (20531 bytes)
c0de7ab1 - Code Table (aka PCtoLineNum) (12642 bytes)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:35 initialize (pc 268435456)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:36 initialize (pc 268435460)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:38 initialize (pc 268435480)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:39 initialize (pc 268435496)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:40 initialize (pc 268435512)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:41 initialize (pc 268435528)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:42 initialize (pc 268435544)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:44 initialize (pc 268435560)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:47 getWidth (pc 268435580)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:48 getWidth (pc 268435584)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:50 getWidth (pc 268435584)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:51 getWidth (pc 268435584)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:53 getWidth (pc 268435584)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:54 getWidth (pc 268435601)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:55 getWidth (pc 268435630)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:56 getWidth (pc 268435643)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:57 getWidth (pc 268435726)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:59 getWidth (pc 268435750)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:62 getWidth (pc 268435760)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:65 setValues (pc 268435764)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:70 setValues (pc 268435768)
    /home/dan/dev/MonkeyC/samples/crystal-face/source/GoalMeter.mc:71 setValues (pc 268435782)
... //SNIP

c1a557b1 - Class Table (class imports) (1770 bytes)
    module: Toybox_Test, class: Logger
    module: Toybox_Test, class: AssertException
    module: Toybox_Test, class: Test
    module: Toybox_System, class: Stats
    module: Toybox_System, class: DeviceSettings
    module: Toybox_System, class: Intent
    module: Toybox_System, class: ClockTime
    module: Toybox_System, class: PreviousOperationNotCompleteException
    module: Toybox_System, class: ConnectionInfo
    module: Toybox_System, class: ServiceDelegate
    module: Toybox_System, class: AppNotInstalledException
    module: Toybox_System, class: UnexpectedAppTypeException
    module: Toybox_SensorHistory, class: SensorHistoryIterator
    module: Toybox_SensorHistory, class: SensorSample
    module: Toybox_Activity, class: Info
    module: Toybox_PersistedContent, class: Waypoint
    module: Toybox_PersistedContent, class: Iterator
    module: Toybox_PersistedContent, class: Track
    module: Toybox_PersistedContent, class: Workout
... //SNIP

f00d600d - Resources (25199 bytes)
6000db01 - Permissions (10 bytes)
    Toybox_SensorHistory
    Toybox_UserProfile

0ece7105 - Exceptions (2 bytes)

5717b015 - Symbols (60133 bytes)
    236: LeftFieldIcon
    8388920: Moment
    8389131: setColor
    8389245: deviceType
    8388718: error
    8388831: SPORT_RUNNING
    8389539: mInitialText
    8390196: TimerRunMultiple
    8390666: BikeSpeedCadence
    8388850: SUB_SPORT_GENERIC
    8389586: TooManyEventsException
    8389308: toRadians
    8390022: restoreHeadlightsNetworkModeControl
... //SNIP

5e771465 - Settings (563 bytes)
    HoursColourOverride: 4294967295
    AlwaysShowMoveBar: false
    HideSeconds: true
    Theme: 0
    RightGoalType: 1
    ThemeColour: 43775
    AppVersion: 1.6.0
    LeftFieldType: 0
    RightFieldType: 2
    MinutesColourOverride: 4294967295
    MonoDarkColour: 11184810
    BackgroundColour: 0
    CaloriesGoal: 2000
    MonoLightColour: 16777215
    MinutesColour: 43775
    HideHoursLeadingZero: true
    LeftGoalType: 0
    MeterBackgroundColour: 5592405
    HoursColour: 43775
    CenterFieldType: 1

e1c0de12 - Developer Signature (1028 bytes)
    signature: 8adb0475cd0dc282ec27...
    modulus: c1c90eeff1c61957e2c3...
    exponent: 65537
00000000 - End (0 bytes)

```


Future
-------

Currently, CIQDB is just able to print information about _most_ of a PRG's file's information. There are still some sections being reverse-engineered (code and resources).

Future plans for this project are to completely parse all prg sections, and act as an interactive debugger, similar to `gdb`. CIQDB will tap into the existing Connect IQ simulator, and the integrated debugger that ships the the SDK. It will pass breakpoint/continue/etc commands through to the simulator. But through parsing the prg, it should be able to have much richer information and display.

License
-------

This code is licensed under the MIT License, Copyright 2018 Dan Panzarella. See the `LICENSE` file for full license.

---

This repository and code are not affiliated with or supported by Garmin in any way. Connect IQ name and title is owned by Garmin Ltd. and its subsidiaries.