
# ntgrrc (Netgear Remote Control) CHANGELOG

## v0.8.1

* CHANGE: using "-v" parameter instead of "-d" for verbose output
* CHANGE: using OS specific temp-directory (typically $TMP or %TEMP%) to store login token (see #17) - thank you for your contribution @Sylensky
* fix the Windows build

----

## v0.8.0

* CHANGE: using "-v" parameter instead of "-d" for verbose output
* CHANGE: using OS specific temp-directory (typically $TMP or %TEMP%) to store login token (see #17) - thank you for your contribution @Sylensky

----

## v0.7.1

* add prompt for a password if not specified via command line argument (see #9)
* fixes in Github Actions to cover compiler errors earlier

----

## v0.6.0

* add new command for power cycling ports (see #10) - thank you for your contribution @davidk

----

## v0.5.1

* add support for JSON response format (alternative to Markdown) (issue #5)

----

## v0.?.?

* change using Go '1.20'

----

## v0.5.0

* add feature to set/change PoE settings at/to the switch (issue #2) thank you for your contribution @davidk
* add show help when no parameter given (issue #6)
* add more help description and --help-all flag (issue #7)
* change using Go 1.19

----

## v0.4.0

* add print POW status and settings as Markdown table, looks better and potentially could be rendered to HTML
* change, minor rename in table header "Temperature" -> "Temp."

----

## v0.3.1

* fix version information in the binaries

----

## v0.3.0

* support logins to multiple host at one moment in time

----

## v0.2.0

* show POE port settings
* fix detection of not logged in

----

## v0.1.0

* POE status of GS305EP switches can be shown

----

## v0.0.5

* testing 

----
