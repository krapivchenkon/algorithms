package p002_chocolate_feast;

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


public class solution {

     public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int t = in.nextInt();
        for(int i = 0; i < t; i++){
            System.out.println(Solve(in.nextInt(), in.nextInt(), in.nextInt()));
        }
    }
    
    private static long Solve(int N, int C, int M){
        int _total = N/C;
        int _wraps = N/C;
        while (_wraps >= M) {
            _total += _wraps/M;
            _wraps = _wraps/M + _wraps%M;
        }

        return _total;
    }
    
    
}
