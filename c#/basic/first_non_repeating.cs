// Online C# Editor for free
// Write, Edit and Run your C# code using C# Online Compiler

using System;
using System.Collections.Generic;

public class HelloWorld
{
	public static void Main(string[] args)
	{
		string input = "leetcode";
		var result = FirstNonRepeatingCharacter(input);
		Console.WriteLine(result);
	}

	static char FirstNonRepeatingCharacter(string input)
	{
		Dictionary<char, int> tracker = new Dictionary<char, int>();
		foreach (char i in input)
		{
			if (tracker.ContainsKey(i))
			{
				tracker[i] = tracker[i]+1;
			}
			else
			{
				tracker[i] = 1;
			}
		}

		foreach (char c in input)
		{
			if (tracker[c] == 1)
			{
				return c;
			}
		}

		return ' ';
	}
}

