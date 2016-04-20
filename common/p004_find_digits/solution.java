package p004_find_digits;

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


public class solution {

	public static int getDividers(long N) {
		HashMap<Integer, Integer> map = new HashMap<Integer, Integer>(9);
		for (int i = 1; i < 10; i++) {
			map.put(i, 0);
		}
		int _res = 0;
		String long_str = Long.toString(N);
		char digit;
		int int_digit;
		for (int i = 0; i < long_str.length(); i++) {
			digit =long_str.charAt(i);
			int_digit = Character.getNumericValue(digit);
			if (digit == '1'){
				map.put(1, map.get(1)+1);
			}else if (digit != '0' && N%int_digit==0 ){
				map.put(int_digit, map.get(int_digit)+1);
			}
		}
		for (Integer sum: map.values()) {
			_res += sum;
		}
		return _res;
	}
	public static void main(String[] args) {
		Scanner in = new Scanner(System.in);
		
		int T = Integer.parseInt(in.nextLine());
		for (int i = 0; i < T; i++) {
			long N = Long.parseLong(in.nextLine());
			System.out.println(getDividers(N));
		}
	}

}
