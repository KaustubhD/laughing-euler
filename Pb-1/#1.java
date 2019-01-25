//#1

  public class HelloWorld{
  public static void main(String []args){
        
  System.out.println(summ(19564));
         
         
     }
  private static int summ(int n){
  int num =0;
        
  for(int i=3;i < n;i++){
    if(i % 3 == 0 || i % 5 == 0){
        num+=i;
                
        }
            
    }
      return num;
        
 }
} 