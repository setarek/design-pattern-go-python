import copy

# prototype interface
class DocumentPrototype:
    def clone(self):
        return copy.deepcopy(self)

# concerete prototype implementation
class Invoive(DocumentPrototype):
    def __init__(self, item, price, amount):
        self.type = "Invoice"
        self.item = item
        self.price = price
        self.amount = amount
    
    def __str__(self):
        return f"{self.type} - Amount: ${self.amount}"
    
class Report(DocumentPrototype):
    def __init__(self, id, total_price):
        self.type = "Report"
        self.id = id
        self.total_price = total_price

    def __str__(self):
        return f"{self.type} - Id: {self.id}, Total Price: {self.total_price}"
    
if __name__ == "__main__":

    invoice = Invoive("default", 0, 0)
    report = Report(0, 0)

    new_invoice = invoice.clone()
    new_invoice.item = "PR02"
    new_invoice.amount = 25

    print(new_invoice)

    new_report = report.clone()
    new_report.id = 2
    new_report.total_price = 5000000

    print(new_report)



    
        
        