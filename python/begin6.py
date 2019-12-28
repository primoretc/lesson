#
#найти длину окружности L и площадь круга S заданного радиуса R
#
import math
R = float(input("Введите радиус круга: "))
while R==0 :
    print("Вы ввели ноль!!!")
    R=float(input("Введите радиус круга: "))
else:
    L= 2*math.pi*R
    print ("Длина окружности равна ",L)
    S=math.pi*math.pow(R,2)
    print("Площадь круга равна ", S)
print("Досвидос!")       