# algorithms
Solutions for some algorithms

add result of folowing command to your bash profile:
`echo "export PATH=$PWD/.helper:\$PATH"`

### Test solution against tests:
`$ run-tests.sh js|py|java`

### Create template directory for new problem
- **For a single lang:**

   `$ create-dir.sh -n PROBLEM_NAME -l py`
- **For multiple langs:**  

   `$ create-dir.sh -n PROBLEM_NAME -l py -l js`

Note that problem shouldn't contain illegal symbols for java packages("-") and should start with [a-zA-Z]