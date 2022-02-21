#!/bin/bash

# NOTE!!! Running this will open 6 seperate vscode windows

# chmod u+x ./test.sh

EXAMPLES=(
  'fiber_no_modules' 
  'fiber_with_modules' 
  'fiber_with_modules_and_sqlc' 
  'gin_no_modules' 
  'gin_with_modules' 
  'gin_with_modules_and_sqlc'
)

CURRENT_DIR=$PWD

for example in "${EXAMPLES[@]}"; do
  # echo "example: $example"

  echo "----- Creating $example!"

  BASE=./examples/$example/    
  GOOS=darwin GOARCH=arm64 go build -o $BASE/gomakeme main.go   # build binary to the example dir

  cd $BASE            # cd into the dir that holds the binary + gomakeme config
  rm -rf ./my_server  # delete the dir, so that this works if debug_mode = true
  ./gomakeme          # exec the built binary
  cd my_server        # cd into the generated dir
  go mod tidy         # clean up
  go mod download     # download dependencies
  gofmt -s -w .       # format generated code
  code .              # open in vscode to check if there are errors 
  # cd ../../..       # go back the dir from which the script was run
  # echo $PWD
  cd $CURRENT_DIR     # go back the dir from which the script was run

done


# build for the local example
: '

go run main.go
cd change_my_name   # cd into the generated dir

go mod tidy         # clean up
go mod download     # download dependencies
gofmt -s -w .       # format generated code
code .              # open in vscode to check if there are errors 

# build binary
GOOS=darwin GOARCH=arm64 go build -o gomakeme main.go
'