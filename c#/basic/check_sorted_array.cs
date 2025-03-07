using System;

public class HelloWorld
{
    public static void Main(string[] args)
    {
       int[] input = new int[]{2,8,5,6};
     
      if (checkIsSorted(input))
      {
          Console.WriteLine("Sorted");
      }
      else
      {
          Console.WriteLine("Not Sorted");
      }
    }
    
    static bool checkIsSorted(int[] input)
    {
        var firstValue = input[0];
       
        
        for (int i=1;i<input.Length;i++)
        {
            if (input[i] < input[i-1])
            {
                return false;
            }
            
        }
        
        return true;
    }
    
}
