// Copyright 2013 Google Inc. All Rights Reserved.
// Copyright 2015 Cockroach Labs.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code originated in the github.com/golang/glog package.

// Package log implements logging.
// There are three logging styles: named, V-style, events.
//
// Named Functions
//
// The functions Info, Warning, Error, and Fatal log their arguments at the
// specified level. All include formatting variants like Infof.
//
// Examples:
//
//	log.Info(ctx, "Prepare to repel boarders")
//	log.Fatal(ctx, "Initialization failed", err)
//	log.Infof(ctx, "client error: %s", err)
//
// V-Style
//
// The V functions can be used to selectively enable logging at a call
// site. Invoking the binary with --verbosity=N will enable V functions
// at level N or higher. Invoking the binary with --vmodule="glob=N" will
// enable V functions at level N or higher with a filename matching glob.
//
// Examples:
//
//	if log.V(2) {
//		log.Info(ctx, "Starting transaction...")
//	}
//
// Events
//
// The Event functions log messages to an existing trace if one exists. The
// VEvent functions logs the message to a trace and also the log file based
// on the V level.
//
// Examples:
//
//	log.VEventf(ctx, 2, "client error; %s", err)
//
// Output
//
// Log output is buffered and written periodically using Flush. Programs
// should call Flush before exiting to guarantee all log output is written.
//
// By default, all log statements write to files in a temporary directory.
// This package provides several flags that modify this behavior.
// These are provided via the util/log/logflags package; see InitFlags.
//
//  --logtostderr=LEVEL
//    Logs are written to standard error as well as to files.
//    Entries with severity below LEVEL are not written to stderr.
//    "true" and "false" are also supported (everything / nothing).
//  --log-dir="..."
//    Log files will be written to this directory instead of the
//    default target directory.
//  --log-file-verbosity=LEVEL
//    Entries with severity below LEVEL are not written to the log file.
//    "true" and "false" are also supported (everything / nothing).
//  --log-file-max-size=N
//    Log files are rotated after reaching that size.
//  --log-dir-max-size=N
//    Log files are removed after log directory reaches that size.
//
// Other flags provide aids to debugging.
//
//  --log-backtrace-at=""
//    When set to a file and line number holding a logging statement,
//    such as
//      -log_backtrace_at=gopherflakes.go:234
//    a stack trace will be written to the Info log whenever execution
//    hits that statement. (Unlike with --vmodule, the ".go" must be
//    present.)
//  --verbosity=0
//    Enable V-leveled logging at the specified level.
//  --vmodule=""
//    The syntax of the argument is a comma-separated list of pattern=N,
//    where pattern is a literal file name (minus the ".go" suffix) or
//    "glob" pattern and N is a V level. For instance,
//      --vmodule=gopher*=3
//    sets the V level to 3 in all Go files whose names begin "gopher".
//
// Protobuf
//
// Autogenerated:
//
package log
