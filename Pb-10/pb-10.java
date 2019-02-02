//#pb-10
public class HelloWorld{
  public static void main(String []args){
    long startTime = System.nanoTime();
    System.out.println(primeSummation(140759)); 
    long endTime = System.nanoTime();
    long timeElapsed = endTime - startTime;
    System.out.println("Execution time in microseconds : " + timeElapsed / 1000);
  }
   public static int primeSummation(int n){
    int sum=0;
    for (int i = 2; i < n; i++) {
      int count = 0;
      for (int j = 2; j <= i / 2; j++) {
        if (i % j == 0) {
          count++;
          break;
        }
      }
      if(count ==0){
  //System.out.println(i);
        sum+=i;
      }
    }   
        return sum;
  }
    
}