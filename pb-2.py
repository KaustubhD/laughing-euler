def fiboEvenSum(n):
  n1=0
  n2=1
  count=0
  sum=0
  while count<n:
      count=n1+n2
      n1=n2
      n2=count
      if count%2==0:
        sum+=count
  return sum

if __name__== "__main__":
  print(fiboEvenSum(10))
