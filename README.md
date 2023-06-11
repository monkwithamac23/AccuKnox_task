# AccuKnox_task

# Menu Item Consumption

This project is a menu item consumption program that reads a log file containing eater IDs and food menu IDs, and provides functionality to calculate the top consumed menu items.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/AccuKnox_task.git
Navigate to the project directory:

bash
Copy code
cd AccuKnox_task
Build the Docker image:

bash
Copy code
docker build -t AccuKnox_task .
Usage
Run the Docker container:

bash
Copy code
docker run -d -p 8080:8080 AccuKnox_task
Access the application by opening a web browser and navigating to http://localhost:8080.

Follow the on-screen instructions to provide a log file for processing and view the top consumed menu items.

Testing
To run the tests for this project, execute the following command:

bash
Copy code
go test -v
The tests validate the functionality of the log file processing and the accuracy of the top menu item retrieval.

Contributing
Contributions to this project are welcome! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.
