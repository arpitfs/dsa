using System;

public class HelloWorld
{
    public static void Main(string[] args)
    {
       int[] input = new int[]{10, 23, 45, 67, 89};
       int number = 89;
       
       int left = 0;
       int right = input.Length - 1;
       
       
       while (left <= right)
       {
           int mid = (left + right)/2;
           if (input[mid] == number)
           {
               Console.WriteLine(mid);
               break;
           }
           else if (number > input[mid] )
           {
               left = mid + 1;
           }
           else 
           {
               right = mid - 1;
           }
       }
    }
    
}
