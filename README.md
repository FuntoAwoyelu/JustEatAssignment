#
## Overview
This CLI application fetches restaurant data based on a given postcode and displays it in the console. It provides information about the restaurant name, cuisines, star rating, and address.


## How to Build, Compile, and Run the Solution

## Clone the Repository:
`git clone <repository-url>`
## Navigate to the Project Directory:
`cd restaurant-cli`
## Compile and Run the Application:
`go run main.go`
## Follow the Prompt:
- After running the command, the application will prompt you to enter a postcode.
- Input the postcode when prompted and press Enter.
- The application will then fetch and display information about restaurants in the specified area.

#
## Assumptions and Clarifications
- The application prompts the user to input the postcode when executed, making it interactive.
- It only displays information for the first 10 restaurants to avoid overwhelming the console with too much data.
- Cuisine information is formatted as a comma-separated string for better readability.

#
## Possible Improvements
- **Pagination:** Implement pagination for fetching and displaying restaurant data, enabling users to navigate through multiple pages of results.
- **Error Handling:** Enhance error handling txo provide more informative messages to users in case of API request failures or other issues.
- **Unit Testing:** Add unit tests to ensure the reliability and stability of the application, covering different scenarios and edge cases.
- **Enhanced Display:** Improve the display format to make it more visually appealing and user-friendly. For example, presenting data in a tabular format or adding color coding for different elements.