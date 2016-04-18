package p003_cut_the_sticks;

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


public class solution {

	public static int getMin(int[] arr) {
		int _min = Integer.MAX_VALUE;
		for (int i = 0; i < arr.length; i++) {
			if (arr[i]<_min && arr[i]!=0) _min = arr[i];			
		}
		return _min;
	}
	public static void cutTheSticks(int[] arr) {
		int _min = getMin(arr);
		int _cnt = 0;
		while (_min != Integer.MAX_VALUE) {
//			System.out.println("arr:"+Arrays.toString(arr));
			_cnt = 0;
			for (int i = 0; i < arr.length; i++) {
				if (arr[i] > 0) {
					arr[i] -= _min;
					_cnt++;
				}
			}
			System.out.println(_cnt);
			_min = getMin(arr);			
		}
		
	}

       public static void main(String[] args) {
            Scanner in = new Scanner(System.in);
    		int _a_size = Integer.parseInt(in.nextLine());
    		// read to array from space separated string
    		int[] _a = new int[_a_size];
    		int _a_item;
    		String next = in.nextLine();
    		String[] next_split = next.split(" ");
    		for(int _a_i = 0; _a_i < _a_size; _a_i++) {
    			_a_item = Integer.parseInt(next_split[_a_i]);
    			_a[_a_i] = _a_item;
    		}
    		cutTheSticks(_a);
    }
}
