public class HelloWorld{

  public static void main(String []args){
        
  System.out.println(smallestMult(10));
       
  }
     
  public static int smallestMult(int n){
    int x=6;
    for(int i=4;i<=n;i++){
      x=lcm(x,i);
          
    }
      return x;  
  }
  public static int lcm(int n1,int n2){
  int gcd=0;   
    for(int i = 1; i <= n1 && i <= n2; ++i){
      if(n1 % i == 0 && n2 % i == 0){
        gcd = i;
      }
    }
    int lcm = (n1 * n2) / gcd;
    return lcm;
  }
}     
  