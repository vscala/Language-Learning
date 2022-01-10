import scala.io.StdIn.readLine

object Fib {
    def main(args: Array[String]) = {
    	print("Enter some value n: ")
    	val n = readLine().toInt
    	var a = 0;
    	var b = 1;
    	var i = 0;
    	
    	for (i <- 1 to n) {
    		b += a
    		a = b - a
    	}
    	
    	
        println("The nth fib number is: " + a)
    }
}

