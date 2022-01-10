import scala.io.StdIn.readLine
import scala.math.sqrt

object Factors {
	def main(args: Array[String]) = {
		print("Enter some number n: ");
		var n = readLine().toInt
		while (n % 2 == 0) {
			print("2 ");
			n /= 2;
		}
		
		var i = 3;
		while (i <= sqrt(n)) {
			while (n % i == 0) {
				print(i + " ");
				n /= i;
			}
			i += 2;
		}
		
		if (n > 1) print(n + " ");
		println();
	}

}
