1. https://go.dev/doc/tutorial/getting-started
a. mkdir 
b. go mod init <some->
- what is go module? 
  - one project or library and contain a collection of Go packages that are then released together
- go mod tidy 
  - go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.
- go mod vendor
  - The go mod vendor command constructs a directory named vendor in the main module’s root directory that contains copies of all packages needed to support builds and tests of packages in the main module. 

  2. database
  install docker

  3. code copy-paste
    1. database connect
    2. http server
    3. breakpoints
    4. make tests
  
  4. make api calls
  