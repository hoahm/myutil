# MyUtil

MyUtil is a simple library written in Golang for manipulate string, number, ... for quickly use in the future. It's also my very first Golang project that integrates BDD testing.

# Features

* numutil
    * Xrand(min, max int) int {}
    * AtoSI(s, sep string) []int {}
    * IsDigit(v string) bool {}

* cmdutil
    * ReadLinePmt(prompt string) (string, error) {}
    * ReadLine() (string, error) {}

* fileutil
    * CurrentDir() string {}
    * HomeDir() string {}
    * Exists(name string) bool {}
    * FilePath(dir, fileName, ext string) string {}

* stringutil
    * RandString(length int) string {}
    * Reverse(s string) string {}

# How to use

Get it:

    go get github.com/hoahm/myutil

Import

    import "github.com/hoahm/myutil"

# Credit

I collect and make references to many different sources on the Internet. I eventually do not know who is the author, so that sometimes I may not attach the reference link. If you find somethings that was licensed to you or someone else, please don't hesitate to let me know.

I would like to give a big thank to all of you, guys. :)        

# Contribution

All contributions are welcome. Please make a Pull Request or send an email to [Nobi Hoang]<nobi.hoa@gmail.com>.

# License

MIT License

Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.