def sumsquarediff(n):
  a=0
  b=0
  a= (n*(n+1))/2
  a=a*a
  b=(n*(n+1)*(2*n+1))/6
  c=a-b
  return c
print(sumsquarediff(10))
