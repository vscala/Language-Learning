import scala.collection.mutable.HashMap

class Graph() {
	
	var nodes = new HashMap[Int, List[Int]].withDefaultValue(Nil)
	
	// Undirected
	def addEdge(a: Int, b: Int) {
		nodes(a) ::= b
		nodes(b) ::= a
	}
	
	def print() {
		nodes.foreach {  
            case (key, value) => println (key + " -> " + value)         
        } 
	}
}

object test {
	def main(args: Array[String]) = {
		val graph = new Graph();
		graph.addEdge(1, 2);
		graph.addEdge(5, 2);
		graph.addEdge(1, 7);
		graph.print()
	}
}
