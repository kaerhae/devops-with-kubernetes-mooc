#/bin/bash

INPUT=$1
OUTPUT=$2

if [[ -z $AGE_FILE_PATH ]]; then
    echo "AGE_FILE_PATH environment variable is not set"
    exit 1
fi

if [[ -z $INPUT ]]; then
    echo "Please provide input file"
    exit 1
fi

if [[ -z $OUTPUT ]]; then
    echo "Please provide output file"
    exit 1
fi

# Command get second line of file, and deletes '# public key:' string from it
# Returns age string
AGE=$(sed '2q;d' $AGE_FILE_PATH | sed 's/# public key: //g')

sops --encrypt \
       --age $AGE \
       $INPUT > $OUTPUT