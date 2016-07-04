govet
=====

Single main file that invokes `go vet` and passes through any arguments that were provided. stdout and stderr are passed
through directly. Exits with an exit code of 1 if underlying command exits with a non-0 exit code or causes any other
errors.

This project exists so that `govet` can be treated like other Go code analysis tools that are standalone executables.
