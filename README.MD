# MyUtil

MyUtil is a simple library written in Golang for manipulate string, number, ... for quickly use in the future. It's also my very first Golang project that integrates BDD testing.

* [Utils](#utils)
  * [Numutil](#numutil)
  * [Cmdutil](#cmdutil)
  * [fileutil](#fileutil)
  * [Stringutil](#stringutil)
* [Data structures](#data-structures)
  * [Linked list](#linked-list)
  * [Queue](#queue)
  * [Stack](#stack)

# Features

## Utils

### numutil

    import "github.com/hoahm/myutil/util/numutil"

    numutil.Xrand(min, max int) int {}
    numutil.AtoSI(s, sep string) []int {}
    numutil.IsDigit(v string) bool {}

### cmdutil

    import "github.com/hoahm/myutil/util/cmdutil"

    cmdutil.ReadLinePmt(prompt string) (string, error) {}
    cmdutil.ReadLine() (string, error) {}

### fileutil

    import "github.com/hoahm/myutil/util/fileutil"

    fileutil.CurrentDir() string {}
    fileutil.HomeDir() string {}
    fileutil.Exists(name string) bool {}
    fileutil.FilePath(dir, fileName, ext string) string {}

### stringutil

    import "github.com/hoahm/myutil/util/stringutil"

    stringutil.RandString(length int) string {}
    stringutil.Reverse(s string) string {}

## Data structures

### Linked list

    import "github.com/hoahm/myutil/datastructures/list"

    // Initialize an empty list
    l := list.New()
    
    // Check if list is empty
    l.IsEmpty()

    // Get the length of list
    l.Len()
    l.Size()

    // Insert element to the head
    l.PushFront(1)

    // Insert element to the tail
    l.PushBack(1)

    // Remove an element
    l.Remove(1)

    // Merge two lists
    l.Merge(l2)

### Queue

    import "github.com/hoahm/myutil/datastructures/queue"

    // Initialize empty queue
    q := queue.New()

    // Check if queue is empty
    q.IsEmpty()

    // Get the length of queue
    q.Len()
    q.Size()

    // Get head element
    q.Front()

    // Get and remove head element
    q.Pop()
    q.DeQueue()

    // Add new element to queue
    q.Push(1)
    q.EnQueue(1)

### Stack

    import "github.com/hoahm/myutil/datastructures/stack"

    // Initialize empty stack
    s := stack.New()    

    // Check if stack is empty
    s.IsEmpty()

    // Get the length of stack
    s.Len()
    s.Size()

    // Get peak element
    s.Top()

    // Get and remove peak element
    s.Pop()

    // Add new element to the peak
    s.Push(1)

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