public class HelloWorld{

  public static void main(String []args){
    long startTime = System.nanoTime();
    System.out.println(sumSquareDifference(20));      
  	long endTime = System.nanoTime();
  	long timeElapsed = endTime - startTime;
		System.out.println("Execution time in microseconds : " + timeElapsed / 1000);
  }

  public static int sumSquareDifference(int n) {
    int a= n*(n+1)/2;
    a=a*a;
    int b=n*(n+1)*(2*n+1)/6;
    int c=a-b;
    return c;
    
  } 
}