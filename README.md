LibGuides Tools
===============

A Golang package for working with LibGuides exported XML.


[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg?style=flat-square)](https://choosealicense.com/licenses/bsd-3-clause)
[![Latest release](https://img.shields.io/github/v/release/caltechlibrary/springytools.svg?style=flat-square&color=b44e88)](https://github.com/caltechlibrary/libguildestools/releases)
<!-- [![DOI](https://data.caltech.edu/badge/201106666.svg)](https://data.caltech.edu/badge/latestdoi/201106666) -->


Table of contents
-----------------

* [Introduction](#introduction)
* [Installation](#installation)
* [Known issues and limitations](#known-issues-and-limitations)
* [Getting help](#getting-help)
* [License](#license)
* [Authors and history](#authors-and-history)
* [Acknowledgments](#authors-and-acknowledgments)


Introduction
------------

There is a periodic need to work with exported LibGuides XML in Caltech Library.  This is a Golang
package for working with the exported data. Go provides a robust may of mapping simple data structures
to and from XML (or JSON). This makes working with XML very easy in a consistent fashion. It seem time to move beyond my usual Bash/sed/python scripts.

One program is currently provided with springytools, __lgxml2sjon__ which converts a LibGuides
XML export file into JSON.


Installation
------------

This is a Golang package providing two commands for working with LibGuides' exported XML. To
compile you will need Go 1.16 or better, GNU Make and Stephen Dolan's [jq](https://stedolan.github.io/jq/) for browser JSON output.

### Steps to compile from source

1. clone the [repository](https://github.com/caltechlibrary/springytools)
2. change into the clone directory
3. test
4. build the command line tool __lgxml2json__
5. use __lgxml2json__ and test output with __jq__
  - Replace "LibGuides_export_XXXXX.xml" with the file path to your exported LibGuides XML file
6. install __lgxml2json__

Example commands to execute in the shell (e.g. Terminal on macOS, xterm on Linux)

~~~
git clone git@github.com:caltechlibrary/springytools
cd springytools
make
make test
make install
~~~

By default installation is to your `$HOME/bin` directory. This directory should be in 
your shell's "PATH".

You can get a brief description of the commands using the `-h` option with the command.

~~~
lgxml2json -h
lglinkreport -h
~~~


Known issues and limitations
----------------------------

This library is currently written to perform the LibGuides link analysis.
It only provides the commands I needed to do the data analysis. It will grow as needed.

The exported XML output from the LibGuides may not be valid UTF-8.  UTF-8 encoding
is required to successfully parse the export file. Looking at the raw XML markup in __vim__
I noticed a number of control code sequences. This corresponded to the errors on parsing
the unsanitized XML file. The problem characters appear as `^A, ^K, ^L, ^S, ^C, ^R`. These
maybe non-UTF-8 characters embedded as UTF-8 when the rich text documents were pasted in via
the LibGuides edit UI. My hunch is these were pasted in/imported from Word documents. Remove
the offending characters allowed the export to parse successfully. These edits are destructive
as some of the codes probably represent UTF-8 characters used in non-English European names or
terminology.




Getting help
------------

File an [issue](https://github.com/caltechlibrary/springytools/issues) on GitHub.



License
-------

Software produced by the Caltech Library is Copyright © 2021 California Institute of Technology.  This software is freely distributed under a BSD/MIT type license.  Please see the [LICENSE](LICENSE) file for more information.


Authors and history
---------------------------

- R. S. Doiel, Software Developer, Digital Library Development, Caltech Library


Acknowledgments
---------------

This work was funded by the California Institute of Technology Library.

(If this work was also supported by other organizations, acknowledge them here.  In addition, if your work relies on software libraries, or was inspired by looking at other work, it is appropriate to acknowledge this intellectual debt too.)

<div align="center">
  <br>
  <a href="https://www.caltech.edu">
    <img width="100" height="100" src="https://raw.githubusercontent.com/caltechlibrary/springytools/main/.graphics/caltech-round.png">
  </a>
</div>
