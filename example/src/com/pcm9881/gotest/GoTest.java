package com.pcm9881.gotest;

public class GoTest {

  public static void main(String args[]) {
    printString("using java library in golang");
    
    String text = "Hello";
    String subText = "GoTest";
    printString(text, subText);
  }
  
  public static String printString(String text) {
    System.out.println(text);
    return text;
  }
  
  public static String printString(String text, String subText) {
    String result = text + " " +subText;
    System.out.println(result);
    return result;
  }
}
