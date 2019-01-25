//#1

  public class HelloWorld{
  public static void main(String []args){
        
  summ();
         
         
     }
  private static void summ(){
  int num =0;
        
  for(int i=3;i < 10000;i++){
    if(i % 3 == 0 || i % 5 == 0){
        num+=i;
                
        }
            
    }
  System.out.println(num);
        
 }
} 