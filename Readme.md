# ⌯Go Conference Booking App

This is a simple booking application for the Go Conference 2023. It allows users to book tickets for the conference, tracks the remaining tickets, and sends a confirmation email to the user.

## Getting Started

To run the application, follow these steps:

1. Ensure you have Go installed on your machine.
2. Clone the repository or download the code files.
3. Open a terminal or command prompt and navigate to the project directory.
4. Run the following command to build and run the application:

```shell
go run .
```

## Features

The Go Conference 2023 Booking App offers the following features:

1. User Input: Users are prompted to enter their first name, last name, email, and the number of tickets they want to book.
2. Validation: User inputs are validated to ensure the name is not too short, the email contains an "@" sign, and the number of tickets is valid.
3. Booking Tickets: If the user input is valid, the application books the tickets, updates the remaining ticket count, and stores the user data.
4. Email Confirmation: A separate goroutine is used to send a confirmation email to the user, simulating a delay with a sleep function.
5. Display Bookings: After each successful booking, the application displays the list of all bookings made so far, showing the first names of the attendees.
6. Ticket Availability: The application checks the remaining ticket count and informs the user if all tickets have been sold out.
7. Concurrency: The program uses the `sync.WaitGroup` to ensure that the email goroutine completes before the program exits.

## Dependencies

The Go Conference 2023 Booking App has no external dependencies.

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions to this project are welcome! If you find any issues or want to add new features, please open an issue or submit a pull request.
