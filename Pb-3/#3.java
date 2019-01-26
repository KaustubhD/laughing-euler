public class HelloWorld{
  public static void main(String []args){

    System.out.println(largestPrimeFactor(13195));  
  }

  public static long largestPrimeFactor(long n){
    
    for(long i=2;i<n;++i){
      if(n % i ==0){
         n = n/i;
         i=1;
      }
    }    
     return n;
  }
}