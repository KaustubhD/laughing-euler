//pb-14
public class Main{
  public static void main(String []args){
    long startTime = System.nanoTime();
    System.out.println(longestCollatzSequence(14));
    long endTime = System.nanoTime();
    long timeElapsed = endTime - startTime;
    System.out.println("Execution time in microseconds : " + timeElapsed / 1000);
  }
  static int count=0;
  static int max=0;
  static int  maxint =0;
  public static int longestCollatzSequence(int n){
    for(int i=1;i<=n;i++){
      count=0;
      eve(i);
      if(count> max){
        max=count;
        maxint=i;
      }
    }
    return maxint;
  }
  public static void eve(int n){ 
    while(n!=1){
      if(n % 2 ==0){
      n =n/2;                
      }
      else{
        n=3*n+1;
      }
      count++;
    }
  }
}