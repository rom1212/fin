sp_dec = [ (2017,       21.83),
       (2016,   11.96),
       (2015,   1.38),
       (2014,   13.69),
       (2013,   32.39),
       (2012,   16.00),
       (2011,   2.11),
       (2010,   15.06),
       (2009,   26.46),
       (2008,   -37.00),
       (2007,   5.49),
       (2006,   15.79),
       (2005,   4.91),
       (2004,   10.88),
       (2003,   28.68),
       (2002,   -22.10),
       (2001,   -11.89),
       (2000,   -9.10),
       (1999,   21.04),
       (1998,   28.58),
       (1997,   33.36),
       (1996,   22.96),
       (1995,   37.58),
       (1994,   1.32),
       (1993,   10.08),
       (1992,   7.62),
       (1991,   30.47),
       (1990,   -3.10),
       (1989,   31.69),
       (1988,   16.61), ]

sp_inc = list(reversed(sp_dec))
print [x for x in sp_inc]
print 'len(sp_inc):', len(sp_inc)

start = 100000

normal = start
for r in sp_inc:
    rate = r[1] / 100
    print 'rate:', rate
    tax = 0
    if rate > 0:
        tax = rate * normal * 0.00
    normal = (1 + rate) * normal - tax

print "normal:", normal, ", after tax: ", normal * 0.80

cap = 0.15
bot = 0.007
box = start
for r in sp_inc:
    rate = r[1] / 100
    real_rate = rate
    if rate < 0:
        real_rate = bot
    elif rate > cap:
        real_rate = cap

    print 'rate:', rate, ', real rate:', real_rate
    box = (1 + real_rate) * box
print 'box:', box
