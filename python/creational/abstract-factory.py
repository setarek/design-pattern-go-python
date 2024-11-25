from abc import ABC, abstractmethod

####################################### product interface 
class PaymentProcessor(ABC):
    @abstractmethod
    def process_payment(self, amount: float):
        pass

####################################### concrete products

class StripeCreditCardProcessor(PaymentProcessor):
    def process_payment(self, amount: float):
        print(f"Stripe: Processing credit card payment of ${amount:.2f}")


class StripeBankTransferProcessor(PaymentProcessor):
    def process_payment(self, amount: float):
        print(f"Stripe: Processing bank transfer of ${amount:.2f}")


class PayPalCreditCardProcessor(PaymentProcessor):
    def process_payment(self, amount: float):
        print(f"PayPal: Processing credit card payment of ${amount:.2f}")


class PayPalBankTransferProcessor(PaymentProcessor):
    def process_payment(self, amount: float):
        print(f"PayPal: Processing bank transfer of ${amount:.2f}")

    
####################################### abstrac factory interface

class PaymentFactory(ABC):
    @abstractmethod
    def create_credit_card_processor(self):
        pass

    @abstractmethod
    def create_bank_transfer_processor(self):
        pass

####################################### concrete factories

class StripeFactory(PaymentFactory):
    def create_credit_card_processor(self):
        return StripeCreditCardProcessor()

    def create_bank_transfer_processor(self):
        return StripeBankTransferProcessor()


class PayPalFactory(PaymentFactory):
    def create_credit_card_processor(self):
        return PayPalCreditCardProcessor()

    def create_bank_transfer_processor(self):
        return PayPalBankTransferProcessor()
    

####################################### client
def process_payments(factory: PaymentFactory, amount: float):
    credit_card_processor = factory.create_credit_card_processor()
    bank_transfer_processor = factory.create_bank_transfer_processor()

    credit_card_processor.process_payment(amount)
    bank_transfer_processor.process_payment(amount * 2)

if __name__ == "__main__":
    gateway = "stripe" 

    if gateway == "stripe":
        factory = StripeFactory()
    elif gateway == "paypal":
        factory = PayPalFactory()
    else:
        raise ValueError("Unsupported payment gateway")

    process_payments(factory, 100.00)
