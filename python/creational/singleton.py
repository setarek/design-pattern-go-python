class Singleton:
    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super(Singleton, cls).__new__(cls, *args, **kwargs)
        return cls._instance
    
if __name__ == "__main__":

    singleton = Singleton()
    singleton2 = Singleton()

    print(singleton is singleton2)