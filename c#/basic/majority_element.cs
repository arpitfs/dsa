using System;
using System.Collections.Generic;

public class HelloWorld
{
	public static void Main(string[] args)
	{
		int[] input = new int[]{3, 3, 4, 2, 4, 4, 2, 4, 4};
		var result = MajorityElement(input);
		Console.WriteLine(result.Item1);
		Console.WriteLine(result.Item2);
	}

	static (int, int) MajorityElement(int[] input)
	{
		Dictionary<int, int> tracker = new Dictionary<int, int>();
		foreach (int i in input)
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
		
		var max = 0;
		var element = 0;

		foreach (var c in tracker)
		{
			if (c.Value > max)
			{
				max = c.Value;
				element = c.Key;
			}
		}

		return (max, element);
	}
}