public class HelloWorld{

  public static void main(String []args){
        
    System.out.println(largestPalindromeProduct(2));  
  }

  public static double largestPalindromeProduct(int n){

    double max=0;
    
    double b= (Math.pow(10,n)-1);
    for(int i=(int)b;i>=1;i--){
        
      for(int j=i;j>=1;j--){
        n =i*j;
	if(n>max){

          String c=String.valueOf(n);
          StringBuilder c1 = new StringBuilder(); 
          c1.append(c); 
          c1 = c1.reverse();
          if(c.equals(c1.toString())){

	    max=n;

          }
	}      
      }
    }     
    return max; 
  }
}