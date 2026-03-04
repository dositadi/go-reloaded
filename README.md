[![Go Report Card](https://goreportcard.com/badge/github.com/dositadi/go-reloaded)](https://goreportcard.com/report/github.com/dositadi/go-reloaded)
![Go Version](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

# 📝 TextEdit Tool: Automated Text Processor

# 📖 Table of Contents

[1. Project Structure](#-project-structure)

[2. Overview](#-overview)

[3. Features](#-features)

[4. Architectural and Logical Flow](#-architecture--logic-flow)

[5. Technical Function Breakdown](#-technical-function-breakdown)

[6. Logic Flow: Step-by-Step](#-logic-flow-step-by-step)

[7. Usage](#-usage)

[8. Error Handling](#️-error-handling)

[9. Getting Started](#-getting-started)

[10. Testing](#-testing)

[11. License](#-license)

#

### 📂 Project Structure

```plaintext
A professional README should show the user how the code is organized. This helps non-technical users see the "order" and developers find the "logic."
```

```plaintext
.
├── asset/                  # Sample input text files
├── cmd/
│   └── main.go             # Application entry point
├── internal/
│   ├── config/             # Engine orchestration and App configuration
│   └── worker_handlers/    # Core transformation logic (The "Workers")
├── output/                 # Generated results after processing
├── pkg/
│   ├── models/             # Custom error structs and data models
│   └── utils/              # Linguistic and mathematical helper functions
├── test/                   # Comprehensive unit and system tests
├── go.mod                  # Go module definition
└── README.md               # Project documentation
```

### 📖 Overview

```plaintext
The TextEdit Tool acts as an automated editor. It reads an input file, processes the text through a series of "Engine" handlers, and writes the corrected version to an output file. It is specifically designed to handle numerical base conversions, casing adjustments, and complex punctuation rules.
```

#


### 🚀 Features

```plaintext
1. Advanced Casing Control
Beyond simple single-word changes, the tool supports indexed transformations:

a. (up, n): Converts the previous n words to UPPERCASE.

b. (low, n): Converts the previous n words to lowercase.

c. (cap, n): Capitalizes the first letter of the previous n words.

2. Numerical Conversion
a. (hex): Automatically detects the preceding word as Hexadecimal and converts it to Decimal.

b. (bin): Detects the preceding word as Binary and converts it to Decimal.

3. Smart Punctuation & Grammar
a. Quote Wrapping: Identifies pairs of single quotes (') and intelligently wraps them around the content without internal spaces (e.g., ' word ' becomes 'word').

b. Vowel Correction: Automatically changes "a" to "an" if the following word starts with a vowel or 'h'.

c. Punctuation Spacing: Ensures symbols like . , ! ? : ; are correctly attached to the preceding word.
```

#

### 🏗 Architecture & Logic Flow

```plaintext
The application follows a modular "Worker/Engine" pattern to ensure clean data processing.

The "Step-by-Step" Flow:
1. Initialization (main.go & app.go): The application validates command-line arguments to ensure valid input/output paths are provided.

2. File Ingestion (read_file.go): The ReadFile function converts the raw file buffer into a "slice of words" (Tokenization) for easier manipulation.

3. The Engine Sequence (engine.go): The EngineWorker runs the text through a strict pipeline of handlers:

a. ActionWithIndex: Handles multi-word casing.

b. StringBinHex: Handles base conversions.

c. PunctuationMark: Handles single quote logic.

d. VowelManipulator: Fixes grammar.

e. Final Assembly: The processed tokens are joined and written to the destination file.
```

#

### 🛠 Technical Function Breakdown

```plaintext
1. The Transformation Engine (workerhandlers)
This layer implements the Engine interface to modify text slices directly.

a. ActionWithIndex: Handles multi-word casing.

- Scans the text for tags like (up, n), (low, n), or (cap, n).

- Uses SplitAndGetIndices to determine how many words to look back.

- Modifies the specified range and deletes the trigger tag from the output.

b. StringBinHexManipulator: Handles single-word conversions.

- Identifies standard tags: (up), (low), (cap), (bin), and (hex).

- Converts Hexadecimal and Binary strings into their Decimal equivalents.

c. SingleAndGroupOfPunctuationManipulator: Manages grammar symbols.

- Ensures punctuation (single or groups like ...) is attached to the word before it.

- If a word starts with punctuation, it moves that mark to the previous word.

d. VowelManipulator: Implements the "a" vs "an" rule.

- Checks if a word starts with a vowel or 'h'.

- If the preceding word is "a", it automatically updates it to "an".

e. PuntuationMarkManipulator: Specifically handles single quote (') pairs by identifying their indices and wrapping them tightly around the enclosed text.

2. Utility & Helper Logic (utils)
These functions perform the granular operations required by the Workers.

a. MergeActionSeparatedWithSpaces: A pre-processing step that rejoins fragmented tags (e.g., merging ( and up, 2) into one string).

b. ToUpper / ToLower / Cap: Directly modifies string pointers to change their case.

c. Bin / Hex: Uses strconv.ParseInt to convert numerical bases into base-10 strings.

d. NotePMIndices: Scans for the positions of single quotes to determine opening and closing pairs.

e. FixPunctMarks: Determines the distance between quote pairs to apply the correct formatting logic (ModifyMid for single words vs. ModifyFirst/SecondMid for phrases).
```
#


### 🔄 Logic Flow: Step-by-Step
```plaintext
The EngineWorker manages the execution order to ensure that one transformation doesn't interfere with another:

1. Ingestion: ReadFile tokenizes the file into a slice of words.

2. Step 1: ActionWithIndex: Complex casing (multi-word) is resolved first.

3. Step 2: StringBinHexManipulator: Standard conversions and single-word casing are applied.

4. Step 3: PuntuationMarkManipulator: Single quotes are wrapped around their content.

5. Step 4: SingleAndGroupOfPunctuationManipulator: Standard punctuation is correctly spaced.

6. Step 5: VowelManipulator: Final grammatical check for vowel-start words.

7. Finalization: WriteResult joins the tokens back into a string and saves the file.
```

#

### 💻 Usage

```bash
go run . sample.txt result.txt
```

```plaintext
Input Example:

"it is a amazing rock ! 10 (bin) years ago , ' i ' was in 1E (hex) mode (up) ."

Output Result:
"It is an amazing rock! 2 years ago, 'i' was in 30 MODE."
```

#

### ⚠️ Error Handling

```plaintext
The tool uses a custom Error model to provide descriptive feedback. It will alert you if:

1. The input file is missing or unreadable.

2. The file is empty.

3. Incorrect syntax is used inside a tag (e.g., (up, non-number)).
```

#

### 🏁 Getting Started
#### Prerequisites
#### Go: Version 1.21 or higher. You can check your version by running go version.

#### Git: To clone and manage the repository.

### Installation
#### 1. Clone the repository:
```bash
git clone https://github.com/dositadi/go-reloaded.git
cd go-reloaded
```

#### 2. Download dependencies:
#### (Even if you only use the standard library, running this ensures your go.mod is synced):
```bash
go mod tidy
```

###  Running the Application
#### To process a file, use the following command structure from the root directory:
```bash
go run cmd/main.go <input_file_path> <output_file_path>
```

#### Example

```bash
go run cmd/main.go asset/sample.txt output/result.txt
```
#

### 🧪 Testing
#### The project includes a suite of tests to ensure the reliability of the transformation workers. To run the tests, execute:
```bash
go test ./test/...
```

#

### 📜 License

```plaintext
This project is licensed under the MIT License.

MIT License Summary:
✅ Commercial Use: You can use this for-profit.

✅ Modification: You can change the code.

✅ Distribution: You can share the code.

⚠️ Liability: The author provides no warranty.

⚠️ Attribute: You must include the original copyright notice in any copy of the software.
```

See the [LICENSE](LICENSE) file for the full text.

#