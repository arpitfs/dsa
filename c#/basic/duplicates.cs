using System;
using System.Collections.Generic;

public class HelloWorld
{
	public static void Main(string[] args)
	{
		int[] input = new int[] {4, 3, 2, 7, 8, 2, 3, 1};

		var result = FindDuplicates(input);
		
		foreach (var r in result)
		{
		    Console.WriteLine(r);
		}
	}

	static List<int> FindDuplicates(int[] input)
	{
		Dictionary<int, bool> tracker = new Dictionary<int, bool>();
		List<int> result = new List<int>();
		foreach (int i in input)
		{
			if (!tracker.ContainsKey(i))
			{
				tracker[i]=true;
			}
			else
			{
				result.Add(i);
			}
		}
		
		return result;
	}
}

