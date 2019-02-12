//#pb-12
public class HelloWorld{
  public static void main(String []args){
    long startTime = System.nanoTime();
    System.out.println(divisibleTriangleNumber(23));
    long endTime = System.nanoTime();
    long timeElapsed = endTime - startTime;
    System.out.println("Execution time in microseconds : " + timeElapsed / 1000);
  }
  public static int divisibleTriangleNumber(int n ){
    int i=1;
    int a=0;
    while(true){
      a=i*(i+1) /2 ;
      if(factor(a)>=n){
        break;
            
      }    
        i++;
    }
        return a;
  }
  public static int factor(int n){
    int count =0;
    for(int j=1;j<=n;j++){
      if(n % j ==0){
        count++;
      }
    }    
        return count;
    
  }
}





















