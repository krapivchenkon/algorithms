#!/usr/bin/env python
def solution():
    t = int(raw_input())
    for test in xrange(t):
        N =int(raw_input())
        N3 = (N-1)//3
        N5 = (N-1)//5    
        N15 = (N-1)//15
        print 3*N3*(N3+1)/2+5*N5*(N5+1)/2-15*N15*(N15+1)/2
            

if __name__ == '__main__':
    solution()