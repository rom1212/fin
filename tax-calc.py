print 'tax calculator'

brackets = [
        (0.0,  0),
        (0.10,  18650),
        (0.15,  75900),
        (0.25, 153100),
        (0.28, 233350),
        (0.33, 416700),
        (0.35, 470700),
        (0.396, 10000000)]

print 'brackets:', brackets

income = 500000
print 'income:', income

total_tax = 0.0
rem = income
for index in range(1, len(brackets)):
    prev_limit = brackets[index-1][1]
    rate, limit = brackets[index]

    print '===================='
    print index, rate, limit, prev_limit

    taxable = (limit - prev_limit) if income > limit else (income - prev_limit)
    print 'taxble:', taxable, 'tax:', taxable * rate

    total_tax += taxable * rate

    if income < limit:
        print 'this is the last bracket for income:', income
        break

print 'total tax:', total_tax
