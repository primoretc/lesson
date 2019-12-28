#Напишите три функции, которые вычисляют сумму чисел в данном списке, используя цикл 
# for, цикл while и рекурсию
#
sp = [12,34,543,676,778] # для проверки, сумма равна 2043
sum=0
i=0

def sumFor(sp):
    sum=0
    for i in sp:
        sum=sum+i
    return sum
       
def sumWhile(sp):
    sum=0
    i=0
    while i<len(sp):
        sum=sum+sp[i]
        i=i+1
    return sum    

i=0
sum=0


def sumRek (sp):
    global i
    global sum
    
      


    while i <= len(sp):
        if i <= len(sp):
            sum=sum+sp[i]
            i+=1   
            print("sum=",sum)
            sumRek(sp)
      
    return sum
    

print(sumFor(sp))
print(sumWhile(sp))
print(sumRek(sp))