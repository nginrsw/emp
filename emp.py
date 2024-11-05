import os
import time

def run_command(command):
    """Run a command in the terminal."""
    os.system(command)

def start_services():
    """Start Nginx and MySQL."""
    os.system("clear")
    print("Starting The Nginx with MySQL...")
    print("--------------------------------->")
    run_command("sudo systemctl enable nginx mysql.service")
    
    run_command("sudo systemctl start nginx mysql.service")
    print("---------------------------------")
    run_command("sudo systemctl status nginx")
    print("---------------------------------")
    run_command("sudo systemctl status mysql.service")


def stop_services():
    """Stop Nginx and MySQL."""
    os.system("clear")
    print("Stopping The Nginx and MySQL...")
    print("--------------------------------->")
    run_command("sudo systemctl stop nginx mysql.service")
    
    run_command("sudo systemctl disable nginx mysql.service")
    print("---------------------------------")
    run_command("sudo systemctl status nginx")
    print("---------------------------------")
    run_command("sudo systemctl status mysql.service")

def sites_enable_enginx():
    os.system("clear")
    """Monitoring Nginx using multitail """
    print("Checking Sites-Enable Config...")
    print("---------------------------------")
    time.sleep(0.5)
    run_command("ls -all /etc/nginx/sites-enabled")
    time.sleep(0.5)
    print("\nChecking Preps...")
    print("---------------------------------")
    run_command("sudo nginx -t")
    print("---------------------------------")
    print(" ")
    """kill the function and return to the main loop in main()"""
    while True:  # Loop to ensure input is valid
        sub_choice = input("\033[F Press 'q' to go back to menu: ")

        if sub_choice == "q":
            break  # If input is 'q', exit the loop and return to main
        else:
            print("\033[F Invalid, Please Try Again....")
            time.sleep(0.7)

def check_services():
    """Check Nginx and MySQL."""
    os.system("clear")
    print("Checking The Nginx and MySQL...")
    time.sleep(1.7)
    print("---------------------------------")
    run_command("sudo systemctl status nginx")
    time.sleep(0.8)
    print("---------------------------------")
    run_command("sudo systemctl status mysql.service")
    
def monitor_nginx():
    """Monitoring Nginx using multitail """
    os.system("clear")
    print("Monitoring Nginx...")
    time.sleep(1)
    run_command("sudo multitail /var/log/nginx/access.log /var/log/nginx/error.log")


def main():
    """Main menu of the program."""
    while True:
        os.system('clear')
        print("+===========+")
        print("Nginx, Mysql Panel")
        print("-------------")
        print("EMP-Menu")
        print("+===========+")
        print("1. Turn On Engine")
        print("2. Turn Off Engine")
        print("3. Check Preps")
        print("4. Check Engine")
        # If you want to use option number 4,
        # make sure you have previously used option number 1, or 1 then 2
        # or option number 4 will look like it's not working properly
        print("5. View Logs")
        print("0. Bail Out")
        print("-------------")
        print("Always press 'q' to back menu")
        print("+===========+")
        
        choice = input("Choose an option (1/2/3/4/5/0): ")
        
        if choice == '1':
            start_services()
        elif choice == '2':
            stop_services()
        elif choice == '3':
            sites_enable_enginx()
        elif choice == '4':
            check_services()
        elif choice == '5':
            monitor_nginx()
        elif choice == '0':
            print("Exiting the program.")
            time.sleep(0.4)
            os.system('clear')
            break
        else:
            print("Invalid choice. Please try again.")
            time.sleep(0.4)

if __name__ == "__main__":
    main()
