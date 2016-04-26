package p005_even_tree;

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;


public class solution {

    public static int edges_to_del = 0;

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int V = in.nextInt();int E = in.nextInt();
        int[][] adj = new int[V][V];
        for (int i = 0; i < E; i++) {
            setEdge(adj, in.nextInt(), in.nextInt());
        }
        getNumOfChilds(adj, 1, 0);
        System.out.println(edges_to_del);
    }

    public static int getNumOfChilds(int[][] matr,int node, int par) {
        int sublings = 1;
        List<Integer> childs = getChilds(matr, node, par); 
        if (childs != null) {
            for (Integer child : childs) {
                sublings += getNumOfChilds(matr, child, node);
            }
        }
        if (sublings%2 == 0 && (node !=0 && par !=0)) {
            delEdge(matr, node, par);
            incDel();

            sublings = 0;
        }
        return sublings;
    }
    
    public static void incDel() {
        edges_to_del++;
    }
    public static List<Integer> getChilds(int[][] matr, int node,int par) {
        ArrayList<Integer> arr = null;
        for (int i = 0; i < matr.length; i++) {
            if (matr[node-1][i] == 1 && i+1 != par){
                if (arr == null) {
                    arr = new ArrayList<Integer>();
                } 
                arr.add(i+1);
            }
        }
        return arr;
    }


    public static void setEdge(int[][] adj, int x, int y) {
        adj[x-1][y-1] = 1;
        adj[y-1][x-1] = 1;
    }

    public static void delEdge(int[][] adj, int x, int y) {
        adj[x-1][y-1] = 0;
        adj[y-1][x-1] = 0;
    }

}
