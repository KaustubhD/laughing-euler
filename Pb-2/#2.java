public class HelloWorld{

  public static void main(String []args){
    
    System.out.println(fiboEvenSum(23));  
  }

  public static int fiboEvenSum(int n){
    int a=0;
    int b=1;
    int c=0;
    int sum=0;
    while(c < n){
      c=a+b;
      a=b;
      b=c;
         
      if(c % 2 == 0){
        sum+=c;
      }  
           
    }
    return sum;
  }
}
