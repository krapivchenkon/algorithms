#!/usr/bin/env bash


PROBLEM=""
SOLUTION="solution"
READMENAME="README.md"
TEMPLATE_DIR="/.lib"
REPO_DIR=$(dirname `which $0` |sed 's/\/[^\/]*$//g')
SCRNAME=$(basename $0)
declare -a LANGT

function usage(){
  # Print a usage statement and exit 1.
# TODO switch to use  cmdline debugger
cat <<USAGE
$SCRNAME - Creates directory with templates for problem solution
SYNOPSIS
$SCRNAME [l] -n PROBLEMNAME 
OPTIONS
-n|--name               - name -f the problem [a-zA-Z0-9] and [_] are allowed 
-l|--lang [py|java|js]  - specify language which will be used for solution
Multiple -l options can be specified
-   Create template dir for solving problem with python:
> $SCRNAME -l py -n 001_multiplies_of_3_and_5
will create directory with templates for py solutions
-   Add Java solution template for current problem
> $SCRNAME -l java .

USAGE
exit 1
}

if [[ $# -eq 0 ]]; then
    usage
fi
############################
# README.md template
############################
read -d '' readme <<- EOF
# Problem Name
---

### Input Format 

### Output Format 

### Constraints

### Sample Input
```
```
### Sample Output
```
```
### Explanation

EOF
############################
# log helpers
############################
bold=$(tput bold)   
function log(){
  echo "$SCRNAME:`date +'%d%m %H%M%S'`:I:"$1
}
function log_col(){
  echo "$SCRNAME:`date +'%d%m %H%M%S'`:I:"$(tput setaf $1)$2"$(tput sgr 0)" 
}
function log_emph(){
  echo "$SCRNAME:`date +'%d%m %H%M%S'`:I:$(tput setaf 6)"$1"$(tput sgr 0)"
}
function log_err(){
  echo "$SCRNAME:`date +'%d%m %H%M%S'`$(tput setaf 1):E:"$1"$(tput sgr 0)"
}
function log_warn(){
  echo "$SCRNAME:`date +'%d%m %H%M%S'`$(tput setaf 1)$(tput setab 7):W:"$1"$(tput sgr 0)"
}
function log_bold(){
 echo "$SCRNAME:`date +'%d%m %H%M%S'`:I:"${bold}$(tput setaf $1)$2"$(tput sgr 0)" 
}
############################
# checking args
############################
while [[ $# > 0 ]]
do
    key="$1"
    case $key in
        -l|--lang) 
LANGT=( "${LANGT[@]}" "$2" );
    shift # past argument
    ;;
    -n|--name) 
PROBLEM=$2
    shift # past argument
    ;;
    *)
usage
    # unknown option
    ;;
esac
shift # past argument or value
done

function contains() {
    local value=$1
    for ((i=0;i < ${#LANGT[@]};i++)) {
        if [ "${LANGT[i]}" == ${value} ]; then
            return 0
        fi
    }
    return 1
}

function copy_templates(){ # 
for tmpl in $(find $REPO_DIR -name "template*");
do
    filename=$(basename "$tmpl")
    EXT="${filename##*.}"
    # eval "echo \"${template_str}\""
    # [[ $? -eq 0 ]] && { log "Copying [$EXT] template"; export PRM="WEFWEFW"; cat $tmpl >$PROBLEM/$SOLUTION.$EXT; }
    contains $EXT
    [[ $? -eq 0 ]] && { 
        template_str=$(cat "${tmpl}")
        log "Copying [$EXT] template";  
        eval "echo \"${template_str}\" >$PROBLEM/$SOLUTION.$EXT"; 
    }
done;


}

############################
# MAIN
############################
# if directory exists than just add template for specified language
if [ -d "$PROBLEM" ]; then
    # if solution for specified language exists - add new file with incremental number
    log_bold 7 "directory: $PROBLEM already exists. Copying specified templates only"
    copy_templates
else
    # create new dir
    log_bold 7 "Creating new directory: $PROBLEM"
    mkdir $PROBLEM
    # Set Vars to be used in Solution templates

    # copy template for specified language
    copy_templates

    echo "$readme" >$PROBLEM/$READMENAME
    # TODO create couple empty test files
    # Thingsto do
    # TODO request if user wants to create empty tests for fill them from stdin
fi;


