using System;

public class HelloWorld
{
    public static void Main(string[] args)
    {
      var input = "programming";
      var checkInput = 'g';
     
      var result = CountOccurences(input, checkInput);
      Console.WriteLine(result);
    }
    
    static int CountOccurences(string input,char checkInput)
    {
        var count = 0;
       foreach (char c in input)
       {
           if (c == checkInput)
           {
               count++;
           }
       }
       
       return count;
    }
    
}
