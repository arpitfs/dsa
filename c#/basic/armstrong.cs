using System;

public class HelloWorld
{
	public static void Main(string[] args)
	{
		int input1 = 153;
		
		if (isArmStrong(input1))
		{
			Console.WriteLine("Yes Armstrong");
		}
		else
		{
			Console.WriteLine("Not Armstrong");
		}
	}

	static bool isArmStrong(int input1)
	{
		int result = 0;
		var length = input1.ToString().Length;
		var originalNumber = input1;

		while(input1 > 0)
		{
			var lastDigit = input1 % 10;
			result = result + (int)Math.Pow(lastDigit, length);
			input1 = input1/10;
		}

		return result == originalNumber;
	}
}

