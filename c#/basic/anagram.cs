using System;

public class HelloWorld
{
    public static void Main(string[] args)
    {
        string input1 = "lisn";
        string input2 = "siln";
        
        if(checkAnagram(input1, input2))
        {
            Console.WriteLine("Anagram");
        }
        else
        {
            Console.WriteLine("Not an anagram");
        }
    }
    
    static bool checkAnagram(string input1 ,  string input2)
{
        int[] result = new int[256];
        input1 = input1.Replace(" ", "").ToLower();
        input2 = input2.Replace(" ", "").ToLower();
        
        if (input1.Length != input2.Length)
        {
            return false;
        }
        
        foreach (char c in input1)
        {
            result[c]++;
        }
        
        foreach (char c in input2)
        {
            result[c]--;
        }
        
        foreach (var c in result)
        {
            if (c!=0)
            {
                return false;
            }
        }
        return true;
}
}

