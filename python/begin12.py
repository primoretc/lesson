# Даны каеты прямоугольного треугольника a b 
# найти его гипотенузу c и периметр
import math
a = float(input("Введите катет a "))
b = float(input("Введите катет b "))
c=math.sqrt(math.pow(a,2)+math.pow(b,2))
print("c=", c)
print("P=", a+b+c)