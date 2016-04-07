#!/usr/bin/env bash
# USAGE:
# run-tests.sh  - will run all tests against python solution
# run-tests.sh java test-1 - 
#
#
#
#
#
# 
LANG="py" # default language
PY_SOLUTION="solution.py"
JS_SOLUTION="solution.js"
JAVA_SOLUTION="solution.java"
JAVAC_SOLUTION="solution.class"
JAVA_CLASS="solution"
SUFIN=".in"
SUFOUT=".out"
SUFTMP=".tmp"
SCRNAME=$(basename $0)
TET=""
# TERMINAL settings
bold=$(tput bold)
sep="##############################################"
reset=$(tput sgr 0)
# offbold=`tput rmso`


declare -a TESTS=();


function usage(){
  # Print a usage statement and exit 1.
# TODO switch to use  cmdline debugger
  cat <<USAGE
   $SCRNAME - Tests problem solution against tests
  SYNOPSIS
   $SCRNAME [py|js|java] [options] [tests]
  OPTIONS
    -h--verbose--solution               show current information
-   -v|--verbose       enable solution debug
-   -s|--solution    if multiple solutions for specified language are present - concreate solution can be specified
-   -o|--output    Print failed tests input/output to stdout
-   -t|--test      specify test name to run with $SUFIN or without it

  PARAMETERS
   py|js|java     Run tests against solution in specified language
  USAGE
   Run all tests against python solution:
     > $SCRNAME

USAGE
  exit 1
}

if [[ $# -eq 0 ]]; then
    usage
fi

while [[ $# > 0 ]]
do
key="$1"

case $key in
    py|java|js)
    LANG=$key
    ;;
    -s|--searchpath)
    SEARCHPATH="$2"
    shift # past argument
    ;;
    -l|--lib)
    LIBPATH="$2"
    shift # past argument
    ;;
    -D|--debug)
    # option
    DEBUG="ON"
    # shift # past argument
    ;;

    --default)
    DEFAULT=YES
    DEBUG
    ;;
    *)
      usage 
            # unknown option
    ;;
esac
shift # past argument or value
done

# [M]* ) log_col 2 "   ""$line";; # green
# [C]* ) log_col 1 "   ""$line";; # red
# [D]* ) log_col 6 "   ""$line";; # cyan
# [A]* ) log_col 3 "   ""$line";; # yellow
# [I]* ) log_col 7 "   ""$line";; # white
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

collect_tests(){
    TESTS=($(ls *$SUFIN | sed 's/'"$SUFIN"'//g'))
}

run_python(){
    TET=$(cat $1$SUFIN |time -p python $PY_SOLUTION 2>&1>$1$SUFTMP)
}

run_java(){
    CPATH=$(basename $PWD).$JAVA_CLASS
    TET=$(cat $1$SUFIN |time -p java $CPATH 2>&1>$1$SUFTMP)
}

run_js(){
    echo "JS"
}

collect_tests
log "Run tests against $LANG solution"
for test in ${TESTS[@]}; do
    log_col 7 $sep
    log_bold 7 "Running [$test]"
    # log_col 7 $sep
    case $LANG in
        py)
        # TET=$(run_python $test)
        run_python $test
        ;;
        java)
        
        [ -f $JAVAC_SOLUTION ] || { log_bold 3 "File should be compiled first"; javac $JAVA_SOLUTION; }
        run_java $test
        ;;
        js)
        run_js
        ;;
        *)
        ;;
    esac
    log "TET: [${TET}]"
    # check if result is correct
    diff_out=$(diff -ub $test$SUFOUT $test$SUFTMP)
    if [ "$?" -ne "0" ]; then
        log_bold 1 "$test FAILED"
    # as a bonus, make our script exit with the right error code.
    else
        log_bold 2 "$test PASSED"
    fi
    # log $(echo $TET| tr -d '\n')
done;
log $sep