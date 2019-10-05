//#1

class HelloWorld{
  public static void main(String []args){

    System.out.println(multiplesOf3and5(19564));  

  }
  public static long multiplesOf3and5(long n){
    long num=0;

      for(long i=3;i<n;i++){
        if(i % 3 == 0 || i % 5 == 0){
          num+=i;

        }

      }
      return num;
  }
} 
