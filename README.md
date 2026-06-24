# ASCII Art Web - Zone01

An interactive web application built in Go that transforms user-inputted text into stylized graphic ASCII art using distinct typography banner layouts. 

## 🚀 Features
* **Live Generation:** Convert standard English characters into block-style ASCII visual text.
* **Banner Styles:** Supports three official core assets: `Standard`, `Shadow`, and `Thinkertoy`.
* **Safe Input Filtering:** Robust error handling preventing dangerous system payloads or unsupported characters.
* **Responsive Layout:** Clean UI design displaying persistent output fields and dedicated state handling.

---

## ⚙️ Project Architecture & Design Pattern

The project relies on a clean, scalable architectural split to enforce standard web software paradigms:

```text
ascii-art-web-export-file/
├── banners/               # Target layout text fonts (.txt assets)
├── internal/              # Core proprietary runtime execution packages
│   ├── banner/            # Safe text file system input parsing
│   ├── handlers/          # HTTP request control pipelines & state evaluation
│   ├── render/            # Multi-layer string graphic rendering algorithms
│   ├── server/            # Endpoint initialization & asset distribution routing
│   └── validation/        # Payload structural health constraints
├── static/                # Public assets, browser styles (CSS), & favicons
├── templates/             # Front-end layout configurations (HTML templates)
├── cmd/main.go            # Central operational system startup entry point
└── go.mod                 # Go operational workspace manifest dependency configuration
```

---

## 🔄 Logical Execution Flow (Order of Operations)

When a browser connects or invokes actions within the server environment, the software stack executes dependencies down a specific functional pipeline:

### 1. Application Initialization (`main.go`)
* **Role:** The main operational execution layer.
* **Process:** Prints starting diagnostics on the host terminal and signals the core server layer to initiate continuous web socket listening.

### 2. Networking and Endpoint Setup (`internal/server/`)
* **Role:** Establishes communication rules and operational endpoints.
* **Process:** Builds exact routing paths, configures security handlers for resource paths, serves external stylesheets from `/static/` directories via a static file server, and monitors TCP networks over port `:8080`.

### 3. Request Orchestration (`internal/handlers/`)
* **Role:** Evaluates user interaction contexts and routes network traffic statuses.
* **Process:** 
  * Rejects unsafe methods (e.g., throwing a `405 Method Not Allowed` header if endpoints receive unsupported request actions).
  * Safely reads form parameters and determines which client responses are required depending on success or internal runtime structural problems.

### 4. Payload Interception (`internal/validation/`)
* **Role:** The defensive security gateway.
* **Process:** Inspects text bodies character-by-character to protect execution memory pipelines from anomalies or malicious paths, ensuring strings stick strictly to printable bounds (ASCII characters 32 to 126). It also sanitizes asset calls against unexpected style properties.

### 5. Storage Access (`internal/banner/`)
* **Role:** File system read pipelines.
* **Process:** Resolves localization variables into path directions, opens the server font text database on runtime demand, and handles translation rules safely separating variations between platforms (e.g., stripping down hidden Windows Carriage Returns `\r`).

### 6. Typographic Render Engine (`internal/render/`)
* **Role:** Algorithmic calculation layer.
* **Process:** Maps system array lines directly to character index formulas (`(char - 32) * 9 + 1`). It dynamically converts layout inputs into clean 8-row structural blocks, handles inner lines, and trims trailing components cleanly to avoid failures during automated audit steps.

### 7. Interface Execution (`templates/`)
* **Role:** Client-side interface rendering.
* **Process:** Merges algorithmic string builds straight into responsive `<pre>` output environments inside the HTML layouts, preserving your data settings and highlighting error boxes instantly whenever standard criteria fail.

---

## 💻 How to Run & Use

### Prerequisites
Make sure you have **Go** installed on your system (version 1.20 or higher recommended).

### 1. Clone the repository
```bash
git clone <your-repository-url>
cd ascii-art-web
```

### 2. Start the Local Server
Execute the application entry point using the standard Go environment command:
```bash
go run ./cmd
```

You should see an confirmation terminal diagnostic logging out:
```text
Server starting at http://localhost:8080
```

### 3. Open in Browser
Open your preferred web browser and navigate to:
```text
http://localhost:8080
```

### 4. Create Art
1. Enter your text inside the input box (supports multi-line inputs).
2. Choose a banner type style preference (`Standard`, `Shadow`, or `Thinkertoy`).
3. Press **Generate ASCII Art** to instantly display your output.

---

## 🛠️ Error Codes Standard Map

This project adheres tightly to standard HTTP protocol metrics during verification checks:
* **`200 OK`**: Layout strings resolved cleanly without errors.
* **`400 Bad Request`**: Submissions included non-ASCII entities, unsupported arguments, or corrupt fields.
* **`404 Not Found`**: Request directed to an unregistered path.
* **`405 Method Not Allowed`**: Request targeted endpoints with unsupported HTTP methods.
* **`500 Internal Server Error`**: Core system dependencies or font system resources are missing or broken.

---

## Run Test Files 

From the root
```
go test ./... -v
```

## 👥 Authors
* **elgeorgiou** - Developer / UI & Render Engineering / Frontend Engineer
* **gpapadaki** - Developer / Security Optimization / Backend Engineer
