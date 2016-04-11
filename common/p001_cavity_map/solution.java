package p001_cavity_map;

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


public class solution {

	   public static void main(String[] args) {
	        Scanner in = new Scanner(System.in);
	        int n = in.nextInt();
	        String grid[] = new String[n];
	        for(int grid_i=0; grid_i < n; grid_i++){
	            grid[grid_i] = in.next();
	        }
		printCavityMap(grid);
	}

	public static void printCavityMap(String[] map) {
		String line;
		System.out.println(map[0]);
		if (map.length>1) {								
			for (int i = 1; i < map.length - 1; i++) {
				line = map[i];

				for (int j = 1; j < map.length - 1; j++) {
					if (map[i].charAt(j) > map[i - 1].charAt(j)
							&& map[i].charAt(j) > map[i + 1].charAt(j)
							&& map[i].charAt(j) > map[i].charAt(j - 1)
							&& map[i].charAt(j) > map[i].charAt(j + 1)) {
						line = line.substring(0, j) + 'X'
								+ line.substring(j + 1, line.length());
					}
				}
				System.out.println(line);
			}
			System.out.println(map[map.length - 1]);
		}
	}
}