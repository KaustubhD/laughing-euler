//#1

  public class HelloWorld{
    public static void main(String []args){
        
      System.out.println(summ(19564));
         
         
    }
    private static long summ(long n){
      long num =0;
        
        for(long i=3;i < n;i++){
          if(i % 3 == 0 || i % 5 == 0){
            num+=i;
                
        }
            
        }
           return num;
        
    }
  } 