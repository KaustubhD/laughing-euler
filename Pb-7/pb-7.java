public class HelloWorld{
  public static void main(String []args){
    //long startTime = System.nanoTime();
    System.out.println(nthPrime(6));      
  	//long endTime = System.nanoTime();
  	//long timeElapsed = endTime - startTime;
		//System.out.println("Execution time in microseconds : " + timeElapsed / 1000);
  }

  public static int nthPrime(int n) {
    int i=0; int x = 1;
    while(i < n){
      x++;
      if(abc(x)){
        i++;
      }
    }
        return x;
  } 
  public static boolean abc(long n) {
    if (n < 2) return false;
      else if (n == 2) return true;
        for (int i = 2; i < n/2 + 1; i++)
          if (n % i == 0)
            return false;
        return true;
    }
}