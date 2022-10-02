from pb.my.my_pb2 import *

report = SellerParams()

first_seller = report.result.add()
first_seller.seller_id = 1
first_seller.rating = 11
first_seller.params["some_double"].double_name = 4.2

second_seller = report.result.add()
second_seller.seller_id = 2
second_seller.params["name"].string_name = 'David'

print('Report: ')
print(report.SerializeToString())

with open('../output.bin', 'wb') as f:
    f.write(report.SerializeToString())



