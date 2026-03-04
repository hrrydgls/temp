
from database.connect import connect

def register():
    
    connect()

    print("User registered successfully.")

if __name__ == "__main__":
    register()