from datetime import datetime
start_time = datetime.now()
def evenFibSum(limit) : 
    if (limit < 2) : 
        return 0
  
    # Assigning value to first two members of series
    num1 = 0
    num2 = 2
    sm= num1 + num2 
      
    # calculating sum of even Fibonacci value 
    while (num2 <= limit) : 
  
        # formula for next even value of Fibonacci  
         
        num3 = 4 * num2 + num1
  
        # loop break if we num3 is greter than limit
        if (num3 > limit) : 
            break
  
        # Move to next even number and update 
        # sum 
        num1 = num2 
        num2 = num3 
        sm = sm + num2
      
    return sm 
  
print(evenFibSum(23))
end_time = datetime.now()
print('Duration: {}'.format(end_time - start_time))
  
