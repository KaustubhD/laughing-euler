public class HelloWorld{

     public static void main(String []args){
        
   System.out.println(largestPrimeFactor(2));  
  }

  public static double largestPrimeFactor(double n){

    double reverse=0;
    
    double b= (Math.pow(10,n)-1);
    for(int i=1;i<=b;i++){
        
        //System.out.println(i);
        for(int j=1;j<=i;j++){
            n =i*j;
            
            String c=String.valueOf(n);
             StringBuilder c1 = new StringBuilder(); 
             c1.append(c); 
             c1 = c1.reverse();
             if(c==c1.toString()){
                 System.out.println(c1);
             }
             
        
            
        }
        
    }     
        return n; 
  }
}
    
    

s