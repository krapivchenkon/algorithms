package p006_k_difference;

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


public class solution {

    public static int getPairs(int[] arr, int K) {
        int _cnt = 0;
        int _low, _high;
        boolean _low_found, _high_found;
        for (int i = 0; i < arr.length; i++) {
            _low = arr[i] - K;
            _high = arr[i] + K;
            _low_found = false;
            _high_found = false;
            for (int j = i+1; j < arr.length; j++) {
                if (_low_found && _high_found) break;
                            
                if (_low >= 0){
                    if (_low == arr[j]) {
                        _cnt++;         
                        _low_found = true;
                        continue;
                    }
                }else _low_found = true; 
                    
                if (_high == arr[j]) {
                    _cnt++; 
                    _high_found = true;
                    continue;
                }                       
            }
        }
        return _cnt;
    }
    
    
    

    public static void main(String[] args) {        

        Scanner in = new Scanner(System.in);
        
        int N = in.nextInt();int K = in.nextInt();              
        in.nextLine();

        int[] _a = new int[N];
        int _a_item;
        String next = in.nextLine();
        String[] next_split = next.split(" ");
        for(int _a_i = 0; _a_i < N; _a_i++) {
            _a_item = Integer.parseInt(next_split[_a_i]);
            _a[_a_i] = _a_item;
        }
        
        System.out.println(getPairs(_a, K));

    }
}
