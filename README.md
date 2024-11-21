# Tweet Analysis and Visualization in Go

This Go-based project analyzes a CSV file of tweets and presents its data through various charts and visualizations. The analysis includes aspects such as word frequency, sentiment analysis, and other metrics that help to gain insights from the tweet data. The results are displayed using different chart types (e.g., word clouds, bar charts, pie charts) to make the data easier to understand.

## Features

- **CSV Parsing**: Reads tweet data from CSV files.
- **Text Analysis**: Processes tweets to analyze word frequencies, symbol frequencies, and tweet frequency per year, each month
- **Chart Generation**: Generates interactive charts including word cloud, frequency bar chart, and pie chart.
- **Built with Go**: Utilizes the Go programming language for fast, efficient data processing.

## Prerequisites

To run this project, ensure that you have the following installed:

- **Go**: The Go programming language (version 1.18 or later).
- **Go Modules**: This project uses Go modules for dependency management.

You also need a CSV file containing tweet data. The CSV file should follow the format:
headers: [tweet_id, user_id, date_created, text, language]
wherein date_created: dd/mm/yyyy h:mm

## Installation

1. **Clone the repository**:

```powershell
git clone https://github.com/KaoM3/Corpus-Analysis-and-Visualization-in-Go.git
```

2. **Install dependencies**:

Ensure you have Go installed, then run:

```powershell
go mod tidy
```

## Usage
**Step 1: Build the program**
Make sure your CSV file is properly formatted and contains the tweet data you want to analyze.
You can place the file in the project's directory or specify the file path when running the program. To build the program, use the following command:

```powershell
go build -o MCO2_9_Go.exe ./src/main
```

**Step 2: Run the Program**
To run the Go program, use the following command:

```powershell
./MCO2_9_Go.exe
```

**Step 3: Give the necessary input**
The program will prompt you to enter the file path of the CSV file. Include the .csv file extension and remove the "".

**Step 4: View the Charts**
After running the program, the charts will be generated and saved to the disk as html files. You can view these charts in your browser.


## Code Structure
`main`
- main package of the program
- MCO2_9_Go.go: The entry point of the application. It handles reading of the CSV file.
- cleaner.go: Contains helper functions for data parsing and cleanup.
- datavis.go: Contains the functions for rendering the data visualization.

## Acknowledgements
- Go ECharts: For generating interactive charts in Go. "github.com/go-echarts/go-echarts/v2"

