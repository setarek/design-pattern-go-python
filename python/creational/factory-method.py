from abc import ABC, abstractmethod

# Product interface
class Notification(ABC):

    @abstractmethod
    def send(self):
        pass

# Product implementation
class SMSNotification(Notification):
    def send(self, message):
        print(f"Send SMS: {message}")

class EmailNotitication(Notification):
    def send(self, message):
        print(f"Send Email: {message}")

# Creator interface
class NotificationFactory(ABC):
    @abstractmethod
    def create_notification(self):
        pass

# Concerte creators
class SMSNotificationFactory(NotificationFactory):
    def create_notification(self) -> Notification:
        return SMSNotification()
    
class EmailNotificationFactory(NotificationFactory):
    def create_notification(self) -> Notification:
        return EmailNotitication()
    

def notify(factory: NotificationFactory, message: str):
    notification = factory.create_notification()
    notification.send(message)
if __name__ == "__main__":
    
    factories = {
        "email": EmailNotificationFactory(),
        "sms": SMSNotificationFactory(),
    }

  
    notify(factories["sms"], "Hello ...")
   