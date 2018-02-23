def total_months(debt, apr, monthly_payment):
    month_num = 0
    while debt > 0:
        month_num += 1

        interest = debt * apr / 12
        principal = monthly_payment - interest
        new_debt = debt - principal

        print "year:", month_num/12, ", month:", month_num % 12, ", month_num", month_num, ", new debt:", new_debt

        debt = new_debt

    return month_num

def total_months_bimonthly(debt, apr, monthly_payment):
    month_num = 0
    while debt > 0:
        month_num += 1

        interest = debt * apr / 24
        principal = monthly_payment/2 - interest
        new_debt = debt - principal

        print "year:", month_num/24, ", month:", month_num % 24 / 2, ", month_num", month_num/2, ", new debt:", new_debt

        debt = new_debt

    return month_num/2

def total_months_heloc(debt, apr, monthly_payment):
    month_num = 0
    while debt > 0:
        month_num += 1

        interest = (debt - 10000) * apr / 12
        principal = monthly_payment - interest
        new_debt = debt - principal

        print "year:", month_num/12, ", month:", month_num % 12, ", month_num", month_num, ", new debt:", new_debt

        debt = new_debt

    return month_num
