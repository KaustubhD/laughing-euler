-from datetime import datetime
def multiplesof3and5(n):
  sum=0
  for i in range(n):
    if i%3==0 or i%5==0:
      sum+=i
  return sum
        
        
if __name__== "__main__":
  start_time = datetime.now()

  print(multiplesof3and5(10))
	end_time = datetime.now()
	print('Duration in microseconds: {}'.format(end_time - start_time))
